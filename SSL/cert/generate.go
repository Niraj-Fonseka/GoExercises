package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

func main() {

	pvtkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		log.Fatal(err)
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My Corp"},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	//each cert needs a unique serial number, random 128 bit number
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pvtkey.PublicKey, pvtkey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	//serialize to files
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("Failed to encode certificate to PEM")
	}
	if err := ioutil.WriteFile("./../cert.pem", pemCert, 0644); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote cert.pem\n")

	//Serialize key.pem
	privBytes, err := x509.MarshalPKCS8PrivateKey(pvtkey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		log.Fatal("Failed to encode key to PEM")
	}
	if err := ioutil.WriteFile("./../key.pem", pemKey, 0600); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote key.pem\n")
}
