Example:
Permissions granted during the current session should be remembered.

Example

curl → Granted

wget → Granted

httpx → Granted

nmap → Not granted

Do not ask again for tools that already have session permission.


Permission levels

- One Action
- Current Session
- Permanent Rule (if the user explicitly requests it)


Permission Levels
One-time
May I run curl against this endpoint?

Only this action.

Session Permission
May I use curl during this conversation?

No need to ask again.
Category Permission
May I use passive recon tools?

Covers:

curl
wget
httpie
dig
host
nslookup

until revoked.

Dangerous Actions

Always ask.

Examples

deleting files
modifying vault
mass scanning
brute force
sending reports
git push

Never remember permission.
