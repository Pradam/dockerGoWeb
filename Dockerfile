FROM golang:1.16

LABEL maintainer="pradamabhilash@gmail.com"

WORKDIR /dockerGoWeb

COPY . .

EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]