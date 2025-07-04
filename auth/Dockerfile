FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8082

CMD [ "./main" ]

# docker build -t auth .
# docker run --rm --name auth-cont -p 8082:8082 auth