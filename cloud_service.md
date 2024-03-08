Flexibility (low-high)
FssA -> PaaS -> CaaS -> IaaS

Compute Engine (IaaS:
    - chose any region + zone + image

GKE (CaaS):
App Engine (PaaS):
    - serverless platform

Cloud Funtion (FaaS):
    - single-purpose functions

Compute Engine:

    Family: 
        - general purpose: E2, N2, N2D, N1
        - high memory: M2, M1
        - Compute optimized: C2 for gaming
    e2-standard-2:
        'e2' -> machine family
        'standard' -> wordload
        '2' -> Number of vCPUs

    internal_ip: only use in gcloud network
    external_ip: connect from external internet

    create Compute Engine with template:
    - startup script
    - instance template
    - custom image

    Custom machin type:
    - select number of core and memory
    
    GPU: add graphic handle, AI/ML processing


