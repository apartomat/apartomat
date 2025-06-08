FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /apartomat

COPY ../../../bin/apartomat-lunux-amd64 .

COPY ../../../migration migration/

COPY ../../../apartomat.key .

ENV PATH="/apartomat:${PATH}"

EXPOSE 80

ENTRYPOINT ["apartomat-lunux-amd64", "run"]
