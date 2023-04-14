FROM golang:1.20 AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build/
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o start .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app/
COPY --from=builder /build/start .
RUN chmod +x /app/start
CMD ["/app/start"]
