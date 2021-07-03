FROM caddy:2.3.0-alpine

RUN apk --no-cache add wget=1.20.3-r1 unzip=6.0-r9 && rm -rf /var/cache/apk/*
WORKDIR /webapp
RUN wget --progress=dot:giga https://github.com/RaboliotLeGris/nienna/releases/download/latest/webapp.zip && \
    unzip webapp.zip && rm webapp.zip && \
    cp -R -- * /usr/share/caddy/