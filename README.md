# BPJS Foreground Service

## How to use

1. Download these 2 file from this repository and put it in the same directory
  > - [BPJSForeground.exe](BPJSForeground.exe)
  > - [BPJSForegroundService.exe](BPJSForegroundService.exe)

2. Run the executable [**BPJSForegroundService.exe**](BPJSForegroundService.exe)
3. send GET request to (change &lt;localhost&gt; to &lt;client IP&gt; for trigger from third party):
  > - http://localhost:8089/ to see if the service is active
  > - http://localhost:8089/activate-bpjs/ to trigger the action (set fingerprint app to foreground)

## How it works

1. [**BPJSForegroundService.exe**](BPJSForegroundService.exe)
  > - create a window indicating that the app is running
  > - create a service using [gin-gonic/gin](https://pkg.go.dev/github.com/gin-gonic/gin)
  >   - on endpoint / that response with message
  >   - on endpoint /activate-bpjs that trigger the action and send reponse of either success or error

2. /activate-bpjs trigger the action that execute [**BPJSForeground.exe**](BPJSForeground.exe) on client or targeted computer
3. [**BPJSForeground.exe**](BPJSForeground.exe) set targeted app (i.e BPJS Fingerprint App) to foreground. [**BPJSForeground.exe**](BPJSForeground.exe) is a compiled version of AutoHotkey script [**BPJSForeground.ahk**](BPJSForeground.ahk)
4. [**BPJSForeground.ahk**](BPJSForeground.ahk) is a script that do all of the action on client computer using [AutoHotKey](https://www.autohotkey.com) utility
