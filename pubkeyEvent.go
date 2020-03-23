package main

type PubkeyEvent struct {
	User    string `json:"user"`
	Pubkey  string `json:"pubkey"`
	Ipaddr  string `json:"ipaddr"`
	Keytype string `json:"keytype"`
}
