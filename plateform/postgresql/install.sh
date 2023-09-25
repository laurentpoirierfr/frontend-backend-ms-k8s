
export NAME_SPACE=postgresql

kubectl create namespace ${NAME_SPACE}
kubectl apply -n ${NAME_SPACE} -f postgresql.yml 

export ADMIN_PASSWORD=password
export USER_NAME=postgres
export USER_PASWWORD=password
export DATABASE=homezone

helm repo add bitnami https://charts.bitnami.com/bitnami

helm -n ${NAME_SPACE} install postgres bitnami/postgresql \
    --set primary.persistence.existingClaim=postgres-pvc \
    --set volumePermissions.enabled=true \
    --set global.postgresql.auth.postgresPassword=${ADMIN_PASSWORD} \
    --set global.postgresql.auth.username=${USER_NAME} \
    --set global.postgresql.auth.password=${USER_PASWWORD} \
    --set global.postgresql.auth.database=${DATABASE}


# kubectl port-forward --namespace postgresql svc/postgres-postgresql 5432:5432