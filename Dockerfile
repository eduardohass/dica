# Create Builder Image, to compile the source code into an executable
FROM golang:1.20.14-alpine3.19 as builder
RUN apk update && apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN go build -a -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the final image, running the API and exposing port 8080
FROM alpine:latest as binary
WORKDIR /root/
COPY --from=builder /app/main .
# ARG PORT
ENV PORT=8080 GIN_MODE=release
EXPOSE $PORT
CMD ["./main"]