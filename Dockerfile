FROM golang:1.15.4-alpine3.12 as builder

WORKDIR /app

RUN apk add make git

COPY Makefile .
COPY go.mod .
COPY cmd cmd

RUN make build


FROM alpine:3.12.1

WORKDIR /app

COPY --from=builder /app/abeja .

EXPOSE 9000

CMD ./abeja
