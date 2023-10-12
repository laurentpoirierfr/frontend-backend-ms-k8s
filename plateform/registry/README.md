# Registry

``` bash
$ kubectl port-forward --namespace kube-system service/registry 5000:80
```

``` bash
$ sudo vi /etc/docker/daemon.json
``` 

``` json
{
    "insecure-registries" : [ "registry.local:5000" ]
}
``` 

``` bash
$ sudo service docker restart 
```

```bash
$ mkdir -p certs

$ openssl req \
  -newkey rsa:4096 -nodes -sha256 -keyout certs/registry.local.key \
  -addext "subjectAltName = DNS:registry.local" \
  -x509 -days 365 -out certs/registry.local.crt

$ sudo cp certs/*.crt /usr/local/share/ca-certificates

$ sudo update-ca-certificates
```


### References

* https://docs.docker.com/registry/insecure/#deploy-a-plain-http-registry
* https://github.com/kubernetes/ingress-nginx/issues/4644