package testcases

import (
	"github.com/yafred/asn1-go-test/dummy"
	"testing"
)

func TestHello(t *testing.T) {
	value := dummy.Hello()

	if value != "hello" {
		t.Fatal("Wrong")
	}
}
