FROM scratch
ARG bin=x-golang-linux-amd64
COPY $bin /hello
ENTRYPOINT ["/hello"]
