package pubip

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Master - services used:
//   http://ident.me
//   http://ipecho.net/plain
//   https://ifconfig.co
//   https://ipinfo.io/ip
//   https://icanhazip.com
//   http://bot.whatismyipaddress.com
//   https://myexternalip.com/raw
//   http://checkip.amazonaws.com
//   dig: myip.opendns.com @resolver1.opendns.com
//   dig: o-o.myaddr.l.google.com @ns1.google.com
type Master struct {
	Parallel int // number of service to try in parallel
}

// IPFn is an alias for a function which returns an ip address.
type IPFn func() (string, error)

// IPFuncs is an alias for a slice of IPFn functions.
type IPFuncs []IPFn

// ErrIPUnknown is returned if all services fail to return the public ip address.
var ErrIPUnknown = errors.New("Public ip address unknown")

// NewMaster object. Allows repetitive request to get the public ip address.
func NewMaster() *Master {
	rand.Seed(time.Now().Unix())
	return &Master{Parallel: 2}
}

// Address returns the public IPv4 or IPv6 address, using a random service.
// Services are queried parallel if Master.Parallel is set > 1 (default 2).
// The maximal timeout for a service-query is 1 second.
func (m *Master) Address() (string, error) {
	a := AllFuncs()
	if m.Parallel < 1 {
		m.Parallel = 1
	} else if m.Parallel > len(a) {
		m.Parallel = len(a)
	}
	return m.addressParallel(a)
}

func (m *Master) addressParallel(a IPFuncs) (string, error) {
	var wg sync.WaitGroup
	inpc := make(chan IPFn, len(a))
	resc := make(chan string)
	// start worker
	for i := 0; i < m.Parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// worker
			for fn := range inpc {
				if ip, ok := runIPFn(fn); ok {
					resc <- ip
					break
				}
			}
		}()
	}
	// feed work to worker
	go func() {
		defer close(inpc)
		for _, fn := range a {
			inpc <- fn
		}
	}()
	// wait for worker to finish
	go func() {
		defer close(resc)
		wg.Wait()
	}()
	// return first ip address, if any
	for ip := range resc {
		return ip, nil
	}
	return "", ErrIPUnknown
}

func runIPFn(fn IPFn) (string, bool) {
	ip, err := fn()
	if err != nil || !IsValid(ip) {
		return "", false
	}
	return ip, true
}
