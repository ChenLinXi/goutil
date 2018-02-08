package sshconnect

import (
	"os"
	"log"

	"github.com/goutil/common/json"
)

// JSON Binding protocol
var serverPath [] struct {
	Server   string `json:"server"`
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

const confPath = "./conf/server.json"

// class initialize func
func init() {
	if err := json.Parse(confPath, &serverPath); err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	log.Fatalln("hero config goes, %s", serverPath)
	if len(serverPath) == 0 {
		log.Fatalln("server is nil")
		os.Exit(-1)
	}
}

// run command
func RunCommand(target Target, cmd string) (string, error) {
	result, err := target.Run(cmd)
	if err != nil {
		return `Err: 执行命令出错`, err
	}
	return result, nil
}

// run command and keep tcp window open
func RunCommandWait(target Target, cmd string) (string, error) {
	result, err := target.RunWait(cmd)
	if err != nil {
		return `Err: 执行命令出错`, err
	}
	return result, nil
}

// build SSH target
func BuildSSHTarget(address, user, password string, port int) *Target {
	return &Target{
		IP:       address,
		Port:     port,
		Username: user,
		Password: password,
	}
}

// build SSH connect
// @param serverName target definition
func BuildSSHConnect(serverName string) *Target {
	for i := 0; i < len(serverPath); i++ {
		if serverPath[i].Server == serverName {
			return BuildSSHTarget(serverPath[i].Address, serverPath[i].User, serverPath[i].Password, serverPath[i].Port)
		}
	}
	return nil
}
