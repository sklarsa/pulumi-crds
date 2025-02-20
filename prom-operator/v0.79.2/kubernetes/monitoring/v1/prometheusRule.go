// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `PrometheusRule` custom resource definition (CRD) defines [alerting](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/) and [recording](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) rules to be evaluated by `Prometheus` or `ThanosRuler` objects.
//
// `Prometheus` and `ThanosRuler` objects select `PrometheusRule` objects using label and namespace selectors.
type PrometheusRule struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput  `pulumi:"metadata"`
	Spec     PrometheusRuleSpecOutput `pulumi:"spec"`
}

// NewPrometheusRule registers a new resource with the given unique name, arguments, and options.
func NewPrometheusRule(ctx *pulumi.Context,
	name string, args *PrometheusRuleArgs, opts ...pulumi.ResourceOption) (*PrometheusRule, error) {
	if args == nil {
		args = &PrometheusRuleArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("monitoring.coreos.com/v1")
	args.Kind = pulumi.StringPtr("PrometheusRule")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrometheusRule
	err := ctx.RegisterResource("kubernetes:monitoring.coreos.com/v1:PrometheusRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusRule gets an existing PrometheusRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusRuleState, opts ...pulumi.ResourceOption) (*PrometheusRule, error) {
	var resource PrometheusRule
	err := ctx.ReadResource("kubernetes:monitoring.coreos.com/v1:PrometheusRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusRule resources.
type prometheusRuleState struct {
}

type PrometheusRuleState struct {
}

func (PrometheusRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleState)(nil)).Elem()
}

type prometheusRuleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta  `pulumi:"metadata"`
	Spec     *PrometheusRuleSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PrometheusRule resource.
type PrometheusRuleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     PrometheusRuleSpecPtrInput
}

func (PrometheusRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleArgs)(nil)).Elem()
}

type PrometheusRuleInput interface {
	pulumi.Input

	ToPrometheusRuleOutput() PrometheusRuleOutput
	ToPrometheusRuleOutputWithContext(ctx context.Context) PrometheusRuleOutput
}

func (*PrometheusRule) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRule)(nil)).Elem()
}

func (i *PrometheusRule) ToPrometheusRuleOutput() PrometheusRuleOutput {
	return i.ToPrometheusRuleOutputWithContext(context.Background())
}

func (i *PrometheusRule) ToPrometheusRuleOutputWithContext(ctx context.Context) PrometheusRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleOutput)
}

// PrometheusRuleArrayInput is an input type that accepts PrometheusRuleArray and PrometheusRuleArrayOutput values.
// You can construct a concrete instance of `PrometheusRuleArrayInput` via:
//
//	PrometheusRuleArray{ PrometheusRuleArgs{...} }
type PrometheusRuleArrayInput interface {
	pulumi.Input

	ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput
	ToPrometheusRuleArrayOutputWithContext(context.Context) PrometheusRuleArrayOutput
}

type PrometheusRuleArray []PrometheusRuleInput

func (PrometheusRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusRule)(nil)).Elem()
}

func (i PrometheusRuleArray) ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput {
	return i.ToPrometheusRuleArrayOutputWithContext(context.Background())
}

func (i PrometheusRuleArray) ToPrometheusRuleArrayOutputWithContext(ctx context.Context) PrometheusRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleArrayOutput)
}

// PrometheusRuleMapInput is an input type that accepts PrometheusRuleMap and PrometheusRuleMapOutput values.
// You can construct a concrete instance of `PrometheusRuleMapInput` via:
//
//	PrometheusRuleMap{ "key": PrometheusRuleArgs{...} }
type PrometheusRuleMapInput interface {
	pulumi.Input

	ToPrometheusRuleMapOutput() PrometheusRuleMapOutput
	ToPrometheusRuleMapOutputWithContext(context.Context) PrometheusRuleMapOutput
}

type PrometheusRuleMap map[string]PrometheusRuleInput

func (PrometheusRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusRule)(nil)).Elem()
}

func (i PrometheusRuleMap) ToPrometheusRuleMapOutput() PrometheusRuleMapOutput {
	return i.ToPrometheusRuleMapOutputWithContext(context.Background())
}

func (i PrometheusRuleMap) ToPrometheusRuleMapOutputWithContext(ctx context.Context) PrometheusRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleMapOutput)
}

type PrometheusRuleOutput struct{ *pulumi.OutputState }

func (PrometheusRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRule)(nil)).Elem()
}

func (o PrometheusRuleOutput) ToPrometheusRuleOutput() PrometheusRuleOutput {
	return o
}

func (o PrometheusRuleOutput) ToPrometheusRuleOutputWithContext(ctx context.Context) PrometheusRuleOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PrometheusRuleOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRule) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PrometheusRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PrometheusRuleOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PrometheusRule) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o PrometheusRuleOutput) Spec() PrometheusRuleSpecOutput {
	return o.ApplyT(func(v *PrometheusRule) PrometheusRuleSpecOutput { return v.Spec }).(PrometheusRuleSpecOutput)
}

type PrometheusRuleArrayOutput struct{ *pulumi.OutputState }

func (PrometheusRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusRule)(nil)).Elem()
}

func (o PrometheusRuleArrayOutput) ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput {
	return o
}

func (o PrometheusRuleArrayOutput) ToPrometheusRuleArrayOutputWithContext(ctx context.Context) PrometheusRuleArrayOutput {
	return o
}

func (o PrometheusRuleArrayOutput) Index(i pulumi.IntInput) PrometheusRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrometheusRule {
		return vs[0].([]*PrometheusRule)[vs[1].(int)]
	}).(PrometheusRuleOutput)
}

type PrometheusRuleMapOutput struct{ *pulumi.OutputState }

func (PrometheusRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusRule)(nil)).Elem()
}

func (o PrometheusRuleMapOutput) ToPrometheusRuleMapOutput() PrometheusRuleMapOutput {
	return o
}

func (o PrometheusRuleMapOutput) ToPrometheusRuleMapOutputWithContext(ctx context.Context) PrometheusRuleMapOutput {
	return o
}

func (o PrometheusRuleMapOutput) MapIndex(k pulumi.StringInput) PrometheusRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrometheusRule {
		return vs[0].(map[string]*PrometheusRule)[vs[1].(string)]
	}).(PrometheusRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusRuleInput)(nil)).Elem(), &PrometheusRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusRuleArrayInput)(nil)).Elem(), PrometheusRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusRuleMapInput)(nil)).Elem(), PrometheusRuleMap{})
	pulumi.RegisterOutputType(PrometheusRuleOutput{})
	pulumi.RegisterOutputType(PrometheusRuleArrayOutput{})
	pulumi.RegisterOutputType(PrometheusRuleMapOutput{})
}
