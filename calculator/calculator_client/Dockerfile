FROM golang:1.15
COPY . /client
WORKDIR /client
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build /client