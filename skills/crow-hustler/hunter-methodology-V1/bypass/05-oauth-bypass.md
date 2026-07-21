> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]] · Related: [[01-auth|Auth playbook]]

# OAuth redirect_uri Bypass

## Bypasses (when server partially-validates)

```
target.com%2523evil.com              # double-encoded #
target.com%09evil.com                 # tab
target.com&evil.com                   # param injection
////evil◎com                          # Chinese fullwidth dot U+FF0E
evil%E3%80%82com.target.com           # Chinese dot
target.com/fake/../../foo             # path traversal
target.com\evil.com                   # backslash
target.com%2e%2e%2f%2e%2e%2fevil      # encoded ../
```

## Chains

- redirect_uri bypass + open-redirect on same site = token theft
- redirect_uri full validation + same-domain open redirect = token theft
- Email-skip bypass: create provider account with phone → add victim's email
