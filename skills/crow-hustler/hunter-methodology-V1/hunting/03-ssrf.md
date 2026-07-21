> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]] · Related: [[../../bypass/03-ssrf-bypass|SSRF bypass]]

# SSRF

## Sensitive Parameters

```
url, uri, path, src, dest, redirect, callback, webhook, endpoint, imageUrl, proxy, fetch, load
```

## Entry Points

- Link preview / unfurl
- PDF export / image proxy
- RSS / Atom import
- OAuth callback
- Webhook registration
- "Import from URL" file uploaders
- Server-side image fetch

## Test with Out-of-Band

Use `interactsh` or Burp Collaborator. Confirm DNS + HTTP callback.

## Escalation

- Internal port scan: `http://10.0.0.5:8080/admin`
- Cloud metadata: `http://169.254.169.254/latest/meta-data/`
- File read: `file:///etc/passwd`
- Chain to RCE: metadata → AWS keys → EC2 control

## Filter Bypass

- IP literal: `http://2130706433/` = `127.0.0.1`
- Octal, hex, decimal variants
- `http://[::1]/`
- DNS rebinding
- nio.io, sslip.io
- `gopher://`, `dict://`, `file://`, `ftp://`, `jar://`
