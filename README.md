# cmoli.es-deploy

## Introduction

Project to automate the creation of the [cmoli.es](https://cmoli.es) website content:

- Download the website content (MD, CSS, JS, media, etc).
- Convert MD files to HTML.
- Copy the images and videos to the required paths.
- Etc.

## Configuration

### Media content

The media content (images, videos, etc.) must be in the `$HOME/Software/cmoli-media-content` folder using the same paths as the markdown web files. This is required because the media content will be copied from this path to the web content path with the `cp -r` command.

### Required software

- [Docker](https://www.docker.com/). You need to have Docker installed and run it as [rootless](https://docs.docker.com/engine/security/rootless/).
- [Git](https://git-scm.com/).
- [Go](https://go.dev/).

### VPS configuration

First, compile the binary:

```bash
make build
```

Send the binary to your VPS:

```bash
make send
```

You can create an alias in your VPS `~/.bashrc`:

```bash
alias deploy='$HOME/Software/cmoli-es-deploy'
```

Run in the VPS:

```bash
deploy
```

## Run

```bash
make run
```
