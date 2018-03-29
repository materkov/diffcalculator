FROM golang:1.10.0
WORKDIR /go/src/github.com/materkov/diffcalculator
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /app cmd/main.go


FROM alpine:3.7
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 8000
CMD ["/app"]
COPY --from=0 /app /app
