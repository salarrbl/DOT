> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Cache Poisoning

## Payloads

- `X-Forwarded-Host: evil.com` → reflected into `<link>` / `<script src=>` → cache poisoned response
- `X-Forwarded-Scheme: http` → canonical URL SSL bypass
- Akamai/Cloudflare `Transfer-Encoding: chunked` smuggling
- Check `__cfduid` / `cf-cache-status` headers for cacheable pages
