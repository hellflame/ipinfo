# ipinfo

little tool for ipinfo lookup

ref API: `https://ipinfo.io`

## install

```bash
go get github.com/hellflame/ipinfo
```

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

