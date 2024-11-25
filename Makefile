ARGOCD_VERSION=v2.13.1
CERT_MANAGER_VERSION=v1.16.2
VMKS_VERSION=v0.49.0


.PHONY:argocd
argocd:
	crd2pulumi --goPath ./argocd/$(ARGOCD_VERSION) https://raw.githubusercontent.com/argoproj/argo-cd/$(ARGOCD_VERSION)/manifests/install.yaml

.PHONY:cert-manager
cert-manager:
	crd2pulumi --goPath ./cert-manager/$(CERT_MANAGER_VERSION) https://github.com/cert-manager/cert-manager/releases/download/$(CERT_MANAGER_VERSION)/cert-manager.yaml

.PHONY:vmks
vmks:
	crd2pulumi --goPath ./vmks/$(VMKS_VERSION) https://raw.githubusercontent.com/VictoriaMetrics/operator/refs/tags/$(VMKS_VERSION)/config/crd/overlay/crd.yaml
