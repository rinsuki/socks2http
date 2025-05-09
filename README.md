# [EXPERIMENTAL] socks2http

[![Build](https://github.com/rinsuki/socks2http/actions/workflows/build.yml/badge.svg)](https://github.com/rinsuki/socks2http/actions/workflows/build.yml)

```
[CLIENT]
   | ^
   | |       protocol: HTTP (proxy)
   v |
[socks2http] <- this
   | ^
   | |       protocol: SOCKS5 (proxy)
   v |
[ANY SOCKS5 PROXY]
   | ^
   | |       protocol: TLS over HTTP (https) 
   v |
[SERVER]
```
