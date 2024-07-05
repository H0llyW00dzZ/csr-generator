// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/H0llyW00dzZ/csr-generator/csr"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the common name (domain name): ")
	commonName, _ := reader.ReadString('\n')
	commonName = strings.TrimSpace(commonName)

	fmt.Print("Enter the number of Subject Alternative Names (SANs): ")
	sanCountStr, _ := reader.ReadString('\n')
	sanCountStr = strings.TrimSpace(sanCountStr)
	sanCount, err := strconv.Atoi(sanCountStr)
	if err != nil {
		log.Fatalf("Invalid number of SANs: %v", err)
	}

	var dnsNames []string
	for i := 0; i < sanCount; i++ {
		fmt.Printf("Enter SAN %d (DNS name): ", i+1)
		dnsName, _ := reader.ReadString('\n')
		dnsName = strings.TrimSpace(dnsName)
		dnsNames = append(dnsNames, dnsName)
	}

	csrPEM, privateKeyPEM, err := csr.GenerateCSRAndPrivateKey(commonName, dnsNames)
	if err != nil {
		log.Fatalf("Failed to generate CSR and private key: %v", err)
	}

	// Save the CSR to a file based on the domain name
	csrFileName := fmt.Sprintf("%s.csr.pem", commonName)
	err = os.WriteFile(csrFileName, csrPEM, 0644)
	if err != nil {
		log.Fatalf("Failed to write CSR to file: %v", err)
	}
	fmt.Printf("CSR saved to %s\n", csrFileName)

	// Save the private key to a file based on the domain name
	privateKeyFileName := fmt.Sprintf("%s.key.pem", commonName)
	err = os.WriteFile(privateKeyFileName, privateKeyPEM, 0600)
	if err != nil {
		log.Fatalf("Failed to write private key to file: %v", err)
	}
	fmt.Printf("Private key saved to %s\n", privateKeyFileName)
}
