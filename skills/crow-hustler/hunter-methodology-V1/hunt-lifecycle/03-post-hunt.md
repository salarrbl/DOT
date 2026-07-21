> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|← Back to Hunt Lifecycle]] · Related: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/submission/00-index|Submission]]

# Post-Hunt

After finishing a session, run this checklist:

## Immediate (end of session)

- [ ] Save all request/response logs to `hunt-logs/` directory
- [ ] Save what worked (payloads, endpoints, techniques)
- [ ] Save what didn't work (save time next session)
- [ ] Note todos for next session
- [ ] Note what you learned with source references

## File Updates

After every hunt, update:
- `../hunting/` — new bypass technique? add it. New payload? add it.
- `../bypass/` — new bypass discovered? document it.
- `../payloads/` — new payload that worked? save it.
- `../thinking/` — new stuck scenario? add to unblock protocol.

## Weekly

- [ ] Review submission status (resolved/dup/NA/OOS)
- [ ] Update target state
- [ ] Check if >30% of submissions are NAs → methodology issue
