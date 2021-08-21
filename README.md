# tcpecho

A simple TCP Echo Server written in golang.

Useful for testing firewalls, etc.

Based on an example from the [Black Hat Go](https://nostarch.com/blackhatgo)
book from no starch press.

## Install

```
go install github.com/mokytis/tcpecho@latest
```

## Usage

`tcpecho` takes a host and port. These are both optional.

```
tcpecho -host 127.0.0.1 -port 8080
```
