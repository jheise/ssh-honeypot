# SSH Honeypot
---

Run a SSH server and collect passwords and fingerprints of public keys. The server itself is completely stripped down to only log connection parameters and always reject any connection attempt

##Config options
```
HONEYPOT_ADDR what address to bind to, default is 0.0.0.0
HONEYPOT_PORT what port to bind on, default is 22
HONEYPOT_LOGTYPE log to file or standard out, values "stdout" or "file", default is "stdout"
HONEYPOT_LOGDEST log file path, if LOGTYPE is "file", default is "honeypot.log"
HONEYPOT_LOGFORMAT how to log data, formated sring or json, values "string" or "json", default is string
HONEYPOT_SERVERKEY path to ssh server key, key must be unencrypted, default is "honeypot_rsa"
```

##Sample Logs
Formatted string
```
2020/03/23 22:34:15 ssh-honeypot: user: root pubkey: SHA256:is4gTi9Lzdi4zRaW8MdKnLSGAETWb8cfwc8oht2usMo ipaddr: 127.0.0.1 type: ssh-rsa
2020/03/23 22:34:15 ssh-honeypot: user: root password: foobar ipaddr: 127.0.0.1

```
JSON
```
2020/03/23 22:38:52 ssh-honeypot: {"user":"root","pubkey":"SHA256:is4gTi9Lzdi4zRaW8MdKnLSGAETWb8cfwc8oht2usMo","ipaddr":"127.0.0.1","keytype":"ssh-rsa"}
2020/03/23 22:38:52 ssh-honeypot: {"user":"root","password":"foobar","ipaddr":"127.0.0.1"}

```