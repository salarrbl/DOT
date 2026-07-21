> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]] · Related: [[../../bypass/05-oauth-bypass|OAuth bypass]] · [[../../thinking/04-interesting-finding|Chaining]]

# Authentication Classes

## Registration

Tests in order:
- Email parameter manipulation (pollution, array, CC, separator, null byte, template injection)
- Duplicate registration takeover
- Post-registration directory probe
- Undocumented signup endpoints (`POST /api/users`, `/_signup`, `/register.json`)

## Forgot Password — 9 Universal Tests

1. **Token reuse** — request twice, use first link after second
2. **Token in response** — inspect JSON responses carefully
3. **Host-header injection** — `Host: evil.com` poisons reset link
4. **Referer poison** — replace referer with attacker URL
5. **OTP brute + email-alias collision** — `victim+1@gmail.com` same OTP?
6. **Null password** — `password=` bypasses update
7. **XXE** — forgot-password accepts XML
8. **IDN Homograph** — `abc@gmail.com` vs `abc@gmáil.com`
9. **Broken session mgmt** — change email, then use old reset link

## MFA / OTP Bypass

| Test | Behavior |
|------|----------|
| `000000`, `null`, blank | Default stubs fail-open |
| Old OTP reuse | Same OTP accepted across rotations |
| Array trick | `otp[]=`, `otp=[""]`, duplicate params |
| Brute-force no limit | 4-digit OTP owned in 30 min |
| Same code for many users | `a@gmail.com`, `a+1@gmail.com` same OTP |
| Response manipulation | Swap `success: false` → `true` |
| Disable MFA via CSRF | Link via password reset flow |
| Brute SMS cost | P3/P4 even without ATO |

## OAuth

Recon: `/.well-known/oauth-authorization-server`, `/.well-known/openid-configuration`

**State parameter** bugs: Missing, predictable, not-verified, reusable, tamperable.

**Scope escalation**: Change scope during token exchange.

**`redirect_uri` bypass**: double-encode, tab, param injection, Unicode dots, backslash, path traversal.

## ATO Chain Model

| Low bug | Chain target | Outcome |
|---------|-------------|---------|
| IDOR + email-flow | Trust on first login | Critical ATO |
| OAuth state missing | Forcing victim to link | Critical ATO |
| OAuth redirect_uri | Phishing + code steal | Critical ATO |
| SSRF + cloud metadata | EC2/IAM tokens | Critical env |
| Subdomain takeover + cookie scope | Main domain session | Critical cookie theft |
