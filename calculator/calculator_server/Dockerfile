FROM golang:1.15
COPY . /server
WORKDIR /server
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build /server