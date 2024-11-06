FROM golang:latest

COPY ../ /notifyme

WORKDIR /notifyme
RUN go mod tidy
ENV API_URL=http://notifyme-app:8080/notification

CMD ["sh", "-c", "cd notifymepooling && go run main.go"]