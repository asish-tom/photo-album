FROM golang:1.15
RUN mkdir /app
WORKDIR /app
CMD ["/app/entry_point.sh"]