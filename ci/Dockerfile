# wiki-app Dockerfile used to build docker image on concourse ci

FROM golang:1.10.0

MAINTAINER Mohammad Dabour (https://github.com/JeffDeCola/hello-go)
# Put the binary into the container
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build -o wiki-app wiki.go wikiserver.go

# Add the ailibty to /bin/bash
#RUN apk add --update bash && rm -rf /var/cache/apk/*

# Runs the binary when someone uses it.
CMD ["./wiki-app"]
