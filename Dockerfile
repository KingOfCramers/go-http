FROM golang:1.16.3

WORKDIR /go/src/go-http

COPY . .

# Fetch dependencies for the project...
# RUN go get -d [deps]

# Installs our program (using go.mod file) at /go/bin/go-http
RUN go install -v

# The file is automatically placed in the $PATH
CMD ["go-http"]
EXPOSE 8090
