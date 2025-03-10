FROM golang:1.18.3 as build

WORKDIR /go/src/github.com/cyruzin/hexagony

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/hexagony

FROM alpine:latest  

RUN apk add ca-certificates

COPY --from=build /go/bin/hexagony .

EXPOSE 8000

CMD ["./hexagony"]