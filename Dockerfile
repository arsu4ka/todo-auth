FROM golang:1.20 as goapi

WORKDIR /app

COPY go.mod go.sum .env ./
COPY ./internal/ ./internal/
COPY ./cmd/ ./cmd/

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/server/main.go
CMD ["/app/main"]
