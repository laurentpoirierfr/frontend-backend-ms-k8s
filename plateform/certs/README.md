# README

## Création d'un certificat 10 ans expiration

```bash
# non-interactive and 10 years expiration
$ openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 \
  -nodes -keyout homezone.dev.key -out homezone.dev.cert -extensions san -config \
  <(echo "[req]"; 
    echo distinguished_name=req; 
    echo "[san]"; 
    echo subjectAltName=DNS:homezone.dev,DNS:*.homezone.dev,IP:10.0.0.1
    ) \
  -subj "/CN=homezone.dev"
```

``` bash
$ kubectl -n kube-system create secret tls mkcert --key homezone.dev.key --cert homezone.dev.cert
$ minikube addons configure ingress
$ minikube addons disable ingress
$ minikube addons enable ingress
$ kubectl -n ingress-nginx get deployment ingress-nginx-controller -o yaml | grep "kube-system"
```



















<!-- ## Copy du certificat pour minikube

```bash
mkdir -p $HOME/.minikube/certs
cp *.cert $HOME/.minikube/certs
```

## Démarrage minikube

```bash
minikube start --embed-certs
``` -->