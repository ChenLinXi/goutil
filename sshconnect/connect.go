package sshconnect

import (
	"net"
	"fmt"
	"time"
	"bytes"
	"errors"

	"golang.org/x/crypto/ssh"
)

// initialize func
func (t *Target) init() {}

// ssh client config
func (t *Target) Config() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: t.Username,
		Auth: []ssh.AuthMethod{ssh.Password(t.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
}

// execute command after connection
// @param cmd string: shell command
// @return string: sync response after executing command
func (t *Target) Run(cmd string) (string, error) {

	// create ssh connect
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", t.IP, t.Password), t.Config())
	if err != nil {
		return "", err
	}
	defer client.Close()

	// create ssh session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// output redirect
	var out bytes.Buffer
	var eout bytes.Buffer
	session.Stdout = &out
	session.Stderr = &eout

	// start session
	err = session.Start(cmd)
	if err != nil {
		return "Error: session start failed", err
	}

	// manual control till finish executing
	// it is optional to set time duration
	time.Sleep(2 * time.Second)

	if eout.Len() == 0 {
		return out.String(), nil
	} else {
		return out.String(), errors.New(eout.String())
	}
}

// execute command after connection
// @param cmd string: shell command
// @return string: sync response after executing command
func (t *Target) RunWait(cmd string) (string, error) {

	// create ssh connect
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", t.IP, t.Password), t.Config())
	if err != nil {
		return "", err
	}
	defer client.Close()

	// create ssh session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// output redirect
	var out bytes.Buffer
	var eout bytes.Buffer
	session.Stdout = &out
	session.Stderr = &eout

	// start session
	err = session.Start(cmd)
	if err != nil {
		return "Error: session start failed", err
	}

	// keep tcp window open
	session.Wait()

	if eout.Len() == 0 {
		return out.String(), nil
	} else {
		return out.String(), errors.New(eout.String())
	}
}
