package main

import (
	"os"
	"strconv"

	"github.com/admpub/sftps"
)

func main() {
	sshHost := `127.0.0.1`
	sshPort := 22
	sshUser := `root`
	sshPasswd := ``
	if port := os.Getenv(`ssh_port`); len(port) > 0 {
		portN, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
		sshPort = portN
	}
	if host := os.Getenv(`ssh_host`); len(host) > 0 {
		sshHost = host
	}
	if user := os.Getenv(`ssh_user`); len(user) > 0 {
		sshUser = user
	}
	if pwd := os.Getenv(`ssh_passwd`); len(pwd) > 0 {
		sshPasswd = pwd
	}
	paramSFTP := sftps.NewSftpParameters(sshHost, sshPort, sshUser, sshPasswd, false)
	sftp, err := sftps.New(sftps.SFTP, paramSFTP)
	if err != nil {
		panic(err)
	}
	_, err = sftp.Connect()
	if err != nil {
		panic(err)
	}
	res, length, err := sftp.Upload(`./test.txt`, `/root/text.txt`)
	if err != nil {
		panic(err)
	}
	_ = res
	_ = length
}
