FROM alpine:3.16

RUN adduser -D nonroot

WORKDIR /

COPY ./hs110-exporter .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/hs110-exporter"]
