> [!abstract] Hub: [[../../Crow-Hustler|← Crow-Hustler]]

# Hunter Methodology V1
Crow-Hustler is a continuously improving bug bounty operating system. Every hunt should either produce a valid finding or improve the methodology for future hunts.


Modular knowledge base for bug bounty hunting. Each directory = one domain. Start here, drill into what you need.

## Agent
- [[01-operating-rules]]
- [[02-permissions]]
- [[03-knowledge-management]]
- [[04-tool-usage]]
- [[05-decision-rules]]
## Knowledge Management
- [[01-learning-during-hunt]]
- [[02-post-hunt-learning]]
- [[03-internet-research]]
- [[04-knowledge-base]]

## Teach-Back (Human Learning Bridge — NEW)
- [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|teach-back-index]]
- [[01-teach-back-template]]
- [[02-teach-back-example]]
- [[03-sessions]]

## Loading Order (first time)

1. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/mindset/00-index|mindset]] — calibrate your brain before touching any target
2. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/cadence/00-index|cadence]] — daily rhythm, anti-burnout, time blocks, human review
3. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|hunt-lifecycle]] — how a hunt runs from start to finish
4. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|recon]] — find the perimeter, then map one asset deeply
5. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|hunting]] — bug-class playbooks (auth, XSS, SSRF, IDOR, etc.)
6. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index|thinking]] — what to do when stuck, first-time, or finding something interesting
7. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|bypass]] — filter/403/WAF bypass techniques
8. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index|payloads]] — payload collections by class
9. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/tools/00-index|tools]] — tool reference by category
10. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/templates/00-index|templates]] — daily note, finding template, post-hunt log
11. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/submission/00-index|submission]] — reporting discipline
12. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/learning/00-index|learning]] — how to learn new bug classes
13. [[4-Methodology/Crow-Hustler/hunter-methodology-V1/teach-back/00-index|teach-back]] — agent writes what you need to learn (READ THIS AFTER EVERY SESSION)

## Directory Map

| Directory | What's inside |
|-----------|---------------|
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/mindset/00-index\|mindset]] | Mental models, iron rules, ethical boundaries |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index\|recon]] | Wide + narrow recon pipelines, origin-IP, 60-min recipe |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index\|hunting]] | Bug-class playbooks (auth, XSS, SSRF, IDOR, etc.) |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index\|thinking]] | Stuck protocol, first-time encounters, interesting findings, stop rules |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index\|hunt-lifecycle]] | Setup → hunt → post-hunt: file org, logs, todos, saves |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index\|bypass]] | 403, XSS filter, SSRF, rate-limit, OAuth bypass catalogs |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index\|payloads]] | XSS, SQLi, SSTI, and common payload collections |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/tools/00-index\|tools]] | Tool reference: recon, XSS, fuzzing, params, cloud |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/templates/00-index\|templates]] | Daily note, finding report, post-hunt log templates |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/cadence/00-index\|cadence]] | Daily template, time blocks, burnout protocol, human review |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/learning/00-index\|learning]] | Learning loop, resources, mastery progression |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/submission/00-index\|submission]] | Golden rules, pre-sub checklist, anti-patterns |
| [[4-Methodology/Crow-Hustler/hunter-methodology-V1/teach-back/00-index\|teach-back]] | NEW: Agent writes human-readable summary after every session |

## How to Use

Open the directory relevant to what you're doing right now. Each directory has its own index. When you finish a hunt, update [[03-post-hunt|post-hunt.md]] and save logs per [[03-post-hunt-log|post-hunt-log.md]].

## Modes

| Mode | When | Depth |
|------|------|-------|
| `quick` | New target, 60 min | Wide recon → pick 1 endpoint → test 2 classes |
| `deep` | Known target, 3-5h | Narrow recon → full auth + IDOR + one more class |
| `learn` | New bug class, 2h | Read theory → test-bed → apply on live |
| `maintain` | Burnout/low energy | 1h light recon, note cleanup, template updates |

## Rule

This directory is alive. Update it after every hunt. Add new payloads, new bypasses, new lessons. If you don't update it, it decays.

# WorkFlow
```
Start

↓

Load Agent Rules

↓

Load Mindset

↓

Load Hunt Lifecycle

↓

Recon

↓

Choose Attack Surface

↓

Hunt

↓

Finding?

├── Yes
│      ↓
│  Validate
│      ↓
│  Report

└── No
      ↓
 Stuck?

      ├── No → Continue
      │
      └── Yes
              ↓
Internet Research
              ↓
Knowledge Update
              ↓
Resume Hunt

↓

Post Hunt

↓

Knowledge Review

↓

TEACH-BACK (Agent writes human summary)

↓

HUMAN READS TEACH-BACK (5-10 min, now)

↓

Human confirms: [x] Understood, [x] Can apply

↓

Finish
```
