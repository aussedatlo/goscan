package app

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Ullaakut/nmap"
	"golang.org/x/crypto/ssh"
)

func getTargetsUp(target string) []string {
	var l []string

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPingScan(),
	)

	if err != nil {
		Error("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		Error("nmap scan failed: %v", err)
	}

	for _, host := range result.Hosts {
		Debug("Host %s is up \n", host.Addresses[0])
		l = append(l, host.Addresses[0].Addr)
	}

	return l
}

func getProductVersion(host string) (string, string) {

	pk, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")

	signer, err := ssh.ParsePrivateKey(pk)

	if err != nil {
		Error("%s", err)
	}

	config := &ssh.ClientConfig{
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		User: "louis",
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 5,
	}

	Debug("Trying to connect with " + host + ":22")
	client, err := ssh.Dial("tcp", host+":22", config)

	if err != nil {
		// Error("Failed to dial: ")
		Error("Failed to dial: " + err.Error())
		return "-", "-"
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		Error("Failed to create session: " + err.Error())
		return "-", "-"
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("cat /etc/ucineo_version"); err != nil {
		Error("Failed to run: " + err.Error())
		return "-", "-"
	}
	Debug(strings.Replace(b.String(), "\n", " ", -1))
	Debug(b.String())

	r := strings.Split(b.String(), "_")
	if len(r) > 1 {
		return r[0], r[1]
	}

	return "", ""
}
