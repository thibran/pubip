package pubip

import (
	"os/exec"
	"strings"
)

// DigUsingOpendns resolver.
func DigUsingOpendns() (string, error) {
	return dig("myip.opendns.com", "@resolver1.opendns.com")
}

// DigUsingGoogle resolver.
func DigUsingGoogle() (string, error) {
	ip, err := dig("TXT", "o-o.myaddr.l.google.com", "@ns1.google.com")
	if err != nil {
		return "", err
	}
	return strings.Replace(ip, `"`, "", -1), nil
}

// dig runs the dig command with a timeout of 1 second.
func dig(args ...string) (string, error) {
	args = append([]string{"+time=1", "+short"}, args...)
	b, err := exec.Command("dig", args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
