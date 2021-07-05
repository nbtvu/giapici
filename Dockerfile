FROM golang:1.15.13
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main "./greater_server/main.go"

CMD ["./main"]