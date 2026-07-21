> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/tools/00-index|← Back to Tools]] · Related: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|Recon]]

# Recon Tools

## Subdomain Discovery

| Tool | Use |
|------|-----|
| `subfinder -d target -all -silent` | Passive subdomain discovery |
| `puredns resolve subs.txt -r ~/.resolver` | Wildcard-resilient resolution |
| `shuffledns -d domain -w wordlist.txt -r resolvers.txt` | DNS brute-force |
| `massdns` | Bulk DNS resolution |
| `amass enum -d target` | Passive + active |

## Service Discovery

| Tool | Use |
|------|-----|
| `httpx -l subs.txt -flc 12,0,11 -silent` | Alive HTTP detection |
| `naabu -p 80,443,8080,8443 -silent` | Port scanning |

## Crawling

| Tool | Use |
|------|-----|
| `katana -list alive.txt -d 3 -jc -ef css,png,svg,jpg,woff` | Modern crawler |
| `gospider` | Backup crawler |
| `hakrawler` | Lightweight crawler |

## Permutation

| Tool | Use |
|------|-----|
| `dnsgen` | Permutation from existing subs |
| `alterx` | Advanced permutations |
