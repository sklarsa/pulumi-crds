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

// VMAuth is the Schema for the vmauths API
type VMAuth struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     VMAuthSpecOutput        `pulumi:"spec"`
	Status   VMAuthStatusPtrOutput   `pulumi:"status"`
}

// NewVMAuth registers a new resource with the given unique name, arguments, and options.
func NewVMAuth(ctx *pulumi.Context,
	name string, args *VMAuthArgs, opts ...pulumi.ResourceOption) (*VMAuth, error) {
	if args == nil {
		args = &VMAuthArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMAuth")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMAuth
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMAuth gets an existing VMAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMAuthState, opts ...pulumi.ResourceOption) (*VMAuth, error) {
	var resource VMAuth
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMAuth resources.
type vmauthState struct {
}

type VMAuthState struct {
}

func (VMAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmauthState)(nil)).Elem()
}

type vmauthArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *VMAuthSpec        `pulumi:"spec"`
}

// The set of arguments for constructing a VMAuth resource.
type VMAuthArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     VMAuthSpecPtrInput
}

func (VMAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmauthArgs)(nil)).Elem()
}

type VMAuthInput interface {
	pulumi.Input

	ToVMAuthOutput() VMAuthOutput
	ToVMAuthOutputWithContext(ctx context.Context) VMAuthOutput
}

func (*VMAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAuth)(nil)).Elem()
}

func (i *VMAuth) ToVMAuthOutput() VMAuthOutput {
	return i.ToVMAuthOutputWithContext(context.Background())
}

func (i *VMAuth) ToVMAuthOutputWithContext(ctx context.Context) VMAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAuthOutput)
}

// VMAuthArrayInput is an input type that accepts VMAuthArray and VMAuthArrayOutput values.
// You can construct a concrete instance of `VMAuthArrayInput` via:
//
//	VMAuthArray{ VMAuthArgs{...} }
type VMAuthArrayInput interface {
	pulumi.Input

	ToVMAuthArrayOutput() VMAuthArrayOutput
	ToVMAuthArrayOutputWithContext(context.Context) VMAuthArrayOutput
}

type VMAuthArray []VMAuthInput

func (VMAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAuth)(nil)).Elem()
}

func (i VMAuthArray) ToVMAuthArrayOutput() VMAuthArrayOutput {
	return i.ToVMAuthArrayOutputWithContext(context.Background())
}

func (i VMAuthArray) ToVMAuthArrayOutputWithContext(ctx context.Context) VMAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAuthArrayOutput)
}

// VMAuthMapInput is an input type that accepts VMAuthMap and VMAuthMapOutput values.
// You can construct a concrete instance of `VMAuthMapInput` via:
//
//	VMAuthMap{ "key": VMAuthArgs{...} }
type VMAuthMapInput interface {
	pulumi.Input

	ToVMAuthMapOutput() VMAuthMapOutput
	ToVMAuthMapOutputWithContext(context.Context) VMAuthMapOutput
}

type VMAuthMap map[string]VMAuthInput

func (VMAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAuth)(nil)).Elem()
}

func (i VMAuthMap) ToVMAuthMapOutput() VMAuthMapOutput {
	return i.ToVMAuthMapOutputWithContext(context.Background())
}

func (i VMAuthMap) ToVMAuthMapOutputWithContext(ctx context.Context) VMAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAuthMapOutput)
}

type VMAuthOutput struct{ *pulumi.OutputState }

func (VMAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAuth)(nil)).Elem()
}

func (o VMAuthOutput) ToVMAuthOutput() VMAuthOutput {
	return o
}

func (o VMAuthOutput) ToVMAuthOutputWithContext(ctx context.Context) VMAuthOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMAuthOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAuth) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMAuthOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAuth) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMAuthOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *VMAuth) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o VMAuthOutput) Spec() VMAuthSpecOutput {
	return o.ApplyT(func(v *VMAuth) VMAuthSpecOutput { return v.Spec }).(VMAuthSpecOutput)
}

func (o VMAuthOutput) Status() VMAuthStatusPtrOutput {
	return o.ApplyT(func(v *VMAuth) VMAuthStatusPtrOutput { return v.Status }).(VMAuthStatusPtrOutput)
}

type VMAuthArrayOutput struct{ *pulumi.OutputState }

func (VMAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAuth)(nil)).Elem()
}

func (o VMAuthArrayOutput) ToVMAuthArrayOutput() VMAuthArrayOutput {
	return o
}

func (o VMAuthArrayOutput) ToVMAuthArrayOutputWithContext(ctx context.Context) VMAuthArrayOutput {
	return o
}

func (o VMAuthArrayOutput) Index(i pulumi.IntInput) VMAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMAuth {
		return vs[0].([]*VMAuth)[vs[1].(int)]
	}).(VMAuthOutput)
}

type VMAuthMapOutput struct{ *pulumi.OutputState }

func (VMAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAuth)(nil)).Elem()
}

func (o VMAuthMapOutput) ToVMAuthMapOutput() VMAuthMapOutput {
	return o
}

func (o VMAuthMapOutput) ToVMAuthMapOutputWithContext(ctx context.Context) VMAuthMapOutput {
	return o
}

func (o VMAuthMapOutput) MapIndex(k pulumi.StringInput) VMAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMAuth {
		return vs[0].(map[string]*VMAuth)[vs[1].(string)]
	}).(VMAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMAuthInput)(nil)).Elem(), &VMAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAuthArrayInput)(nil)).Elem(), VMAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAuthMapInput)(nil)).Elem(), VMAuthMap{})
	pulumi.RegisterOutputType(VMAuthOutput{})
	pulumi.RegisterOutputType(VMAuthArrayOutput{})
	pulumi.RegisterOutputType(VMAuthMapOutput{})
}