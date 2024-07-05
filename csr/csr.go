// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package csr

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
)

// GenerateCSRAndPrivateKey generates a Certificate Signing Request (CSR) and a private key.
//
// The function takes the following parameters:
//   - commonName: The common name (domain name) for the certificate.
//   - dnsNames: A slice of Subject Alternative Names (SANs) for the certificate.
//
// The function returns the following:
//   - csrPEM: The generated CSR in PEM format.
//   - privateKeyPEM: The generated private key in PEM format.
//   - err: An error, if any occurred during the generation process.
//
// The function generates an ECDSA private key using the P-256 curve and creates a CSR template
// with the provided common name and SANs. The CSR is signed using the private key and encoded
// in PEM format. The private key is also encoded in PEM format.
//
// Example usage:
//
//	csrPEM, privateKeyPEM, err := GenerateCSRAndPrivateKey("example.com", []string{"www.example.com", "api.example.com"})
//	if err != nil {
//	    // Handle the error
//	}
//	// Use csrPEM and privateKeyPEM for further processing or saving to files
func GenerateCSRAndPrivateKey(commonName string, dnsNames []string) (csrPEM, privateKeyPEM []byte, err error) {
	// Generate a new ECDSA key pair using the P-256 curve
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Generate a CSR template with the provided common name and SANs
	subject := pkix.Name{
		CommonName: commonName,
	}
	csrTemplate := x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
		DNSNames:           dnsNames,
	}

	// Create and sign the CSR using the private key
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode the CSR in PEM format
	csrPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	// Encode the private key in PEM format
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}
	privateKeyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})

	return csrPEM, privateKeyPEM, nil
}
