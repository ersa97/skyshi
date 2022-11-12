FROM golang:alpine as builder
WORKDIR /application

COPY . .
RUN go build -o /application/application ./main.go


FROM alpine:latest

WORKDIR /usr/bin

COPY .env .env

COPY --from=builder /application/application /usr/bin/application

EXPOSE 3030

CMD application

