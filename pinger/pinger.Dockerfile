FROM golang

WORKDIR /app

RUN rm /etc/localtime

RUN ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

USER root

CMD go mod tidy; go run .