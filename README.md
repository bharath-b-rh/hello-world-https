# hello-world-https
A primitive HTTPS server which has provision to configure a server through command line arguments.

# Quick Start

## Build
```shell
make build
```

## Build docker image
```
make image
```
Note: Image build makes use of [imagebuilder](https://github.com/openshift/imagebuilder) and requires binary to be present in PATH.

## Push docker image
```
make push
```
Note: Set IMAGE and TAG to required image name and tag.

# Usage
Accepts below command line arguments are supported
1. cacert: CA certificate bundle to verify client certificates.
2. cert: Certificate for the server.
3. key: Private key for the server.
4. port: Port server should listen on.
5. servername: Server Name to be used for SNI validation.
