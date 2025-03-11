## Cloud IAP

# OAuth

  - authorization framework
  - allow users to grant access to server resources without sharing credentials 
  - [Google oauth playground](https://developers.google.com/oauthplayground/)
  - only authorization, don't care about how authentication implement
# OpenID Connect
  - extend of oauth2.0

# Identity aware proxy

```
- single point of control for managing user access to web or cloud resource
- manage ssh and http(s) based on resources
- easiest way to implement authen and author for your cloud base (app engine, cloud run, GKE)
```

# Usage

```
- ssh from private IP
- secure app engine http resource
- firewall rule: Allow ssh to VM just from browser
```
# Setup, demo with app engine

### Step 1

  Deploy simple app engine and test

  example
  ```sh
  curl https://compute-engine-452904.as.r.appspot.com
  # hello world
  ```

### Step 2

  - Enable IAP (Home -> IAP -> Enable IAP)
  - Config consent screen
  - 
### Step 3 (enable from app engine)

  - App engine -> settings -> turn on IAP

