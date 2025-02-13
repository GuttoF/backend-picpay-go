FROM golang:1.24.0-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY cmd/ cmd/
COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api cmd/main.go

###############################################
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/api /app/api

RUN chmod +x /app/api

EXPOSE 3000

CMD ["/app/api"]