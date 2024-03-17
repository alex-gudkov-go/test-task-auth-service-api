FROM golang:1.22-alpine as build
    WORKDIR /app/
    COPY go.mod go.sum .env ./
    COPY cmd/ cmd/
    COPY internal/ internal/
    COPY pkg/ pkg/
    RUN go mod download
    RUN go build -o bin/a.out cmd/app/main.go

FROM alpine:3.19 AS release
    WORKDIR /app/
    COPY --from=build /app/bin/a.out bin/
    COPY --from=build /app/.env ./
    EXPOSE 8080
    CMD [ "bin/a.out" ]
