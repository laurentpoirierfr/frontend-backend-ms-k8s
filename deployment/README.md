# README

## Helm template 

```bash
$ helm template  ./helm/ms-app -f backend-values.yaml 
```

## Frontend installation

```bash
$ export NAMESPACE=homezone-dev
$ kubectl create namespace ${NAMESPACE}
$ helm install frontend ./helm/ms-app -f backend-values.yaml --namespace ${NAMESPACE} --replace

# Uninstall backend
$ helm uninstall frontend ./helm/ms-app --namespace ${NAMESPACE}
```




