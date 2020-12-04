DOCKER_CMD := "docker"
INFLUXDB_VERSION := v2.0.2
INFLUXDB_IMG_TAG := fesaille/influxdb:$(INFLUXDB_VERSION)

CONTAINER_NAME := influx2

USER_ID := 1000

all: run

.influxdb: Dockerfile
	$(DOCKER_CMD) build \
		-t $(INFLUXDB_IMG_TAG) \
		--build-arg VERSION=$(INFLUXDB_VERSION) \
		--build-arg USER_ID=$(USER_ID) \
	.
	@touch $@

run: .influxdb
	$(DOCKER_CMD) run -d --rm --name $(CONTAINER_NAME) \
		-p 8086:8086 \
		-v $(shell pwd)/init:/my \
		$(INFLUXDB_IMG_TAG)

config:
	$(DOCKER_CMD) exec $(CONTAINER_NAME) sh /my/influxdb.sh
