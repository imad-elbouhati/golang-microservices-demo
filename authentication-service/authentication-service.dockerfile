FROM golang:1.19-alpine3.17 AS builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
RUN chmod +x /app/authApp

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/authApp /app/


CMD [ "/app/authApp" ]