FROM edoburu/pgbouncer:1.11.0

ARG uid
ARG gid

USER root

RUN apk add --no-cache shadow \
    && usermod -u $uid postgres && groupmod -g $gid postgres \
    && apk del shadow

RUN chown postgres:postgres -R /etc/pgbouncer

USER postgres