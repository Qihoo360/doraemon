# requiring Docker 17.05 or higher on the daemon and client
# see https://docs.docker.com/develop/develop-images/multistage-build/
# BUILD COMMAND :
# docker --build-arg RELEASE_VERSION=v1.0.0 -t infra/wayne:v1.0.0 .

# build server
FROM openresty/openresty:1.15.8.1-1-centos

COPY web/dist/ /usr/local/openresty/nginx/html/

RUN sed -i '/server_name/a\    location /api/v1/\n     {\n     proxy_redirect off; \n     proxy_set_header Host $host; \n     proxy_set_header X-Real-IP $remote_addr; \n     proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; \n     proxy_pass http://localhost:8080;\n    }' /etc/nginx/conf.d/default.conf

RUN sed -i '/index/a\        try_files $uri $uri/ /index.html;' /etc/nginx/conf.d/default.conf 

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 

CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]
