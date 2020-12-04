#!/bin/bash

influx setup \
  --bucket MYBUCKET \
  --username influx \
  --password password \
  --org MYORG \
  --force
