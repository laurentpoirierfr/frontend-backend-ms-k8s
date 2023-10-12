# Loki

```bash
$ export NAMESPACE=loki
$ kubectl get secret --namespace ${NAMESPACE} loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
$ kubectl port-forward --namespace ${NAMESPACE} service/loki-grafana 3000:80
```

Navigate to **http://localhost:3000** and login with admin and the password output above. 

login : admin