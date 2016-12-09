package pubip

import "testing"

func TestFromIdentMe(t *testing.T) {
	check(t, FromIdentMe)
}

func TestFromIPecho(t *testing.T) {
	check(t, FromIPecho)
}

func TestFromIfconfig(t *testing.T) {
	check(t, FromIfconfig)
}

func TestFromIPinfo(t *testing.T) {
	check(t, FromIPinfo)
}

func TestFromIcanhazip(t *testing.T) {
	check(t, FromIcanhazip)
}

func TestFromWhatismyipaddress(t *testing.T) {
	check(t, FromWhatismyipaddress)
}

func TestFromMyexternalIP(t *testing.T) {
	check(t, FromMyexternalIP)
}

func TestFromAmazon(t *testing.T) {
	check(t, FromAmazon)
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
