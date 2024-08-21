I. Storage

    oganized and persistent storage

Category:

    - big query
    - cloud sql
    - cloud spaner 
    - datastore
    - filestore

Database type:

    Relational db OLTP -> cloud SQL, cloud spanner 
    relational OLAP db -> columnar storage, for admin and DA
    nosql db -> filestore (serverless transaction docs, mobile + web), bigtable (large db), not severless
    inmem database (cache) -> memstore, access in microsecond
    OLPT: Online Transactional processing

Scenarios:

    start up quick -> datastore, filestore, filestore = datastore++ (multi devices)
    nosql, < 10Gi -> cloud datastore
    process 1Mi TPS -> cloud spanner
    process 1k TPS -> cloud SQL
    store ~1 peta bytes ->  big table nosql
    huge volume of IoT devices -> bigtable, nosql
    stream huge time series data (logs) -> bigtable, nosql

Cloud SQL

    - MySQL and SQL server

# NoSQL:

Datastore:
    
    - Document base database
    - auto scale and auto partitions
    - recommended upto a few TBs
    - If bigger volume bigtable
    - support transaction + index but not support join + sum + count
    - flexible schema

Filestore: (Datastore++)
    
    - support multi devices 
    - offline mode or data sync across multiple devices - mobile, iot
    - provide client-side lib (web, ios, android)
    - collections <=> tables

Bigtable:

    - petabytes and more 
    - not serverless

Bigquery - datawarehouse

    - SQL-like query
    - realtime + serverless
    - export: cloud storage or data studio (visualization), formats csv, json, 
    - import: various src
    - query from external src, cloud storage, cloud sql, bigtable
    - remember bq can be expensive as you are running them on large data sets, est price with --dry-run flag
    - price cost based on 1MB scaned data
    - dry-run (https://cloud.google.com/bigquery/docs/running-queries#dry-run)

Cloud Storage

    - suitable for both structure and unstructure data

Cloud Spanner

    - growth with minimum config
    - global distributed relational db
    - management complexity to a minimum

Notes:

    - growth with minimum config => Cloud Spanner

