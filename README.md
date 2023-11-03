# clash-ui

- Visit [clash-dashboard](https://github.com/Dreamacro/clash-dashboard) ui: http://localhost:8088/c/
- Visit [yacd](https://github.com/haishanh/yacd) ui: http://localhost:8088/y/

# Usage

- run it in command `./clash-ui`

```
Usage of clash-ui:
  -c string
        https://github.com/Dreamacro/clash-dashboard path  (default "/c")
  -l string
        Listen address (default ":8088")
  -y string
        https://github.com/haishanh/yacd path (default "/y")

```

- run it by docker:

```
docker run -p 8088:8088 ghcr.io/cxjava/clash-ui
```

# Install

* install from [jpillora/installer](https://github.com/jpillora/installer)

```sh
#install to /usr/local/bin
curl https://i.jpillora.com/cxjava/clash-ui! | bash
```

## Via [goblin.run](https://goblin.run)

```shell
# binary will be /usr/local/bin/clash-ui
curl -sSfL https://goblin.run/github.com/cxjava/clash-ui | sh

# to put to a custom path
curl -sSfL https://goblin.run/github.com/cxjava/clash-ui | PREFIX=/tmp sh
```


