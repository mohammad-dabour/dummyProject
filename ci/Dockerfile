# wiki-app Dockerfile used to build docker image on concourse ci

FROM golang:1.10.0

MAINTAINER Mohammad Dabour git@github.com:mohammad-dabour/dummyProject.git
# Put the binary into the container
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build -o wiki-app wiki.go wikiserver.go


# Runs the binary when someone uses it.
CMD ["./wiki-app"]
