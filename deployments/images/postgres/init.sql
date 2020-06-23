REVOKE ALL ON DATABASE demo FROM public;
REVOKE ALL ON SCHEMA public FROM public;

CREATE USER web with nosuperuser nocreatedb nocreaterole noinherit noreplication;

GRANT CONNECT ON DATABASE demo TO web;
GRANT USAGE ON SCHEMA public TO web;
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO web;

ALTER USER web PASSWORD 'postgres';