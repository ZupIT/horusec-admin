FROM golang:1.17 as build

WORKDIR /go/src/app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before copying source
RUN go mod download

RUN go get github.com/google/wire/cmd/wire@v0.5.0

# Copy the go source
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Generating the dependency injection containers
RUN go generate ./...

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o /go/bin/app ./cmd/app

# Using distroless as minimal base image
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY web/ web/
COPY --from=build /go/bin/app /app
USER 65532:65532

ENTRYPOINT ["/app"]
