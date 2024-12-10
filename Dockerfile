FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build -o planejja cmd/api/main.go

EXPOSE 8000

RUN chmod +x planejja
CMD ["./planejja", "air", "-c", ".air.toml"]
