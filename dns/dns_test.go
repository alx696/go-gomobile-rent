package dns

import (
	"log"
	"net/url"
	"testing"
)

func TestA(t *testing.T) {
	str := DigShort("rent.app.lilu.red", 16)
	log.Println(str)
}

func TestB(t *testing.T) {
	str := LookupTXT("rent.app.lilu.red")
	log.Println(str)
}

func TestC(t *testing.T) {
	str, err := url.PathUnescape("buy=%E5%8A%9F%E8%83%BD%E9%9C%80%E8%A6%81%E6%BF%80%E6%B4%BB%E5%90%8E%E6%89%8D%E8%83%BD%E4%BD%BF%E7%94%A8%EF%BC%8C%E8%AF%B7%E5%A4%8D%E5%88%B6%E7%94%B3%E8%AF%B7%E4%BF%A1%E6%81%AF%E5%8A%A0%E5%BE%AE%E4%BF%A1957388815%E8%B4%AD%E4%B9%B0%E3%80%82")
	if err != nil {
		log.Println(err)
	}
	log.Println(str)
}
