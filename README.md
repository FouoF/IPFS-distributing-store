# IPFS-Store
This is a project based on IPFS cluster and IPFS, it uses cloud native design and can be esaliy deployed in K8s cluster.

## Prepare
This project use local docker register, so if you want to use it correctly, you must have you own register and build images youself.
To reduce network issue, all image is build with pre-build binary file, build before docker build is necessary.

## Deploy
```
# deploy IPFS cluster
cd ./deploy/ipfs-operator
make deploy

# /deploy contains all necessary K8s resource this project need.
# NOTICE: The project use cilium as CNI, if you use other CNI or use cloud-provider load-balancer, you need to configure it yourself.
kubectl apply -f /deploy *
