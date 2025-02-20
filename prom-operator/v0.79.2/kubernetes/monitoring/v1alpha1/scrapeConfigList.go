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

// ScrapeConfigList is a list of ScrapeConfig
type ScrapeConfigList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of scrapeconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ScrapeConfigTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewScrapeConfigList registers a new resource with the given unique name, arguments, and options.
func NewScrapeConfigList(ctx *pulumi.Context,
	name string, args *ScrapeConfigListArgs, opts ...pulumi.ResourceOption) (*ScrapeConfigList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("monitoring.coreos.com/v1alpha1")
	args.Kind = pulumi.StringPtr("ScrapeConfigList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScrapeConfigList
	err := ctx.RegisterResource("kubernetes:monitoring.coreos.com/v1alpha1:ScrapeConfigList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScrapeConfigList gets an existing ScrapeConfigList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScrapeConfigList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScrapeConfigListState, opts ...pulumi.ResourceOption) (*ScrapeConfigList, error) {
	var resource ScrapeConfigList
	err := ctx.ReadResource("kubernetes:monitoring.coreos.com/v1alpha1:ScrapeConfigList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScrapeConfigList resources.
type scrapeConfigListState struct {
}

type ScrapeConfigListState struct {
}

func (ScrapeConfigListState) ElementType() reflect.Type {
	return reflect.TypeOf((*scrapeConfigListState)(nil)).Elem()
}

type scrapeConfigListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of scrapeconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ScrapeConfigType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ScrapeConfigList resource.
type ScrapeConfigListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of scrapeconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ScrapeConfigTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ScrapeConfigListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scrapeConfigListArgs)(nil)).Elem()
}

type ScrapeConfigListInput interface {
	pulumi.Input

	ToScrapeConfigListOutput() ScrapeConfigListOutput
	ToScrapeConfigListOutputWithContext(ctx context.Context) ScrapeConfigListOutput
}

func (*ScrapeConfigList) ElementType() reflect.Type {
	return reflect.TypeOf((**ScrapeConfigList)(nil)).Elem()
}

func (i *ScrapeConfigList) ToScrapeConfigListOutput() ScrapeConfigListOutput {
	return i.ToScrapeConfigListOutputWithContext(context.Background())
}

func (i *ScrapeConfigList) ToScrapeConfigListOutputWithContext(ctx context.Context) ScrapeConfigListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScrapeConfigListOutput)
}

// ScrapeConfigListArrayInput is an input type that accepts ScrapeConfigListArray and ScrapeConfigListArrayOutput values.
// You can construct a concrete instance of `ScrapeConfigListArrayInput` via:
//
//	ScrapeConfigListArray{ ScrapeConfigListArgs{...} }
type ScrapeConfigListArrayInput interface {
	pulumi.Input

	ToScrapeConfigListArrayOutput() ScrapeConfigListArrayOutput
	ToScrapeConfigListArrayOutputWithContext(context.Context) ScrapeConfigListArrayOutput
}

type ScrapeConfigListArray []ScrapeConfigListInput

func (ScrapeConfigListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScrapeConfigList)(nil)).Elem()
}

func (i ScrapeConfigListArray) ToScrapeConfigListArrayOutput() ScrapeConfigListArrayOutput {
	return i.ToScrapeConfigListArrayOutputWithContext(context.Background())
}

func (i ScrapeConfigListArray) ToScrapeConfigListArrayOutputWithContext(ctx context.Context) ScrapeConfigListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScrapeConfigListArrayOutput)
}

// ScrapeConfigListMapInput is an input type that accepts ScrapeConfigListMap and ScrapeConfigListMapOutput values.
// You can construct a concrete instance of `ScrapeConfigListMapInput` via:
//
//	ScrapeConfigListMap{ "key": ScrapeConfigListArgs{...} }
type ScrapeConfigListMapInput interface {
	pulumi.Input

	ToScrapeConfigListMapOutput() ScrapeConfigListMapOutput
	ToScrapeConfigListMapOutputWithContext(context.Context) ScrapeConfigListMapOutput
}

type ScrapeConfigListMap map[string]ScrapeConfigListInput

func (ScrapeConfigListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScrapeConfigList)(nil)).Elem()
}

func (i ScrapeConfigListMap) ToScrapeConfigListMapOutput() ScrapeConfigListMapOutput {
	return i.ToScrapeConfigListMapOutputWithContext(context.Background())
}

func (i ScrapeConfigListMap) ToScrapeConfigListMapOutputWithContext(ctx context.Context) ScrapeConfigListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScrapeConfigListMapOutput)
}

type ScrapeConfigListOutput struct{ *pulumi.OutputState }

func (ScrapeConfigListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScrapeConfigList)(nil)).Elem()
}

func (o ScrapeConfigListOutput) ToScrapeConfigListOutput() ScrapeConfigListOutput {
	return o
}

func (o ScrapeConfigListOutput) ToScrapeConfigListOutputWithContext(ctx context.Context) ScrapeConfigListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ScrapeConfigListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ScrapeConfigList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of scrapeconfigs. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ScrapeConfigListOutput) Items() ScrapeConfigTypeArrayOutput {
	return o.ApplyT(func(v *ScrapeConfigList) ScrapeConfigTypeArrayOutput { return v.Items }).(ScrapeConfigTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ScrapeConfigListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ScrapeConfigList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ScrapeConfigListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ScrapeConfigList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ScrapeConfigListArrayOutput struct{ *pulumi.OutputState }

func (ScrapeConfigListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScrapeConfigList)(nil)).Elem()
}

func (o ScrapeConfigListArrayOutput) ToScrapeConfigListArrayOutput() ScrapeConfigListArrayOutput {
	return o
}

func (o ScrapeConfigListArrayOutput) ToScrapeConfigListArrayOutputWithContext(ctx context.Context) ScrapeConfigListArrayOutput {
	return o
}

func (o ScrapeConfigListArrayOutput) Index(i pulumi.IntInput) ScrapeConfigListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScrapeConfigList {
		return vs[0].([]*ScrapeConfigList)[vs[1].(int)]
	}).(ScrapeConfigListOutput)
}

type ScrapeConfigListMapOutput struct{ *pulumi.OutputState }

func (ScrapeConfigListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScrapeConfigList)(nil)).Elem()
}

func (o ScrapeConfigListMapOutput) ToScrapeConfigListMapOutput() ScrapeConfigListMapOutput {
	return o
}

func (o ScrapeConfigListMapOutput) ToScrapeConfigListMapOutputWithContext(ctx context.Context) ScrapeConfigListMapOutput {
	return o
}

func (o ScrapeConfigListMapOutput) MapIndex(k pulumi.StringInput) ScrapeConfigListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScrapeConfigList {
		return vs[0].(map[string]*ScrapeConfigList)[vs[1].(string)]
	}).(ScrapeConfigListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScrapeConfigListInput)(nil)).Elem(), &ScrapeConfigList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScrapeConfigListArrayInput)(nil)).Elem(), ScrapeConfigListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScrapeConfigListMapInput)(nil)).Elem(), ScrapeConfigListMap{})
	pulumi.RegisterOutputType(ScrapeConfigListOutput{})
	pulumi.RegisterOutputType(ScrapeConfigListArrayOutput{})
	pulumi.RegisterOutputType(ScrapeConfigListMapOutput{})
}
