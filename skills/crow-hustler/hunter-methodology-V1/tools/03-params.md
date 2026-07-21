> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/tools/00-index|← Back to Tools]]

# Parameter Discovery Tools

| Tool | Use |
|------|-----|
| `paramspider -d target.com --exclude js,png,svg` | Wayback-based param extraction |
| `x8 -w params -u URL -X GET POST` | Hidden param brute-force |
| `arjun -u URL` | Param brute-force (backup) |
| `ffuf -w wordlist.txt -u URL/FUZZ -ac -fs 0` | Content discovery fuzzing |
| Burp `GAP` | Auto-parameter discovery |
