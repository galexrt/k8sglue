/*
Copyright (c) 2018 Alexander Trost <galexrt@googlemail.com>. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cert The original code was taken from https://golang.org/src/crypto/tls/generate_cert.go.
package cert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"time"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) (*pem.Block, error) {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}, nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal ECDSA private key: %v", err)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}, nil
	default:
		return nil, nil
	}
}

// Generate generates a PEM encoded certificate and key for the given hosts
// hosts      - Hostnames and IPs to generate a certificate for
// validFrom  - Creation date formatted as Jan 1 15:04:05 2011
// validFor   - Duration that certificate is valid for
// isCA       - whether this cert should be its own Certificate Authority
// rsaBits    - Size of RSA key to generate. Ignored if --ecdsa-curve is set
// ecdsaCurve - ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521
func Generate(hosts []string, validFrom string, validFor time.Duration, isCA bool, rsaBits int, ecdsaCurve string) (string, string, error) {
	var priv interface{}
	var err error
	switch ecdsaCurve {
	case "":
		fallthrough
	case "RSA":
		priv, err = rsa.GenerateKey(rand.Reader, rsaBits)
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		return "", "", fmt.Errorf("unrecognized elliptic curve: %q", ecdsaCurve)
	}
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %s", err)
	}

	var notBefore time.Time
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", validFrom)
		if err != nil {
			return "", "", fmt.Errorf("failed to parse creation date: %s", err)
		}
	}

	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		return "", "", fmt.Errorf("failed to create certificate: %s", err)
	}

	var certOut bytes.Buffer
	if err = pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return "", "", fmt.Errorf("failed to pem encode certificate: %s", err)
	}
	keyPEMBlock, err := pemBlockForKey(priv)
	if err != nil {
		return "", "", fmt.Errorf("failed to pem block key: %s", err)
	}
	var keyOut bytes.Buffer
	if err = pem.Encode(&keyOut, keyPEMBlock); err != nil {
		return "", "", fmt.Errorf("failed to pem encode key: %s", err)
	}

	return certOut.String(), keyOut.String(), nil
}
