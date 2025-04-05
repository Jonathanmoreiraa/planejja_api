FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o planejja cmd/api/main.go

EXPOSE 8080

RUN chmod +x planejja
CMD ["./planejja"]