> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]]

# Header Injection Tests

Headers to test for auth bypass, cache poisoning, and routing manipulation:

```
X-Forwarded-For: 127.0.0.1
X-Forwarded-Host: evil.com
X-Original-URL: /admin
X-Rewrite-URL: /admin
X-Custom-IP-Authorization: 127.0.0.1
X-Host: evil.com
X-Forwarded-Server: evil.com
X-Forwarded-Scheme: http
Referer: https://target.com/admin
```

## When to Use

- 403 bypass (auth bypass)
- Cache poisoning (unkeyed headers)
- SSRF redirect control
- Host header injection
