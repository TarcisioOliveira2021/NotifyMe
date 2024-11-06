FROM golang:latest

COPY ../ /notifyme

WORKDIR /notifyme
RUN go mod tidy

EXPOSE 8080
CMD ["sh", "-c", "cd app && go run main.go"]


