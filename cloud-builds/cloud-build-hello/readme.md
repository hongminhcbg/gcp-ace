$ gcloud artifacts repositories create quickstart-docker-repo2 --repository-format=docker \
    --location=asia-east1 --description="Docker repository"

$ gcloud builds submit --region=asia-east1 --tag asia-east1-docker.pkg.dev/disco-plane-418206/quickstart-docker-repo2/quickstart-image:v1.0.1
