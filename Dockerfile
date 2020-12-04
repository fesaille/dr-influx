ARG VERSION=v2.0.2
FROM quay.io/influxdb/influxdb:${VERSION}

ARG USER_NAME=influx
ARG USER_ID=1000
ARG USER_HOME=/var/lib/influxdb

RUN adduser \
  --quiet \
  --home ${USER_HOME} \
  --uid ${USER_ID} \
  --disabled-password \
  --gecos "" \
  ${USER_NAME}

RUN chown -Rf ${USER_NAME} ${USER_HOME}

USER ${USER_NAME}
