# Create Builder Image, to compile the source code into an executable
FROM golang:1.20.14-alpine3.19 as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main main.go

# Create the final image, running the API and exposing port 8080
FROM scratch
WORKDIR /root/
COPY --from=builder /app/main .
# ENV CON_STR $CON_STR
ARG PORT
ENV PORT=8080
EXPOSE $PORT
CMD ["./main"]