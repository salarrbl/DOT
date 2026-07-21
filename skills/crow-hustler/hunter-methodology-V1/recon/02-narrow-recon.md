> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/recon/00-index|← Back to Recon]]

# Narrow Recon

Pick the most promising live subdomain. Map it until you know its auth flow, state-changes, and API surface.

## Map Features, Not URLs

1. **Be a normal user for 3–7 days.** Read docs. Build a client. Click every feature.
2. **Extract traffic.** Burp + DevTools → preserve headers, cookies, tokens, full request bodies.
3. **Map flows.** Mindmap in Obsidian. Note every state-changing endpoint.

## Find Endpoints

**Passive**:
- Wayback Machine — endpoints, parameters, files, JS
- `gau`, `waybackurls`
  ```bash
  waybackurls https://domain.com | unfurl keys | sort -u
  ```

**Active**:
- `katana` (best crawler)
- Manual view-source on every JS file
- `gospider`, `hakrawler`

## Parameter Discovery (most important step)

```bash
python3 paramspider.py -d target.com --exclude js,png,svg
x8 -w params -u https://target.com/path?a=b -X GET POST
```

- Burp `GAP` extension
- `fallparams`, `arjun`
- Source params from wayback, JS, forms, URL path templates

**Critical**: Test the same parameter on many pages. A param ignored on page 1 may reflect on page 5.
