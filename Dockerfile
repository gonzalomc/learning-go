FROM golang:alpine as builder
WORKDIR /app
RUN apk add --no-cache git 
RUN go get -d github.com/gorilla/mux
COPY . .
RUN go build -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["/app/app"]