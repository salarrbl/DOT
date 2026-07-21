> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|← Back to Recon]]

# 60-Minute Recon Recipe

For a brand new target when time is limited.

1. `subfinder -d target -all -silent` + `crt.sh`
2. `puredns resolve subdomains.txt -r ~/.resolver`
3. `httpx -l resolved.txt -flc 12 -silent -mc 200,301,302,403,500`
4. `katana -list alive.txt -d 3 -jc -ef css,png,svg,jpg,woff`
5. `ParamSpider` on each domain
6. Manually visit top 10 endpoints with Burp — extract OAuth `.well-known/openid-configuration`, API paths, JS files
7. **Stop. Think.** Form a hypothesis before scanning.
