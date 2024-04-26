# # stage de build
# FROM golang:1.20.14-alpine3.19 AS build
# WORKDIR /app
# COPY . /app
# RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go
# # stage imagem final
# FROM scratch
# WORKDIR /app
# COPY --from=build /app/api ./
# ARG CON_STR=postgres://postgres:changeme@postgres_container:5432/dica
# ENV CON_STR $CON_STR
# EXPOSE 8080
# CMD [ "./api" ]


# Create Builder Image, to compile the source code into an executable
FROM golang:1.20.14-alpine3.19 as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the final image, running the API and exposing port 8080
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
ARG CON_STR="postgres://postgres:changeme@postgres_container:5432/dica"
ENV CON_STR $CON_STR
ARG PORT
ENV PORT=8080
EXPOSE $PORT
CMD ["./main"]


# # Create Builder Image, to compile the source code into an executable
# FROM golang:1.17-alpine as builder
# RUN apk add --no-cache gcc musl-dev
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# # Create the final image, running the API and exposing port 8080
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/main .
# ARG PORT
# ENV PORT=$PORT
# EXPOSE $PORT
# CMD ["./main"]