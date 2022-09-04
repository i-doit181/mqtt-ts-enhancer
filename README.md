# mqtt-ts-enhancer
Enhance mqtt messages with a timestamp

### Build image

#### arm64
```
docker build --platform linux/arm64 --build-arg goarch=arm64 --progress=plain -t <imageName> -f Dockerfile .
```

#### amd64
```
docker build --platform linux/amd64 --progress=plain -t <imageName> -f Dockerfile .
```