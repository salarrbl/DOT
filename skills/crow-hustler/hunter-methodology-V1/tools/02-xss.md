> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/tools/00-index|← Back to Tools]] · Related: [[04-xss|XSS playbook]]

# XSS Tools

## Custom: Cexss

Go-based parameter-driven XSS scanner (the hunter's own tool).
See `../../CEXSS.md` or `road.md` for architecture.

Current status: Phase 14/15, ~75% done. Pending: param ranking, result parser, notifications, reports.

## Other XSS Tools

| Tool | Use |
|------|-----|
| `x8 -w params -u URL -X GET POST` | Hidden parameter discovery + reflection |
| `reflex -f urls.txt` | XSS reflection checker |
| Burp `Reflector` extension | Auto-find reflections |
| Burp `DOM Invader` | PostMessage + DOM XSS |
| Burp `GAP` extension | Parameter discovery |

## Scanning Rule

Cexss is for XSS parameter discovery at scale. **Never auto-submit Cexss output.**
Manual confirm before submission: clean session, different browser, chain impact.
