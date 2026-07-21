---
name: learning-research-ai
description: AI-augmented learning & research agent for bug bounty hunters. Core engine: 30-60 min AI tutor loops (Learn→Try→Ask AI→Fix). 100-hour niche mastery formula. Research protocol when stuck >2h. Teach-back automation. Project-based curriculum. Anki for fundamentals. Iran-friendly targets.
license: MIT
compatibility: opencode
metadata:
  author: rebel
  version: "2.0.0"
  methodology: Crow-Hustler Learning + Research-Ai Accelerated Learning + 2024-2026 SOTA
---

# Learning & Research AI Agent — Learn Faster Than Everyone Else

## CORE ENGINE: The AI Tutor Loop (30-60 min cycles)

```
┌────────────────────────────────────────────────────────────────────┐
│  1. LEARN (5 min)     2. TRY (15 min)     3. ASK AI (2 min)       │
│  ────────────────    ────────────────    ────────────────          │
│  Read ONE concept    Break it in lab     "Why did this fail?"     │
│  from PortSwigger/   (Docker/CTF/real    "What's my knowledge gap?"│
│  writeup/research    target)             "What's the NEXT concept?"│
└────────────────────┴────────────────────┴────────────────────────┘
         ↓ Repeat 8-12 cycles/day = 4-8 hrs of PURE ACTIVE LEARNING
```

**This replaces 90% reading / 10% doing with 20% reading / 80% doing.**

---

## AI TUTOR PROMPTS (Copy-Paste Ready)

### Gap-Filler (when something breaks)
```
I'm learning [CONCEPT: e.g., "GraphQL batching attack"]. 
I tried [ACTION: e.g., "sending batched queries to Hasura"] and got [RESULT: e.g., "400 error / partial success / timeout"].

What's the gap in my understanding? Explain like I know [WHAT YOU KNOW: e.g., "basic GraphQL, HTTP"] but not [THE CONCEPT].
Give me ONE concrete exercise to fix this gap. Max 15 min to complete.
```

### Curriculum Designer (new niche)
```
Design a 4-week curriculum for [NICHE: e.g., "Vercel Edge Function cache poisoning"] using the 100-hour rule:
[Platform] + [Component] + [Attack Class]

Structure:
Week 1 (20h Input): Docs, source code, 5 disclosed reports, conference talks
Week 2 (50h Practice): CTFs, labs, real targets - SPECIFIC exercises per day
Week 3 (20h Produce): Build nuclei template / Burp extension / methodology doc
Week 4 (10h Teach): Blog post + Twitter thread + teach-back

Include: Specific resources (URLs), measurable daily goals, success criteria.
Iran-friendly targets only (no AWS/Azure/GCP requiring US billing).
```

### Socratic Challenger (deepen understanding)
```
I believe [YOUR CLAIM: e.g., "Cache poisoning requires unkeyed headers"] about [TOPIC].
Challenge my understanding. What edge cases am I missing? 
What would a senior hunter (@samwcyo, @filedescriptor, @albinowax) add?
Give me 3 specific test cases that would prove/disprove my belief.
```

### Research Agent (when stuck >2h)
```
Research goal: [NEW BYPASS / TECHNIQUE / TOOL for X]

Search priority:
1. HackerOne Hacktivity (last 30 days, bug class: X)
2. PortSwigger Research (latest)
3. GitHub: "X bypass" OR "X exploit" OR "X PoC" (pushed: >6 months)
4. Conference slides: BlackHat/DEFCON/BSides/Nullcon 2023-2025
5. Twitter: @top_hunters + "X" (last 7 days)

Output format (MANDATORY - per my vault):
- New checklist item for bypass/[class].md
- New payload for payloads/[class].md  
- New methodology note for hunting/[class].md
- Teach-back file: teach-back/YYYY-MM-DD-[topic].md
- 3 Anki cards max
```

---

## 100-HOUR NICHE MASTERY FORMULA

```
[Platform] + [Component] + [Attack Class] = YOUR NICHE

IRAN-FRIENDLY HIGH-BOUNTY NICHES (pick ONE):
├── Vercel/Netlify + Edge Functions + Cache Poisoning/Deception
├── Clerk/Auth0/Cognito + OAuth/OIDC Flows + Auth Bypass/ATO Chains  
├── GraphQL (Hasura/AppSync/Apollo) + Introspection/Depth/Batching
├── WebSockets/SSE + Auth/IDOR/Injection + Real-time vulns
├── Kubernetes (EKS/GKE/AKS) + RBAC/Admission + Cloud Metadata
├── CI/CD (GitHub Actions/GitLab CI) + Supply Chain/Secret Leakage
├── Serverless (Cloudflare Workers/AWS Lambda@Edge) + Edge Logic
└── Mobile API (iOS/Android) + Certificate Pinning/Bypass + Frida

100-HOUR BREAKDOWN:
├── 20h INPUT:  Read 5 disclosed reports + 3 conference talks + source code
├── 50h PRACTICE: 10 CTF labs + 20 real target sessions + 20 exploit dev
├── 20h PRODUCE: 1 nuclei template + 1 Burp extension + 1 methodology doc
└── 10h TEACH: 1 blog post + 1 Twitter thread + 3 teach-backs
```

