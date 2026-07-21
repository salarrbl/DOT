> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/submission/00-index|← Back to Submission]]

# Report Structure & Golden Rules

## Universal Report Structure

```
1. Title        — [Vuln Class] at [Endpoint] leads to [Impact]
2. Summary      — 2-3 sentences: what, where, why it matters
3. Severity     — CVSS 3.1 with justification
4. Steps to Reproduce — numbered, exact, copy-pasteable
5. Proof of Concept — HTTP request/response, screenshots, video
6. Impact       — business consequence in plain English
7. Remediation  — actionable fix
8. References   — OWASP / CWE links
```

## Title Formula

`[Bug type] in [specific feature/endpoint] allows [specific attacker action]`

**Good:** "Stored XSS in profile bio field allows session hijacking of any user viewing the profile"
**Bad:** "XSS on target.com"

## Impact Writing

Impact is what a real attacker could do, not a textbook definition.

**Weak:** "SSRF vulnerability allows internal port scanning."
**Strong:** "SSRF in the image import feature allows an unauthenticated attacker to read AWS IAM credentials from the instance metadata service, enabling full compromise of the production AWS account."

## 6 Golden Rules

1. **Never send two bugs in one report.** Two bugs, two reports, two payments.
2. **No "maybe" in attack scenario.** "An attacker might" = N/A. "An attacker can" = report.
3. **Two-sided reproduction.** Attacker steps + victim steps (for session/cookie/CSRF bugs).
4. **No jargon for triage.** They are not engineers. Plain-language impact.
5. **PoC video < 2 minutes.** Long videos → triagers give up.
6. **Demonstrate chain worth.** XSS alone is P3/P4. Chain it to ATO before submission.
