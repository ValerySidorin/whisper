FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /whisper

FROM alpine:3.14

WORKDIR /

COPY --from=build /whisper /whisper

COPY --from=build /app/internal/infrastructure/config/config.yaml /config.yaml

EXPOSE 8080

ENTRYPOINT [ "/whisper", "serve", "--config", "/config.yaml"]