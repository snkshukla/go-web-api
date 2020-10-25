FROM golang:alpine AS build_base

WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .

# Build the Go app
RUN go build -o ./out/app .

# Start fresh from a smaller image
FROM alpine

COPY --from=build_base /tmp/app/out/app /

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app"]
