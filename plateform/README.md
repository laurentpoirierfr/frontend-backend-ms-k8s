# README

## Installation de Minikube

* https://kubernetes.io/fr/docs/tasks/tools/install-minikube/

```bash
$ minikube start --memory 8192 --cpus 4
$ minikube addons enable metrics-server
$ minikube addons enable ingress
$ minikube addons enable ingress-dns
$ minikube ip
```

* minikube ip ==> **192.168.49.2**

```bash
$ sudo vi /etc/resolv.conf
```

Ajout du domaine **local** en point sur le résolveur dns de minikube.

```bash
search .
nameserver 127.0.0.53
options edns0 trust-ad

search local
nameserver 192.168.49.2
timeout 5
```

ou bien dans **/etc/hosts**

```bash
192.168.49.2   keycloak.local
192.168.49.2   backend.local
192.168.49.2   frontend.local
```

## Installation du cert-manager

```bash
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.13.0/cert-manager.yaml
```


## Keycloack

* [Keycloak](./docs/keycloak.md)

## Postgresql

* [Postgresql](./docs/postgresql.md)


## Reference

* https://alxrodav.medium.com/keycloak-custom-login-theme-21be10ad3f4
* https://github.com/briantward/ansible-keycloak/tree/main
* https://www.talkingquickly.co.uk/webapp-authentication-keycloak-OAuth2-proxy-nginx-ingress-kubernetes


* https://ibrahimhkoyuncu.medium.com/kubernetes-ingress-external-authentication-with-oauth2-proxy-and-keycloak-9924a3b2d34a
* https://symbiosis.host/blog/oauth2-proxy