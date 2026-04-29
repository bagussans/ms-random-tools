## 🚀 PRODUCTION BUILD

```bash
# on root, run script below, so it will build locally
docker build -t ms-random-tools:local -f deployment_config/Dockerfile .

# run script below to import docker image to k3s containerd, so it can be read and deploy via k3s
docker save ms-random-tools:local | k3s ctr images import -

# apply all yaml from deployment_config as usual
```