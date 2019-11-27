# ipinfo

little tool for ipinfo lookup

ref API: `https://ipinfo.io`

> this project is a migration from python version of 
> iplocate https://github.com/hellflame/iplocate

## install

```bash
go get -u github.com/hellflame/ipinfo
```

be sure `$GOPATH/bin` is in `$PATH`, 
executable program will be installed there

> if you can't use go get

download source code

```bash
git clone https://github.com/hellflame/ipinfo
cd ipinfo
go install
```

you can't just use `go build` to get executable binary, 
it will told you (because there is a `ipinfo` dir inside it)

```
go build: build output "ipinfo" already exists and is a directory
```

use `go build -o your-preferred-name` instead

## usage

```
tool for ip info lookup

Usage:
  ipinfo [flags]

Flags:
  -h, --help      help for ipinfo
      --version   version for ipinfo
```

### 1. current host info

```bash
ipinfo
```

### 2. given ip info

```bash
ipinfo 8.8.8.8
```

output

```

     8.8.8.8

City                Mountain View
Country             US
Hostname            dns.google
Loc                 37.3860,-122.0838
Org                 AS15169 Google LLC
Postal              94035
Region              California
Timezone            America/Los_Angeles
```

### 3. given host ip

```bash
ipinfo www.bing.com
```

> may be different from yours

output 

```

     202.89.233.100

City                Beijing
Country             CN
Loc                 39.9075,116.3972
Org                 AS59067 Microsoft Mobile Alliance Internet Services Co., Ltd
Region              Beijing
Timezone            Asia/Shanghai

     202.89.233.101

City                Beijing
Country             CN
Loc                 39.9075,116.3972
Org                 AS59067 Microsoft Mobile Alliance Internet Services Co., Ltd
Region              Beijing
Timezone            Asia/Shanghai
```

### 4. host/ip list

```bash
ipinfo www.bing.com 8.8.8.8 
```

### 5. as part of your program

```go
package main
import (
    "github.com/hellflame/ipinfo/ipinfo"
    "fmt"
    "net"
)

func main() {
    fmt.Println(ipinfo.IpInfo(net.ParseIP("8.8.8.8")))
}
```