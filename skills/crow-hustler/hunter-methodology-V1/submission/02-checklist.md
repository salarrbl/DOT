> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/submission/00-index|← Back to Submission]]

# Pre-Submission Checklist

| Check | Required |
|-------|----------|
| Reproducible from clean session? | Yes |
| Tested authenticated AND unauthenticated? | Yes (stateful bugs) |
| Chains documented explicitly? | Yes (if P3/P4 base) |
| Not duplicates disclosed before? | Search H1 / Bugcrowd disclosed |
| Not OOS by program rules? | Re-read entire policy |
| Token named correctly? | Yes |
| PoC uses placeholder data (no real users)? | Yes |
| Cleaned up test data after? | Yes |

## Severity Scoring (CVSS 3.1 Quick Ref)

| Severity | CVSS Range | Examples |
|----------|-----------|----------|
| **Critical** | 9.0–10.0 | RCE, full ATO, massive data leak, full cloud compromise |
| **High** | 7.0–8.9 | SQLi without RCE, stored XSS with session theft, IDOR on PII, SSRF to internal |
| **Medium** | 4.0–6.9 | Reflected XSS, limited IDOR, CSRF with significant action |
| **Low** | 0.1–3.9 | CSRF on low-risk action, minor info disclosure, open redirect alone |

Don't inflate. Triagers see through it. Downgrades hurt your reputation.

## Always-Rejected Findings (NA List)

- Self-XSS (victim attacks themselves)
- Missing security headers without exploitation
- SPF / DMARC / DKIM issues
- Rate limit absence on non-sensitive endpoints
- Username/email enumeration without impact
- Clickjacking on pages without state-changing actions
- CSRF on logout / unauthenticated endpoints
- Information disclosure in error messages (unless secrets)
- HTTPS / TLS configuration nits
- DoS via resource exhaustion
- CVEs in outdated libraries without working exploit on THIS app

## Status Tracking

When you submit, tag it:

- **Resolved** — paid / triaged-confirmed
- **Duplicate** — already known
- **NA** — not applicable
- **Out-Of-Scope** — valid but OOS
- **Info-P5** — valid, P5

If NA/dup > 30% of submissions → hunting-misalignment signal. Fix methodology.
