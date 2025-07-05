run_nginx:
	docker run  --name $(name) \
  --add-host=host.docker.internal:host-gateway \
  -p 8080:8080 \
  -v /Users/abhi1060/Dev/System\ Design/SSLPinning/proxy/nginx.conf:/etc/nginx/nginx.conf:ro \
  -d nginx 