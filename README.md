WIP

```sh
docker build -t my-go-docker .
docker run -p 80:8080 --rm -it -v $(pwd):/app -v /app/tmp --name my-go-docker-run my-go-docker
```
