#build stage
FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN chmod +x /app/main

#run stage
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/main /app

CMD [ "/app/main" ]
