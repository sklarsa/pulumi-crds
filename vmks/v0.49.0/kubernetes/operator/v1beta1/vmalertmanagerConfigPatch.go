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
// VMAlertmanagerConfig is the Schema for the vmalertmanagerconfigs API
type VMAlertmanagerConfigPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput          `pulumi:"metadata"`
	Spec     VMAlertmanagerConfigSpecPatchPtrOutput   `pulumi:"spec"`
	Status   VMAlertmanagerConfigStatusPatchPtrOutput `pulumi:"status"`
}

// NewVMAlertmanagerConfigPatch registers a new resource with the given unique name, arguments, and options.
func NewVMAlertmanagerConfigPatch(ctx *pulumi.Context,
	name string, args *VMAlertmanagerConfigPatchArgs, opts ...pulumi.ResourceOption) (*VMAlertmanagerConfigPatch, error) {
	if args == nil {
		args = &VMAlertmanagerConfigPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMAlertmanagerConfig")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMAlertmanagerConfigPatch
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerConfigPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMAlertmanagerConfigPatch gets an existing VMAlertmanagerConfigPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMAlertmanagerConfigPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMAlertmanagerConfigPatchState, opts ...pulumi.ResourceOption) (*VMAlertmanagerConfigPatch, error) {
	var resource VMAlertmanagerConfigPatch
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlertmanagerConfigPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMAlertmanagerConfigPatch resources.
type vmalertmanagerConfigPatchState struct {
}

type VMAlertmanagerConfigPatchState struct {
}

func (VMAlertmanagerConfigPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertmanagerConfigPatchState)(nil)).Elem()
}

type vmalertmanagerConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch        `pulumi:"metadata"`
	Spec     *VMAlertmanagerConfigSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a VMAlertmanagerConfigPatch resource.
type VMAlertmanagerConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     VMAlertmanagerConfigSpecPatchPtrInput
}

func (VMAlertmanagerConfigPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertmanagerConfigPatchArgs)(nil)).Elem()
}

type VMAlertmanagerConfigPatchInput interface {
	pulumi.Input

	ToVMAlertmanagerConfigPatchOutput() VMAlertmanagerConfigPatchOutput
	ToVMAlertmanagerConfigPatchOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchOutput
}

func (*VMAlertmanagerConfigPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (i *VMAlertmanagerConfigPatch) ToVMAlertmanagerConfigPatchOutput() VMAlertmanagerConfigPatchOutput {
	return i.ToVMAlertmanagerConfigPatchOutputWithContext(context.Background())
}

func (i *VMAlertmanagerConfigPatch) ToVMAlertmanagerConfigPatchOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerConfigPatchOutput)
}

// VMAlertmanagerConfigPatchArrayInput is an input type that accepts VMAlertmanagerConfigPatchArray and VMAlertmanagerConfigPatchArrayOutput values.
// You can construct a concrete instance of `VMAlertmanagerConfigPatchArrayInput` via:
//
//	VMAlertmanagerConfigPatchArray{ VMAlertmanagerConfigPatchArgs{...} }
type VMAlertmanagerConfigPatchArrayInput interface {
	pulumi.Input

	ToVMAlertmanagerConfigPatchArrayOutput() VMAlertmanagerConfigPatchArrayOutput
	ToVMAlertmanagerConfigPatchArrayOutputWithContext(context.Context) VMAlertmanagerConfigPatchArrayOutput
}

type VMAlertmanagerConfigPatchArray []VMAlertmanagerConfigPatchInput

func (VMAlertmanagerConfigPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (i VMAlertmanagerConfigPatchArray) ToVMAlertmanagerConfigPatchArrayOutput() VMAlertmanagerConfigPatchArrayOutput {
	return i.ToVMAlertmanagerConfigPatchArrayOutputWithContext(context.Background())
}

func (i VMAlertmanagerConfigPatchArray) ToVMAlertmanagerConfigPatchArrayOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerConfigPatchArrayOutput)
}

// VMAlertmanagerConfigPatchMapInput is an input type that accepts VMAlertmanagerConfigPatchMap and VMAlertmanagerConfigPatchMapOutput values.
// You can construct a concrete instance of `VMAlertmanagerConfigPatchMapInput` via:
//
//	VMAlertmanagerConfigPatchMap{ "key": VMAlertmanagerConfigPatchArgs{...} }
type VMAlertmanagerConfigPatchMapInput interface {
	pulumi.Input

	ToVMAlertmanagerConfigPatchMapOutput() VMAlertmanagerConfigPatchMapOutput
	ToVMAlertmanagerConfigPatchMapOutputWithContext(context.Context) VMAlertmanagerConfigPatchMapOutput
}

