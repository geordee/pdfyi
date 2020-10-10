FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/geordee/pdfyi
COPY . .
RUN go get -d -v
RUN export CGO_ENABLED=0 && go build -o /go/bin/pdfyi

FROM alpine:latest

ENV APP_UPLOAD_LIMIT=
ENV APP_ALLOW_INSECURE=
ENV S3_LOCATION=
ENV S3_ENDPOINT=
ENV S3_ACCESS_KEY=
ENV S3_SECRET_KEY=
ENV S3_USE_SSL=

RUN apk add wkhtmltopdf
COPY --from=builder /go/bin/pdfyi /go/bin/pdfyi
ENTRYPOINT ["/go/bin/pdfyi"]
EXPOSE 9090
