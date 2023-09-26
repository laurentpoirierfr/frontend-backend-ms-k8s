# Mongodb


## Installation

```bash
$ export NAMESPACE=mongodb
$ kubectl create ns ${NAMESPACE}
$ helm show values bitnami/mongodb > values-bitnami.yaml
$ helm install keycloack bitnami/mongodb -f values-bitnami.yaml -n ${NAMESPACE}
```

```bash
NAME: keycloack
LAST DEPLOYED: Tue Sep 26 12:11:43 2023
NAMESPACE: mongodb
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: mongodb
CHART VERSION: 14.0.0
APP VERSION: 7.0.1

** Please be patient while the chart is being deployed **

MongoDB&reg; can be accessed on the following DNS name(s) and ports from within your cluster:

    keycloack-mongodb.mongodb.svc.mongodb.local

To get the root password run:

    export MONGODB_ROOT_PASSWORD=$(kubectl get secret --namespace mongodb keycloack-mongodb -o jsonpath="{.data.mongodb-root-password}" | base64 -d)

To get the password for "mongodb" run:

    export MONGODB_PASSWORD=$(kubectl get secret --namespace mongodb keycloack-mongodb -o jsonpath="{.data.mongodb-passwords}" | base64 -d | awk -F',' '{print $1}')

To connect to your database, create a MongoDB&reg; client container:

    kubectl run --namespace mongodb keycloack-mongodb-client --rm --tty -i --restart='Never' --env="MONGODB_ROOT_PASSWORD=$MONGODB_ROOT_PASSWORD" --image docker.io/bitnami/mongodb:7.0.1-debian-11-r0 --command -- bash

Then, run the following command:
    mongosh admin --host "keycloack-mongodb" --authenticationDatabase admin -u $MONGODB_ROOT_USER -p $MONGODB_ROOT_PASSWORD

To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace mongodb svc/keycloack-mongodb 27017:27017 &
    mongosh --host 127.0.0.1 --authenticationDatabase admin -p $MONGODB_ROOT_PASSWORD
```  