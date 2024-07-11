FROM golang:1.22.5 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o menus-server ./cmd/menus-server

FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/menus-server /usr/local/bin/menus-server
COPY --from=builder /app/data /data
CMD ["/usr/local/bin/menus-server"]
