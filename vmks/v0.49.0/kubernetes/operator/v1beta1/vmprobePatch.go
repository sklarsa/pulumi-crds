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
// VMProbe defines a probe for targets, that will be executed with prober,
// like blackbox exporter.
// It helps to monitor reachability of target with various checks.
type VMProbePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     VMProbeSpecPatchPtrOutput       `pulumi:"spec"`
	Status   VMProbeStatusPatchPtrOutput     `pulumi:"status"`
}

// NewVMProbePatch registers a new resource with the given unique name, arguments, and options.
func NewVMProbePatch(ctx *pulumi.Context,
	name string, args *VMProbePatchArgs, opts ...pulumi.ResourceOption) (*VMProbePatch, error) {
	if args == nil {
		args = &VMProbePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMProbe")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMProbePatch
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMProbePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMProbePatch gets an existing VMProbePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMProbePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMProbePatchState, opts ...pulumi.ResourceOption) (*VMProbePatch, error) {
	var resource VMProbePatch
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMProbePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMProbePatch resources.
type vmprobePatchState struct {
}

type VMProbePatchState struct {
}

func (VMProbePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmprobePatchState)(nil)).Elem()
}

type vmprobePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *VMProbeSpecPatch       `pulumi:"spec"`
}

// The set of arguments for constructing a VMProbePatch resource.
type VMProbePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     VMProbeSpecPatchPtrInput
}

func (VMProbePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmprobePatchArgs)(nil)).Elem()
}

type VMProbePatchInput interface {
	pulumi.Input

	ToVMProbePatchOutput() VMProbePatchOutput
	ToVMProbePatchOutputWithContext(ctx context.Context) VMProbePatchOutput
}

func (*VMProbePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VMProbePatch)(nil)).Elem()
}

func (i *VMProbePatch) ToVMProbePatchOutput() VMProbePatchOutput {
	return i.ToVMProbePatchOutputWithContext(context.Background())
}

func (i *VMProbePatch) ToVMProbePatchOutputWithContext(ctx context.Context) VMProbePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMProbePatchOutput)
}

// VMProbePatchArrayInput is an input type that accepts VMProbePatchArray and VMProbePatchArrayOutput values.
// You can construct a concrete instance of `VMProbePatchArrayInput` via:
//
//	VMProbePatchArray{ VMProbePatchArgs{...} }
type VMProbePatchArrayInput interface {
	pulumi.Input

	ToVMProbePatchArrayOutput() VMProbePatchArrayOutput
	ToVMProbePatchArrayOutputWithContext(context.Context) VMProbePatchArrayOutput
}

type VMProbePatchArray []VMProbePatchInput

func (VMProbePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMProbePatch)(nil)).Elem()
}

func (i VMProbePatchArray) ToVMProbePatchArrayOutput() VMProbePatchArrayOutput {
	return i.ToVMProbePatchArrayOutputWithContext(context.Background())
}

func (i VMProbePatchArray) ToVMProbePatchArrayOutputWithContext(ctx context.Context) VMProbePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMProbePatchArrayOutput)
}

// VMProbePatchMapInput is an input type that accepts VMProbePatchMap and VMProbePatchMapOutput values.
// You can construct a concrete instance of `VMProbePatchMapInput` via:
//
//	VMProbePatchMap{ "key": VMProbePatchArgs{...} }
type VMProbePatchMapInput interface {
	pulumi.Input

	ToVMProbePatchMapOutput() VMProbePatchMapOutput
	ToVMProbePatchMapOutputWithContext(context.Context) VMProbePatchMapOutput
}

type VMProbePatchMap map[string]VMProbePatchInput

func (VMProbePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMProbePatch)(nil)).Elem()
}

func (i VMProbePatchMap) ToVMProbePatchMapOutput() VMProbePatchMapOutput {
	return i.ToVMProbePatchMapOutputWithContext(context.Background())
}

func (i VMProbePatchMap) ToVMProbePatchMapOutputWithContext(ctx context.Context) VMProbePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMProbePatchMapOutput)
}

type VMProbePatchOutput struct{ *pulumi.OutputState }

func (VMProbePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMProbePatch)(nil)).Elem()
}

func (o VMProbePatchOutput) ToVMProbePatchOutput() VMProbePatchOutput {
	return o
}

func (o VMProbePatchOutput) ToVMProbePatchOutputWithContext(ctx context.Context) VMProbePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMProbePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMProbePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMProbePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMProbePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMProbePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VMProbePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o VMProbePatchOutput) Spec() VMProbeSpecPatchPtrOutput {
	return o.ApplyT(func(v *VMProbePatch) VMProbeSpecPatchPtrOutput { return v.Spec }).(VMProbeSpecPatchPtrOutput)
}

func (o VMProbePatchOutput) Status() VMProbeStatusPatchPtrOutput {
	return o.ApplyT(func(v *VMProbePatch) VMProbeStatusPatchPtrOutput { return v.Status }).(VMProbeStatusPatchPtrOutput)
}

type VMProbePatchArrayOutput struct{ *pulumi.OutputState }

func (VMProbePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMProbePatch)(nil)).Elem()
}

func (o VMProbePatchArrayOutput) ToVMProbePatchArrayOutput() VMProbePatchArrayOutput {
	return o
}

func (o VMProbePatchArrayOutput) ToVMProbePatchArrayOutputWithContext(ctx context.Context) VMProbePatchArrayOutput {
	return o
}

func (o VMProbePatchArrayOutput) Index(i pulumi.IntInput) VMProbePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMProbePatch {
		return vs[0].([]*VMProbePatch)[vs[1].(int)]
	}).(VMProbePatchOutput)
}

type VMProbePatchMapOutput struct{ *pulumi.OutputState }

func (VMProbePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMProbePatch)(nil)).Elem()
}

func (o VMProbePatchMapOutput) ToVMProbePatchMapOutput() VMProbePatchMapOutput {
	return o
}

func (o VMProbePatchMapOutput) ToVMProbePatchMapOutputWithContext(ctx context.Context) VMProbePatchMapOutput {
	return o
}

func (o VMProbePatchMapOutput) MapIndex(k pulumi.StringInput) VMProbePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMProbePatch {
		return vs[0].(map[string]*VMProbePatch)[vs[1].(string)]
	}).(VMProbePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMProbePatchInput)(nil)).Elem(), &VMProbePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMProbePatchArrayInput)(nil)).Elem(), VMProbePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMProbePatchMapInput)(nil)).Elem(), VMProbePatchMap{})
	pulumi.RegisterOutputType(VMProbePatchOutput{})
	pulumi.RegisterOutputType(VMProbePatchArrayOutput{})
	pulumi.RegisterOutputType(VMProbePatchMapOutput{})
}