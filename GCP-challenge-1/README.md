# GCP Challenge 1
```bash
gcloud init

gcloud config set compute/region eu-west1

gcloud config set compute/zone eu-west1-b

gcloud compute instances create hackajob-vm-1 hackajob-vm-2 \
    --machine-type=f1-micro \
    --image-family=debian-11 \
    --image-project=debian-cloud

gcloud compute health-checks create tcp hackajob-health-check \
    --port=80

gcloud compute instance-groups unmanaged create hackajob-instance-group

gcloud compute instance-groups unmanaged add-instances hackajob-instance-group \
    --instances=hackajob-vm-1,hackajob-vm-2

gcloud compute instance-groups set-named-ports hackajob-instance-group \
    --named-ports=http:80

gcloud compute backend-services create hackajob-backend-service \
    --global \
    --protocol=HTTP \
    --port-name=http \
    --health-checks=hackajob-health-check

gcloud compute url-maps create hackajob-lb \
    --default-service=hackajob-backend-service

gcloud compute target-http-proxies create hackajob-http-proxy \
    --url-map=hackajob-lb

gcloud compute forwarding-rules create hackajob-lb \
    --global \
    --target-http-proxy=hackajob-http-proxy \
    --ports=80
```