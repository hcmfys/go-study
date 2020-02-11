package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func parsePemFile(path string) {
	certPemBlock, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	certBlock, _ := pem.Decode(certPemBlock)
	x509Cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println(x509Cert.Subject)

	fmt.Println(x509Cert.EmailAddresses)
	fmt.Println(x509Cert.PublicKey)
	fmt.Println(x509Cert.Issuer)
}
func main() {

	pwd, _ := os.Getwd()
	p := pwd + "/tls/cert/cert.pem"
	parsePemFile(p)
}
