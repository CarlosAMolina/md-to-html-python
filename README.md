# cmoli.es-deploy

## Introduction

Project to automate the deploy of cmoli.es.

## VPS configuration

First, build the binary:

```bash
make build
```

Send the binary to your VPS.

You can create an alias in your VPS `~/.bashrc`:

```bash
alias deploy='$HOME/Software/cmoli-es-deploy'
```

Usage:

```bash
deploy
```
