FROM debian:buster-slim@sha256:f6ed7ce6e3264649e1d4f40585247c50e32faaf268984c5c5cbf0e67cf7f0ec7

RUN apt update && \
    apt-get install -y git build-essential ffmpeg libpcre3 libpcre3-dev libssl-dev zlib1g-dev wget tar make

WORKDIR /tmp/nginx_tmp
RUN git clone https://github.com/sergey-dryabzhinsky/nginx-rtmp-module.git
RUN wget http://nginx.org/download/nginx-1.21.4.tar.gz && \
    tar -xf nginx-1.21.4.tar.gz

WORKDIR /tmp/nginx_tmp/nginx-1.21.4
RUN ./configure --prefix=/usr/local/nginx --with-http_ssl_module --add-module=../nginx-rtmp-module && \
    make && \
    make install

COPY nginx.conf /usr/local/nginx/conf/nginx.conf

CMD ["/usr/local/nginx/sbin/nginx", "-g", "daemon off;"]