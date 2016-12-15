package pubip

import "testing"
import "strings"

func TestAllFuncs_v4(t *testing.T) {
	for _, f := range AllFuncs(IPv4) {
		if _, err := f(); err != nil && err != errNotV4Address {
			if e := err.Error(); strings.HasPrefix(e, "status code") ||
				strings.Contains(e, "Client.Timeout exceeded") {
				continue
			}
			t.Error(err)
		}
	}
}

func TestAllFuncs_v6(t *testing.T) {
	for _, f := range AllFuncs(IPv6) {
		if _, err := f(); err != nil && err != errNotV6Address {
			if e := err.Error(); strings.HasPrefix(e, "status code") ||
				strings.Contains(e, "Client.Timeout exceeded") {
				continue
			}
			t.Error(err)
		}

	}
}
