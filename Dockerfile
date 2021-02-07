FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

EXPOSE 8080 8080

CMD ["./main"]