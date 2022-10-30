FROM golang:1.18-alpine3.15
# using alpine linux which is slim, so we need to install build dependencies
RUN apk update && apk add --no-cache -U build-base git curl bash make ca-certificates

COPY . /pog
WORKDIR /pog
RUN go build .

EXPOSE 8080

ENTRYPOINT [ "./pog" ]