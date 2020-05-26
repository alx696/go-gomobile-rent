package dns

import (
	"log"
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
