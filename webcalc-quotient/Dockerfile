FROM golang:1.13.4-buster as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN GOOS=linux CGO_ENABLED=0 go build -installsuffix cgo main.go

FROM alpine:3.10.3

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"] 
