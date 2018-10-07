# [EXPERIMENTAL] socks2http

[![CircleCI](https://circleci.com/gh/rinsuki/socks2http.svg?style=svg)](https://circleci.com/gh/rinsuki/socks2http)

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