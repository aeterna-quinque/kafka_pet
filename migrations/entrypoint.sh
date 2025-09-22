#!/bin/bash

DBSTRING="host=$POSTGRES_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_NAME sslmode=$POSTGRES_SSL"

goose postgres "$DBSTRING" up