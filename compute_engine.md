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
Review:
    
    Static IP Address: Get a const IP address for VM intances
    Instance template: preconfig templates simplifying the creation of VMs
