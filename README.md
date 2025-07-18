# generate204

simple http server always respond http 204 for connectivity check.

default at 8080 port

used for filling sing-box urltest, run container on landing proxy server, urltest type outbound, for best route.

How to use

```
podman pull ghcr.io/zhaokunqi/generate204:latest
```

```
podman run -d -p 8080:8080 --name http204 ghcr.io/zhaokunqi/generate204:latest
```
