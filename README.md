# X-ClusterLink

X-ClusterLink: An Efficient Cross-Cluster CommunicationFramework in Multi-Kubernetes Clusters

This repository contains codes for a research paper that was submitted to WWW 2025.

### About This Repository

**This repository has been prepared for anonymous review. To protect proprietary and organizational information, certain sensitive details (e.g., internal identifiers) have been replaced or removed.**

## Overview
X-ClusterLink is a framework designed to enable efficient cross-cluster communication in multi-Kubernetes cluster environments. It aims to achieve low latency, high throughput, strong robustness, and scalability, meeting the demands of modern distributed systems.

## Environment
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

## Build

In Kubernetes-based environments, containerized applications (like controllers or services) are typically packaged into Docker images, pushed to a container registry, and then deployed on the worker nodes as Pods. Here's how you can package and deploy your controller

The first step is to build the Docker image for your controller. This is done using the Dockerfile present in your project directory. The make docker-build command automates the process of building the image.

```
make docker-build
```

## Start 
### Deployment (on Kubernetes cluster)

Once your Docker image is built and stored, the next step is to deploy it to the Kubernetes cluster. Kubernetes abstracts the deployment process and ensures that the application runs on a set of worker nodes (depending on your configurations).

The make deploy command is used to apply Kubernetes deployment configurations, which typically include a Deployment resource, Pod definitions, and any other required Kubernetes objects

```
make deploy
kubectl apply -f deploy.yaml
```
### Deployment gateway pod

These command deploys the gateway and network configurations to the Kubernetes cluster. The gateway-config.yaml sets up the gateway configuration, external-network.yaml defines the external network settings, and gateway.yaml deploys the gateway pod.

```
kubectl apply gateway-config.yaml
kubectl apply external-network.yaml
kubectl apply -f gateway.yaml
```

### Multi-Cluster Connection

This command applies the `connect.yaml` configuration to establish a connection between Kubernetes clusters via tunnel between gateway pods, enabling cross-cluster communication.
```
kubectl apply -f connect.yaml
```