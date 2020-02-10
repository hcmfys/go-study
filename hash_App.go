package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5_string32(data string) string {
	d := md5.Sum([]byte(data))
	return fmt.Sprintf("%X", d)
}

func md5_string16(data string) string {
	return md5_string32(data)[8:24]
}

func md5_str(data string) string {

	m := md5.New()
	m.Write([]byte(data))
	p := m.Sum(nil)
	return hex.EncodeToString(p)
}

func main() {

	s := md5_string32("java")
	fmt.Println(s)

	s = md5_string16("java")
	fmt.Println(s)

	s = md5_str("java")

	fmt.Println(s)
}
