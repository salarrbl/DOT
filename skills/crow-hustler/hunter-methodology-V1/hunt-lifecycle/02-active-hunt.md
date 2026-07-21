> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|← Back to Hunt Lifecycle]] · Related: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index|Thinking]]

# Active Hunt

## During the Hunt

### One Hypothesis at a Time

1. State hypothesis: "Parameter X reflects in endpoint Y"
2. Test with 3 requests
3. Result: confirmed / rejected
4. Log result → move to next hypothesis

### Track Everything

- Every request that matters → save to hunt log
- Every response that matters → save raw
- Every dead end → note why it was wrong

### Stuck Timer

If >30 min on one hypothesis → it's probably wrong.
If >2h total without a finding → activate unblock protocol.

### Automation Philosophy

Your automation is a **research assistant**, not a noise machine. Build a pipeline that:
1. Finds new subdomains daily
2. Screenshots and fingerprints tech
3. Highlights ones matching high-value patterns (`admin`, `api`, `dev`)
4. Alerts you via Discord/Slack
5. Queues for your manual review

Automate **intelligence**, not **discovery**.

### What to Save

- Request/response pairs (curl format)
- Screenshots (if browser-based)
- Payloads that worked
- Payloads that didn't (and why)
