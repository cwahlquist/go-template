FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-template"]
COPY ./bin/ /