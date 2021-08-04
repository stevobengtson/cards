FROM golang:1.16-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Hot reloading mod
RUN go get -u github.com/cosmtrek/air \
    && go get github.com/go-delve/delve/cmd/dlv@latest \
    && go get golang.org/x/tools/gopls \
    && go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get github.com/ramya-rao-a/go-outline \
    && go get honnef.co/go/tools/cmd/staticcheck

EXPOSE 8080
EXPOSE 2345

CMD air

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o cards -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/cards /usr/local/bin/cards
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/cards"]
