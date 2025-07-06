run_nginx:
	docker run  --name $(name) \
  --add-host=host.docker.internal:host-gateway \
  -p 8080:8080 \
  -v /Users/abhi1060/Dev/System\ Design/SSLPinning/proxy/nginx.conf:/etc/nginx/nginx.conf:ro \
  -v /Users/abhi1060/Dev/System\ Design/SSLPinning/proxy/ssl:/etc/nginx/ssl \
  -d nginx 


rm:
	docker stop $(name)
	docker rm $(name)
