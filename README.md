GCP ACE docs

  - GCE: IAAS
    - provide:
      - raw compute (compute engine)
      - raw storage
      - network capabilities
      - pay for resources they allocate a head of time
  - GKE: CAAS
  - cloud func: FAAS
  - cloud run: CAAS
  - App engine: PAAS, serverless
    - bind code to lib, forcus on application logic
    - pay for the resources they actually use

# Serverless

  - Cloud Functions: Function as a Service
  - Cloud Run: Serverless containers
  - App Engine: Flexible and Standard
  - BigQuery: Serverless Data Warehouse
  - Cloud Storage: Serverless Storage
  - Pub/Sub: Serverless Messaging”

# TODOS

  - Cloud endpoints
  - count API gateway
  - cloud profiler
  - You have a GKE cluster with Istio service mesh deployed. One of your services, Service A, cannot invoke another service, Service B, over HTTPS. Both services have Istio sidecars and you have confirmed that Service A's requests to Service B are leaving the sidecar. What could be the cause of this issue?
  - Private Google Access
  - Handling sessions with Firestore (https://cloud.google.com/go/getting-started/session-handling-with-firestore)
  - cloud enpoints -> cloud funcs
  - cloud funcs -> (private gg access) -> cloud SQL 
  - serverless VPC access connector
  - IAP and grant at resource level, read and verify token
  - cloud shell -> terraform -> resource
  - custom SA, compute engine -> trigger cloud func
