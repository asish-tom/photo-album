FROM golang:latest
RUN mkdir /app
WORKDIR /app
CMD ["/app/entry_point.sh"]