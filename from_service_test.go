package pubip

import "testing"

func TestAllFuncs2_v4(t *testing.T) {
	for _, f := range AllFuncs(IPv4) {
		check(t, f)
	}
}

func TestAllFuncs2_v6(t *testing.T) {
	for _, f := range AllFuncs(IPv6) {
		check(t, f)
	}
}

func check(t *testing.T, fn IPFn) {
	ip, err := fn()
	//fmt.Printf("%q\n", ip)
	if err != nil {
		t.Error(err)
	}
	if !IsValid(ip) {
		t.Fail()
	}
}
