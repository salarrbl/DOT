> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|← Back to Recon]] · Related: [[01-recon|Recon tools]]

# Wide Recon

## Find Organization / Domain

```bash
curl -s "https://crt.sh/?q=%25.<company>&output=json" | jq -r '.[].name_value' | sed 's/\*\.//g' | sort -u
```

## Passive Subdomain Discovery

- crt.sh, Censys, subdomainfinder.c99.nl
- `subfinder -d target -all -silent`
- `https://chaos.projectdiscovery.io`
- Sourcegraph, GitHub code search, Wayback
- Favicon hash → Shodan:
  ```bash
  curl -s "<URL>/favicon.ico" | base64 | python3 -c "import mmh3,sys;print(mmh3.hash(sys.stdin.buffer.read()))"
  # https://www.shodan.io/search?query=http.favicon.hash:<hash>
  ```
- CSP header extraction:
  ```bash
  curl -I -s https://target.com | grep -iE "content-security-policy|CSP" | tr " " "\n" | grep "\." | tr -d ";" | sed 's/\*\.//g' | sort -u
  ```

## Active Subdomain Discovery

```bash
shuffledns -d domain.com -w sub-wordlist.txt -r resolvers.txt -m $(which massdns) -mode resolve -silent
gobuster dns -w sublist.txt -d domain.com -t 50
puredns resolve subs.txt -r ~/.resolver
```

## Permutation Generation

- `dnsgen`, `atldns`, `alterx`
- Patterns: `api`, `api-test`, `api-dev`, `dev-api`, `test-api`, `v1-api`, `dev-api`, `staging-api`, `qa-api`, `uat-api`, `internal-api`, `admin-api`, `mgmt-api`, `backend-api`

## Service Discovery

```bash
httpx -l subs.txt -flc 12,0,11 -silent
naabu -p 80,8000,8080,8880,2052,2082,2086,2095,443,2053,2083,2087,2096,8443,10443 -silent
```

## CDN Cut

Separate CDN IPs from origin IPs. Use ASN mapping, BGP heuristics (`bgp.he.net`).
