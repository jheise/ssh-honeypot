package main

import (
	// standard
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	"strings"

	// x
	"golang.org/x/crypto/ssh"
)

func parseIp(addr net.Addr) string {
	addrStr := addr.String()
	ipAddr := strings.Split(addrStr, ":")[0]
	return ipAddr
}

func logPasswordEvent(logEvent *PasswordEvent) {
	switch logFormat {
	case "json":
		event, err := json.Marshal(logEvent)
		if err != nil {
			panic(err)
		}
		logger.Printf("%s\n", event)
	default:
		logger.Printf("user: %s password: %s ipaddr: %s\n", logEvent.User, logEvent.Password, logEvent.Ipaddr)
	}
}

func logPubkeyEvent(logEvent *PubkeyEvent) {
	switch logFormat {
	case "json":
		event, err := json.Marshal(logEvent)
		if err != nil {
			panic(err)
		}
		logger.Printf("%s\n", event)
	default:
		logger.Printf("user: %s pubkey: %s ipaddr: %s type: %s\n", logEvent.User, logEvent.Pubkey, logEvent.Ipaddr, logEvent.Keytype)
	}
}

func getEnv(keyName string, defaultValue string) string {
	value, exists := os.LookupEnv(keyName)
	if !exists {
		return defaultValue
	}
	return value
}

func loadPrivateKey() ssh.Signer {
	privateBytes, err := ioutil.ReadFile(serverKey)
	if err != nil {
		panic(err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		panic(err)
	}

	return privateKey
}
