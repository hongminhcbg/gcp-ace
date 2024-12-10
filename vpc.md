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
    - a isolated private cloud and hosted within a public cloud

VPC subnet:

    - 1 subnet belong to 1 region
    - 1 VPC must has minimum 1 subnet

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

## Type of VPC

### default

```
  - create when CE api created
  - every project has default VPC
```

### Auto
```
- auto created
- fix subnetwork range created per region
- can expand from /20 to /16
```

### Custom
```
- no subnet auto created
- can create subnet manual per region
```

## VM to VM comunication

```
protocol:
  ssh: port 22
  http: port 80
  icmp: Internet control message protocol, ping-pong method
```

## Firewall

```
- priority: 0-65536, 0 is highest
- internet -> firewall -> VPC
- control the traffic from out side world
- IP ragne 0.0.0.0/0 is match all
```
Allow ping (ICMP) firewall
```json
{
  "allowed": [
    {
      "IPProtocol": "icmp"
    }
  ],
  "creationTimestamp": "2024-12-08T22:58:48.797-08:00",
  "description": "Allow ICMP from anywhere",
  "direction": "INGRESS",
  "disabled": false,
  "enableLogging": false,
  "id": "4033887356866735047",
  "kind": "compute#firewall",
  "logConfig": {
    "enable": false
  },
  "name": "default-allow-icmp",
  "network": "projects/sqlx-444206/global/networks/default",
  "priority": 65534,
  "selfLink": "projects/sqlx-444206/global/firewalls/default-allow-icmp",
  "sourceRanges": [
    "0.0.0.0/0"
  ]
}
```

Allow ssh firewall rule

```json
{
  "allowed": [
    {
      "IPProtocol": "tcp",
      "ports": [
        "22"
      ]
    }
  ],
  "creationTimestamp": "2024-12-09T23:18:38.099-08:00",
  "description": "allow ssh",
  "direction": "INGRESS",
  "disabled": false,
  "enableLogging": false,
  "id": "9064343930878449057",
  "kind": "compute#firewall",
  "logConfig": {
    "enable": false
  },
  "name": "allow-22",
  "network": "projects/sqlx-444206/global/networks/dbvpc",
  "priority": 999,
  "selfLink": "projects/sqlx-444206/global/firewalls/allow-22",
  "sourceRanges": [
    "0.0.0.0/0"
  ]
}
```
