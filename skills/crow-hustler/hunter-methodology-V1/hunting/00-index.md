> [!abstract] Hub: [[README-Crow-Hustler-V1|← Back to README]]

# Hunting

Feature-centric bug-class playbooks. Each file is a standalone testing guide.

## Priority Order (hunt in this sequence)

1. [[01-auth]] — Auth / Password Reset (highest yield)
2. [[02-broken-access-control]] — IDOR / RBAC
3. [[03-ssrf]] — SSRF + cloud metadata
4. [[04-xss]] — XSS (as chain material, not standalone)
5. [[05-other-classes]] — SQLi, SSTI, RCE, CORS, etc.
6. [[06-business-logic]] — Business logic flaws (highest payout)
7. [[07-api-security]] — REST, GraphQL, gRPC, SOAP, Webhooks
8. [[08-cloud-misconfig]] — S3, Azure, GCP, CloudFront, IAM

## Gap Closers (bug classes to learn)

| Class | File |
|-------|------|
| JWT | [[jwt-attacks]] |
| GraphQL | [[graphql]] |
| NoSQLi | [[nosqli]] |
| Race Conditions | [[race-condition]] |
| Cache Poisoning | [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/reference/cache-poisoning]] |
| Prototype Pollution | [[prototype-pollution]] |
| File Upload | [[file-upload]] |
| LLM/AI Security | [[llm-security]] |
| SAML 2.0 | [[saml-security]] |
| WebAuthn/Passkey | [[webauthn-passkey]] |
| WebAssembly (WASM) | [[wasm-security]] |

## Quick Rule

Reflected XSS, CSRF, open redirect alone are NOT your targets. They are chain material.
