FROM golang:1.16 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get github.com/google/wire/cmd/wire && wire ./...
RUN go get -d -v ./...

RUN go build -o /go/bin/app ./cmd/app

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/app /app
COPY --from=build /go/src/app/web /web

CMD ["/app"]
