FROM alpine:3.7
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 8000
CMD ["/diffcalculator"]
COPY /diffcalculator /diffcalculator
