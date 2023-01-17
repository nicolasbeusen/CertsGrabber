package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var hostport string

	// read the host and port from pipeline
	for scanner.Scan() {
		hostport = scanner.Text()

		hostportSplit := strings.Split(hostport, ":")
		if len(hostportSplit) != 2 {
			continue
		}

		// Connect to the host and port
		conn, err := tls.Dial("tcp", hostport, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			continue
		}
		defer conn.Close()

		// Get the certificate from the connection
		cert := conn.ConnectionState().PeerCertificates[0]

		// Extract CommonName
		commonName := cert.Subject.CommonName
		fmt.Println(commonName)

		// Extract DNS Names
		dnsNames := cert.DNSNames

		for _, dnsName := range dnsNames {
			fmt.Printf("%s\n", dnsName)
		}
	}
}
