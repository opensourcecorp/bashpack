FROM debian:12

RUN apt-get update && apt-get install -y \
      bats \
      git \
      shellcheck

RUN useradd --create-home nonroot

RUN mkdir /tmp/bashpack
WORKDIR /tmp/bashpack

COPY ./Makefile ./manifest.bashpack /tmp/bashpack/
COPY ./src /tmp/bashpack/src

CMD ["bats", "/tmp/bashpack/src/test.bats"]
