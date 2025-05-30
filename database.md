I. Storage

    oganized and persistent storage

# Category:

    - big query
    - cloud sql
    - cloud spaner 
    - datastore
    - filestore

# Database type:

    Relational db OLTP -> cloud SQL, cloud spanner 
    relational OLAP db -> columnar storage, for admin and DA
    nosql db -> filestore (serverless transaction docs, mobile + web), bigtable (large db), not severless
    inmem database (cache) -> memstore, access in microsecond
    OLTP: Online Transactional processing
    OLAP: Online Analytics processing -> reporting
# OLAP

    - columnar storage

# OLTP

    - row storage, each row is stored together

# Scenarios:

    start up quick -> datastore, filestore, filestore = datastore++ (multi devices)
    nosql, < 10Gi -> cloud datastore
    process 1Mi TPS -> cloud spanner
    process 1k TPS -> cloud SQL
    store ~1 peta bytes ->  big table nosql
    huge volume of IoT devices -> bigtable, nosql
    stream huge time series data (logs) -> bigtable, nosql

# Cloud SQL

    - MySQL and SQL server
    - only support read replicas

## Cloud Spanner

    - growth with minimum config
    - global distributed relational db
    - management complexity to a minimum

# NoSQL:

    - trafe-off "strong consistency and SQL features" to "scalability and high performance"
    - horizontal scale to Petabytes, hundreds nodes and milions TPS 

## Datastore:

    - Document base database
    - auto scale and auto partitions
    - recommended upto a few TBs
    - If bigger volume bigtable
    - support transaction + index but not support join + sum + count
    - flexible schema

## Filestore: (Datastore++)
    
    - support multi devices 
    - offline mode or data sync across multiple devices - mobile, iot
    - provide client-side lib (web, ios, android)
    - collections <=> tables
    - cursors = query limit + offset, optimize RAM when query

## Bigtable:

    - petabytes and more 
    - not serverless
    - https://cloud.google.com/bigquery/docs/running-queries#batch => save time

## Bigquery - datawarehouse

    - SQL-like query
    - realtime + serverless
    - export: cloud storage or data studio (visualization), formats csv, json, 
    - import: various src
    - query from external src, cloud storage, cloud sql, bigtable
    - remember bq can be expensive as you are running them on large data sets, est price with --dry-run flag
    - price cost based on 1MB scaned data
    - dry-run (https://cloud.google.com/bigquery/docs/running-queries#dry-run)
    - can create separate viewer on same datasheet

## Inmem db (redis + memory store)

    - access in nano seconds


## Cloud Storage

    - suitable for both structure and unstructure data



Datastudio 

    - build dashboard

Notes:

    - growth with minimum config => Cloud Spanner

Vertical scaling:
    
    - 8 vcpu - RAM 16Gi -> 12 vcpu - RAM 24Gi

Hoziontal scaling:

    - 1 node -> 3 nodes and same config

IOPS:

    - the number of read/write operation per second

## Guides:

### Compute Engine connect to CloudSQL

docs: 

https://cloud.google.com/sql/docs/mysql/connect-instance-compute-engine
https://cloud.google.com/sql/docs/mysql/configure-ip#add
https://cloud.google.com/sql/docs/mysql/connect-admin-ip

#### With public IP

```yaml

    step 1: Create cloud sql
    step 2: Create VM
    step 3: White list IP, cloudSQL -> Connections -> Public IP -> Add Network -> Add public IP of VM to whitelist
    step 4: Test connection from VM, $mysql -u mint -h 34.124.242.225 # 34.124.242.225 is public IP of cloudSQL  
```

#### With private IP

    https://cloud.google.com/sql/docs/mysql/private-ip

### Appengine connect to CloudSQL
#### With connection name
```
    step1: Deploy app engine, example code (https://github.com/hongminhcbg/gcp-ace/blob/main/app-engines/app-login/main.go)
    step2: Grant role 'Secret Manager Secret Accessor' to default app engine SA 
    step3: Setup secret value, important key 'sql_cname' = connection name
    step4: Enable 'Cloud SQL Admin API'
```

### Cloud SQL Auth Proxy (https://cloud.google.com/sql/docs/mysql/connect-instance-auth-proxy)

```yaml
    - provides secure access to your instances with out a need for auth-networks or configuring SSL 
    - secure connection: auto encrypt traffic 
    - easy: The Cloud SQL Auth Proxy uses IAM permissions to control who and what can connect to your Cloud SQL instances. Thus, the Cloud SQL Auth Proxy handles authentication with Cloud SQL, removing the need to provide static IP addresses.
    - auth: auto refresh access token
    - need install proxy client on the client machine
    - [APP -> auth proxy client] -> (internet) -> [auth proxy server -> SQL instances]
    - On GKE, SQL Auth Proxy run as sidecar container
```

```sh
## Install sql proxy client
wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
chmod +x ./cloud_sql_proxy
./cloud_sql_proxy -instances=glowing-furnace-304608:us-central1:my-first-cloud-sql-instance=tcp:3306
mysql -u root -p --host 127.0.0.1
## Run sql proxy client

```
## Cloud Spanner

    - cloudSQL limitation 30TB
    - distributed and scalable for GCP
    - Hoziontal scalebility (appending more node)
    - use when data volume > 2TB
    - expensive than cloudSQL
    - read/write across region (no need read replicas)
    - cloud spanner = cloudSQL + hoziontal scale (no limit RAM and CPU)

## Decrease spanner unit

  In the latter case, you might try reducing the compute capacity by progressively smaller amounts until you find the minimum capacity that Spanner needs to manage all of the instance's splits. If the instance no longer requires so many splits due to a change in usage patterns, Spanner might eventually merge some splits together and allow you to try reducing the instance's compute capacity further after a week or two.
  When removing compute capacity, monitor your CPU utilization and request latencies in Cloud Monitoring to ensure CPU utilization stays under 65% for regional instances and 45% for each region in multi-region instances. You might experience a temporary increase in request latencies while removing compute capacity.
