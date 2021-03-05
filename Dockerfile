FROM golang:1.16 as build

WORKDIR /go/src/app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before copying source
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Generating the dependency injection containers
RUN go generate ./...

# Build
RUN go build -o /go/bin/app ./cmd/app

# Using distroless as minimal base image
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY web/ web/
COPY --from=build /go/bin/app /app

ENTRYPOINT ["/app"]
