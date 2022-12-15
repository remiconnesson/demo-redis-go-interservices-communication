FROM golang:1.19-alpine as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./main.go .

CMD ["go", "run", "main.go"]
