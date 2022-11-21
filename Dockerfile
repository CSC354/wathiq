FROM golang:1.19

RUN mkdir -p /usr/src/wathiq
COPY . /usr/src/wathiq
WORKDIR /usr/src/wathiq
RUN go mod tidy
ENTRYPOINT go run cmd/wathiq.go
