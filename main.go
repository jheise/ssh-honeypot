package main

import (
	// standard
	"io/ioutil"
	"log"
	"net"
	"os"

	// x
	"golang.org/x/crypto/ssh"
)

var (
	logger      *log.Logger
	servicePort string
	serviceAddr string
	serviceUri  string
	logType     string
	logDest     string
	logFormat   string
	serverKey   string
	ADDRENV     = "HONEYPOT_ADDR"
	PORTENV     = "HONEYPOT_PORT"
	LOGTYPE     = "HONEYPOT_LOGTYPE"
	LOGDEST     = "HONEYPOT_LOGDEST"
	LOGFORMAT   = "HONEYPOT_LOGFORMAT"
	SERVERKEY   = "HONEYPOT_SERVERKEY"
)

func init() {
	// collect address and port to build connection string
	servicePort = getEnv(PORTENV, "22")
	serviceAddr = getEnv(ADDRENV, "0.0.0.0")
	serviceUri = serviceAddr + ":" + servicePort

	// collect logging settings
	logType = getEnv(LOGTYPE, "stdout")
	logDest = getEnv(LOGDEST, "honeypot.log")
	logFormat = getEnv(LOGFORMAT, "string")

	// path to server key
	serverKey = getEnv(SERVERKEY, "honeypot_rsa")
}

func main() {

	switch logType {
	case "file":
		fileHandle, err := os.OpenFile(logDest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		logger = log.New(fileHandle, "ssh-honeypot: ", log.LstdFlags|log.LUTC|log.Lmsgprefix)
	default:
		logger = log.New(os.Stdout, "ssh-honeypot: ", log.LstdFlags|log.LUTC|log.Lmsgprefix)
	}

	// create an sshd config handling public keys and password auth
	config := &ssh.ServerConfig{
		PasswordCallback:  passwordCallback,
		PublicKeyCallback: publicKeyCallback,
		// KeyboardInteractiveCallback: interactiveCallback,
		MaxAuthTries: 0,
	}

	// load server private key
	privateKey := loadPrivateKey()
	config.AddHostKey(privateKey)

	// create listener and wait for a connection
	listener, err := net.Listen("tcp", serviceUri)
	if err != nil {
		panic(err)
	}

	// loop
	for {
		nConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// no connection is ever accepted, nothing
		_, _, _, _ = ssh.NewServerConn(nConn, config)
	}
}
