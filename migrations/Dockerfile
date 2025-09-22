FROM gomicro/goose

COPY ./migrations/*.sql /migrations/
COPY ./migrations/entrypoint.sh /migrations/

ENTRYPOINT ["/migrations/entrypoint.sh"]