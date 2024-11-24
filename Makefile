ARGOCD_VERSION=v2.13.1
VMKS_VERSION=v0.49.0

.PHONY:argocd
argocd:
	crd2pulumi --goPath ./argocd/$(ARGOCD_VERSION) https://raw.githubusercontent.com/argoproj/argo-cd/$(ARGOCD_VERSION)/manifests/install.yaml

.PHONY:vmks
vmks:
	crd2pulumi --goPath ./vmks/$(VMKS_VERSION) https://raw.githubusercontent.com/VictoriaMetrics/operator/refs/tags/$(VMKS_VERSION)/config/crd/overlay/crd.yaml
