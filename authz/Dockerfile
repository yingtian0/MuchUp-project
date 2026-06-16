FROM golang:1.26-alpine


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/main.go
EXPOSE 9000
CMD ["./server"]
