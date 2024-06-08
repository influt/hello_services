# hello_services
Concurrent calls to microservices in Go
+ZAP test setup for gokit service

Service in gokit expects a non-empty json body and responds with latin alphabet characters

Commands for ZAP:

- docker pull zaproxy/zap-stable

Interception proxy: 
    - docker run -p 8090:8090 -i zaproxy/zap-stable zap.sh -daemon -port 8090 -config api.addrs.addr.name=.* -config api.addrs.addr.regex=true -config api.disablekey=true -host 0.0.0.0
    - curl 172.17.0.1:8080/chars -x http://localhost:8090 -kiv --data "{}"

Full scan: docker run -v .:/zap/wrk/:rw -p 8090:8090 -i zaproxy/zap-stable zap-full-scan.py -t http://172.17.0.1:8080/chars -r testreport.html

Note: do a proper network setup and replace zap and chars service IP address with localhost 