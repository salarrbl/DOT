> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# Cloud Misconfigurations

| Service | Common bugs |
|---------|-------------|
| **AWS S3** | Public buckets, world-writable, object listing, predictable names |
| **Azure Blob** | Public containers, storage key leaks |
| **GCP Storage** | `allUsers` permission, IAM misconfig |
| **CloudFront / Akamai / Cloudflare** | Origin IP leaks, cache poisoning, origin access control bypass |
| **IAM** | Overly permissive roles, STS token leaks via SSRF |
| **Cognito / Auth0** | User enum, unauthenticated SignUp, unverified self-sign-up |

## S3 Testing

```bash
aws s3 ls s3://bucket-name --no-sign-request
aws s3api get-bucket-acl --bucket bucket-name --no-sign-request
```

Tools: `lazys3`, `s3enum`, `cloud_enum`, `pacu`
