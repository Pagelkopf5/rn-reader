# syntax=docker/dockerfile:1
FROM golang:1.20

WORKDIR /app
COPY . .

RUN go mod init backend && go mod tidy
RUN go build -o main .

CMD ["./main"]