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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// VMUser is the Schema for the vmusers API
type VMUserPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     VMUserSpecPatchPtrOutput        `pulumi:"spec"`
	Status   VMUserStatusPatchPtrOutput      `pulumi:"status"`
}

// NewVMUserPatch registers a new resource with the given unique name, arguments, and options.
func NewVMUserPatch(ctx *pulumi.Context,
	name string, args *VMUserPatchArgs, opts ...pulumi.ResourceOption) (*VMUserPatch, error) {
	if args == nil {
		args = &VMUserPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMUser")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMUserPatch
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMUserPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMUserPatch gets an existing VMUserPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMUserPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMUserPatchState, opts ...pulumi.ResourceOption) (*VMUserPatch, error) {
	var resource VMUserPatch
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMUserPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMUserPatch resources.
type vmuserPatchState struct {
}

type VMUserPatchState struct {
}

func (VMUserPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmuserPatchState)(nil)).Elem()
}

type vmuserPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *VMUserSpecPatch        `pulumi:"spec"`
}

// The set of arguments for constructing a VMUserPatch resource.
type VMUserPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     VMUserSpecPatchPtrInput
}

func (VMUserPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmuserPatchArgs)(nil)).Elem()
}

type VMUserPatchInput interface {
	pulumi.Input

	ToVMUserPatchOutput() VMUserPatchOutput
	ToVMUserPatchOutputWithContext(ctx context.Context) VMUserPatchOutput
}

func (*VMUserPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VMUserPatch)(nil)).Elem()
}

func (i *VMUserPatch) ToVMUserPatchOutput() VMUserPatchOutput {
	return i.ToVMUserPatchOutputWithContext(context.Background())
}

func (i *VMUserPatch) ToVMUserPatchOutputWithContext(ctx context.Context) VMUserPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMUserPatchOutput)
}

// VMUserPatchArrayInput is an input type that accepts VMUserPatchArray and VMUserPatchArrayOutput values.
// You can construct a concrete instance of `VMUserPatchArrayInput` via:
//
//	VMUserPatchArray{ VMUserPatchArgs{...} }
type VMUserPatchArrayInput interface {
	pulumi.Input

	ToVMUserPatchArrayOutput() VMUserPatchArrayOutput
	ToVMUserPatchArrayOutputWithContext(context.Context) VMUserPatchArrayOutput
}

type VMUserPatchArray []VMUserPatchInput

func (VMUserPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMUserPatch)(nil)).Elem()
}

func (i VMUserPatchArray) ToVMUserPatchArrayOutput() VMUserPatchArrayOutput {
	return i.ToVMUserPatchArrayOutputWithContext(context.Background())
}

func (i VMUserPatchArray) ToVMUserPatchArrayOutputWithContext(ctx context.Context) VMUserPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMUserPatchArrayOutput)
}

// VMUserPatchMapInput is an input type that accepts VMUserPatchMap and VMUserPatchMapOutput values.
// You can construct a concrete instance of `VMUserPatchMapInput` via:
//
//	VMUserPatchMap{ "key": VMUserPatchArgs{...} }
type VMUserPatchMapInput interface {
	pulumi.Input

	ToVMUserPatchMapOutput() VMUserPatchMapOutput
	ToVMUserPatchMapOutputWithContext(context.Context) VMUserPatchMapOutput
}

type VMUserPatchMap map[string]VMUserPatchInput

func (VMUserPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMUserPatch)(nil)).Elem()
}

func (i VMUserPatchMap) ToVMUserPatchMapOutput() VMUserPatchMapOutput {
	return i.ToVMUserPatchMapOutputWithContext(context.Background())
}

func (i VMUserPatchMap) ToVMUserPatchMapOutputWithContext(ctx context.Context) VMUserPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMUserPatchMapOutput)
}

type VMUserPatchOutput struct{ *pulumi.OutputState }

func (VMUserPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMUserPatch)(nil)).Elem()
}

func (o VMUserPatchOutput) ToVMUserPatchOutput() VMUserPatchOutput {
	return o
}

func (o VMUserPatchOutput) ToVMUserPatchOutputWithContext(ctx context.Context) VMUserPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMUserPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMUserPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMUserPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMUserPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMUserPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VMUserPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o VMUserPatchOutput) Spec() VMUserSpecPatchPtrOutput {
	return o.ApplyT(func(v *VMUserPatch) VMUserSpecPatchPtrOutput { return v.Spec }).(VMUserSpecPatchPtrOutput)
}

func (o VMUserPatchOutput) Status() VMUserStatusPatchPtrOutput {
	return o.ApplyT(func(v *VMUserPatch) VMUserStatusPatchPtrOutput { return v.Status }).(VMUserStatusPatchPtrOutput)
}

type VMUserPatchArrayOutput struct{ *pulumi.OutputState }

func (VMUserPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMUserPatch)(nil)).Elem()
}

func (o VMUserPatchArrayOutput) ToVMUserPatchArrayOutput() VMUserPatchArrayOutput {
	return o
}

func (o VMUserPatchArrayOutput) ToVMUserPatchArrayOutputWithContext(ctx context.Context) VMUserPatchArrayOutput {
	return o
}

func (o VMUserPatchArrayOutput) Index(i pulumi.IntInput) VMUserPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMUserPatch {
		return vs[0].([]*VMUserPatch)[vs[1].(int)]
	}).(VMUserPatchOutput)
}

type VMUserPatchMapOutput struct{ *pulumi.OutputState }

func (VMUserPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMUserPatch)(nil)).Elem()
}

func (o VMUserPatchMapOutput) ToVMUserPatchMapOutput() VMUserPatchMapOutput {
	return o
}

func (o VMUserPatchMapOutput) ToVMUserPatchMapOutputWithContext(ctx context.Context) VMUserPatchMapOutput {
	return o
}

func (o VMUserPatchMapOutput) MapIndex(k pulumi.StringInput) VMUserPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMUserPatch {
		return vs[0].(map[string]*VMUserPatch)[vs[1].(string)]
	}).(VMUserPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMUserPatchInput)(nil)).Elem(), &VMUserPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMUserPatchArrayInput)(nil)).Elem(), VMUserPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMUserPatchMapInput)(nil)).Elem(), VMUserPatchMap{})
	pulumi.RegisterOutputType(VMUserPatchOutput{})
	pulumi.RegisterOutputType(VMUserPatchArrayOutput{})
	pulumi.RegisterOutputType(VMUserPatchMapOutput{})
}
