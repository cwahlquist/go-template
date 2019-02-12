FROM golang:latest
EXPOSE 8080
ENTRYPOINT ["/go-template"]
COPY ./bin/ /
