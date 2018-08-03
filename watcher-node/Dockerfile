# Build
FROM golang:1.9.2-alpine as builder

WORKDIR /go/src/github.com/aitkenster/file-watcher/watcher-node
COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o watcher-node .

# Run
FROM scratch

COPY --from=builder /go/src/github.com/aitkenster/file-watcher/watcher-node/watcher-node .

CMD ["./watcher-node"]