type VMAlertmanagerConfigPatchMap map[string]VMAlertmanagerConfigPatchInput

func (VMAlertmanagerConfigPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (i VMAlertmanagerConfigPatchMap) ToVMAlertmanagerConfigPatchMapOutput() VMAlertmanagerConfigPatchMapOutput {
	return i.ToVMAlertmanagerConfigPatchMapOutputWithContext(context.Background())
}

func (i VMAlertmanagerConfigPatchMap) ToVMAlertmanagerConfigPatchMapOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertmanagerConfigPatchMapOutput)
}

type VMAlertmanagerConfigPatchOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerConfigPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (o VMAlertmanagerConfigPatchOutput) ToVMAlertmanagerConfigPatchOutput() VMAlertmanagerConfigPatchOutput {
	return o
}

func (o VMAlertmanagerConfigPatchOutput) ToVMAlertmanagerConfigPatchOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMAlertmanagerConfigPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMAlertmanagerConfigPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMAlertmanagerConfigPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMAlertmanagerConfigPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMAlertmanagerConfigPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VMAlertmanagerConfigPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o VMAlertmanagerConfigPatchOutput) Spec() VMAlertmanagerConfigSpecPatchPtrOutput {
	return o.ApplyT(func(v *VMAlertmanagerConfigPatch) VMAlertmanagerConfigSpecPatchPtrOutput { return v.Spec }).(VMAlertmanagerConfigSpecPatchPtrOutput)
}

func (o VMAlertmanagerConfigPatchOutput) Status() VMAlertmanagerConfigStatusPatchPtrOutput {
	return o.ApplyT(func(v *VMAlertmanagerConfigPatch) VMAlertmanagerConfigStatusPatchPtrOutput { return v.Status }).(VMAlertmanagerConfigStatusPatchPtrOutput)
}

type VMAlertmanagerConfigPatchArrayOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerConfigPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (o VMAlertmanagerConfigPatchArrayOutput) ToVMAlertmanagerConfigPatchArrayOutput() VMAlertmanagerConfigPatchArrayOutput {
	return o
}

func (o VMAlertmanagerConfigPatchArrayOutput) ToVMAlertmanagerConfigPatchArrayOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchArrayOutput {
	return o
}

func (o VMAlertmanagerConfigPatchArrayOutput) Index(i pulumi.IntInput) VMAlertmanagerConfigPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMAlertmanagerConfigPatch {
		return vs[0].([]*VMAlertmanagerConfigPatch)[vs[1].(int)]
	}).(VMAlertmanagerConfigPatchOutput)
}

type VMAlertmanagerConfigPatchMapOutput struct{ *pulumi.OutputState }

func (VMAlertmanagerConfigPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlertmanagerConfigPatch)(nil)).Elem()
}

func (o VMAlertmanagerConfigPatchMapOutput) ToVMAlertmanagerConfigPatchMapOutput() VMAlertmanagerConfigPatchMapOutput {
	return o
}

func (o VMAlertmanagerConfigPatchMapOutput) ToVMAlertmanagerConfigPatchMapOutputWithContext(ctx context.Context) VMAlertmanagerConfigPatchMapOutput {
	return o
}

func (o VMAlertmanagerConfigPatchMapOutput) MapIndex(k pulumi.StringInput) VMAlertmanagerConfigPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMAlertmanagerConfigPatch {
		return vs[0].(map[string]*VMAlertmanagerConfigPatch)[vs[1].(string)]
	}).(VMAlertmanagerConfigPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerConfigPatchInput)(nil)).Elem(), &VMAlertmanagerConfigPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerConfigPatchArrayInput)(nil)).Elem(), VMAlertmanagerConfigPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertmanagerConfigPatchMapInput)(nil)).Elem(), VMAlertmanagerConfigPatchMap{})
	pulumi.RegisterOutputType(VMAlertmanagerConfigPatchOutput{})
	pulumi.RegisterOutputType(VMAlertmanagerConfigPatchArrayOutput{})
	pulumi.RegisterOutputType(VMAlertmanagerConfigPatchMapOutput{})
}