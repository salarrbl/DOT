> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Race Conditions

Goal: bypass single-use / rate-limit invariants by parallelizing requests.
Tool: Burp Turbo Intruder (single-packet for HTTP/2, last-byte for HTTP/1).

## Hunt Patterns

1. Coupon reuse — 50 parallel `/redeem` requests
2. Gift-card balance — withdraw max, parallel additional withdraw → negative
3. Rate-limit bypass — N logins in single-packet
4. Email verification skip — 50 parallel verification requests
5. Email-alias collision — run OTP tests in parallel
6. Invitation seat exhaustion — 50 invites to same user
