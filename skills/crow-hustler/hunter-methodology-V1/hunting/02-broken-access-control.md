> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Broken Access Control / IDOR / RBAC

## IDOR Parameters (always check)

```
id, user, account, number, order, no, doc, key, email, group, profile, edit
REST: /api/users/123, /orders/456
```

Test by registering two roles (low + high). Every endpoint. Swap IDs.

## RBAC Testing Matrix

For each feature × every role:

| Feature | Admin | Manager | Read-only | Unauth |
|---------|-------|---------|-----------|--------|
| Create user | ✓ | ✗ | ✗ | ✗ |
| Delete user | ✓ | **VULN** | ✗ | ✗ |
| Read report | ✓ | ✓ | ✓ | **VULN** |

## Sensitive Data Endpoints

Fuzz: `/api/users/all`, `/api/users/list`, `/api/admin/users`, `/api/users/active`

## Account Deletion

- Skip confirm flow, change IDs in URL
- After deletion — email released for re-registration? (permanent ownership)
