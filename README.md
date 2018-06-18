# Hello TLS

This repository contains source code and Dockerfile for demo ELS.

It is a simple application written in `go` that runs a http server to handle any request passed.

ONLY FOR Demo Purpose.

To build a static linked binary:

    $ ➜ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

To build Docker Image:

    $ ➜ docker build -t smile0x90/hello-els:1.0 .

