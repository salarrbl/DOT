> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunt-lifecycle/00-index|← Back to Hunt Lifecycle]]

# Common Sensitive Paths

Always check these during narrow recon:

```
/.git/config
/.git/HEAD
/.svn/entries
/.env
/.DS_Store
/.htaccess
/wp-config.php.bak
/backup.zip
/admin
/api/swagger.json
/api/docs
/openapi.json
/graphql
/actuator
/actuator/env
/actuator/heapdump
```

## Source Code Disclosure

Classic `.git/config` exposure on production lets attackers clone the repo with `git-dumper`, recovering source code, credentials, and internal paths. Still prevalent in 2026.
