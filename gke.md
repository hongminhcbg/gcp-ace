https://cloud.google.com/kubernetes-engine/docs/concepts/node-pools

features:

    - exe untrusted 3rd party 
    - efficient + auto scaling

GKE apha mode: max 30 days

Notes:

    - GKE loadbalancer, Add an annotation: cloud.google.com/load-balancer-type: Internal => auto create loadbalancer
    - GKE auto pilot = optimize cost and ops

# Cloud Connector:

    - k8s can manage cloud resource

# [Workload Identity gke](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)

## Usecase

    - k8s pods can read data from pub/sub or cloud storage, but no need to create single sa per pods. Because quota limit is 100 sa per project

### Enable for the cluster

```
gcloud container clusters update autopilot-cluster-2 --location=asia-southeast1 --workload-pool=compute-engine-452904.svc.id.goog
```

### Create k8s sa

    kubectl create serviceaccount gsviewer --namespace my-workload

### Binding sa

    gcloud projects add-iam-policy-binding projects/compute-engine-452904 \
        --role=roles/container.clusterViewer \
        --member=principal://iam.googleapis.com/projects/1056296113540/locations/global/workloadIdentityPools/compute-engine-452904.svc.id.goog/subject/ns/my-workload/sa/gsviewer \
        --condition=None

    gcloud projects add-iam-policy-binding projects/compute-engine-452904 \
        --role=roles/storage.objectViewer \
        --member=principal://iam.googleapis.com/projects/1056296113540/locations/global/workloadIdentityPools/compute-engine-452904.svc.id.goog/subject/ns/my-workload/sa/gsviewer \
        --condition=None

### Create a pod with your sa

```yaml
apiVersion: v1
kind: Pod
metadata:
  namespace: my-workload
  name: cloud-sdk-pod
  labels:
    app: cloud-sdk
  annotations:
    proxy.istio.io/config: '{ "holdApplicationUntilProxyStarts": true }'    
spec:
  serviceAccountName: gsviewer
  containers:
    - name: cloud-sdk-container
      image: google/cloud-sdk:489.0.0-stable
      command: ["sleep", "7200"]  # Keeps the container running for interaction
      resources:
        limits:
          cpu: "500m"
          memory: "512Mi"
        requests:
          cpu: "250m"
          memory: "256Mi"
```

### Exec into the pods and get result

```bash
k exec -it -n my-workload cloud-sdk-pod -- /bin/bash

gsutil ls gs://xxxyyymint
# results: 
# gs://xxxyyymint/.DS_Store
```

## Engress

    - a set of rules for routing external HTTP(s) to svcs