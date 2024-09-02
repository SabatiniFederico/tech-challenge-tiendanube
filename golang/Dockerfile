FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main ./app/main.go
EXPOSE 8000

CMD ["./main"]