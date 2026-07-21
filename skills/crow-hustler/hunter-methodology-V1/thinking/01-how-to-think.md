> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/thinking/00-index|← Back to Thinking]]

# How to Think

## Core Questions

Before any request, ask:
- "How is this *not* supposed to work?"
- "What happens if I skip this step?"
- "What if I'm a different user?"
- "What if I do this in the wrong order?"

## Chain Thinking

Never report a chain-worthy bug in isolation. Build the full story:
- IDOR + email-flow → ATO
- OAuth state missing + CSRF → ATO
- SSRF + cloud metadata → keys → RCE

## Hypothesis-Driven

1. Form a hypothesis: "This endpoint is vulnerable to IDOR because..."
2. Test with 3 requests max
3. If false → switch hypothesis, don't dig deeper on wrong track
4. If true → build the chain

## The 4-Question Framework (for Business Logic)

For every endpoint:
1. **What is this supposed to do?** — understand normal flow
2. **What data can I control?** — params, headers, filenames, body fields
3. **Where does this data go?** — reflected, stored, queried, forwarded
4. **What if I break the expected flow?** — swap IDs, skip steps, reorder, send arrays

## AI-Assisted Hypothesis Generation

Feed an HTTP request/response pair to an LLM and ask for vulnerability hypotheses. Use it to:
- Analyze JavaScript files for DOM sinks
- Generate business logic flaw ideas from feature descriptions
- Draft report narratives from raw notes
- Generate CVSS justifications

Rule: AI is a copilot, never the pilot. It produces leads; you verify.

## Read the Target

The target tells you what to do. Don't look for XSS because "it's fun." Read the application. What does it do? Where is the money? What are the state changes? Hunt there.
