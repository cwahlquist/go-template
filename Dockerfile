FROM scratch
EXPOSE 8080
EXPOSE 8090
ENTRYPOINT ["/go-template"]
COPY ./bin/ /
