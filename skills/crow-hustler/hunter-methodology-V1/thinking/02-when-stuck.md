> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index|← Back to Thinking]] · Related: [[02-anti-burnout|Anti-burnout]]

# When Stuck

## The 5-Step Unblock Loop

When >2h without progress, do these in order:

### Step 1 — Zoom Out
You're at the wrong abstraction level. Stop examining ONE request. Re-read the app as a whole. Re-draw the flow diagram. Where is the most broken / least defended part?

### Step 2 — Change Your User Perspective
- Unauthenticated? Make an account.
- Low-priv? Get an admin account.
- Switch browsers, IP, User-Agent — routing changes.

### Step 3 — Revisit Narrow Recon
- Re-read JS files. New endpoints appear.
- Look at API docs for internal calls.
- Undocumented endpoints: `/_debug`, `/internal`, `/_/`, `/admin`, `/actuator`, `/swagger`
- Response headers: `Server`, `X-Powered-By`, `X-AspNet-Version`
- HTML comments, `<meta>` tags — build numbers, dev URLs

### Step 4 — Single-Packet Deceleration
Stop scanners. Burp Repeater. ONE request. Read EVERY byte.
- Set-Cookie headers
- Stack traces, error messages
- HTML comments

### Step 5 — Hypothesis Pivot
Switch bug class entirely. Not different endpoint — different class.
- "Been deep on XSS? Check password reset for token reuse."
- "Hidden IDOR? Register two users, swap IDs."

## Stuck Tracker

```
Stuck on: (hypothesis & target)
Stuck minutes: (when hit wall)
Pivot to: (new plan)
```

If stuck >3 days in a row → switch target entirely.
