// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package csr_test

import (
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/H0llyW00dzZ/csr-generator/csr"
)

func TestGenerateCSRAndPrivateKey(t *testing.T) {
	commonName := "example.com"
	dnsNames := []string{"www.example.com", "api.example.com"}

	csrPEM, privateKeyPEM, err := csr.GenerateCSRAndPrivateKey(commonName, dnsNames)
	if err != nil {
		t.Fatalf("GenerateCSRAndPrivateKey failed: %v", err)
	}

	// Parse the generated CSR
	csrBlock, _ := pem.Decode(csrPEM)
	if csrBlock == nil {
		t.Fatal("Failed to decode CSR PEM")
	}
	parsedCSR, err := x509.ParseCertificateRequest(csrBlock.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse CSR: %v", err)
	}

	// Verify the common name and SANs in the CSR
	if parsedCSR.Subject.CommonName != commonName {
		t.Errorf("Unexpected common name in CSR. Got: %s, Want: %s", parsedCSR.Subject.CommonName, commonName)
	}
	if len(parsedCSR.DNSNames) != len(dnsNames) {
		t.Errorf("Unexpected number of SANs in CSR. Got: %d, Want: %d", len(parsedCSR.DNSNames), len(dnsNames))
	}
	for i, dnsName := range dnsNames {
		if parsedCSR.DNSNames[i] != dnsName {
			t.Errorf("Unexpected SAN at index %d. Got: %s, Want: %s", i, parsedCSR.DNSNames[i], dnsName)
		}
	}

	// Parse the generated private key
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	if privateKeyBlock == nil {
		t.Fatal("Failed to decode private key PEM")
	}
	if privateKeyBlock.Type != "EC PRIVATE KEY" {
		t.Errorf("Unexpected private key type. Got: %s, Want: EC PRIVATE KEY", privateKeyBlock.Type)
	}
	_, err = x509.ParseECPrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}
}
