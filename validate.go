package pubip

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// IsValid returns false if the ip address format is faulty.
func IsValid(ip string) bool {
	ipLen := len(ip)
	if ipLen == 0 {
		return false
	}
	if ipLen <= 15 { // max IPv4 len is 15 bytes: 255.255.255.255
		if IsIPv4(ip) {
			return true
		}
	}
	return IsIPv6(ip)
}

// IsIPv6 is true if the passed string is a valid IPv6 address.
func IsIPv6(ip string) bool {
	// is empty or exceeds max length
	if l := len(ip); l == 0 || l > 45 {
		return false
	}
	a := strings.Split(ip, ":")
	// check if last block is IPv4 notation
	if strings.Contains(ip, ".") {
		// check IPv4 block
		if !IsIPv4(a[len(a)-1]) {
			return false
		}
		// remove IPv4 block from slice
		a = a[:len(a)-1]
	}
	// check block count
	if l := len(a); l < 1 || l > 8 {
		return false
	}
	var hasContent bool
	for _, block := range a {
		blockLen := len(block)
		if !hasContent && blockLen > 0 {
			hasContent = true
		}
		// skip block containing only one zero – it is a valid value
		if block == "0" {
			continue
		}
		// maybe add leading zeros to block, eg db8 to 0db8
		if blockLen != 4 {
			block = fmt.Sprintf("%#04s", block)
		}
		if b, err := hex.DecodeString(block); err != nil || len(b) == 0 {
			return false
		}
	}
	return hasContent
}

// IsIPv4 is true if the passed string is a valid IPv4 address.
func IsIPv4(ip string) bool {
	// is empty
	if len(ip) == 0 {
		return false
	}
	a := strings.Split(ip, ".")
	// check block count
	if len(a) != 4 {
		return false
	}
	for _, block := range a {
		if i, err := strconv.Atoi(block); err != nil || i < 0 || i > 255 {
			return false
		}
	}
	return true
}
