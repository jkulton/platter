# Platter

![Platter](./platter.png)

Create a file server from a directory, fast.

Platter is a simple wrapper around the Go standard library [`net/http` `FileServer`](https://golang.org/pkg/net/http/#FileServer).

## Install

Before installing please have Go installed and `$GOPATH/bin` in your `$PATH`. ([More info here](https://golang.org/doc/install#install))

```sh
go get github.com/jkulton/platter/cmd/platter
```

## Usage

```sh
platter
```

Once you've installed, just run `platter` to serve the files in your current directory over HTTP.

## Options/Flags

(all options are as their name implies optional)

| Flag | Default | Description |
|------|---------|-------------|
| `-p 8080` | `8000`  | Port to bind to |
| `-d ./files` | `.` | Local directory of files to serve |
| `-a 192.168.1.2` | `0.0.0.0` | Address to serve on |
| `-auth username:password` | No default | `user:pass` formatted authorization (for HTTP basic auth) |
