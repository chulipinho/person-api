FROM golang:latest
WORKDIR /app/person-api
COPY . .
RUN go build
CMD ["./person-api"]
EXPOSE 1234