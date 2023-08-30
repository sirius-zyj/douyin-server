FROM ubuntu:22.04
FROM golang:1.20.0
WORKDIR /app
COPY . .
# RUN go mod tidy
# RUN bash run.sh
RUN go build