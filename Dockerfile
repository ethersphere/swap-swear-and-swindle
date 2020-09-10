FROM debian:jessie-slim

# Geth and solidity version to use.
ARG GETH_VERSION="1.9.20-979fc968"
ARG SOLIDITY_VERSION="0.6.12"

# Tell apt-get we're not giving interactive feedback.
ARG DEBIAN_FRONTEND=noninteractive

# Install Debian dependencies with apt-get.
RUN apt-get update && \ 
    apt-get -y upgrade && \
    apt-get -y install --no-install-recommends \
        ca-certificates \
        wget \
        jq \
        gettext-base && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Instal geth and solc by pulling official builds.
WORKDIR sw3
COPY ./scripts/install-deps.sh /sw3/scripts/
RUN ./scripts/install-deps.sh

COPY . /sw3/

CMD ["/sw3/scripts/abigen.sh", "ERC20SimpleSwap", "SimpleSwapFactory", "Migrations"]
