FROM golang:1.23

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o planejja cmd/api/main.go

EXPOSE 8000

RUN chmod +x planejja
CMD ["air", "-c", ".air.toml"]