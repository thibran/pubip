# pubip
go library to receive the public IPv4 or IPv6 IP address

Version: 0.1  
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
    ip, _ := m.Address()
    fmt.Println("pub ip:", ip)
}
```