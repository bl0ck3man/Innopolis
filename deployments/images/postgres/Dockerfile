FROM postgres:12.1-alpine

ARG uid
ARG gid

RUN apk add --no-cache shadow \
    && usermod -u $uid postgres && groupmod -g $gid postgres \
    && apk del shadow

#COPY pg_hba.conf /var/lib/postgresql/data/
COPY pg_hba.conf /data/database
COPY *.sql /docker-entrypoint-initdb.d/

RUN chmod a+r /docker-entrypoint-initdb.d/*