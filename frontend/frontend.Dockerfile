FROM golang

WORKDIR /app

RUN rm /etc/localtime

RUN ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

EXPOSE 8080

CMD go mod tidy; go run .