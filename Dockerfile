FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app
COPY . /app
RUN go build -o main main.go

# build small image
FROM alpine:3.18

# setup timezone
ENV TZ=Asia/Bangkok
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]