package safe

import (
	"log"
	"testing"
)

func TestA(t *testing.T) {
	err := GenerateKey("/home/m/下载/ed25519-public", "/home/m/下载/ed25519-private")
	log.Println(err)
}

func TestB(t *testing.T) {
	err := Sign("/home/m/下载/ed25519-private", "/home/m/下载/test.txt", "/home/m/下载/test.sign")
	log.Println(err)
}

func TestC(t *testing.T) {
	ok := Verify("/home/m/下载/ed25519-public", "/home/m/下载/test.txt", "/home/m/下载/test.sign")
	log.Println(ok)
}

func TestD(t *testing.T) {
	str, err := ToPem("/home/m/下载/ed25519-public", "KEY")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(str)

	err = FromPem(str, "KEY", "/home/m/下载/ed25519-public-from_pem")
	log.Println(err)
}

func TestE(t *testing.T) {
	str, err := SignText("/home/m/下载/ed25519-private", "a")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(str)

	ok := VerifyText(appRentPublic, str, "a")
	log.Println(ok)
}
