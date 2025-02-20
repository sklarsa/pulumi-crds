// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `PrometheusAgent` custom resource definition (CRD) defines a desired [Prometheus Agent](https://prometheus.io/blog/2021/11/16/agent/) setup to run in a Kubernetes cluster.
//
// The CRD is very similar to the `Prometheus` CRD except for features which aren't available in agent mode like rule evaluation, persistent storage and Thanos sidecar.
type PrometheusAgent struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput        `pulumi:"metadata"`
	Spec     PrometheusAgentSpecOutput      `pulumi:"spec"`
	Status   PrometheusAgentStatusPtrOutput `pulumi:"status"`
}

// NewPrometheusAgent registers a new resource with the given unique name, arguments, and options.
func NewPrometheusAgent(ctx *pulumi.Context,
	name string, args *PrometheusAgentArgs, opts ...pulumi.ResourceOption) (*PrometheusAgent, error) {
	if args == nil {
		args = &PrometheusAgentArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("monitoring.coreos.com/v1alpha1")
	args.Kind = pulumi.StringPtr("PrometheusAgent")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrometheusAgent
	err := ctx.RegisterResource("kubernetes:monitoring.coreos.com/v1alpha1:PrometheusAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusAgent gets an existing PrometheusAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusAgentState, opts ...pulumi.ResourceOption) (*PrometheusAgent, error) {
	var resource PrometheusAgent
	err := ctx.ReadResource("kubernetes:monitoring.coreos.com/v1alpha1:PrometheusAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusAgent resources.
type prometheusAgentState struct {
}

type PrometheusAgentState struct {
}

func (PrometheusAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAgentState)(nil)).Elem()
}

type prometheusAgentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta   `pulumi:"metadata"`
	Spec     *PrometheusAgentSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PrometheusAgent resource.
type PrometheusAgentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     PrometheusAgentSpecPtrInput
}

func (PrometheusAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAgentArgs)(nil)).Elem()
}

type PrometheusAgentInput interface {
	pulumi.Input

	ToPrometheusAgentOutput() PrometheusAgentOutput
	ToPrometheusAgentOutputWithContext(ctx context.Context) PrometheusAgentOutput
}

func (*PrometheusAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAgent)(nil)).Elem()
}

func (i *PrometheusAgent) ToPrometheusAgentOutput() PrometheusAgentOutput {
	return i.ToPrometheusAgentOutputWithContext(context.Background())
}

func (i *PrometheusAgent) ToPrometheusAgentOutputWithContext(ctx context.Context) PrometheusAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentOutput)
}

// PrometheusAgentArrayInput is an input type that accepts PrometheusAgentArray and PrometheusAgentArrayOutput values.
// You can construct a concrete instance of `PrometheusAgentArrayInput` via:
//
//	PrometheusAgentArray{ PrometheusAgentArgs{...} }
type PrometheusAgentArrayInput interface {
	pulumi.Input

	ToPrometheusAgentArrayOutput() PrometheusAgentArrayOutput
	ToPrometheusAgentArrayOutputWithContext(context.Context) PrometheusAgentArrayOutput
}

type PrometheusAgentArray []PrometheusAgentInput

func (PrometheusAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAgent)(nil)).Elem()
}

func (i PrometheusAgentArray) ToPrometheusAgentArrayOutput() PrometheusAgentArrayOutput {
	return i.ToPrometheusAgentArrayOutputWithContext(context.Background())
}

func (i PrometheusAgentArray) ToPrometheusAgentArrayOutputWithContext(ctx context.Context) PrometheusAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentArrayOutput)
}

// PrometheusAgentMapInput is an input type that accepts PrometheusAgentMap and PrometheusAgentMapOutput values.
// You can construct a concrete instance of `PrometheusAgentMapInput` via:
//
//	PrometheusAgentMap{ "key": PrometheusAgentArgs{...} }
type PrometheusAgentMapInput interface {
	pulumi.Input

	ToPrometheusAgentMapOutput() PrometheusAgentMapOutput
	ToPrometheusAgentMapOutputWithContext(context.Context) PrometheusAgentMapOutput
}

type PrometheusAgentMap map[string]PrometheusAgentInput

func (PrometheusAgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAgent)(nil)).Elem()
}

func (i PrometheusAgentMap) ToPrometheusAgentMapOutput() PrometheusAgentMapOutput {
	return i.ToPrometheusAgentMapOutputWithContext(context.Background())
}

func (i PrometheusAgentMap) ToPrometheusAgentMapOutputWithContext(ctx context.Context) PrometheusAgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentMapOutput)
}

type PrometheusAgentOutput struct{ *pulumi.OutputState }

func (PrometheusAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAgent)(nil)).Elem()
}

func (o PrometheusAgentOutput) ToPrometheusAgentOutput() PrometheusAgentOutput {
	return o
}

func (o PrometheusAgentOutput) ToPrometheusAgentOutputWithContext(ctx context.Context) PrometheusAgentOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PrometheusAgentOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAgent) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PrometheusAgentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAgent) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PrometheusAgentOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PrometheusAgent) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o PrometheusAgentOutput) Spec() PrometheusAgentSpecOutput {
	return o.ApplyT(func(v *PrometheusAgent) PrometheusAgentSpecOutput { return v.Spec }).(PrometheusAgentSpecOutput)
}

func (o PrometheusAgentOutput) Status() PrometheusAgentStatusPtrOutput {
	return o.ApplyT(func(v *PrometheusAgent) PrometheusAgentStatusPtrOutput { return v.Status }).(PrometheusAgentStatusPtrOutput)
}

type PrometheusAgentArrayOutput struct{ *pulumi.OutputState }

func (PrometheusAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAgent)(nil)).Elem()
}

func (o PrometheusAgentArrayOutput) ToPrometheusAgentArrayOutput() PrometheusAgentArrayOutput {
	return o
}

func (o PrometheusAgentArrayOutput) ToPrometheusAgentArrayOutputWithContext(ctx context.Context) PrometheusAgentArrayOutput {
	return o
}

func (o PrometheusAgentArrayOutput) Index(i pulumi.IntInput) PrometheusAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrometheusAgent {
		return vs[0].([]*PrometheusAgent)[vs[1].(int)]
	}).(PrometheusAgentOutput)
}

type PrometheusAgentMapOutput struct{ *pulumi.OutputState }

func (PrometheusAgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAgent)(nil)).Elem()
}

func (o PrometheusAgentMapOutput) ToPrometheusAgentMapOutput() PrometheusAgentMapOutput {
	return o
}

func (o PrometheusAgentMapOutput) ToPrometheusAgentMapOutputWithContext(ctx context.Context) PrometheusAgentMapOutput {
	return o
}

func (o PrometheusAgentMapOutput) MapIndex(k pulumi.StringInput) PrometheusAgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrometheusAgent {
		return vs[0].(map[string]*PrometheusAgent)[vs[1].(string)]
	}).(PrometheusAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentInput)(nil)).Elem(), &PrometheusAgent{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentArrayInput)(nil)).Elem(), PrometheusAgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentMapInput)(nil)).Elem(), PrometheusAgentMap{})
	pulumi.RegisterOutputType(PrometheusAgentOutput{})
	pulumi.RegisterOutputType(PrometheusAgentArrayOutput{})
	pulumi.RegisterOutputType(PrometheusAgentMapOutput{})
}
