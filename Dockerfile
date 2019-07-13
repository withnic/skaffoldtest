FROM golang:1.12.6-alpine3.9 as builder
COPY main.go .
RUN go build -o /app

FROM alpine:3.9
EXPOSE 8080
COPY --from=builder /app .
ENTRYPOINT ["./app"]