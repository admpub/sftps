package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/admpub/sftps"
	"github.com/webx-top/com"
)

func main() {
	sshHost := `127.0.0.1`
	sshPort := 22
	sshUser := `root`
	sshPasswd := ``
	sshKeyFile := ``
	sshUsePassphrase := false
	sshPassphrase := ``
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
	if keyfile := os.Getenv(`ssh_keyfile`); len(keyfile) > 0 {
		sshKeyFile = keyfile
	}
	if usePassphrase := os.Getenv(`ssh_use_passphrase`); len(usePassphrase) > 0 {
		sshUsePassphrase, _ = strconv.ParseBool(usePassphrase)
	}
	if passphrase := os.Getenv(`ssh_passphrase`); len(passphrase) > 0 {
		sshPassphrase = passphrase
	}
	com.Dump(map[string]interface{}{
		`sshHost`:          sshHost,
		`sshPort`:          sshPort,
		`sshUser`:          sshUser,
		`sshPasswd`:        sshPasswd,
		`sshKeyFile`:       sshKeyFile,
		`sshUsePassphrase`: sshUsePassphrase,
		`sshPassphrase`:    sshPassphrase,
	})
	paramSFTP := sftps.NewSftpParameters(sshHost, sshPort, sshUser, sshPasswd, false)
	if len(sshKeyFile) > 0 {
		pemBytes, err := ioutil.ReadFile(sshKeyFile)
		if err != nil {
			panic(fmt.Errorf(`Private Key File "%v": %v`, sshKeyFile, err))
		}
		paramSFTP.Keys(string(pemBytes), sshUsePassphrase, sshPassphrase)
	}
	sftp, err := sftps.New(sftps.SFTP, paramSFTP)
	if err != nil {
		panic(err)
	}
	_, err = sftp.Connect()
	if err != nil {
		panic(err)
	}
	_, list, err := sftp.List(".")
	if err != nil {
		panic(err)
	}
	//println(list)
	ents, err := sftp.StringToEntities(list)
	if err != nil {
		panic(err)
	}
	com.Dump(ents)
	/*
		res, length, err := sftp.Upload(`./test.txt`, `/root/text.txt`)
		if err != nil {
			panic(err)
		}
		_ = res
		_ = length
	*/
}
