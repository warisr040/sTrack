# Building sTrack.exe

The exe is a ~7 MB native Windows window (Microsoft Edge WebView2) around the same `index.html`
used for the website. It serves the page only on 127.0.0.1 — nothing goes to the internet.
It also gives the page access to the team database file by path (read, write, and a `.lock` file while merging); only `.json` files given by full path are allowed. The window loads nothing but its own local page and opens no pop-ups (see `third_party/go-webview2`, a copy of github.com/jchv/go-webview2 with bindings run off the UI thread, a request filter and a pop-up block).
The local copy of the plan lives in `%LOCALAPPDATA%\sTrack`. Requires WebView2 Runtime (built into Windows 10/11).

To rebuild after editing the site:
0. Run `python3 tools/csp.py` from the repo root. It stamps the page's Content-Security-Policy with the hash of its script; without it the page refuses to run its own (edited) script.
1. Copy `web/index.html` here as `web_index.html`.
2. Install Go (https://go.dev). Resources (icon, manifest, version info 2.0) are in
   `rsrc_windows_amd64.syso`. To regenerate it after changing `versioninfo.json`, `icon.ico` or `app.manifest`:
       go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
       goversioninfo -64 -o rsrc_windows_amd64.syso versioninfo.json
3. Build:
       go build -trimpath -ldflags "-H windowsgui" -o sTrack.exe .
   (Keep the debug information: do not add -s -w. Stripped Go programs are more often flagged by antivirus heuristics.)
   (From Linux/macOS add: GOOS=windows GOARCH=amd64 CGO_ENABLED=0)
