# syntax=docker/dockerfile:1.7

# ETAP 1
FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

# Informacja o autorze
LABEL org.opencontainers.image.authors="Antoni Kędra"

WORKDIR /app

# Build arguments przekazywane przez buildx
ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=0

RUN apk add --no-cache ca-certificates

# Cache zależności GO
COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Kopiowanie całego kodu
COPY . .

# Cache kompilacji GO
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -o weather-app


# ETAP 2
FROM scratch

LABEL org.opencontainers.image.authors="Antoni Kędra"

WORKDIR /app

# Kopiowanie certyfikatów CA
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Kopiowanie binarki i pliku HTML
COPY --from=builder /app/weather-app /app/weather-app
COPY --from=builder /app/index.html /app/index.html

# Port aplikacji
EXPOSE 8080

# Healthcheck sprawdza czy serwer działa
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD ["/app/weather-app", "healthcheck"]

# Uruchomienie aplikacji
ENTRYPOINT ["/app/weather-app"]