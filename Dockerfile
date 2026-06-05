# ETAP 1
FROM golang:1.22-alpine AS builder

# Informacja o autorze
LABEL org.opencontainers.image.authors="Antoni Kędra"

# Zmienne środowiskowe pod minimalny binarny plik (wymusza kompilacje statycznej binarki, wymusza kompilację na Linuxie, kompilacja na x86_64)
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Instalacja certyfikatów CA
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Kopiowanie pliku modułów
COPY go.mod ./
RUN go mod download

# Kopiowanie całego kodu
COPY . .

# Budowa pliku binarnego
RUN go build -ldflags="-s -w" -o weather-app


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