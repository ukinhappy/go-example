openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=127.0.0.1" -days 5000 -out ca.crt
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=127.0.0.1" -out server.csr
echo subjectAltName = IP:127.0.0.1 > extfile.cnf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -extfile extfile.cnf -CAcreateserial  -out server.crt -days 5000