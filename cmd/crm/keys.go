package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func genPairToFile(fileNamePrefix string) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("can't generate key pair: %s", err)
	}

	{
		bpriv, err := x509.MarshalPKCS8PrivateKey(privKey)
		if err != nil {
			log.Fatalf("failed to marshal private key: %s", err)
		}

		privateBlock := &pem.Block{
			Type:  "ED25519 PRIVATE KEY",
			Bytes: bpriv,
		}

		privKeyFileName := fileNamePrefix
		privKeyFile, err := os.OpenFile(privKeyFileName, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return nil, nil, fmt.Errorf("can't write to file %s: %s", privKeyFileName, err)
		}

		err = pem.Encode(privKeyFile, privateBlock)
		if err != nil {
			return nil, nil, fmt.Errorf("can't encode to %s: %s", privKeyFileName, err)
		}
	}

	{
		bpub, err := x509.MarshalPKIXPublicKey(pubKey)
		if err != nil {
			log.Fatalf("can't marshal private key: %s", err)
		}

		publicBlock := &pem.Block{
			Type:  "ED25519 PUBLIC KEY",
			Bytes: bpub,
		}

		pubKeyFileName := fileNamePrefix + ".pub"
		pubKeyFile, err := os.OpenFile(pubKeyFileName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, nil, fmt.Errorf("can't write to file %s: %s", pubKeyFileName, err)
		}

		err = pem.Encode(pubKeyFile, publicBlock)
		if err != nil {
			return nil, nil, fmt.Errorf("can't encode to %s: %s", pubKeyFileName, err)
		}
	}

	return pubKey, privKey, nil
}

func readPrivateKeyFromFile(fileName string) (ed25519.PrivateKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't read private key from file: %s", err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can't read private key from file: %s", err)
	}

	block, _ := pem.Decode(b)

	bb, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	return bb.(ed25519.PrivateKey), nil
}
