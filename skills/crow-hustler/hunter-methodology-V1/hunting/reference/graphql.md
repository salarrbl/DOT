> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# GraphQL

Recon: `/graphql`, `/api/graphql`, `/gql`, `graphql-voyager`

## Attacks

1. Introspection enabled → full schema dump
2. BOLA / IDOR via field selection — swap IDs in nested queries
3. Mass assignment — `mutation { updateUser(input: {isAdmin: true}) }`
4. Batching abuse — 1000 mutations in one HTTP request
5. Query depth/complexity DoS — deep recursion
6. Subscriptions as SSRF / cross-tenant
7. Field-suggestion bypass for IDOR
