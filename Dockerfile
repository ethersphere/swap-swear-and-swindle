FROM node:10.16.0-stretch as builder

WORKDIR /sw3
ADD . /sw3

RUN npm install
