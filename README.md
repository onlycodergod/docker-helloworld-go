# go-docker-hello

[![Build Status](https://travis-ci.org/patricksanders/go-docker-hello.svg?branch=master)](https://travis-ci.org/patricksanders/go-docker-hello)

This is a simple example of building a docker image to run a Go program.

Because it uses multi-stage build, Docker CE 17.06 or newer is required.

Go version can be changed by passing the `GOLANG_VERSION` arg from the Makefile or by modifying the default in the Dockerfile.

## Usage

Do what you want. Modify `hello.go`, or replace it with your own Go program.

### building

```
make build       <----- easy to build this project
```

### running

```
make run         <----- easy run this project
```