# Keycloak

## Construire l'image keycloack

```bash
$ cd keycloak
$ make docker-push
```

## Installation

```bash
$ kubectl create ns keycloak
$ helm show values bitnami/keycloak > values-bitnami.yaml
$ helm install keycloack bitnami/keycloak -f values-bitnami.yaml -n keycloak
$ kubectl get all -n keycloak
$ kubectl edit svc keycloak-http -n keycloak
```


NAME: keycloack
LAST DEPLOYED: Mon Sep 25 16:28:20 2023
NAMESPACE: keycloak
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: keycloak
CHART VERSION: 16.1.5
APP VERSION: 22.0.3

** Please be patient while the chart is being deployed **

Keycloak can be accessed through the following DNS name from within your cluster:

    keycloack-keycloak.keycloak.svc.cluster.local (port 80)

To access Keycloak from outside the cluster execute the following commands:

1. Get the Keycloak URL and associate its hostname to your cluster external IP:

   export CLUSTER_IP=$(minikube ip) # On Minikube. Use: `kubectl cluster-info` on others K8s clusters
   echo "Keycloak URL: http://keycloak.local/"
   echo "$CLUSTER_IP  keycloak.local" | sudo tee -a /etc/hosts

2. Access Keycloak using the obtained URL.
3. Access the Administration Console using the following credentials:

  echo Username: admin
  echo Password: $(kubectl get secret --namespace keycloak keycloack-keycloak -o jsonpath="{.data.admin-password}" | base64 -d)