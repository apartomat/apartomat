FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /apartomat

COPY ../../../bin/project-page-api .
ENV PATH="/apartomat:${PATH}"

EXPOSE 8080

ENTRYPOINT ["project-page-api"]
