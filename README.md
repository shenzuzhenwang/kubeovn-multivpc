# X-ClusterLink

X-ClusterLink: An Efficient Cross-Cluster CommunicationFramework in Multi-Kubernetes Clusters

This repository contains codes for a research paper that was submitted to WWW 2025.

# Overview
X-ClusterLink is a framework designed to enable efficient cross-cluster communication in multi-Kubernetes cluster environments. It aims to achieve low latency, high throughput, strong robustness, and scalability, meeting the demands of modern distributed systems.

# Environment
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

# Build
```
make docker-build
make deploy
```

# Start 
## Deployment (on Kubernetes cluster)
```
kubectl apply -f deploy.yaml
```
Once deployed successfully, the output will be:
```
# kubectl get pod -A
multi-vpc-system      multi-vpc-controller-manager-7cf6c6b9d6-q4sq5    2/2     Running            0               73s
```
## Modify Broker RBAC
Modify RBAC in the cluster where the broker is deployed (to synchronize GatewayExIps via the broker):
```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: broker-cluster
  namespace: k8s-broker
rules:
  - apiGroups:
      - example.io
    resources:
      - clusters
      - endpoints
    verbs:
      - create
      - get
      - list
      - watch
      - patch
      - update
      - delete
  - apiGroups:
      - example.io
    resources:
      - brokers
    verbs:
      - get
      - list
  - apiGroups:
      - multicluster.x-k8s.io
    resources:
      - '*'
    verbs:
      - create
      - get
      - list
      - watch
      - patch
      - update
      - delete
  - apiGroups:
      - discovery.k8s.io
    resources:
      - endpointslices
      - endpointslices/restricted
    verbs:
      - create
      - get
      - list
      - watch
      - patch
      - update
      - delete
  - apiGroups:
      - ""
    resources:
      - secrets
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - serviceaccounts
    verbs:
      - get
      - list
  - apiGroups:
      - example.io
    resources:
      - gatewayexips
    verbs:
      - create
      - get
      - list
      - watch
      - patch
      - update
      - delete
```

## Connection
Create tunnel.yaml to connect clusters
```yaml
apiVersion: "example.io/v1"
kind: VpcNatTunnel
metadata:
  name: tn1
  namespace: ns2
spec:
  remoteCluster: "cluster2"      
  remoteVpc: "test-vpc-2"        
  interfaceAddr: "10.100.0.1/24" 
  localVpc: "test-vpc-1"        
  type: "vxlan" 
```
Where `remoteCluster` is the Submariner cluster ID of the remote cluster, and `remoteVpc` is the name of the remote VPC
```
kubectl apply -f tunnel.yaml
```
Then check the connection.

## About This Repository

This repository has been prepared for anonymous review. To protect proprietary and organizational information, certain sensitive details (e.g., internal identifiers) have been replaced or removed.