## how to handle tls for `elasticsearch`?

* run docker containers from elasticsearch and kibana images
* in your host, get certificate of the elasticsearch instance and trust it:
```sh
openssl s_client -showcerts -connect localhost:9200 < /dev/null | openssl x509 -outform PEM > mycertfile.pem
openssl x509 -in mycertfile.pem -out mycertfile.crt -outform PEM

sudo cp mycertfile.pem /usr/local/share/ca-certificates/
sudo cp mycertfile.crt /usr/local/share/ca-certificates/

sudo update-ca-certificates
```
* connect to elasticsearch container shell, then get apikey for kibana and password for `elastic` user:
```sh
elasticsearch-create-enrollment-token -s kibana
elasticsearch-reset-password -u elastic
```