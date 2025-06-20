FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./rest_api

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]

# docker build -t rest-api .
# docker run -d --rm --name rest-api-cont -p 8080:8080 rest-api