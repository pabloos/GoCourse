
test:
	go test -v -race ./server

certs:
	mkdir -p server/cert cert/
	openssl genrsa -out server/cert/private.key 4096
	openssl req -new -x509 -sha256 -days 1825 -key server/cert/private.key -out cert/public.crt