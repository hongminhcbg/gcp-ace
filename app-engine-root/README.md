## App engine with multi svc

### Step 1

Init app engine and deploy first svc (default svc)

### Step 2

```sh
  cd svc1 && gcloud app deploy
```

Access svc1 URL https://py-dot-compute-engine-442508.as.r.appspot.com/

### Step 3

```sh
  cd svc2 && gcloud app deploy
```

Access svc2 URL https://nodejs-dot-compute-engine-442508.as.r.appspot.com/

