# Zadanie 1, dodatkowe 2
## Użyte polecenia
### Tworzenie buildera
```bash
docker buildx create --name labbuilder --driver docker-container --use
```
<img width="751" height="92" alt="image" src="https://github.com/user-attachments/assets/1e655b36-b53d-4f3e-86fe-d756ae9d3b54" />

### Uruchamianie buildera
```bash
docker buildx inspect --bootstrap
```
<img width="936" height="1046" alt="image" src="https://github.com/user-attachments/assets/08b01998-eda4-427f-a784-915aa1422387" />

### Budowanie
```bash
docker buildx build \
        --platform linux/amd64,linux/arm64 \
        -t antek03/pawcho:latest \
        --push \
        --cache-from type=registry,ref=antek03/pawcho:buildcache \
        --cache-to type=registry,ref=antek03/pawcho:buildcache,mode=max \
        --build-arg BUILDKIT_INLINE_CACHE=1 \
        .
```

Multiarch
```bash
--platform linux/amd64,linux/arm64
```
Registry cache
```bash
--cache-to type=registry
```
Inline cache
```bash
--build-arg BUILDKIT_INLINE_CACHE=1
```

Potwierdzenie manifestu multiarch
```bash
docker buildx imagetools inspect antek03/pawcho:latest
```
<img width="1314" height="634" alt="image" src="https://github.com/user-attachments/assets/3501db7f-25fe-47ba-8fa0-afcfa165fc39" />
<img width="933" height="1000" alt="image" src="https://github.com/user-attachments/assets/a0bf86fc-b91d-4286-95dd-5c123e99cd85" />
Ponowne uruchomienie procesu budowania wykorzystywało dane cache zapisane w registry,
co potwierdzają wpisy `CACHED` w logach buildx.

## Działanie aplikacji
<img width="797" height="263" alt="image" src="https://github.com/user-attachments/assets/41c8cca4-0a7b-4cb1-b704-e01f68c4029e" />
<img width="1906" height="219" alt="image" src="https://github.com/user-attachments/assets/07610b1a-2144-4872-952d-119711a9e7dd" />
<img width="628" height="213" alt="image" src="https://github.com/user-attachments/assets/9770501d-34d1-4e7c-9f85-682fe3f0196a" />
