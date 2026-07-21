> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# API Security

## REST

- Missing auth on some methods
- Mass assignment
- Rate limit absence
- Inconsistent IDOR across endpoints

## GraphQL

See `reference/graphql.md`

## gRPC

- Reflection enabled
- No field-level authorization

## SOAP

- WSDL exposure
- XXE
- WS-Security bypass

## Webhooks

- SSRF sink
- Signature bypass
- Replay attacks

Reference: OWASP API Security Top 10 (2023)
