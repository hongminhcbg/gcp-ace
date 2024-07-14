Overview
    
    User -> cloud LB -> Instances, cloud func

layer 7: http, https, smtp
layer 4: tcp, udp, tls
layer 3: ip

TCP -> reliability > performance
UDP -> performance > reliability (video streaming)
TLS -> secure TCP

Serverless:
    
    - aws lamda vs azure cloud func vs google function
    - pay for request, not pay for server
    - if no req, no cost

App engine:

    - PaaS
    - serverless

Important Feature
    
    - health check, route req to healthy instance
    - auto slacing
    - global LB with single IP

3 type of LB (HTTPS, TCP, UDP)

How to choose LB type

    If internet 
        If HTTP traffic
            If regional => Regional External LB
            else => Global or classic LB
        If TCP/SSL => TCP/SSL proxy
    else
        If HTTP => internal LB
        else => internal TCP/UDP LB
            


