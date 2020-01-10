[![Build Status](https://github.com/axetroy/ip-cli/workflows/test/badge.svg)](https://github.com/axetroy/ip-cli/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/ip-cli)](https://goreportcard.com/report/github.com/axetroy/ip-cli)
![Latest Version](https://img.shields.io/github/v/release/axetroy/ip-cli.svg)
![License](https://img.shields.io/github/license/axetroy/ip-cli.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/ip-cli.svg)

Command-line tool to show ip address

### Usage

```bash
$ ip
local 192.168.3.15
public xxx.xxx.xxx
```

### Build from source code

Make sure you have `Golang@v1.13.1` installed.

```shell
$ git clone https://github.com/axetroy/ip-cli.git $GOPATH/src/github.com/axetroy/ip-cli
$ cd $GOPATH/src/github.com/axetroy/ip-cli
$ make build
```

### Test

```bash
$ make test
```

### License

The [MIT License](LICENSE)
