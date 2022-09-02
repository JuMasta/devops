
FROM nginx:1.23-alpine

RUN rm /etc/nginx/conf.d/default.conf
COPY ./awesome.conf /etc/nginx/conf.d/
COPY ./nginx.conf ./proxy.conf /etc/nginx/