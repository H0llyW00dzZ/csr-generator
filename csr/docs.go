// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

// Package csr provides functionality for generating Certificate Signing Requests (CSRs) and private keys.
//
// The package exports a single function, [GenerateCSRAndPrivateKey], which generates an ECDSA private key
// using the P-256 curve and creates a CSR template with the provided common name and Subject Alternative
// Names (SANs). The CSR is signed using the private key and encoded in PEM format. The private key is also
// encoded in PEM format.
//
// # Usage
//
// The [GenerateCSRAndPrivateKey] function takes the following parameters:
//   - commonName: The common name (domain name) for the certificate.
//   - dnsNames: A slice of Subject Alternative Names (SANs) for the certificate.
//
// The function returns the following:
//   - csrPEM: The generated CSR in PEM format.
//   - privateKeyPEM: The generated private key in PEM format.
//   - err: An error, if any occurred during the generation process.
//
// Example:
//
//	commonName := "example.com"
//	dnsNames := []string{"www.example.com", "api.example.com"}
//
//	csrPEM, privateKeyPEM, err := csr.GenerateCSRAndPrivateKey(commonName, dnsNames)
//	if err != nil {
//	    log.Fatalf("Failed to generate CSR and private key: %v", err)
//	}
//
//	// Save the CSR and private key to files
//	err = os.WriteFile("example.com.csr.pem", csrPEM, 0644)
//	if err != nil {
//	    log.Fatalf("Failed to write CSR to file: %v", err)
//	}
//
//	err = os.WriteFile("example.com.key.pem", privateKeyPEM, 0600)
//	if err != nil {
//	    log.Fatalf("Failed to write private key to file: %v", err)
//	}
package csr
