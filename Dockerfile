FROM golang:1.15.2 AS builder

WORKDIR /go/src/github.com/ibreakthecloud/infra-mon
# Copy all files
COPY . .
# Enable Go Modules
ENV GO111MODULE=on
# Fetch dependencies before go build
RUN go mod download
# Build the binary
RUN CGO_ENABLED=0 go build -o metrics-server

FROM alpine AS final
COPY --from=builder /go/src/github.com/ibreakthecloud/infra-mon/metrics-server /
EXPOSE 8080
ENTRYPOINT ["/metrics-server"]