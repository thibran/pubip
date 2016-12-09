package pubip

import "testing"

func TestDigUsingOpendns(t *testing.T) {
	ip, err := DigUsingOpendns()
	if err != nil {
		t.Error(err)
	}
	if !IsValid(ip) {
		t.Fail()
	}
	//fmt.Println("opendns:", ip)
}

func TestDigUsingGooglet(t *testing.T) {
	ip, err := DigUsingGoogle()
	if err != nil {
		t.Error(err)
	}
	if !IsValid(ip) {
		t.Fail()
	}
	//fmt.Println("google:\t", ip)
}
