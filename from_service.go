package pubip

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var services = []service{
	{v6: "http://ident.me"},
	{v4: "http://ipecho.net/plain"},
	{
		v6: "https://v6.ifconfig.co",
		v4: "https://v4.ifconfig.co",
	},
	{v4: "https://ipinfo.io/ip"},
	{
		v6: "https://ipv6.icanhazip.com",
		v4: "https://ipv4.icanhazip.com",
	},
	{v6: "http://bot.whatismyipaddress.com"},
	{
		v6: "https://myexternalip.com/raw",
		v4: "https://ipv4.myexternalip.com/raw",
	},
	{v4: "http://checkip.amazonaws.com"},
	{
		v6: "https://6.ifcfg.me",
		v4: "https://4.ifcfg.me",
	},
	{v6: "https://ip.tyk.nu"},
	{v6: "https://tnx.nl/ip"},
	{
		v6: "https://l2.io/ip",
		v4: "https://www.l2.io/ip",
	},
	{v6: "https://ip.appspot.com"},
	{v4: "https://ipof.in/txt"},
	{v6: "https://wgetip.com"},
	{v4: "http://eth0.me"},
	{v6: "https://tnx.nl/ip"},
}

const (
	contentType    = "Content-Type"
	typeTextPlain  = "text/plain"
	errWrongStatus = "status code %d from %q"
)

type service struct {
	v6 url
	v4 url
}

func (ser service) ipv6func() IPFn {
	return func() (string, error) {
		ip, err := get(ser.v6)
		if err == nil && !IsIPv6(ip) {
			return "", errNotV6Address
		}
		return ip, err
	}
}

func (ser service) ipv4func() IPFn {
	return func() (string, error) {
		ip, err := get(ser.v4)
		if err == nil && !IsIPv4(ip) {
			return "", errNotV4Address
		}
		return ip, err
	}
}

// AllFuncs of IPFn with IPType t in random order.
// Don't forget to call before using this function once rand.Seed().
func AllFuncs(t IPType) IPFuncs {
	var a IPFuncs
	for _, ser := range services {
		switch {
		case (t == IPv6orIPv4 || t == IPv4) && len(ser.v4) > 0:
			a = append(a, ser.ipv4func())
			continue
		case (t == IPv6orIPv4 || t == IPv6) && len(ser.v6) > 0:
			a = append(a, ser.ipv6func())
			continue
		}
	}
	r := make(IPFuncs, len(a))
	for k, v := range rand.Perm(len(a)) {
		r[k] = a[v]
	}
	return r
}

// get returns the body of the response with a request timeout of 1 second.
func get(u url) (string, error) {
	// setup request params
	req, err := http.NewRequest("GET", string(u), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(contentType, typeTextPlain)
	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}
	// query server
	r, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	// check status-code
	if r.StatusCode != http.StatusOK {
		return "", fmt.Errorf(errWrongStatus, r.StatusCode, u)
	}
	// read result
	lr := io.LimitReader(r.Body, 64) // read max 64 bytes from server
	body, err := ioutil.ReadAll(lr)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}
