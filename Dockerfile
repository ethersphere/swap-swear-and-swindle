FROM alpine:latest as build

ARG GETH_VERSION="1.9.20-979fc968"
ARG SOLIDITY_VERSION="0.6.12"

RUN apk update
RUN apk add jq

WORKDIR sw3
COPY ./scripts/install-deps.sh /sw3/scripts/
RUN ./scripts/install-deps.sh

COPY . /sw3/

CMD ./scripts/abigen.sh ERC20SimpleSwap SimpleSwapFactory Migrations
