<img src="https://user-images.githubusercontent.com/37074372/207257285-ea661ff2-7d11-48f2-ae3e-5155f83dcd8a.png" align="left" width="75px"/>
<strong>Quick-tricks</strong> - Bitrix vulnerability scanner based on <a href="https://t.me/webpwn/317">Attacking Bitrix</a> guide

<br clear="left"/>

## Installation
### Requirements
* Go 1.19+
  ```
  go install github.com/indigo-sadland/quick-tricks@latest
  ```
## Usage
```
Bitrix vulnerability scanner

Usage:
  quick-tricks [command]

Available Commands:
  help        Help about any command
  lfi         Module 'lfi' checks if there are endpoints vulnerable to Local File Inclusion.
  quick       Run all quick modules ('recon', 'lfi', 'redirect', 'spoofing' and 'xss')
  rce         Module 'rce' tries to exploit vulnerable components of the target Bitrix.
  recon       Module 'recon' helps to find login page endpoints, local path disclosure and license key.
  redirect    Module 'redirect' checks endpoints vulnerable to Open Redirect.
  spoofing    Module 'spoofing' tests target for possibility of Content Spoofing attack.
  ssrf        Module 'ssrf' helps to check whether the target is vulnerable to SSRF or not.
  xss         Module 'xss' checks target's endpoints that potentially can be vulnerable XSS.

Flags:
  -h, --help   help for quick-tricks

Use "quick-tricks [command] --help" for more information about a command.
```
###
*Author is not responsible for any damages. You are responsible for your own actions. Attacking targets without prior mutual consent is illegal.*