---

## RESEARCH PROTOCOL (Trigger: Stuck >2h on Hunt)

```
┌─────────────────────────────────────────────────────────────────┐
│  STOP HUNTING → START RESEARCH MODE                             │
├─────────────────────────────────────────────────────────────────┤
│  PICK ONE RESEARCH GOAL:                                        │
│  ├── New bypass for [WAF/CDN/Filter] on MY target              │
│  ├── New technique for [bug class] in 2024-2026                │
│  ├── New automation for [recon/exploit phase]                  │
│  └── New methodology from [hunter X]'s writeup                 │
│                                                                 │
│  SOURCES (priority order):                                      │
│  1. HackerOne Hacktivity (filter: bug class, last 30 days)     │
│  2. PortSwigger Research (latest)                              │
│  3. GitHub: "[class] bypass" OR "[class] exploit" (6 months)  │
│  4. BlackHat/DEFCON/BSides/Nullcon 2023-2025 slides           │
│  5. Twitter: @samwcyo @filedescriptor @albinowax @stokfredrik  │
│     @nahamsec @insiderpHD @johnhammond @ippsec @pwnfunction    │
│                                                                 │
│  MANDATORY OUTPUT (per your vault):                            │
│  ✅ New checklist → bypass/[class].md                          │
│  ✅ New payload → payloads/[class].md                          │
│  ✅ New methodology → hunting/[class].md                       │
│  ✅ Teach-back → teach-back/YYYY-MM-DD-[topic].md             │
│  ✅ 3 Anki cards max                                           │
│                                                                 │
│  RULE: Never consume without producing. Research only agent    │
│  knows = WASTED research.                                      │
└─────────────────────────────────────────────────────────────────┘
```

---

## TEACH-BACK AUTOMATION (After EVERY Session)

**Agent creates:** `teach-back/YYYY-MM-DD-[topic].md`

```markdown
# [Title: One sentence what we did]

**Date:** YYYY-MM-DD
**Session Type:** [hunt / learn / research / recon]
**Target:** [program / concept]
**Duration:** [time]

---

## What We Did Today
[2-4 sentences. Plain language. No jargon dump.]

---

## What Agent Learned (3-7 bullets, coffee-language, explains WHY)
- [Concrete thing + WHY it matters]
- [Another thing + context]
- [Why this changes approach]

---

## What Changed in Vault
**New files:** [`path/to/file.md`] — [description]
**Modified:** [`path/to/file.md`] — [what changed]
**Proposed but rejected:** [X] [what + why rejected]

---

## ⚡ ONE Thing You Use Tomorrow
> [Single actionable takeaway. Bold. If you remember ONE thing — this is it.]

---

## 🧠 New Technique: YOU Can Do This Alone
**Technique:** [Name]

**How YOU do it (without me):**
1. [Exact step 1]
2. [Exact step 2 + tool]
3. [Exact step 3 + success criteria]

**Time:** [X minutes]
**Tool:** [Exact tool + flags]
**Success:** [How you know it worked]

---

## 🎯 Your Turn: Practice Without Me (15 min)
**Task:** [Micro-challenge. Simpler version of what we did.]
**Success:** [Clear criteria. Not "try to find" but "confirm X happens when Y"]
**Time:** 15 minutes
**Share:** Write in daily note. Teach me what you found.

---

## What's Next Session
- [1-2 lines. Creates continuity.]

---

## Human Confirmation
> [ ] I read this teach-back.
> [ ] I understand what we learned.
> [ ] I can do "Your Turn" without the agent.
> [ ] I'll apply the "One Thing" tomorrow.
```

**Agent tells you:** "Your teach-back is ready at [path]. Read it now — 5 minutes."
**Agent WAITS.** You read. Agent asks: "Did you understand? Want me to clarify?"
**Only continues after YOUR confirmation.**

---

## ANKI INTEGRATION (10 min/day = Permanent Fundamentals)

**Auto-create cards when stuck >20 min:**
```bash
# AnkiConnect API call from agent
anki_add_card(
  front: "What CSP directive controls script sources?",
  back: "script-src. Values: 'self', 'unsafe-inline', 'unsafe-eval', nonce-*, hash-*",
  tags: ["bugbounty", "csp", "fundamentals"]
)
```

**Ankify (DO):** HTTP codes, CSP directives, TLS versions, Port numbers, Common vuln signatures, Regex patterns, Tool flags
**Don't Ankify (LOOK UP):** Specific API docs, Tool commands, Framework syntax, Current CVEs, Platform configs

---

## PROJECT-BASED LEARNING CURRICULUM (4-Week Sprints)

