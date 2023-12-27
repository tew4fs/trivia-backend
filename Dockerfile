FROM golang:1.21

ENV APP_HOME=/app

COPY . $APP_HOME
WORKDIR $APP_HOME

EXPOSE 3000

RUN ["go", "build", "-o", "bin/trivia-backend", "main.go"]

CMD ["./bin/trivia-backend"]