FROM golang:1.20.3

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o rest-api    ./cmd/main.go

CMD ["./rest-api"]