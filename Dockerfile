FROM golang:latest as builder
LABEL maintainer="Asish Tom"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV KAFKA_HOST 172.17.0.2
ENV DB_HOST db
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9090
CMD ["./main"]