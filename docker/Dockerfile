FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/go-todo ./app/cmd

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build /out/go-todo /go-todo

USER nonroot:nonroot

ENTRYPOINT ["/go-todo"]