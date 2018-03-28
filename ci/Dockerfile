# wiki-app Dockerfile used to build docker image on concourse ci

FROM golang:1.10.0

MAINTAINER Mohammad Dabour git@github.com:mohammad-dabour/dummyProject.git
#Originally i have checked this (https://github.com/JeffDeCola/hello-go) he has good example for who wants ro start concourse with go app.
# Put the binary into the container
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build -o wiki-app wiki.go wikiserver.go


# Runs the binary when someone uses it.
CMD ["./wiki-app"]
