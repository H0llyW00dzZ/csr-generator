// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

// # Installation
//
// To install the CSR Generator package, use the following command:
//
//	go install github.com/H0llyW00dzZ/csr-generator@latest
//
// This will install the csr-generator command-line tool in the $GOPATH/bin directory.
//
// # Usage
//
// To use the CSR Generator, run the csr-generator command:
//
//	csr-generator
//
// The command-line tool will prompt for the common name (domain name) and the number of Subject Alternative Names (SANs).
// Then, it will ask for each SAN (DNS name) to be entered.
//
// After providing the necessary information, the tool will generate a CSR and a corresponding private key.
// The CSR will be saved to a file named <common-name>.csr.pem, and the private key will be saved to a file named <common-name>.key.pem.
//
// For more detailed usage and documentation of the csr package, please refer to the package documentation:
// [github.com/H0llyW00dzZ/csr-generator/csr]
package main
