FROM golang as builder
WORKDIR /go/src/github.com/jmonsinjon/go-system-resources
RUN go get -d -v github.com/mackerelio/go-osstat/memory
COPY main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/jmonsinjon/go-system-resources/main .
CMD ["./main"]