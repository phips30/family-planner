FROM golang:1.24.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build family-planner

EXPOSE 8080

CMD ["./family-planner"]