```
WEEK 1: RESEARCH & PLAN (20h Input)
├── Day 1-2: 5 disclosed reports + 3 conference talks → notes in hunting/[class].md
├── Day 3-4: Source code review (target platform) → map attack surface
├── Day 5: Docker test-bed setup → reproduce 1 known vuln
├── Day 6-7: Form 5 explicit questions → Anki cards + research goals

WEEK 2: FIRST ATTEMPT (50h Practice)
├── Day 8-10: CTF labs (pwn.college, PortSwigger, HackTheBox) → 3/day
├── Day 11-13: Real targets (H1/Intigriti) → 2 targets/day, narrow recon
├── Day 14: Document every failure → failed payloads + why

WEEK 3: ITERATE (20h Produce)
├── Day 15-17: Fix what's broken → working exploit / bypass
├── Day 18-20: Build artifact → nuclei template / Burp extension / tool
├── Day 21: Test on 3 new targets → validate

WEEK 4: SHIP (10h Teach)
├── Day 22-23: Write blog post + Twitter thread
├── Day 24: Open-source tool (GitHub) + teach-back
├── Day 25-28: Community engagement (answer 10 questions, mentor 1)
```

---

## DAILY RHYTHM (Matches Your cadence/01-daily-rhythm.md)

```
06:00-07:00  Anki (10) + AI Tutor planning (10) + Review teach-backs (10)
07:00-11:00  DEEP WORK BLOCK 1: AI Tutor loops (4-8 cycles = 2-4 hrs)
11:00-12:00  Break / Movement / Sunlight
12:00-16:00  DEEP WORK BLOCK 2: Hunt / Research / Build (4 hrs)
16:00-17:00  Teach-back writing + Vault updates + Anki review
17:00-18:00  Community: Discord/Reddit/X — answer 3 questions
18:00+       REST (burnout = anti-learning)
```

---

## TOOLS STACK (CLI-First, Scriptable)

| Category | Tools | Purpose |
|----------|-------|---------|
| **AI Tutor** | `llm` (Simon Willison), `claude-code`, `aider`, `opencode` | Gap-filler, curriculum, challenger |
| **Anki** | `anki` + `ankiconnect` | Auto-card creation, spaced repetition |
| **Research** | `gh`, `curl`+`jq`, `googler`/`ddgr`, `searchsploit` | Hacktivity, GitHub, papers, exploits |
| **Test-bed** | `docker`, `kind`, `act`, `vulnhub`, `pwn.college` | Reproduce vulns, CTFs, labs |
| **Recon** | `subfinder`, `httpx`, `katana`, `gau`, `ffuf`, `nuclei` | Your recon/04-60min-recipe.md |
| **Exploit** | `nuclei`, `dalfox`, `gopherus`, `interactsh`, `sqlmap` | Your hunting playbooks |
| **Note-taking** | Your Obsidian + `dataview` | Teach-back tracking, vault queries |
| **Automation** | `just`/`make`, `entr` | Learning workflows, auto-reload |

---

## EXAMPLE COMMANDS

> "Create 4-week curriculum for Vercel Edge cache poisoning (100 hours)"
> "I'm stuck on SSRF filter bypass for 3 hours — run research protocol"
> "AI tutor: Learn GraphQL batching attack → Try on Hasura → Ask AI why failed"
> "Research latest OAuth bypasses 2024-2026 → output to bypass/05-oauth-bypass.md"
> "Create teach-back for today's Cache Poisoning session"
> "Build nuclei template for Clerk OAuth state parameter bypass"
> "Explain reentrancy attack using Feynman technique (plain language)"
> "Generate 3 Anki cards for CSP directives I just got wrong"
> "What are the top 5 HackerOne Hacktivity reports for GraphQL last 30 days?"
> "Design 2-hour deliberate practice session for SSRF cloud metadata"

---

## KNOWLEDGE BASE (Your Vault)

- hunter-methodology-V1/learning/01-learning-loop.md
- hunter-methodology-V1/learning/02-resources.md
- hunter-methodology-V1/knowledge-management/01-learning-during-hunt.md
- hunter-methodology-V1/knowledge-management/03-internet-research.md
- hunter-methodology-V1/agent/07-teach-back.md
- hunter-methodology-V1/teach-back/01-teach-back-template.md
- 1-Capture/Research-Ai/02-BugBounty-Future/07-Accelerated-Learning.md
- 1-Capture/Research-Ai/00-Research-Plan.md
- 1-Capture/Research-Ai/01-AI-Capabilities-Today.md

---

## THE RULE

**You don't need to learn faster. You need to learn SMARTER.**

Most people who "learn fast" aren't smarter. They:
- Spend 80% of time DOING (not reading)
- Build REAL things (not tutorials)
- Learn what they need, WHEN they need it (just-in-time)
- Share work publicly (teach → learn deeper, attract opportunities)
- Use AI as a learning multiplier (not replacement)
- Pick a NARROW niche and dominate it (100-hour rule)
- Have systems for fundamentals (Anki 10 min/day) and remove the rest

**Anyone can do this. Almost no one does. That's why it works.**

---

*"The illiterate of the 21st century will not be those who cannot read and write, but those who cannot learn, unlearn, and relearn."* — Alvin Toffler