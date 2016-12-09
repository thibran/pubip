package pubip

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	timeoutNormal = time.Duration(600 * time.Millisecond)
	timeoutLong   = time.Duration(800 * time.Millisecond)
)

// FromIdentMe - http://ident.me
func FromIdentMe() (string, error) {
	return get("http://ident.me", timeoutNormal)
}

// FromIPecho - http://ipecho.net/plain
func FromIPecho() (string, error) {
	return get("http://ipecho.net/plain", timeoutNormal)
}

// FromIfconfig - https://ifconfig.co
func FromIfconfig() (string, error) {
	ip, err := get("https://ifconfig.co", timeoutLong)
	return strings.TrimSpace(ip), err
}

// FromIPinfo - https://ipinfo.io/ip
func FromIPinfo() (string, error) {
	ip, err := get("https://ipinfo.io/ip", timeoutLong)
	return strings.TrimSpace(ip), err
}

// FromIcanhazip - https://icanhazip.com
func FromIcanhazip() (string, error) {
	ip, err := get("https://icanhazip.com", timeoutLong)
	return strings.TrimSpace(ip), err
}

// FromWhatismyipaddress - http://bot.whatismyipaddress.com
func FromWhatismyipaddress() (string, error) {
	return get("http://bot.whatismyipaddress.com", timeoutLong)
}

// FromMyexternalIP - https://myexternalip.com/raw
func FromMyexternalIP() (string, error) {
	ip, err := get("https://myexternalip.com/raw", timeoutLong)
	return strings.TrimSpace(ip), err
}

// FromAmazon - http://checkip.amazonaws.com
func FromAmazon() (string, error) {
	ip, err := get("http://checkip.amazonaws.com", timeoutNormal)
	return strings.TrimSpace(ip), err
}

// get returns the body of the response with a request timeout of 1 second.
func get(url string, d time.Duration) (string, error) {
	client := http.Client{
		Timeout: time.Duration(d),
	}
	r, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	lr := io.LimitReader(r.Body, 64) // read max 64 bytes from server
	body, err := ioutil.ReadAll(lr)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
