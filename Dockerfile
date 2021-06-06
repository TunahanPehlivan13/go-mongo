FROM golang:1.16-alpine as go-builder

ARG VERSION=master

WORKDIR /build

# Install the dependencies
COPY . .

# Build the docker-gen executable
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o go-mongo ./cmd/api/main.go

FROM alpine:3.13

ENV DOCKER_HOST unix:///tmp/docker.sock

# Install packages required by the image
RUN apk add --no-cache --virtual .bin-deps openssl

# Install docker-gen from build stage
COPY --from=go-builder /build/go-mongo /usr/local/bin/go-mongo

ENTRYPOINT ["/usr/local/bin/go-mongo"]