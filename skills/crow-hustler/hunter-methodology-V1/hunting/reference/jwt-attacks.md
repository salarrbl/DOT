> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# JWT Attacks

Recon: `/api/auth/verify`, `/api/token`, `.well-known/jwks.json`

## Attacks (run all)

1. `alg=none` — strip signature, send unsigned token
2. `kid` injection — `kid: "../../../dev/null"`, `kid: "/proc/self/environ"`
3. `jku` / `jwk` header — point to your own URL for key
4. Weak secret — `secret`, `secret123`, `qwerty` via hashcat
5. Algorithm switch — HS256 → RS256, use pub key as HMAC secret
6. Expired token accepted — `exp` not checked
7. Issuer/audience not checked — token from app A works on app B

Tools: `jwt_tool`, `portswigger/jwt-attacks`
