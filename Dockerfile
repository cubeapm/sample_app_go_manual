FROM golang:1.22.2-alpine3.19

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download -x

ADD . .

RUN go build -o ./main .

CMD ["./main"]
