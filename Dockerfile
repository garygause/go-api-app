FROM golang:latest as builder
WORKDIR /app
COPY . /app/

FROM ubuntu:latest

RUN apt update -y && apt dist-upgrade -y

WORKDIR /app