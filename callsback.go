package main

import (
	// standard
	"fmt"
	"strings"

	// x
	"golang.org/x/crypto/ssh"
)

func passwordCallback(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
	logEvent := PasswordEvent{c.User(), string(pass), parseIp(c.RemoteAddr())}
	logPasswordEvent(&logEvent)
	// logger.Printf("user: %s password: %s ipaddr: %s\n", c.User(), string(pass), parseIp(c.RemoteAddr()))

	return nil, fmt.Errorf("Invalid Password")
	// return nil, nil
}

func publicKeyCallback(c ssh.ConnMetadata, pubkey ssh.PublicKey) (*ssh.Permissions, error) {
	keyType := pubkey.Type()
	user := c.User()
	// if the public key is actually a cert, grab the key from the cert
	if strings.Contains(keyType, "-cert-") {
		cert := pubkey.(*ssh.Certificate)
		pubkey = cert.Key
	}

	// generate a fingerprint of the key
	fingerprint := ssh.FingerprintSHA256(pubkey)

	// generate log event struct and log the event
	logEvent := PubkeyEvent{user, fingerprint, parseIp(c.RemoteAddr()), keyType}
	logPubkeyEvent(&logEvent)

	// never accept any key
	return nil, fmt.Errorf("Invalid Key")
}
