# pong
Simple webserver with single health endpoint.

## Development

### Run

Run the server locally:

```
go run ./cmd/server
```

### Docker

Use Docker to build the image locally

```
# build images
docker build -t pong .

# run image
docker run pong
docker run -it pong /bin/bash
```

## OpenShift
```
# setup build
oc new-build https://github.com/tobiasdenzler/pong

# start build
oc start-build -F pong
```