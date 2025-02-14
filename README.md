Build docker image
```
docker build --tag gober .
```

List images
```
docker image ls
```

Run container
```
docker run --publish 8080:8080 gober
```