# README


```bash
kubectl apply -f backend.yml
kubectl apply -f frontend.yml
minikube service backend-app-service --url -n homezone-dev
minikube service frontend-app-service --url -n homezone-dev
```
