FROM golang:1.19.4-alpine

ENV GO111MODULE=auto
ENV GOOS=linux

RUN mkdir /files
COPY hw.go /files
WORKDIR /files
RUN go build -o /files/hw hw.go

ENTRYPOINT ["/files/hw"]
