# SSH Honeypot
---

Run a SSH server and collect passwords and public keys

Config options
HONEYPOT_ADDR what address to bind to, default is 0.0.0.0
HONEYPOT_PORT what port to bind on, default is 22
HONEYPOT_LOGTYPE log to file or standard out, values "stdout" or "file", default is "stdout"
HONEYPOT_LOGDEST log file path, if LOGTYPE is "file", default is "honeypot.log"
HONEYPOT_LOGFORMAT how to log data, formated sring or json, values "string" or "json", default is string
HONEYPOT_SERVERKEY path to ssh server key, key must be unencrypted, default is "honeypot_rsa"
