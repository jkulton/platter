# Platter

Create a file server from a directory, fast.

Platter is a simple wrapper around Go's `http.FileServer`.

## Usage

```
platter
```

Once you've installed, just run `platter` to serve the files in your current directory over http.

See below for some simple options.

## Options/Flags

(all options are as their name implies optional)

* `-p 8080` - Port to bind to, defaults to `8000`
* `-d ./files` Directory to serve files for, defaults to `.`
* `-a 192.168.1.2` Address to serve on.
* `-auth username:password` user:pass formatted authorization (for basic auth)