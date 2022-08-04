FROM golang:1.18-buster as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /whisper

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /whisper /whisper

COPY --from=build /app/internal/infrastructure/config/config.yaml /config.yaml

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/whisper", "serve", "--config", "/config.yaml"]