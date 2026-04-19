FROM golang:1.23-alpine 


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/server ./cmd/main.go
EXPOSE 8080
EXPOSE 9000
CMD ["/app/server"]