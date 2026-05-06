FROM golang:1.25-alpine AS build

WORKDIR /src

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /out/go-todo ./app/cmd

FROM alpine:3.22

WORKDIR /

RUN apk add --no-cache ca-certificates

COPY --from=build /out/go-todo /go-todo

EXPOSE 8080

ENTRYPOINT ["/go-todo"]