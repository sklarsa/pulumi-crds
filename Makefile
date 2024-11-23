ARGOCD_VERSION=

.PHONY:argocd
argocd:
	crd2pulumi --goPath ./argocd/$(ARGOCD_VERSION) https://raw.githubusercontent.com/argoproj/argo-cd/$(ARGOCD_VERSION)/manifests/install.yaml
