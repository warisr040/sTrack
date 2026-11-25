#!/usr/bin/env python3
"""Stamp web/index.html with a Content-Security-Policy that allows only sTrack's own script.
Run after every change to the script (the hash must match, or the page will not run)."""
import base64, hashlib, re, sys
p = sys.argv[1] if len(sys.argv) > 1 else 'web/index.html'
s = open(p, encoding='utf-8').read()
m = re.search(r'<script>(.*?)</script>', s, re.S)
h = base64.b64encode(hashlib.sha256(m.group(1).encode('utf-8')).digest()).decode()
pol = ("default-src 'none'; script-src 'sha256-%s'; style-src 'unsafe-inline'; img-src data: blob:; font-src data:; "
       "media-src data: blob:; connect-src 'none'; form-action 'none'; base-uri 'none'; object-src 'none'; "
       "frame-src 'none'; worker-src 'none'; manifest-src 'none'") % h
meta = '<meta http-equiv="Content-Security-Policy" content="%s">' % pol
if 'http-equiv="Content-Security-Policy"' in s:
    s = re.sub(r'<meta http-equiv="Content-Security-Policy" content="[^"]*">', meta, s, count=1)
else:
    s = s.replace('<head>', '<head>\n' + meta, 1)
open(p, 'w', encoding='utf-8').write(s)
print('script sha256-' + h)
