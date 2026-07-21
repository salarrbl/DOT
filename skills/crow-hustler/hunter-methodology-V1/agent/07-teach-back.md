# Teach-Back Rule

**This rule is non-negotiable.** After every session that involves learning, hunting, research, or methodology improvement, you MUST produce a teach-back.

## When to Produce a Teach-Back

Produce a teach-back file at the END of every session:

| Session Type | Teach-Back Required? |
|-------------|---------------------|
| Hunt (any duration) | YES |
| Learning session (new bug class, technique) | YES |
| Research (internet research for a problem) | YES |
| Recon only (no hunting done) | YES (but shorter — just what was discovered) |
| Methodology update (changing vault rules) | YES |
| Tool setup / configuration | YES (if anything was learned) |
| Casual question / quick check | NO (unless user learned something) |
| Pure execution (user knows exactly what to do) | NO |

## How to Produce a Teach-Back

1. **Use the template:** `teach-back/01-teach-back-template.md`
2. **Create the file:** `teach-back/YYYY-MM-DD-short-description.md`
3. **Fill out ALL sections** (skip "Your Turn" only if session was trivial)
4. **Tell the user:** "I've written your teach-back. It's at `teach-back/YYYY-MM-DD-x.md`. Read it now — 5 minutes."

## Rules for Teach-Back Content

### What I Learned (Section 2)
- **Write TO the human, not TO the vault.** "You" not "the hunter."
- **Explain WHY, not just WHAT.** "Instagram strips special chars from `state` BECAUSE they hash it server-side" — not just "state is stripped."
- **No more than 7 bullets.** If you learned 15 things, pick the top 7.
- **Every bullet must be TRUE.** If you're unsure, say "We suspect X but haven't confirmed."
- **No copy-paste from vault files.** Rewrite in human language.

### ONE Thing (Section 4)
- **Single sentence.** Bold. Unmissable.
- **Actionable tomorrow.** Not theory. Not "interesting concept." Something they use on their next target.
- **If the session was all negative results,** the ONE thing is: "What NOT to try" or "How to recognize a dead end faster."

### New Technique (Section 5)
- **The human must be able to do this WITHOUT you.** Include exact steps. Exact tools. Exact time.
- **If you automated it,** explain the manual version. The human learns the fundamentals, not the automation.
- **Always include:** time estimate, tool needed, success criteria.

### Your Turn (Section 6)
- **A micro-challenge.** 10-20 minutes. Simple. Doable immediately.
- **NOT a re-run of what we did.** A SMALLER, SIMPLER version that proves understanding.
- **Success criteria are clear.** Not "try to find" but "confirm that X happens when you do Y."

## What NOT to Include

- **Don't dump raw tool output.** Teach-backs are curated, not verbose.
- **Don't list every command we ran.** Summarize. They can find commands in the hunt log.
- **Don't write for the vault.** The vault gets its own updates. Teach-back is for the HUMAN.
- **Don't make it longer than 10 minutes to read.** If they need to scroll 5 pages, it's too long.
- **Don't make the "Your Turn" too hard.** It should be achievable. Success builds confidence.

## The Iron Rule of Teach-Backs

> **The human MUST learn something they can apply WITHOUT you. If they can only do it WITH you, they haven't learned it.**

Every teach-back is a test: "If I (the agent) disappeared tomorrow, could you (the human) still use what we learned?"

If the answer is no, the teach-back failed. Rewrite it.

## After Writing the Teach-Back

1. Tell the user: "Your teach-back is ready."
2. Give them the file path.
3. WAIT. Do not continue. Let them read.
4. Ask: "Did you understand everything? Do you want me to clarify anything?"
5. Only after they confirm understanding, mark the session complete.

## Scheduled Review

The user should review all teach-backs:
- **Daily**: Skim (5 min)
- **Weekly**: Re-read top 3, practice them (20 min)
- **Monthly**: Full review, update skill tracker (30 min)

Remind them of this schedule. Don't let teach-backs pile up unread.

---

*The vault is for machines. The teach-back is for humans. You serve the human, not the vault.*
