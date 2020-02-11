package main

import (
	"encoding/pem"
	"fmt"
	"log"

	"github.com/fullsailor/pkcs7"
)

var certChain = `-----BEGIN CERTIFICATE-----
MIIEKjCCAhKgAwIBAgIQVUzJj/mbV3n8PS0CHQTfzzANBgkqhkiG9w0BAQsFADBk
MQswCQYDVQQGEwJVUzEWMBQGA1UECBMNTmV3IEhhbXBzaGlyZTETMBEGA1UEBxMK
UG9ydHNtb3V0aDEoMCYGA1UEAxMfU2ltcGxlQ0EgTm9uLVB1YmxpYyBUZXN0IElz
c3VlcjAeFw0xODExMjgyMzMwNDVaFw0xOTAyMjYyMzMwNDVaMBUxEzARBgNVBAMT
CkphbmUgU21pdGgwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC3iiYS
42eqJOJfF81MekU0w+8UTQgL9D9fh0BYAljgvKx1jvlg6l2CX1OPfNMxOWjDo7DB
ICRrPKxq/FH/tsPumzNWLp6Fsu559MgzkudzRHDVdCF0pa/qujeJeiDLXs7mtekc
rKW6SyrKceknSoPCAwnWr/CEBL09pG3lef7vUMfmkuSMW+L7upHpL/sHyqnyoAnh
9b7IeeQ832rg7b7VOX6zmDcCr6qndhXt6L/neidFbX736wH+fF+iKyTw7XJNflqT
YvifUvnI9rGriDZdAPVa1/a94tKHbFzil/2UEDzNo61mcucbcHqhUL4Iezpi/s5I
iyXqM9oMoMQKPSAlAgMBAAGjJzAlMAsGA1UdDwQEAwIDqDAWBgNVHSUBAf8EDDAK
BggrBgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAgEAXe8ipocWVv+Bc78ci6QjAb6N
1DN0n9X7Ai6nvekDP1hqvNUhqKmKhV8pEL7GapyH7Rz3shYJyLEwlV0zUgS3/Uvv
38Ae5xl2uQLl4eoMz4T9NXewZyRmeSyfwz1za93wKGKz6IhoYRepI4EwaPjYowok
nNrocMRFuZHeUysSdLNXtgxKvtRYFI63rjNikJNM7C8mKOzeVobdegZipALxonDb
FcVUhikyu6YkT3Rc5X/oW5I2LfCl9v8mhjbuIPmLsZJyTcDBK81AFIX8g7Iavb8L
buRBIgSTShQISPrunnxbcQg3YimdNCn0n5llDmUP3jxWqiuvEZ7pSAUU6aVY7x8B
fAq9utHTz63FqVCLyl8us/oYZmeYpxpw5HSFhqwujKzJhvO1raaoU+3zsybrsxAy
tzgm05WucSoMTjhOBZFr8OVQxKnMHwgNudwsxuXqqFhUBV5JLkhxWIZxUoH3fhgz
9b2yH0pf/Vgurglfk/onMYR33B0grAT6/NW294oUOKCP9jdwNPQr+eRgoDU6hZ/P
UZMfr+dhVpIHPouSKCrkNXKLrBLFZsg8UiyMXiNB27OBK+mWCH8gUv5EPjW1PUCV
338v19sruqtHRRs7ZnWYMMxGeWosJ4eK/ysTjCespenOeC1HBVEnHQE+H/JvN7Pc
GN3nlK0GRJl5j3y6nmw=
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEKzCCAhOgAwIBAgIRAKju/EHi+Hw4PQZENuj9F1AwDQYJKoZIhvcNAQELBQAw
ZDELMAkGA1UEBhMCVVMxFjAUBgNVBAgTDU5ldyBIYW1wc2hpcmUxEzARBgNVBAcT
ClBvcnRzbW91dGgxKDAmBgNVBAMTH1NpbXBsZUNBIE5vbi1QdWJsaWMgVGVzdCBJ
c3N1ZXIwHhcNMTgxMTI4MjMzMTQ2WhcNMTkwMjI2MjMzMTQ2WjAVMRMwEQYDVQQD
EwpGcmVkIEpvbmVzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt4om
EuNnqiTiXxfNTHpFNMPvFE0IC/Q/X4dAWAJY4LysdY75YOpdgl9Tj3zTMTlow6Ow
wSAkazysavxR/7bD7pszVi6ehbLuefTIM5Lnc0Rw1XQhdKWv6ro3iXogy17O5rXp
HKyluksqynHpJ0qDwgMJ1q/whAS9PaRt5Xn+71DH5pLkjFvi+7qR6S/7B8qp8qAJ
4fW+yHnkPN9q4O2+1Tl+s5g3Aq+qp3YV7ei/53onRW1+9+sB/nxfoisk8O1yTX5a
k2L4n1L5yPaxq4g2XQD1Wtf2veLSh2xc4pf9lBA8zaOtZnLnG3B6oVC+CHs6Yv7O
SIsl6jPaDKDECj0gJQIDAQABoycwJTALBgNVHQ8EBAMCA6gwFgYDVR0lAQH/BAww
CgYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggIBAImzRaStB2tb0tCs0x9d0NTy
Z8Q3Fkexvt816xfBGiRS1AtetnV59sJKJvPe9qRoKQDRk7Q+/HO80Iu1o7dWfp7S
h29g9a1uhIVCY1ijr8cW8La9H/OF2KxgGX8TOldrftUl60sA3riJ7lQYOW0NkU+o
wrRsIlMCLhP8NZYjKn2Wf066JtV6Z4Be2CgVpaXkuTIE3h4BOv0kG9OsMvuRMqZw
n7z6EqhduSujwHevB2dIZgixacnE5v5hnpZm1ujzlgbAAZWh7uFthktXL5fBbkW7
heoOv320GOTvqkGVVlc4Pac5kR+P6JWfCDETfIVvtyTtfei5tDm606rBa3BHtiRu
kO4m63fnYziUzRXrF/bDcAzVx5jHQWugpxeI5UcaF6e5psOvhlP6O0sIeN6ThY/I
f+zZN6kf8eZ0OYk+61fVPbbpkF05AFhKkduIRfywUy5+iXrPK4iuQxhM/kRCO0H3
rxXcDlc+9Ol1JiJNNF2lV8+U1APsPFK0gnrxoYLFyRsejCiV8/D67v8oEk0gfiMm
/uZUhC2jYJ05lfK7aGV72Cf82g46zAuNAiH4zwvmxvC/2tcqcaYFyK15D1dQGqOg
z3LR6viX/ncO72ywQWgjRc5hdR5vLinIEXTRlNKm4EW+AoLY3x+Erhmq69EEnGMr
4/GUXRuNB7Is+lFM3JYE
-----END CERTIFICATE-----`

func main() {

	// Decode each PEM block in the input and append the ASN.1
	// DER bytes for each certificate therein to the data slice.

	input := []byte(certChain)
	data := []byte{}

	for len(input) > 0 {
		var block *pem.Block
		block, input = pem.Decode(input)
		data = append(data, block.Bytes...)
	}

	// Build a PKCS#7 degenerate "certs only" structure from
	// that ASN.1 certificates data.

	var err error
	data, err = pkcs7.DegenerateCertificate(data)
	if err != nil {
		log.Fatalf("couldn't create degenerate PKCS7 object: %v", err)
	}

	// Convert the PKCS#7 structure to a PEM-encoded string.

	pemString := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PKCS7",
		Bytes: data,
	}))

	// Print the string, or do whatever you want with it.

	fmt.Printf("%s", pemString)
}
