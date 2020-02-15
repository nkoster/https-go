package main

import (
	"fmt"
	"os"
	"strings"
)

func args() (string, string, string) {
	key := ""
	crt := ""
	www := "./"
	for i, arg := range os.Args {
		if arg == "-key" {
			if i+1 < len(os.Args) {
				key = os.Args[i+1]
			}
			arg = ""
		}
		if arg == "-crt" {
			if i+1 < len(os.Args) {
				crt = os.Args[i+1]
			}
			arg = ""
		}
		if arg == "-www" {
			if i+1 < len(os.Args) {
				www = os.Args[i+1]
				www = strings.TrimRight(www, "/")
			}
			arg = ""
		}
		if arg == "--help" || arg == "-help" || arg == "/h" {
			fmt.Printf("usage: %s -crt <certificate file> -key <key file> [-www <path>]]\n", os.Args[0])
			os.Exit(0)
		}
	}
	return key, crt, www
}
