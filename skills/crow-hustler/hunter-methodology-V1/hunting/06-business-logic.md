> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Business Logic

Where scanners fail and humans win. Highest-paying class because automation can't catch it.

## The 4-Question Framework

For every endpoint, ask:

1. **What is this supposed to do?** — understand normal flow
2. **What data can I control?** — params, headers, filenames, body fields
3. **Where does this data go?** — reflected, stored, queried, forwarded
4. **What if I break the expected flow?** — swap IDs, skip steps, reorder, send arrays

## Categories

| Category | Examples |
|----------|----------|
| **Price manipulation** | Negative quantity, currency confusion (pay in JPY/INR while priced in USD), floating-point rounding |
| **Coupon abuse** | Reuse single-use via race condition, apply multiple coupons, expired codes still valid |
| **Workflow skipping** | Skip payment step, bypass email verification, submit order without checkout |
| **Quota/limit bypass** | Free tier abuse, API rate limit via header tricks |
| **State machine violations** | Do step C before A, invoke refund on already-refunded order |
| **Parameter pollution** | Duplicate params `?user=A&user=B` resolved differently by frontend/backend |
| **Test card abuse** | Production accepts Stripe test cards (`4242 4242 4242 4242`) |

## Chain Patterns

| Low bugs | Chained into | Result |
|----------|-------------|--------|
| IDOR + email verification bypass | ATO | Critical |
| SSRF + cloud metadata | Cloud account compromise | Critical |
| XSS + CSRF bypass | Full session hijack via admin | High/Critical |
| Open redirect + OAuth flow | OAuth token theft | Critical |
| Subdomain takeover + cookie scope | Session cookie theft | Critical |
| File upload + path traversal | Webroot file write | High/RCE |
| Prototype pollution + HTML injection | DOM XSS with CSP bypass | High |

## 2026 Advanced Chains

| Modern chain | Components | Impact |
|-------------|-----------|--------|
| **AI + SSRF** | Prompt injection → LLM URL fetch → cloud metadata | Cloud compromise |
| **Browser-in-Browser** | CSS/JS fake browser → credential harvesting → MFA bypass | Account takeover |
| **Service Worker Persistence** | XSS → malicious SW install → persistent access | Long-term compromise |
| **SAML + XXE** | XXE in SAML request → file read → signing key theft | Identity system compromise |
| **WebAuthn + CORS** | CORS misconfig → cross-origin WebAuthn → credential theft | Auth bypass |
| **Graph-based Auto-discovery** | Multiple low-severity → automated chain detection → critical path | Automated escalation |
