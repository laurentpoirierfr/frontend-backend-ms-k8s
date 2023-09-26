# README

## Helm template debug

```bash
$ helm template  ./helm/ms-app -f backend-values.yaml 
```

## Frontend installation

```bash
$ export NAMESPACE=homezone-dev
$ kubectl create namespace ${NAMESPACE}

$ helm install backend ./helm/ms-app -f backend-values.yaml --namespace ${NAMESPACE} --replace
$ helm install frontend ./helm/ms-app -f frontend-values.yaml --namespace ${NAMESPACE} --replace

# Uninstall backend
$ helm uninstall backend ./helm/ms-app --namespace ${NAMESPACE}
# Uninstall frontend
$ helm uninstall frontend ./helm/ms-app --namespace ${NAMESPACE}
```




