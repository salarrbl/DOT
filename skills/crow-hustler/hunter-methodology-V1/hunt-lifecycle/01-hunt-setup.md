> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|← Back to Hunt Lifecycle]]

# Hunt Setup

Before you start a session.

## Choose Mode

| Mode | When | Duration | Focus |
|------|------|----------|-------|
| `quick` | New target | 60 min | Wide recon → 1 endpoint → 2 bug classes |
| `deep` | Known target | 3-5h | Narrow recon → full auth + IDOR + 1 more |
| `learn` | New class | 2h | Theory → test-bed → apply |
| `maintain` | Low energy | 1h | Light recon, note cleanup, template update |

## Pre-Hunt Checklist

- [ ] What mode? (quick/deep/learn/maintain)
- [ ] Which target?
- [ ] What's the current scope / attack surface?
- [ ] What bug class am I testing? (single focus)
- [ ] What's my hypothesis?
- [ ] What tools do I need loaded?

## Scope Analysis (Before Anything)

1. **Read the program policy** — in-scope assets, excluded vuln classes, rules of engagement
2. **Check bounty table** — some programs pay more for specific classes (auth bypass, RCE)
3. **Check for new assets** — freshly-added assets pay better (less competition)
4. **Disclosed reports** — search for patterns, avoid duplicates
5. **Identify shared infrastructure** — common dependency = one exploit for N targets

## File Setup

Create a hunt log file at `hunt-logs/YYYY-MM-DD_target.md` with:
- Target, mode, date
- Hypothesis
- Tools loaded
- Any previous state
