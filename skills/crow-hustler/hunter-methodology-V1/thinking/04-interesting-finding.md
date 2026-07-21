> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index|← Back to Thinking]] · Related: [[01-auth|Chaining]]

# When You Find Something Interesting

Found a reflection, an error, an unexpected response? Here's the protocol.

## Immediate Actions

1. **Don't get excited.** Excitement leads to sloppy validation.
2. **Save the raw request+response.** Full HTTP conversation.
3. **Re-test from clean session.** No leftover state.
4. **Test authenticated AND unauthenticated.** (for stateful bugs)

## Validation Checklist

- Can you reproduce it reliably? (3/3 times)
- Is it in scope? (re-read program policy)
- Is it a known issue? (search disclosed reports)
- What's the minimum viable payload?
- Can you chain it? (escalate severity)

## The 7-Question Validation Gate

Before writing any report, run these:

1. **Is it in scope?** — asset in scope, vuln class not excluded
2. **Is there real impact?** — can you state it as a business consequence?
3. **Can I reproduce it?** — fresh browser, clean account, 3 times
4. **Is the PoC clean?** — no real user data, no destructive actions
5. **Is severity defensible?** — CVSS 3.1 justified?
6. **Has it been reported before?** — check disclosed reports for dupes
7. **Does the title sell the impact?** — impact formula applied?

If any answer is "no," don't submit yet.

## Before Telling Anyone

- Document: exact URL, method, headers, body, response
- Test on a different browser / device
- Build the full chain if base bug is P3/P4 alone

## Don't

- Don't submit a "maybe" — "An attacker might" = N/A
- Don't submit without business impact
- Don't share on Discord/Twitter before triage
