# Zadanie 1
## Opis aplikacji
Opracowana aplikacja została napisana w języku Go.
Podczas uruchomienia aplikacja:
- wypisuje datę i godzinę startu,
- wyświetla dane autora,
- uruchamia serwer HTTP na porcie 8080,
- wysyła zapytania HTTP do zewnętrznego API,
- loguje przebieg działania aplikacji na standardowe wyjście.

## Użyte komendy
### Zbudowanie obrazu kontenera
```bash
docker build -t weather-app .
```
<img width="998" height="766" alt="image" src="https://github.com/user-attachments/assets/f784de2b-0097-482d-91f3-7462ce35729c" />

### Uruchomienie kontenera na podstawie zbudowanego obrazu
```bash
docker run -dp 8080:8080 --name weather-container weather-app
```
<img width="658" height="83" alt="image" src="https://github.com/user-attachments/assets/369a76ee-8d89-4850-af34-84e34a1d434a" />

### Uzyskanie informacji z logów aplikacji
```bash
docker logs weather-container
```
<img width="619" height="471" alt="image" src="https://github.com/user-attachments/assets/4055e18e-9fdc-443d-8481-fafd63b474c5" />

### Sprawdzenie liczby warstw oraz rozmiaru obrazu
```bash
docker image history weather-app
```
<img width="1163" height="272" alt="image" src="https://github.com/user-attachments/assets/0e2958cb-d2a8-4100-ab66-a14a99e78133" />

```bash
docker image ls weather-app
```
<img width="802" height="150" alt="image" src="https://github.com/user-attachments/assets/8ae319f0-06f4-479c-8645-d637518e7d22" />

## Działanie aplikacji
<img width="625" height="194" alt="image" src="https://github.com/user-attachments/assets/78e4ed05-029e-48c5-9244-4d403f336cc3" />
<img width="625" height="194" alt="image" src="https://github.com/user-attachments/assets/b8e6ea83-6afb-4c2c-898b-d915f047a665" />
<img width="625" height="194" alt="image" src="https://github.com/user-attachments/assets/21cf1627-9d0a-4838-8e9f-1de7dacaeb49" />



