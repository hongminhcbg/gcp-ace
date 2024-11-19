Virtual Cloud Private

    - network traffic within a VPC is isolated from all others VPC
    - best pactice create all resources within a VPC
    - VPC is a global resource, can contain resource from any region
    - VPC network peering: resource in diff subnet can talk to each others 
    - resource in private subnet can't access from internet
    - resource in public subnet can talk to private subnet
    - CIDR block
    - firewall rule (0-65535) [0: highest]
    - VM in default netword can't ping to internal IP of VM in VPC. But can ping to public IP

VPC subnet:

    1 subnet belong to 1 region

Allow pings both 2 VM

    https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/tutorials/tutorial-troubleshooting-workflow#create-firewall-rule-3

## Hybrid cloud, connect GCP from on-premise network

Cloud VPN:

    - connect on-premise to GCP network
    - implement IPsec VPN tunnel

Cloud Interconnect:

    - high speed physical connection between on premise and VPC networks

Notes:

    - best practices whitelist by SA
    - enable private cloud access in VPC => access Google cloud API and svc

Peer network u must peer A->B and B->A
![Alt text](./imgs/peernetwork.jpeg?raw=true "Title")
