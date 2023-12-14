FROM golang:1.21-alpine

WORKDIR /app

COPY . .

EXPOSE 5000

CMD tail -f /dev/null
