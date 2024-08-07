I. Set up startup script

    https://cloud.google.com/compute/docs/instances/startup-scripts/linux#console
        
    Create Instance -> Advanced options -> Management -> Startup script
    import the script

    #! /bin/bash
    apt-get update
    apt-get -y install apache2
    echo hello. im $(hostname), myip $(hostname -i) > /var/www/html/index.html
    systemctl restart apache2
    # note you must wait the script run completely

three solution to quick init instance

    - startup script
    - create instace template
    - custom image

shielded VM:

    - Shielded VM offers verifiable integrity of your Compute Engine VM instances
    - U can be confident your instances haven't been malware or Rootkit

Google CLoud Active Directory Sync:

    - https://tools.google.com/dlpage/dirsync/thankyou.html
    - sync active directory to google

Preemtive VM:

    - preemptible /priˈɛmptɪv/ VMs can provide up to 80% discount over normal VMs if the workloads are fault-tolerant
    - howrever, CE migh be stop (preemt) there instances if it need to reclaim the compute capacity for allocates another VM

deployment-manager:

    - https://cloud.google.com/deployment-manager/docs
    - --preview option in the same project, and observe the state of interdependent resources.
   
tasks:
    
    - msg ordering

ssh to VM or CE:

    - metadate -> tab ssh key -> copy and paste value
    - OS login and setup manual

connect to VM from local (https://cloud.google.com/compute/docs/connect/add-ssh-keys#console):
    
    - gen ssh key at local $ssh-keygen -t rsa (setup public and pivate like git)
    - edit running VM, Security -> SSH key -> add one
    - get public key from local (cat xxx/rsa_key.pub)
    - paste to text box -> save
    - ssh like normal flow (ssh -i xxx/rsa -l $YOURNAME $YOUR_IP)
