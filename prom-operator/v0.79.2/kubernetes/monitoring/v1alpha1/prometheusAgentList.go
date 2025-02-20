// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PrometheusAgentList is a list of PrometheusAgent
type PrometheusAgentList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of prometheusagents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items PrometheusAgentTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewPrometheusAgentList registers a new resource with the given unique name, arguments, and options.
func NewPrometheusAgentList(ctx *pulumi.Context,
	name string, args *PrometheusAgentListArgs, opts ...pulumi.ResourceOption) (*PrometheusAgentList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("monitoring.coreos.com/v1alpha1")
	args.Kind = pulumi.StringPtr("PrometheusAgentList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrometheusAgentList
	err := ctx.RegisterResource("kubernetes:monitoring.coreos.com/v1alpha1:PrometheusAgentList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusAgentList gets an existing PrometheusAgentList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusAgentList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusAgentListState, opts ...pulumi.ResourceOption) (*PrometheusAgentList, error) {
	var resource PrometheusAgentList
	err := ctx.ReadResource("kubernetes:monitoring.coreos.com/v1alpha1:PrometheusAgentList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusAgentList resources.
type prometheusAgentListState struct {
}

type PrometheusAgentListState struct {
}

func (PrometheusAgentListState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAgentListState)(nil)).Elem()
}

type prometheusAgentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of prometheusagents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []PrometheusAgentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PrometheusAgentList resource.
type PrometheusAgentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of prometheusagents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items PrometheusAgentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (PrometheusAgentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAgentListArgs)(nil)).Elem()
}

type PrometheusAgentListInput interface {
	pulumi.Input

	ToPrometheusAgentListOutput() PrometheusAgentListOutput
	ToPrometheusAgentListOutputWithContext(ctx context.Context) PrometheusAgentListOutput
}

func (*PrometheusAgentList) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAgentList)(nil)).Elem()
}

func (i *PrometheusAgentList) ToPrometheusAgentListOutput() PrometheusAgentListOutput {
	return i.ToPrometheusAgentListOutputWithContext(context.Background())
}

func (i *PrometheusAgentList) ToPrometheusAgentListOutputWithContext(ctx context.Context) PrometheusAgentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentListOutput)
}

// PrometheusAgentListArrayInput is an input type that accepts PrometheusAgentListArray and PrometheusAgentListArrayOutput values.
// You can construct a concrete instance of `PrometheusAgentListArrayInput` via:
//
//	PrometheusAgentListArray{ PrometheusAgentListArgs{...} }
type PrometheusAgentListArrayInput interface {
	pulumi.Input

	ToPrometheusAgentListArrayOutput() PrometheusAgentListArrayOutput
	ToPrometheusAgentListArrayOutputWithContext(context.Context) PrometheusAgentListArrayOutput
}

type PrometheusAgentListArray []PrometheusAgentListInput

func (PrometheusAgentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAgentList)(nil)).Elem()
}

func (i PrometheusAgentListArray) ToPrometheusAgentListArrayOutput() PrometheusAgentListArrayOutput {
	return i.ToPrometheusAgentListArrayOutputWithContext(context.Background())
}

func (i PrometheusAgentListArray) ToPrometheusAgentListArrayOutputWithContext(ctx context.Context) PrometheusAgentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentListArrayOutput)
}

// PrometheusAgentListMapInput is an input type that accepts PrometheusAgentListMap and PrometheusAgentListMapOutput values.
// You can construct a concrete instance of `PrometheusAgentListMapInput` via:
//
//	PrometheusAgentListMap{ "key": PrometheusAgentListArgs{...} }
type PrometheusAgentListMapInput interface {
	pulumi.Input

	ToPrometheusAgentListMapOutput() PrometheusAgentListMapOutput
	ToPrometheusAgentListMapOutputWithContext(context.Context) PrometheusAgentListMapOutput
}

type PrometheusAgentListMap map[string]PrometheusAgentListInput

func (PrometheusAgentListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAgentList)(nil)).Elem()
}

func (i PrometheusAgentListMap) ToPrometheusAgentListMapOutput() PrometheusAgentListMapOutput {
	return i.ToPrometheusAgentListMapOutputWithContext(context.Background())
}

func (i PrometheusAgentListMap) ToPrometheusAgentListMapOutputWithContext(ctx context.Context) PrometheusAgentListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAgentListMapOutput)
}

type PrometheusAgentListOutput struct{ *pulumi.OutputState }

func (PrometheusAgentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAgentList)(nil)).Elem()
}

func (o PrometheusAgentListOutput) ToPrometheusAgentListOutput() PrometheusAgentListOutput {
	return o
}

func (o PrometheusAgentListOutput) ToPrometheusAgentListOutputWithContext(ctx context.Context) PrometheusAgentListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PrometheusAgentListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAgentList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of prometheusagents. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o PrometheusAgentListOutput) Items() PrometheusAgentTypeArrayOutput {
	return o.ApplyT(func(v *PrometheusAgentList) PrometheusAgentTypeArrayOutput { return v.Items }).(PrometheusAgentTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PrometheusAgentListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAgentList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PrometheusAgentListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *PrometheusAgentList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type PrometheusAgentListArrayOutput struct{ *pulumi.OutputState }

func (PrometheusAgentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAgentList)(nil)).Elem()
}

func (o PrometheusAgentListArrayOutput) ToPrometheusAgentListArrayOutput() PrometheusAgentListArrayOutput {
	return o
}

func (o PrometheusAgentListArrayOutput) ToPrometheusAgentListArrayOutputWithContext(ctx context.Context) PrometheusAgentListArrayOutput {
	return o
}

func (o PrometheusAgentListArrayOutput) Index(i pulumi.IntInput) PrometheusAgentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrometheusAgentList {
		return vs[0].([]*PrometheusAgentList)[vs[1].(int)]
	}).(PrometheusAgentListOutput)
}

type PrometheusAgentListMapOutput struct{ *pulumi.OutputState }

func (PrometheusAgentListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAgentList)(nil)).Elem()
}

func (o PrometheusAgentListMapOutput) ToPrometheusAgentListMapOutput() PrometheusAgentListMapOutput {
	return o
}

func (o PrometheusAgentListMapOutput) ToPrometheusAgentListMapOutputWithContext(ctx context.Context) PrometheusAgentListMapOutput {
	return o
}

func (o PrometheusAgentListMapOutput) MapIndex(k pulumi.StringInput) PrometheusAgentListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrometheusAgentList {
		return vs[0].(map[string]*PrometheusAgentList)[vs[1].(string)]
	}).(PrometheusAgentListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentListInput)(nil)).Elem(), &PrometheusAgentList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentListArrayInput)(nil)).Elem(), PrometheusAgentListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAgentListMapInput)(nil)).Elem(), PrometheusAgentListMap{})
	pulumi.RegisterOutputType(PrometheusAgentListOutput{})
	pulumi.RegisterOutputType(PrometheusAgentListArrayOutput{})
	pulumi.RegisterOutputType(PrometheusAgentListMapOutput{})
}
