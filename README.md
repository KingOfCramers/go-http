# 🧮 go-http

Basic http server with Docker and Go, with CircleCI for continuous integration.

## Development

Build the Docker image:

`docker build -t kingofcramers/go-http:latest .`

And run it locally:

`docker run -it --rm -p 8090:8090 kingofcramers/go-http:latest`

## Testing

`go test`

## Production

This application runs in Kubernetes, you'll need to modify these directions to replace `kingofcramers` with your username on Docker Hub:

1. `docker build -t kingofcramers/go-http:latest .`
2. `docker push kingofcramers/go-http:latest`
3. `kubectl apply -f kubernetes/ns.yaml`
4. `kubectl apply -f kubernetes/rc.yaml`
