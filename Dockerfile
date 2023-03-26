FROM ubuntu:latest

# Temporary environment variables so that apt and the php installation does not complain.
# https://github.com/phusion/baseimage-docker/issues/319
ENV DEBIAN_FRONTEND noninteractive
ENV TZ="Europe/Madrid"
ENV DEBCONF_NOWARNINGS="yes"

# Apt: Update the repository sources list, install apt-utils (complains otherwise) and upgrade packages
# 'apt-get' has a stable CLI interface. Do not use 'apt' as it does not.
RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get upgrade -y
RUN apt-get install -y systemd \
                       w3m \
                       nano \
                       vim \
                       curl \
                       wget \
                       git

WORKDIR /go-installer
RUN wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"

WORKDIR /chorizo-with-mustard-api
COPY ./api .
RUN go build -o /init-server
CMD ["/init-server"]
