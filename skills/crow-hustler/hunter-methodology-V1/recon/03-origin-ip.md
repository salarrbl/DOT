> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|← Back to Recon]]

# Origin-IP Discovery

| Technique | Method |
|-----------|--------|
| Historical DNS | viewdns.info, SecurityTrails |
| Subdomain brute → dig A | Forgotten backend dev retains origin IP |
| Favicon hash → Shodan | `curl -s target/favicon.ico \| base64 \| mmh3 hash` → Shodan |
| Copyright / http.html | Shodan / Censys / FOFA |
| Certificate search | SSL cert subject, issuer, SAN on all `*.target.com` |
| CIDR: ASN → entire block | `bgp.he.net` → `echo CIDR \| httpx -silent -duc` |
| SSRF / send-email | Mail header analysis via mxtoolbox |
| Final verification | `echo <ip> <domain>` in /etc/hosts → confirm bypass |

```bash
nmap --script ssl-cert -p 443 <IP>
```
