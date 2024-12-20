// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:operator.victoriametrics.com/v1beta1:VLogs":
		r = &VLogs{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VLogsList":
		r = &VLogsList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VLogsPatch":
		r = &VLogsPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAgent":
		r = &VMAgent{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAgentList":
		r = &VMAgentList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAgentPatch":
		r = &VMAgentPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlert":
		r = &VMAlert{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertList":
		r = &VMAlertList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertPatch":
		r = &VMAlertPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanager":
		r = &VMAlertmanager{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerConfig":
		r = &VMAlertmanagerConfig{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerConfigList":
		r = &VMAlertmanagerConfigList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerConfigPatch":
		r = &VMAlertmanagerConfigPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerList":
		r = &VMAlertmanagerList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerPatch":
		r = &VMAlertmanagerPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAuth":
		r = &VMAuth{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAuthList":
		r = &VMAuthList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMAuthPatch":
		r = &VMAuthPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMCluster":
		r = &VMCluster{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMClusterList":
		r = &VMClusterList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMClusterPatch":
		r = &VMClusterPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMNodeScrape":
		r = &VMNodeScrape{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMNodeScrapeList":
		r = &VMNodeScrapeList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMNodeScrapePatch":
		r = &VMNodeScrapePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMPodScrape":
		r = &VMPodScrape{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMPodScrapeList":
		r = &VMPodScrapeList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMPodScrapePatch":
		r = &VMPodScrapePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMProbe":
		r = &VMProbe{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMProbeList":
		r = &VMProbeList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMProbePatch":
		r = &VMProbePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMRule":
		r = &VMRule{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMRuleList":
		r = &VMRuleList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMRulePatch":
		r = &VMRulePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMScrapeConfig":
		r = &VMScrapeConfig{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMScrapeConfigList":
		r = &VMScrapeConfigList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMScrapeConfigPatch":
		r = &VMScrapeConfigPatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMServiceScrape":
		r = &VMServiceScrape{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMServiceScrapeList":
		r = &VMServiceScrapeList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMServiceScrapePatch":
		r = &VMServiceScrapePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMSingle":
		r = &VMSingle{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMSingleList":
		r = &VMSingleList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMSinglePatch":
		r = &VMSinglePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMStaticScrape":
		r = &VMStaticScrape{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMStaticScrapeList":
		r = &VMStaticScrapeList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMStaticScrapePatch":
		r = &VMStaticScrapePatch{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMUser":
		r = &VMUser{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMUserList":
		r = &VMUserList{}
	case "kubernetes:operator.victoriametrics.com/v1beta1:VMUserPatch":
		r = &VMUserPatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"crds",
		"operator.victoriametrics.com/v1beta1",
		&module{version},
	)
}
