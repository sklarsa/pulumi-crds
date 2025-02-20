// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VMAlertmanagerList is a list of VMAlertmanager
type VMAlertmanagerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of vmalertmanagers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items VMAlertmanagerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewVMAlertmanagerList registers a new resource with the given unique name, arguments, and options.
func NewVMAlertmanagerList(ctx *pulumi.Context,
	name string, args *VMAlertmanagerListArgs, opts ...pulumi.ResourceOption) (*VMAlertmanagerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMAlertmanagerList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMAlertmanagerList
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMAlertmanagerList gets an existing VMAlertmanagerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMAlertmanagerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMAlertmanagerListState, opts ...pulumi.ResourceOption) (*VMAlertmanagerList, error) {
	var resource VMAlertmanagerList
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMAlertmanagerList resources.
type vmalertmanagerListState struct {
}

type VMAlertmanagerListState struct {
}

func (VMAlertmanagerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertmanagerListState)(nil)).Elem()
}

type vmalertmanagerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of vmalertmanagers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []VMAlertmanagerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a VMAlertmanagerList resource.
type VMAlertmanagerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of vmalertmanagers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items VMAlertmanagerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (VMAlertmanagerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertmanagerListArgs)(nil)).Elem()
}

type VMAlertmanagerListInput interface {
	pulumi.Input

	ToVMAlertmanagerListOutput() VMAlertmanagerListOutput
	ToVMAlertmanagerListOutputWithContext(ctx context.Context) VMAlertmanagerListOutput
}

func (*VMAlertmanagerList) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlertmanagerList)(nil)).Elem()
}

func (i *VMAlertmanagerList) ToVMAlertmanagerListOutput() VMAlertmanagerListOutput {
	return i.ToVMAlertmanagerListOutputWithContext(context.Background())
}

func (i *VMAlertmanagerList) ToVMAlertmanagerListOutputWithContext(ctx context.Context) VMAlertmanagerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerListOutput)
}

// VMAlertmanagerListArrayInput is an input type that accepts VMAlertmanagerListArray and VMAlertmanagerListArrayOutput values.
// You can construct a concrete instance of `VMAlertmanagerListArrayInput` via:
//
//	VMAlertmanagerListArray{ VMAlertmanagerListArgs{...} }
type VMAlertmanagerListArrayInput interface {
	pulumi.Input

	ToVMAlertmanagerListArrayOutput() VMAlertmanagerListArrayOutput
	ToVMAlertmanagerListArrayOutputWithContext(context.Context) VMAlertmanagerListArrayOutput
}

type VMAlertmanagerListArray []VMAlertmanagerListInput

func (VMAlertmanagerListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlertmanagerList)(nil)).Elem()
}

func (i VMAlertmanagerListArray) ToVMAlertmanagerListArrayOutput() VMAlertmanagerListArrayOutput {
	return i.ToVMAlertmanagerListArrayOutputWithContext(context.Background())
}

func (i VMAlertmanagerListArray) ToVMAlertmanagerListArrayOutputWithContext(ctx context.Context) VMAlertmanagerListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerListArrayOutput)
}

// VMAlertmanagerListMapInput is an input type that accepts VMAlertmanagerListMap and VMAlertmanagerListMapOutput values.
// You can construct a concrete instance of `VMAlertmanagerListMapInput` via:
//
//	VMAlertmanagerListMap{ "key": VMAlertmanagerListArgs{...} }
type VMAlertmanagerListMapInput interface {
	pulumi.Input

	ToVMAlertmanagerListMapOutput() VMAlertmanagerListMapOutput
	ToVMAlertmanagerListMapOutputWithContext(context.Context) VMAlertmanagerListMapOutput
}

type VMAlertmanagerListMap map[string]VMAlertmanagerListInput

func (VMAlertmanagerListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlertmanagerList)(nil)).Elem()
}

func (i VMAlertmanagerListMap) ToVMAlertmanagerListMapOutput() VMAlertmanagerListMapOutput {
	return i.ToVMAlertmanagerListMapOutputWithContext(context.Background())
}

func (i VMAlertmanagerListMap) ToVMAlertmanagerListMapOutputWithContext(ctx context.Context) VMAlertmanagerListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerListMapOutput)
}

type VMAlertmanagerListOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlertmanagerList)(nil)).Elem()
}

func (o VMAlertmanagerListOutput) ToVMAlertmanagerListOutput() VMAlertmanagerListOutput {
	return o
}

func (o VMAlertmanagerListOutput) ToVMAlertmanagerListOutputWithContext(ctx context.Context) VMAlertmanagerListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMAlertmanagerListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAlertmanagerList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of vmalertmanagers. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o VMAlertmanagerListOutput) Items() VMAlertmanagerTypeArrayOutput {
	return o.ApplyT(func(v *VMAlertmanagerList) VMAlertmanagerTypeArrayOutput { return v.Items }).(VMAlertmanagerTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMAlertmanagerListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAlertmanagerList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMAlertmanagerListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *VMAlertmanagerList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type VMAlertmanagerListArrayOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlertmanagerList)(nil)).Elem()
}

func (o VMAlertmanagerListArrayOutput) ToVMAlertmanagerListArrayOutput() VMAlertmanagerListArrayOutput {
	return o
}

func (o VMAlertmanagerListArrayOutput) ToVMAlertmanagerListArrayOutputWithContext(ctx context.Context) VMAlertmanagerListArrayOutput {
	return o
}

func (o VMAlertmanagerListArrayOutput) Index(i pulumi.IntInput) VMAlertmanagerListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMAlertmanagerList {
		return vs[0].([]*VMAlertmanagerList)[vs[1].(int)]
	}).(VMAlertmanagerListOutput)
}

type VMAlertmanagerListMapOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlertmanagerList)(nil)).Elem()
}

func (o VMAlertmanagerListMapOutput) ToVMAlertmanagerListMapOutput() VMAlertmanagerListMapOutput {
	return o
}

func (o VMAlertmanagerListMapOutput) ToVMAlertmanagerListMapOutputWithContext(ctx context.Context) VMAlertmanagerListMapOutput {
	return o
}

func (o VMAlertmanagerListMapOutput) MapIndex(k pulumi.StringInput) VMAlertmanagerListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMAlertmanagerList {
		return vs[0].(map[string]*VMAlertmanagerList)[vs[1].(string)]
	}).(VMAlertmanagerListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerListInput)(nil)).Elem(), &VMAlertmanagerList{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerListArrayInput)(nil)).Elem(), VMAlertmanagerListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerListMapInput)(nil)).Elem(), VMAlertmanagerListMap{})
	pulumi.RegisterOutputType(VMAlertmanagerListOutput{})
	pulumi.RegisterOutputType(VMAlertmanagerListArrayOutput{})
	pulumi.RegisterOutputType(VMAlertmanagerListMapOutput{})
}
