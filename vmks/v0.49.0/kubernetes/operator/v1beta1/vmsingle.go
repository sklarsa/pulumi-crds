// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VMSingle  is fast, cost-effective and scalable time-series database.
type VMSingle struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     VMSingleSpecOutput      `pulumi:"spec"`
	Status   VMSingleStatusPtrOutput `pulumi:"status"`
}

// NewVMSingle registers a new resource with the given unique name, arguments, and options.
func NewVMSingle(ctx *pulumi.Context,
	name string, args *VMSingleArgs, opts ...pulumi.ResourceOption) (*VMSingle, error) {
	if args == nil {
		args = &VMSingleArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMSingle")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMSingle
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMSingle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMSingle gets an existing VMSingle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMSingle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMSingleState, opts ...pulumi.ResourceOption) (*VMSingle, error) {
	var resource VMSingle
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMSingle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMSingle resources.
type vmsingleState struct {
}

type VMSingleState struct {
}

func (VMSingleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmsingleState)(nil)).Elem()
}

type vmsingleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *VMSingleSpec      `pulumi:"spec"`
}

// The set of arguments for constructing a VMSingle resource.
type VMSingleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     VMSingleSpecPtrInput
}

func (VMSingleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmsingleArgs)(nil)).Elem()
}

type VMSingleInput interface {
	pulumi.Input

	ToVMSingleOutput() VMSingleOutput
	ToVMSingleOutputWithContext(ctx context.Context) VMSingleOutput
}

func (*VMSingle) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSingle)(nil)).Elem()
}

func (i *VMSingle) ToVMSingleOutput() VMSingleOutput {
	return i.ToVMSingleOutputWithContext(context.Background())
}

func (i *VMSingle) ToVMSingleOutputWithContext(ctx context.Context) VMSingleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSingleOutput)
}

// VMSingleArrayInput is an input type that accepts VMSingleArray and VMSingleArrayOutput values.
// You can construct a concrete instance of `VMSingleArrayInput` via:
//
//	VMSingleArray{ VMSingleArgs{...} }
type VMSingleArrayInput interface {
	pulumi.Input

	ToVMSingleArrayOutput() VMSingleArrayOutput
	ToVMSingleArrayOutputWithContext(context.Context) VMSingleArrayOutput
}

type VMSingleArray []VMSingleInput

func (VMSingleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMSingle)(nil)).Elem()
}

func (i VMSingleArray) ToVMSingleArrayOutput() VMSingleArrayOutput {
	return i.ToVMSingleArrayOutputWithContext(context.Background())
}

func (i VMSingleArray) ToVMSingleArrayOutputWithContext(ctx context.Context) VMSingleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSingleArrayOutput)
}

// VMSingleMapInput is an input type that accepts VMSingleMap and VMSingleMapOutput values.
// You can construct a concrete instance of `VMSingleMapInput` via:
//
//	VMSingleMap{ "key": VMSingleArgs{...} }
type VMSingleMapInput interface {
	pulumi.Input

	ToVMSingleMapOutput() VMSingleMapOutput
	ToVMSingleMapOutputWithContext(context.Context) VMSingleMapOutput
}

type VMSingleMap map[string]VMSingleInput

func (VMSingleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMSingle)(nil)).Elem()
}

func (i VMSingleMap) ToVMSingleMapOutput() VMSingleMapOutput {
	return i.ToVMSingleMapOutputWithContext(context.Background())
}

func (i VMSingleMap) ToVMSingleMapOutputWithContext(ctx context.Context) VMSingleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSingleMapOutput)
}

type VMSingleOutput struct{ *pulumi.OutputState }

func (VMSingleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSingle)(nil)).Elem()
}

func (o VMSingleOutput) ToVMSingleOutput() VMSingleOutput {
	return o
}

func (o VMSingleOutput) ToVMSingleOutputWithContext(ctx context.Context) VMSingleOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMSingleOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMSingle) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMSingleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VMSingle) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMSingleOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *VMSingle) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o VMSingleOutput) Spec() VMSingleSpecOutput {
	return o.ApplyT(func(v *VMSingle) VMSingleSpecOutput { return v.Spec }).(VMSingleSpecOutput)
}

func (o VMSingleOutput) Status() VMSingleStatusPtrOutput {
	return o.ApplyT(func(v *VMSingle) VMSingleStatusPtrOutput { return v.Status }).(VMSingleStatusPtrOutput)
}

type VMSingleArrayOutput struct{ *pulumi.OutputState }

func (VMSingleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMSingle)(nil)).Elem()
}

func (o VMSingleArrayOutput) ToVMSingleArrayOutput() VMSingleArrayOutput {
	return o
}

func (o VMSingleArrayOutput) ToVMSingleArrayOutputWithContext(ctx context.Context) VMSingleArrayOutput {
	return o
}

func (o VMSingleArrayOutput) Index(i pulumi.IntInput) VMSingleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMSingle {
		return vs[0].([]*VMSingle)[vs[1].(int)]
	}).(VMSingleOutput)
}

type VMSingleMapOutput struct{ *pulumi.OutputState }

func (VMSingleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMSingle)(nil)).Elem()
}

func (o VMSingleMapOutput) ToVMSingleMapOutput() VMSingleMapOutput {
	return o
}

func (o VMSingleMapOutput) ToVMSingleMapOutputWithContext(ctx context.Context) VMSingleMapOutput {
	return o
}

func (o VMSingleMapOutput) MapIndex(k pulumi.StringInput) VMSingleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMSingle {
		return vs[0].(map[string]*VMSingle)[vs[1].(string)]
	}).(VMSingleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMSingleInput)(nil)).Elem(), &VMSingle{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSingleArrayInput)(nil)).Elem(), VMSingleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSingleMapInput)(nil)).Elem(), VMSingleMap{})
	pulumi.RegisterOutputType(VMSingleOutput{})
	pulumi.RegisterOutputType(VMSingleArrayOutput{})
	pulumi.RegisterOutputType(VMSingleMapOutput{})
}
