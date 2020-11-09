FROM golang:1.15.4-alpine3.12 as builder

WORKDIR /app

COPY go.mod main.go ./

RUN go build


FROM alpine:3.12.1

WORKDIR /app

COPY --from=builder /app/abeja .

CMD ./abeja
