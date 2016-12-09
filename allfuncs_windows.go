package pubip

import "math/rand"

// AllFuncs of IPFn in random order.
// Don't forget to call once rand.Seed() before using this function.
func AllFuncs() IPFuncs {
	a := IPFuncs{
		// using the dig command-line application
		//   non -> only used on unix.
		//
		// using http-get
		FromIdentMe,
		FromIPecho,
		FromIfconfig,
		FromIPinfo,
		FromIcanhazip,
		FromWhatismyipaddress,
		FromMyexternalIP,
		FromAmazon,
	}
	r := make(IPFuncs, len(a))
	for k, v := range rand.Perm(len(a)) {
		r[k] = a[v]
	}
	return r
}
