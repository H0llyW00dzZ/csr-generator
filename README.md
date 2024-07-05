# CSR Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/H0llyW00dzZ/csr-generator/csr.svg)](https://pkg.go.dev/github.com/H0llyW00dzZ/csr-generator)

The CSR Generator is a Go package that provides functionality for generating Certificate Signing Requests (CSRs) and private keys. It simplifies the process of creating CSRs with specified common names and Subject Alternative Names (SANs).

## Features

- Generate CSRs with specified common names and SANs
- Generate corresponding ECDSA private keys
- Return CSRs and private keys in PEM format

## Currently Supported

- [x] Private CAs that rely on domain and DNS names (e.g, [`Google Cloud Private CAs`](https://cloud.google.com/security/products/certificate-authority-service))
- [x] Public CAs that rely on domain and DNS names


## TODO

- [ ] Support other elliptic curves (currently, only the P-256 curve is supported)
- [ ] Implement a configuration file for default values and settings
- [ ] Provide a web-based user interface for generating CSRs

## License

The CSR Generator package is licensed under the BSD 3-Clause License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgements

The CSR Generator package was inspired by the need for a simple and efficient way to generate CSRs and private keys in Go.
