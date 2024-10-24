FROM golang:1.23.1 AS builder
WORKDIR /go/src/shiharai-kun
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o ./cmd/bin/app ./cmd/main.go

# リリース用のステージ
FROM gcr.io/distroless/static AS api
COPY --from=builder /go/src/shiharai-kun/cmd/bin/app /
EXPOSE 8081
ENTRYPOINT ["/app"]
