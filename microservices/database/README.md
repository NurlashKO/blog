Welcome to `database` microservice docs
---

Command to connect to the database:
```bash
gcloud compute ssh --zone "us-west1-b" "instance-1" --tunnel-through-iap --project "kouzoh-p-nurlashko" -- -NL 5432:localhost:5432
```