## OrchestratorEX

This repository is used to allow for easy visualization / understanding of overlay networks in kubernetes and docker swarm.

### Deployment instructions:

**Docker Swarm:**

```
## All commands in PowerShell
curl "https://raw.githubusercontent.com/WAftring/OrchestratorEX/master/SWARM-OrchestratorEX.yaml" -o docker-compose.yml
docker stack deploy --compose-file docker-compose.yml orchestratorex
```

**Kubernetes:**

WIP
