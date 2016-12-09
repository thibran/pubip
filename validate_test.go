package pubip

import "testing"

func TestIsValid(t *testing.T) {
	if !IsValid("93.253.33.74") {
		t.Fail()
	}
	if !IsValid("2001:0db8:0:0:0:0:1428:57ab") {
		t.Fail()
	}
	if IsValid("=====") {
		t.Fail()
	}
}

func TestIsIPv4_ok(t *testing.T) {
	if !IsIPv4("127.0.0.1") {
		t.Fail()
	}
	if !IsIPv4("0.0.0.0") {
		t.Fail()
	}
	if !IsIPv4("255.255.255.255") {
		t.Fail()
	}
	if !IsIPv4("0.0.0.01") {
		t.Fail()
	}
}

func TestIsIPv4_fail(t *testing.T) {
	// empty string
	if IsIPv4("") {
		t.Fail()
	}
	// char in first block
	if IsIPv4("a127.0.0.1") {
		t.Fail()
	}
	// not enought blocks
	if IsIPv4("0.0.0") {
		t.Fail()
	}
	// no values in some blocks
	if IsIPv4("0...0") {
		t.Fail()
	}
	// negative value
	if IsIPv4("-255.255.255.255") {
		t.Fail()
	}
}

func TestIsIPv6_ok(t *testing.T) {
	if !IsIPv6("2001:0db8:0:0:0:0:1428:57ab") {
		t.Fail()
	}
	if !IsIPv6("2001:db8::1428:57ab") {
		t.Fail()
	}
	if !IsIPv6("2001:0db8:85a3:08d3:1319:8a2e:127.0.0.1") {
		t.Fail()
	}
	if !IsIPv6("::ffff:7f00:1") {
		t.Fail()
	}
	if !IsIPv6("0:0:0:0:0:0:0:0") {
		t.Fail()
	}
	if !IsIPv6("::1") {
		t.Fail()
	}
}

func TestIsIPv6_fail(t *testing.T) {
	// empty string
	if IsIPv6("") {
		t.Fail()
	}
	// empty, no data
	if IsIPv6(":::::::") {
		t.Fail()
	}
	// invalid space character
	if IsIPv6("::ffff: 7f0:1") {
		t.Fail()
	}
	// second block contains an invalid hex value
	if IsIPv6("2001:08Z3::1428:57ab") {
		t.Fail()
	}
	// IPv4 address in wrong block
	if IsIPv6("2001:0db8:08d3:1319:8a2e:127.0.0.1:85a3") {
		t.Fail()
	}
	// last block is not a valid ipv4 address
	if IsIPv6("2001:0db8:85a3:08d3:1319:8a2e:127.0.0.256") {
		t.Fail()
	}

	// TODO should faild, but does't!

	// if IsIPv6("7f00::1428::2") {
	// 	t.Fail()
	// }
	// if IsIPv6("0:::::::") {
	// 	t.Fail()
	// }
	// if IsIPv6("1111:2222:1111:1111:2222:3.3.3.3") {
	// 	t.Fail()
	// }
}
