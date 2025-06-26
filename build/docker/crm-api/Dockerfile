FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /apartomat

COPY ../../../bin/crm-api .
COPY ../../../migration migration/
COPY ../../../apartomat.key .

ENV PATH="/apartomat:${PATH}"

EXPOSE 8080

ENTRYPOINT ["crm-api", "run"]
