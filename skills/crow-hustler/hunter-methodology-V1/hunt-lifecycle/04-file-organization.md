> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|← Back to Hunt Lifecycle]]

# File Organization

How to structure every hunt.

## Directory Layout

```
hunter-methodology/          ← this directory (permanent methodology)
hunt-logs/                   ← per-session logs (created by you)
├── 2026-06-29_walmart.md
├── 2026-06-30_centralmayorista.md
└── ...

target-state/                ← per-target notes
├── walmart.md
├── centralmayorista.md
├── vizio.md
└── ...
```

## Per-Session Log Format (`hunt-logs/YYYY-MM-DD_target.md`)

```markdown
# YYYY-MM-DD — Target Name

Mode: quick/deep/learn/maintain
Hypothesis: ...

## Requests
- [ ] req 1: what I tried → result
- [ ] req 2: what I tried → result

## Findings
- finding1: URL, payload, impact
- finding2: URL, payload, impact

## Payloads That Worked
- payload1: context
- payload2: context

## Dead Ends
- dead1: why it didn't work

## Stuck Log
- stuck on: ... at HH:MM
- pivot to: ...

## Todos for Next Session

## What I Learned (with source)
```

## Per-Target State (`target-state/target.md`)

```markdown
# Target: name

Status: wide-recon / narrow-recon / deep-hunt / saturated
Subdomains mapped: N
Auth flows understood: yes/no
API surface mapped: yes/no
Last tested: YYYY-MM-DD
Findings filed: N
```

## Update Cadence

- **After every hunt**: update hunt log + what-learned
- **After every finding**: update payload/bypass/hunting files
- **Weekly**: update target state, review submission status
- **Monthly**: prune stale targets, archive old logs
