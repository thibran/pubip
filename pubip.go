package pubip

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Master - services used:
//   http://ident.me
//   http://ipecho.net
//   https://ifconfig.co
//   https://ipinfo.io
//   https://icanhazip.com
//   http://bot.whatismyipaddress.com
//   https://myexternalip.com
//   http://checkip.amazonaws.com
//   https://ifcfg.me
//   https://ip.tyk.nu
//   https://tnx.nl
//   https://l2.io
//   https://ip.appspot.com
//   https://ipof.in
//   https://wgetip.com
//   http://eth0.me
//   https://tnx.nl
type Master struct {
	Parallel int    // number of service to try in parallel
	Format   IPType // accepted ip address return format
}

// IPFn is an alias for a function which returns an ip address.
type IPFn func() (string, error)

// IPFuncs is an alias for a slice of IPFn functions.
type IPFuncs []IPFn

// IPType of the ip address
type IPType int

const (
	// IPv6orIPv4 accept both
	IPv6orIPv4 IPType = iota
	// IPv6 only
	IPv6
	// IPv4 only
	IPv4
)

func (t IPType) String() string {
	switch t {
	case IPv6orIPv4:
		return "IPv6orIPv4"
	case IPv6:
		return "IPv6"
	case IPv4:
		return "IPv4"
	}
	panic("unknown IPType: " + string(t))
}

type url string

// ErrIPUnknown is returned if all services fail to return the public ip address.
var ErrIPUnknown = errors.New("Public ip address unknown")

// NewMaster object. Allows repetitive request to get the public ip address.
func NewMaster() *Master {
	rand.Seed(time.Now().Unix())
	return &Master{
		Parallel: 2,
		Format:   IPv6orIPv4,
	}
}

// Address returns the public IP address, using a random service.
// Services are queried parallel if Master.Parallel is set > 1 (default 2).
// IPv6 or IPv4 can be set with Master.Format (default IPv6orIPv4).
// The maximal timeout for a service-query is 2 second.
func (m *Master) Address() (string, error) {
	a := AllFuncs(m.Format) // TODO replace with new fn
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
				if ip, err := fn(); err == nil {
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
