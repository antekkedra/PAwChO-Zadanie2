# Zadanie 2

## Opis rozwiązania
Zaimplementowany pipeline realizuje następujące etapy:

1. Pobranie kodu źródłowego z repozytorium
2. Budowa obrazu na podstawie Dockerfile
3. Test CVE
4. Publikacja obrazu do GHCR
5. Wykorzystanie cache w DockerHub w celu przyspieszenia buildów

---

## Etapy pipeline

- Build obrazu Docker (multi-stage build)
- Skanowanie obrazu narzędziem Trivy
- Warunkowa publikacja obrazu (tylko gdy brak HIGH/CRITICAL podatności)
- Push do GHCR

---

## Architektury

Obraz został zbudowany dla architektur:

- linux/amd64
- linux/arm64

Przy użyciu QEMU oraz Docker Buildx.

---

## Cache

W projekcie zastosowano cache typu registry:

- Backend: DockerHub
- Tryb: max
- Repozytorium: `docker.io/antek03/zadanie2-cache`

Cache przechowuje warstwy budowania, co znacząco skraca czas kolejnych uruchomień pipeline.

---

## CVE

Do analizy podatności wykorzystano narzędzie Trivy.

- Analizowane są podatności typu HIGH oraz CRITICAL
- W przypadku wykrycia krytycznych zagrożeń pipeline jest przerywany
- Obraz nie jest wtedy publikowany do GHCR

---

## Tagowanie obrazów

Zastosowano następujący schemat tagowania:

- `latest` – najnowsza stabilna wersja obrazu
- `github.sha` – unikalna identyfikacja wersji odpowiadająca commitowi

---

## Wyniki działania pipeline

Pipeline został uruchomiony automatycznie po każdym pushu do gałęzi `main`.

Efekty działania:

- Obraz został zbudowany poprawnie
- Przeskanowany pod kątem podatności
- Opublikowany w GitHub Container Registry

---

## Linki

- [GHCR image](https://github.com/antekkedra/PAwChO-Zadanie2/pkgs/container/zadanie2)
- [DockerHub cache](https://hub.docker.com/r/antek03/zadanie2-cache)

---

## Screeny

<img width="760" height="805" alt="image" src="https://github.com/user-attachments/assets/ccc3612d-9f92-4656-8e96-2903477244d3" />

<img width="1092" height="552" alt="zadanie1" src="https://github.com/user-attachments/assets/9c736607-226e-431d-928a-e474040d5b6a" />

<img width="850" height="965" alt="image" src="https://github.com/user-attachments/assets/a32253c4-0f77-41e4-8d6c-5c735f39a8fa" />

