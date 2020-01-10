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

### Installation

If you are using Linux/MacOS. you can install it with following command:

```shell
# install latest version
wget -qO- https://raw.githubusercontent.com/axetroy/ip-cli/master/install.sh | bash
# or install specified version
wget -qO- https://raw.githubusercontent.com/axetroy/ip-cli/master/install.sh | bash -s v0.1.0
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/dvs/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:~/bin"
```

finally, try it out.

```bash
dvs --help
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
