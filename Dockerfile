FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app/main.go

FROM alpine:3.22

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

USER 405

CMD [ "/app/main" ]
