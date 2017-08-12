FROM golang:1.8.3-alpine as builder
WORKDIR /go/src/github.com/vasiliy-t/microtrain
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o customer/customer customer/main.go

FROM alpine:3.6
WORKDIR /app
COPY --from=builder /go/src/github.com/vasiliy-t/microtrain/customer/customer ./
CMD ["./customer"]
