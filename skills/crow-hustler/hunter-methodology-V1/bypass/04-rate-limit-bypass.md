> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]]

# Rate Limit Bypass

- `X-Forwarded-For: 1.2.3.4` — rotate IPs
- Trailing slash, dot, encoded chars → cache fragmentation
- HTTP verb switch (POST → GET)
- `Accept-Encoding: gzip` → different code path
- Case toggle URL/path
- Every header: `X-Client-IP`, `X-Real-IP`, `True-Client-IP`, `CF-Connecting-IP`
- Multi-account: 5 accounts, each does 1/5 of brute-force
- API key from another user → quotas per-key
