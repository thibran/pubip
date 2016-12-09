# pubip
go library to receive the public IPv4 or IPv6 IP address

Version: 0.2  
Installation: `go get github.com/thibran/pubip`

```go
import (
    "github.com/thibran/pubip"
    "fmt"
)

func main() {
    m := pubip.NewMaster()

    // optional, set the numbe of parallel requests (default 2)
    // m.Parallel = 4

    // optional limit the result to e.g IPv6 (default IPv6orIPv4)
    // m.Format = pubip.IPv6

    ip, _ := m.Address()
    fmt.Println("pub ip:", ip)
}
```