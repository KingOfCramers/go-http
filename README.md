# go-http

Basic http server with Docker and Go, with CircleCI for continuous integration.

## Development

Build the Docker image:

`docker build -t go-http:latest .`

And run it locally:

`docker run -it -p 8090:8090 go-http`
