# requiring Docker 17.05 or higher on the daemon and client
# see https://docs.docker.com/develop/develop-images/multistage-build/
# BUILD COMMAND :
# docker --build-arg RELEASE_VERSION=v1.0.0 -t infra/wayne:v1.0.0 .

# build server
FROM openresty/openresty:1.15.8.1-1-centos

COPY ./dist/ /usr/local/openresty/nginx/html/

COPY ./default.conf /etc/nginx/conf.d/

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 

CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]
