#!/usr/bin/env bash

docker pull cckwes/shoplist:latest
docker stop shoplist-container || true
docker rm shoplist-container || true
docker run --name shoplist-container \
    -e JWT_AUDIENCE=$JWT_AUDIENCE \
    -e JWT_ISSUER=$JWT_ISSUER \
    -e APP_ENV=production \
    -e SQLITE_FILE=/db/data.sqlite \
    -v /db:/db \
    -d cckwes/shoplist:latest
