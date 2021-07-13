FROM alpine:3.14

COPY ./dummy-external-adapter /usr/local/bin/dummy-external-adapter
RUN chmod +x /usr/local/bin/dummy-external-adapter
EXPOSE 6060

ENTRYPOINT ["./usr/local/bin/dummy-external-adapter"]