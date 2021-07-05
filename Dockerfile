FROM golang:1.15.13 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]


#--------------------------------------------------Final Image---------------------------------------------------
FROM alpine:latest
# create app's own directory
RUN mkdir /app
# copy binary from builder image
COPY --from=builder /app/main /app/main
# expose port
EXPOSE 8800
WORKDIR /app
CMD ["./main"]