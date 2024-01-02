FROM golang:1.21-alpine

WORKDIR /app

COPY . .

EXPOSE 3000

RUN go build -o trivia-backend

CMD ["./trivia-backend"]