FROM golang as builder

COPY . .

RUN CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o /app main.go

FROM scratch

COPY --from=builder /app /app

ENTRYPOINT ["/app"]
