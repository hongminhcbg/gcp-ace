# gcloud command line interface
  - most GCP svc can be managed by gcloud interface
# some gcp svc have specific tool

    - cloud storage: gsutil
    - cloud big query: bq
    - cloud bigtable: cbt

# command structure:

  gcloud GROUP SUBGROUP ACTION ...
  example command to create a VM
  
  gcloud compute instances create my-instance \
  --boot-disk-size=50GB \
  --zone=asia-southeast2-c \
  --machine-type=n1-standard-2 \
  --boot-disk-size=50GB \
  --image ubuntu-1804-bionic-v20190204 \
  --image-project ubuntu-os-cloud
  
  // add default zone and region
  $ gcloud compute project-info add-metadata --metadata=google-compute-default-region=asia-southeast1
  $ gcloud compute project-info add-metadata --metadata=google-compute-default-region=asia-southeast1,google-compute-default-zone=asia-southeast1-a
  // default zone with cmd
  $ gcloud config set compute/zone aaa
  
  // list image
  $ gcloud compute images list
  $ gcloud auth active-svc-account ... // use SA local, because gcloud sdk is not a standard method to auth

# Manage Instance Group (MIG):

    - managed as one unit

MIG scenarios:

    - want MIG managed app to survive zone failures => create multi zone MIG
    - want to create VMs of different config => create un-managed intance group 

Pubsub:
    
    - exactly_once_delivery:
        duplicate_msg_issue:
        Publisher:
                Publisher might have a network failure resulting in not receiving the ack from Cloud Pub/Sub. This would cause the publisher to republish the message.
                Publisher application might crash before receiving acknowledgement on an already published message.
        Subscriber
            Subscriber might also experience network failure post-processing the message, resulting in not acknowledging the message. This would result in redelivery of the message when the message has already been processed.
            Subscriber application might crash after processing the message, but before acknowledging the message. This would again cause redelivery of an already processed message.
        features:
            - No redelivery occurs after the message is successfully acknowledged
            - No redelivery occurs after the msg is outstanding (ack deadline timeout)
# Access GCP

  - web console
  - cli
  - go, python, nodejs sdk
  - mobile
  - cloud shell quota: 50 hours / w
