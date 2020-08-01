#!/usr/bin/env bash

docker pull cckwes/shoplist:latest
docker stop shoplist-container || true
docker rm shoplist-container || true
docker run --name shoplist-container \
    --env-file /home/app/shoplist.env \
    -v /db:/db \
    -p 3000:3000 \
    -d cckwes/shoplist:latest
