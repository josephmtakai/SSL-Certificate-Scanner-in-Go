package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the SSL Certificate Scanner")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a website URL (or type 'exit' to quit): ")
		url, _ := reader.ReadString('\n')
		url = strings.TrimSpace(url)

		if url == "exit" {
			break
		}

		results, err := scanSSL(url)
		if err != nil {
			fmt.Println("Error scanning SSL certificate:", err)
			continue
		}

		fmt.Println("Scan Results:")
		fmt.Println(results)
	}
}

// ScanSSL checks the SSL certificate of a given URL
func scanSSL(url string) (string, error) {
	// Configure HTTP client to skip verification (for scanning purposes)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	// Make a request to get the TLS connection state
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	certificates := resp.TLS.PeerCertificates
	if len(certificates) == 0 {
		return "No SSL certificate found.", nil
	}

	// Process the certificates
	result := ""
	for _, cert := range certificates {
		result += fmt.Sprintf("Certificate Subject: %s\n", cert.Subject)
		result += fmt.Sprintf("Issuer: %s\n", cert.Issuer)
		result += fmt.Sprintf("Not Before: %s\n", cert.NotBefore)
		result += fmt.Sprintf("Not After: %s\n", cert.NotAfter)

		// Check if the certificate is expired
		if time.Now().After(cert.NotAfter) {
			result += "Status: EXPIRED\n\n"
		} else {
			result += "Status: VALID\n\n"
		}

		// Check for weak cipher suites
		if isWeakCipher(resp.TLS.CipherSuite) {
			result += "Warning: Weak Cipher Suite detected.\n"
		} else {
			result += "Cipher Suite: Strong\n"
		}
	}

	return result, nil
}

// isWeakCipher checks if the provided cipher suite is considered weak
func isWeakCipher(cipherSuite uint16) bool {
	weakCiphers := []uint16{
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_RC4_128_SHA,
		tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		// Add other weak cipher suites as necessary
	}

	for _, weak := range weakCiphers {
		if cipherSuite == weak {
			return true
		}
	}
	return false
}
