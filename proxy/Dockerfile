FROM nginx:latest

COPY nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /etc/nginx/ssl

COPY ssl /etc/nginx/ssl

RUN nginx -t

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]

