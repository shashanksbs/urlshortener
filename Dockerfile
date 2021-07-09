FROM golang:1.16-alpine as builder

RUN mkdir /hello
WORKDIR /hello

RUN update-ca-certificates

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o shortener .

WORKDIR /build

RUN cp /hello/shortener .

FROM scratch

COPY --from=builder /build/shortener /app/

WORKDIR /app

CMD ["./shortener"]
