FROM alpine:3.15

WORKDIR /apartomat

COPY bin/apartomat-lunux-amd64 .
COPY bin/migration-lunux-amd64 .

COPY migration migration/

COPY shoppinglist.key .
COPY shoppinglist.key.pub .

ENV PATH="/apartomat:${PATH}"

EXPOSE 80

ENTRYPOINT ["apartomat-lunux-amd64", "run"]