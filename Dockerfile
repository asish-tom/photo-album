FROM golang:1.15
RUN mkdir /app
WORKDIR /app
ENTRYPOINT ["/app/entry_point.sh"]
