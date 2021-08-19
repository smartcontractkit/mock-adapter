FROM golang:alpine as builder

# Setup app folder
RUN mkdir /app
ADD . /app
WORKDIR /app

# Build th app
RUN go mod download && \
    CGO_ENABLED=0 go build -o main . && \
    chmod +x ./main

# Move to small scratch image
FROM scratch
COPY --from=builder /app/main /app/main
EXPOSE 6060

ENTRYPOINT ["/app/main"]