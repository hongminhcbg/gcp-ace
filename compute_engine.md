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

container images to create VM:

  - public image
  - custom image
  - container image
  - not support hybrid image
  - list active VM daily => separate two config, switch config, write script run 'cloud compute list'
  - set metadata 'enable-os-login=true', grant user role/compute.osLogin => easiest way to grant user without ssh key

compute options:

  -  

# Backup

# Attach extra disk

Local disk:
  
  - the disk directly attach with your machine
  - no network delay

Additional storage:

  - network storage

Create disk:
  
  - Compute engine -> storage -> disk -> create new
  - https://cloud.google.com/compute/docs/disks/?_gl=1*1r6byv*_ga*MTAwODAwMzE4OC4xNzMxNzU4MTI1*_ga_WH2QY8WWF5*MTczMjI2NDAwNy40LjEuMTczMjI2NDg0My41OC4wLjA.#introduction

Cmd format and mount disk

```sh
  sudo mkfs.ext4 -m 0 -F -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb // or sudo mkfs.ext4 /dev/sdb
  sudo mkdir -p /mnt/disk-1
  sudo mount -o discard,defaults /dev/sdb /mnt/disk-1
  sudo chmod a+w /mnt/disk-1/
```

Local disk

  - only add when create VM

## Availability Policy (standard or spot)

Onhost maintainance

  - migrate the VM from one rack to another rack
  - if chose "Migrate the VM instance" => migrate without down time
  - if type is spot use can chose 'Delete' or 'Stop'

# Manage Instance Group (MIG):

    - managed as one unit

MIG scenarios:

    - want MIG managed app to survive zone failures => create multi zone MIG
    - want to create VMs of different config => create un-managed intance group 

# Unmaneage Instance Group:

  - Group of VM diff configation
  - Create VM1 and VM2 with diff config
  - Create UMIG and sellect 2VM



