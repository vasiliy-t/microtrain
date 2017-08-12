FROM golang:1.8.3-alpine as builder
WORKDIR /go/src/github.com/vasiliy-t/microtrain/wishlist
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v

FROM alpine:3.6
WORKDIR /app
COPY --from=builder /go/src/github.com/vasiliy-t/microtrain/wishlist/wishlist .
CMD ["/app/wishlist"]
