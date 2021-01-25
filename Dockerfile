# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /go/src/idan-testis/
# Copy the Go Modules manifests
COPY . .

RUN GO111MODULE=on go get .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /go/src/idan-testis/ .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
