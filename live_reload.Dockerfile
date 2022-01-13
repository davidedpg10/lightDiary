# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.17.0-stretch

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

EXPOSE 8080

CMD air