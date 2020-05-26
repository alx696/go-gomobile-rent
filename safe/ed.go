package safe

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

const (
	//小明租赁验证签名公钥
	appRentPublic = `-----BEGIN KEY-----
sdtSp4ZR3CLJEguSCLe3WWrBS6VkOjMgfqwZqxi8BzI=
-----END KEY-----`
)

func pemEncode(bs []byte, t string) (string, error) {
	block := &pem.Block{Type: t, Bytes: bs}
	buf := bytes.NewBufferString("")
	err := pem.Encode(buf, block)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func pemDecode(pemText string, t string) ([]byte, error) {
	block, _ := pem.Decode([]byte(pemText))
	if block == nil || block.Type != t {
		return nil, errors.New("数据错误")
	}
	return block.Bytes, nil
}

func GenerateKey(publicPath, privatePath string) error {
	rr := rand.Reader
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(rr)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(publicPath, publicKeyBytes, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(privatePath, privateKeyBytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Sign(privatePath, filePath, signPath string) error {
	privateKeyBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return err
	}
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	signBytes := ed25519.Sign(privateKeyBytes, fileBytes)
	err = ioutil.WriteFile(signPath, signBytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func SignText(privatePath, text string) (string, error) {
	privateKeyBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return "", err
	}
	signBytes := ed25519.Sign(privateKeyBytes, []byte(text))
	pemText, err := pemEncode(signBytes, "SIGN")
	if err != nil {
		return "", err
	}
	return pemText, nil
}

func Verify(publicPath, filePath, signPath string) bool {
	publicKeyBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return false
	}
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false
	}
	signBytes, err := ioutil.ReadFile(signPath)
	if err != nil {
		return false
	}
	return ed25519.Verify(publicKeyBytes, fileBytes, signBytes)
}

func VerifyText(publicPemText, signPemText, text string) bool {
	publicKeyBytes, err := pemDecode(publicPemText, "KEY")
	if err != nil {
		return false
	}
	signBytes, err := pemDecode(signPemText, "SIGN")
	if err != nil {
		return false
	}
	return ed25519.Verify(publicKeyBytes, []byte(text), signBytes)
}

func AppRentVerifyText(signPemText, text string) bool {
	publicKeyBytes, err := pemDecode(appRentPublic, "KEY")
	if err != nil {
		return false
	}
	signBytes, err := pemDecode(signPemText, "SIGN")
	if err != nil {
		return false
	}
	return ed25519.Verify(publicKeyBytes, []byte(text), signBytes)
}

func ToPem(keyPath string, t string) (string, error) {
	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}
	return pemEncode(keyBytes, t)
}

func FromPem(pemText, t, keyPath string) error {
	bs, err := pemDecode(pemText, t)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(keyPath, bs, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
