# gyaml
Golang YAML Tool

[![CI](https://github.com/takumin/gyaml/actions/workflows/integration.yml/badge.svg)](https://github.com/takumin/gyaml/actions/workflows/integration.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/takumin/gyaml)](https://goreportcard.com/report/github.com/takumin/gyaml)

## Install

### GitHub Releases

Download a prebuilt binary from [GitHub Releases](https://github.com/takumin/gyaml/releases) and install it in $PATH.

### aqua

[aqua](https://aquaproj.github.io/) is a CLI Version Manager.

```bash
aqua g -i takumin/yaml
```

## usage

```
NAME:
   gyaml - Golang YAML Tool

USAGE:
   gyaml [global options] command [command options] [arguments...]

VERSION:
   v0.0.1 (1d21549dade673a075dbcadf91a23ae9d094047b)

AUTHOR:
   Takumi Takahashi

COMMANDS:
   completion                      command completion
   validation, validate, valid, v  yaml file validation
   help, h                         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value, -l value  log level [$LOG_LEVEL]
   --help, -h                   show help
   --version, -v                print the version
```

## validation

```
NAME:
   gyaml validation - yaml file validation

USAGE:
   gyaml validation command [command options] [file or directory...]

OPTIONS:
   --log-level value, -l value                              log level [$LOG_LEVEL]
   --type value, -t value                                   report type (default: "rdjsonl") [$TYPE]
   --include value, -i value [ --include value, -i value ]  include file extension (default: "**/*.yml", "**/*.yaml") [$INCLUDE]
   --exclude value, -e value [ --exclude value, -e value ]  exclude file extension [$EXCLUDE]
   --help, -h                                               show help
```
