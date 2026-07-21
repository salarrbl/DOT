> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/learning/00-index|← Back to Learning]]

# Learning Resources (Ranked by Signal)

1. **PortSwigger Web Security Academy** — free, labs, canonical
2. **HackerOne Hacktivity** — real disclosed reports
3. **Bugcrowd University / Intigriti 101** — beginner-friendly
4. **TryHackMe** (bug bounty path)
5. **HackTheBox Academy** — paid, high quality
6. **PentesterLab** — bite-size
7. OWASP WebGoat, JuiceShop, DVWA — practice platforms
8. **YouTube**: NahamSec, STÖK, InsiderPhD, John Hammond, LiveOverflow, IppSec, PwnFunction, Hakluke

## Books (hunter's list — do not add)

- `Web Hacking Arsenal` (Rafay Baloch) — currently reading
- `The Web Application Hacker's Handbook` (Dafydd Stuttard) — canonical reference
- `Real-World Bug Hunting` (Peter Yaworski) — disclosed report case studies
- `Bug Bounty Bootcamp` (Vickie Li) — modern entry-level guide
- `The Tangled Web` (Michal Zalewski) — browser internals
- `The Body Keeps the Score` (Bessel van der Kolk)
- `Bushido - The Soul Of Japan` (Inazō Nitobe)

## Key Repositories

- `EdOverflow/bugbounty-cheatsheet`
- `nahamsec/Resources-for-Beginner-Bug-Bounty-Hunters`
- `ngalongc/bug-bounty-reference` (curated writeups by class)
- `enaqx/awesome-pentest`
- `sehno/Bug-bounty` (checklists)
- `ZephrFish/BugBountyTemplates`
- `0xPugal/One-Liners` — one-liner collection

## High-Value Writeups to Study

| Researcher | Bug | Takeaway |
|-----------|-----|----------|
| Sam Curry + Shubham Shah | Universal XSS via Netlify Next.js | Shared infrastructure = shared vulns |
| Frans Rosén | OAuth hijacking via "dirty dancing" | postMessage + third-party XSS → ATO |
| James Kettle | Browser-powered desync attacks | HTTP smuggling not dead |
| Jacopo Tediosi | Worldwide Akamai cache poisoning ($50k) | Test every CDN layer |
| Simon Scannell | Zimbra memcache injection | Deep protocol knowledge beats surface testing |
| Felix Wilhelm | Hacking the Cloud with SAML | SAML/SSO has enormous attack surface |
| Neil Madden | Psychic Signatures in Java (CVE-2022-21449) | Classical crypto bugs still ship |
