FROM golang:1.16

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o main .
RUN chmod +x ./main

EXPOSE 6060

ENTRYPOINT ["/app/main"]