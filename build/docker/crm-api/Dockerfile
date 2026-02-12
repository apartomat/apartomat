FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /apartomat

COPY bin/crm-api bin/migration .
COPY apartomat.key .
COPY build/docker/crm-api/entrypoint.sh /apartomat/entrypoint.sh

RUN chmod +x /apartomat/entrypoint.sh

ENV PATH="/apartomat:${PATH}"

EXPOSE 8080

ENTRYPOINT ["/apartomat/entrypoint.sh"]
CMD ["run"]
