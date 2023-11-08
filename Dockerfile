FROM golang:1.19.2 AS builder
RUN useradd -u 10001 prometheus-exporter

LABEL builder=true

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... \
 && go install -v ./... \
 && go test -short ./... \
 && go build \
    -a \
    -ldflags '-s -w -extldflags "-static"' \
    -o /bin/solace_prometheus_exporter

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest


RUN microdnf install -y curl-minimal-7.76.1-26.el9_3.2 libcurl-minimal-7.76.1-26.el9_3.2


LABEL name="solace/pubsubplus-prometheus-exporter"
LABEL vendor="Solace Corporation"
LABEL version="1.0.1"
LABEL release="1.0.1"
LABEL summary="Solace PubSub+ Prometheus Exporter"
LABEL description="The Solace PubSub+ Prometheus Exporter exports Event Broker metrics for Prometheus. It is a modified version of the Community exporter for RedHat certification."

USER 10001:10001

EXPOSE 9628
ENTRYPOINT [ "/solace_prometheus_exporter", "--config-file=/etc/solace/solace_prometheus_exporter.ini" ]
CMD [ ]

COPY LICENSE /licenses/LICENSE
COPY THIRD-PARTY-LICENSES.md /licenses/THIRD-PARTY-LICENSES.md

COPY docker/solace_prometheus_exporter.ini /etc/solace/solace_prometheus_exporter.ini

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /bin/solace_prometheus_exporter /solace_prometheus_exporter
