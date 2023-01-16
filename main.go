package main

import (
	"bufio"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
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
			fmt.Println("host and port should be provided in format host:port")
			continue
		}

		// Connect to the host and port
		conn, err := tls.Dial("tcp", hostport, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			fmt.Println("Error connecting to host: ", err)
			continue
		}
		defer conn.Close()

		// Get the certificate from the connection
		cert := conn.ConnectionState().PeerCertificates[0]

		// Get the serial number of the certificate
		serial := cert.SerialNumber
		serialHex := hex.EncodeToString(serial.Bytes())

		// construct the file name with the serial number
		fileName := serialHex + ".json"

		// convert the certificate to json
		certJson, err := json.MarshalIndent(cert, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling certificate to json: ", err)
			continue
		}

		// save the certificate to file
		certOut, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file: ", err)
			continue
		}
		defer certOut.Close()

		_, err = certOut.Write(certJson)
		if err != nil {
			fmt.Println("Error writing json to file: ", err)
			continue
		}
		fmt.Println("Certificate saved to file: ", fileName)
	}
}
