package addons

const CloudProviderAwsTemplate = `
---
# Source: aws-cloud-controller-manager/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cloud-controller-manager
  namespace: kube-system
  labels:
    helm.sh/chart: "aws-cloud-controller-manager-0.0.7"
---
# Source: aws-cloud-controller-manager/templates/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:cloud-controller-manager
  labels:
    helm.sh/chart: "aws-cloud-controller-manager-0.0.7"
rules: 
  - apiGroups:
    - ""
    resources:
    - events
    verbs:
    - create
    - patch
    - update
  - apiGroups:
    - ""
    resources:
    - nodes
    verbs:
    - '*'
  - apiGroups:
    - ""
    resources:
    - nodes/status
    verbs:
    - patch
  - apiGroups:
    - ""
    resources:
    - services
    verbs:
    - list
    - patch
    - update
    - watch
  - apiGroups:
    - ""
    resources:
    - services/status
    verbs:
    - list
    - patch
    - update
    - watch
  - apiGroups:
    - ""
    resources:
    - serviceaccounts
    verbs:
    - create
  - apiGroups:
    - ""
    resources:
    - persistentvolumes
    verbs:
    - get
    - list
    - update
    - watch
  - apiGroups:
    - ""
    resources:
    - endpoints
    verbs:
    - create
    - get
    - list
    - watch
    - update
  - apiGroups:
    - coordination.k8s.io
    resources:
    - leases
    verbs:
    - create
    - get
    - list
    - watch
    - update
  - apiGroups:
    - ""
    resources:
    - serviceaccounts/token
    verbs:
    - create
---
# Source: aws-cloud-controller-manager/templates/cluserrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:cloud-controller-manager
  labels:
    helm.sh/chart: "aws-cloud-controller-manager-0.0.7"
roleRef:
  kind: ClusterRole
  name: system:cloud-controller-manager
  apiGroup: rbac.authorization.k8s.io
subjects:
- apiGroup: ""
  kind: ServiceAccount
  name: cloud-controller-manager
  namespace: kube-system
---
# Source: aws-cloud-controller-manager/templates/rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: cloud-controller-manager:apiserver-authentication-reader
  namespace: kube-system
  labels: 
    helm.sh/chart: "aws-cloud-controller-manager-0.0.7"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- apiGroup: ""
  kind: ServiceAccount
  name: cloud-controller-manager
  namespace: kube-system
---
# Source: aws-cloud-controller-manager/templates/daemonset.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: aws-cloud-controller-manager
  labels:
    k8s-app:  aws-cloud-controller-manager
    helm.sh/chart: "aws-cloud-controller-manager-0.0.7"
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app:  aws-cloud-controller-manager
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      name: aws-cloud-controller-manager
      labels:
        k8s-app:  aws-cloud-controller-manager
    spec:
      tolerations: 
        - effect: NoSchedule
          key: node.cloudprovider.kubernetes.io/uninitialized
          value: "true"
        - effect: NoSchedule
          key: node-role.kubernetes.io/master
        - effect: NoSchedule
          key: node-role.kubernetes.io/control-plane
      nodeSelector: 
        null
      dnsPolicy: Default
      priorityClassName: system-node-critical
      serviceAccountName: cloud-controller-manager
      securityContext:
        {}
      containers:
        - name: aws-cloud-controller-manager
          image: "registry.k8s.io/provider-aws/cloud-controller-manager:v1.23.0-alpha.0"
          args:
            - --v=2
            - --cloud-provider=aws
            - --controllers=cloud-node,cloud-node-lifecycle
            - --configure-cloud-routes=false
          resources:
            requests:
              cpu: 200m
          env:
            []
          securityContext:
            {}
`
