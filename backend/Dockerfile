FROM golang:1.24.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mv ./internal/common/config/.env.docker ./internal/common/config/.env

RUN go build -o ./family-planner ./cmd/main.go

EXPOSE 8080

CMD ["./family-planner"]