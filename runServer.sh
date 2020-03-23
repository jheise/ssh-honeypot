#!/bin/bash

export HONEYPOT_PORT="2202"
export HONEYPOT_ADDR="0.0.0.0"
export HONEYPOT_LOGTYPE="stdout"
export HONEYPOT_LOGDEST="test-honeypot.log"
export HONEYPOT_LOGFORMAT="string"
export HONEYPOT_SERVERKEY="honeypot_rsa"

./ssh-honeypot
