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
// VMPodScrape is scrape configuration for pods,
// it generates vmagent's config for scraping pod targets
// based on selectors.
type VMPodScrapePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	Spec     VMPodScrapeSpecPatchPtrOutput   `pulumi:"spec"`
	Status   VMPodScrapeStatusPatchPtrOutput `pulumi:"status"`
}

// NewVMPodScrapePatch registers a new resource with the given unique name, arguments, and options.
func NewVMPodScrapePatch(ctx *pulumi.Context,
	name string, args *VMPodScrapePatchArgs, opts ...pulumi.ResourceOption) (*VMPodScrapePatch, error) {
	if args == nil {
		args = &VMPodScrapePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMPodScrape")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMPodScrapePatch
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMPodScrapePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMPodScrapePatch gets an existing VMPodScrapePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMPodScrapePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMPodScrapePatchState, opts ...pulumi.ResourceOption) (*VMPodScrapePatch, error) {
	var resource VMPodScrapePatch
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMPodScrapePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMPodScrapePatch resources.
type vmpodScrapePatchState struct {
}

type VMPodScrapePatchState struct {
}

func (VMPodScrapePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmpodScrapePatchState)(nil)).Elem()
}

type vmpodScrapePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	Spec     *VMPodScrapeSpecPatch   `pulumi:"spec"`
}

// The set of arguments for constructing a VMPodScrapePatch resource.
type VMPodScrapePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     VMPodScrapeSpecPatchPtrInput
}

func (VMPodScrapePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmpodScrapePatchArgs)(nil)).Elem()
}

type VMPodScrapePatchInput interface {
	pulumi.Input

	ToVMPodScrapePatchOutput() VMPodScrapePatchOutput
	ToVMPodScrapePatchOutputWithContext(ctx context.Context) VMPodScrapePatchOutput
}

func (*VMPodScrapePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**VMPodScrapePatch)(nil)).Elem()
}

func (i *VMPodScrapePatch) ToVMPodScrapePatchOutput() VMPodScrapePatchOutput {
	return i.ToVMPodScrapePatchOutputWithContext(context.Background())
}

func (i *VMPodScrapePatch) ToVMPodScrapePatchOutputWithContext(ctx context.Context) VMPodScrapePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMPodScrapePatchOutput)
}

// VMPodScrapePatchArrayInput is an input type that accepts VMPodScrapePatchArray and VMPodScrapePatchArrayOutput values.
// You can construct a concrete instance of `VMPodScrapePatchArrayInput` via:
//
//	VMPodScrapePatchArray{ VMPodScrapePatchArgs{...} }
type VMPodScrapePatchArrayInput interface {
	pulumi.Input

	ToVMPodScrapePatchArrayOutput() VMPodScrapePatchArrayOutput
	ToVMPodScrapePatchArrayOutputWithContext(context.Context) VMPodScrapePatchArrayOutput
}

type VMPodScrapePatchArray []VMPodScrapePatchInput

func (VMPodScrapePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMPodScrapePatch)(nil)).Elem()
}

func (i VMPodScrapePatchArray) ToVMPodScrapePatchArrayOutput() VMPodScrapePatchArrayOutput {
	return i.ToVMPodScrapePatchArrayOutputWithContext(context.Background())
}

func (i VMPodScrapePatchArray) ToVMPodScrapePatchArrayOutputWithContext(ctx context.Context) VMPodScrapePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMPodScrapePatchArrayOutput)
}

// VMPodScrapePatchMapInput is an input type that accepts VMPodScrapePatchMap and VMPodScrapePatchMapOutput values.
// You can construct a concrete instance of `VMPodScrapePatchMapInput` via:
//
//	VMPodScrapePatchMap{ "key": VMPodScrapePatchArgs{...} }
type VMPodScrapePatchMapInput interface {
	pulumi.Input

	ToVMPodScrapePatchMapOutput() VMPodScrapePatchMapOutput
	ToVMPodScrapePatchMapOutputWithContext(context.Context) VMPodScrapePatchMapOutput
}

type VMPodScrapePatchMap map[string]VMPodScrapePatchInput

func (VMPodScrapePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMPodScrapePatch)(nil)).Elem()
}

func (i VMPodScrapePatchMap) ToVMPodScrapePatchMapOutput() VMPodScrapePatchMapOutput {
	return i.ToVMPodScrapePatchMapOutputWithContext(context.Background())
}

func (i VMPodScrapePatchMap) ToVMPodScrapePatchMapOutputWithContext(ctx context.Context) VMPodScrapePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMPodScrapePatchMapOutput)
}

type VMPodScrapePatchOutput struct{ *pulumi.OutputState }

func (VMPodScrapePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMPodScrapePatch)(nil)).Elem()
}

func (o VMPodScrapePatchOutput) ToVMPodScrapePatchOutput() VMPodScrapePatchOutput {
	return o
}

func (o VMPodScrapePatchOutput) ToVMPodScrapePatchOutputWithContext(ctx context.Context) VMPodScrapePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMPodScrapePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMPodScrapePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMPodScrapePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMPodScrapePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMPodScrapePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *VMPodScrapePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o VMPodScrapePatchOutput) Spec() VMPodScrapeSpecPatchPtrOutput {
	return o.ApplyT(func(v *VMPodScrapePatch) VMPodScrapeSpecPatchPtrOutput { return v.Spec }).(VMPodScrapeSpecPatchPtrOutput)
}

func (o VMPodScrapePatchOutput) Status() VMPodScrapeStatusPatchPtrOutput {
	return o.ApplyT(func(v *VMPodScrapePatch) VMPodScrapeStatusPatchPtrOutput { return v.Status }).(VMPodScrapeStatusPatchPtrOutput)
}

type VMPodScrapePatchArrayOutput struct{ *pulumi.OutputState }

func (VMPodScrapePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMPodScrapePatch)(nil)).Elem()
}

func (o VMPodScrapePatchArrayOutput) ToVMPodScrapePatchArrayOutput() VMPodScrapePatchArrayOutput {
	return o
}

func (o VMPodScrapePatchArrayOutput) ToVMPodScrapePatchArrayOutputWithContext(ctx context.Context) VMPodScrapePatchArrayOutput {
	return o
}

func (o VMPodScrapePatchArrayOutput) Index(i pulumi.IntInput) VMPodScrapePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMPodScrapePatch {
		return vs[0].([]*VMPodScrapePatch)[vs[1].(int)]
	}).(VMPodScrapePatchOutput)
}

type VMPodScrapePatchMapOutput struct{ *pulumi.OutputState }

func (VMPodScrapePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMPodScrapePatch)(nil)).Elem()
}

func (o VMPodScrapePatchMapOutput) ToVMPodScrapePatchMapOutput() VMPodScrapePatchMapOutput {
	return o
}

func (o VMPodScrapePatchMapOutput) ToVMPodScrapePatchMapOutputWithContext(ctx context.Context) VMPodScrapePatchMapOutput {
	return o
}

func (o VMPodScrapePatchMapOutput) MapIndex(k pulumi.StringInput) VMPodScrapePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMPodScrapePatch {
		return vs[0].(map[string]*VMPodScrapePatch)[vs[1].(string)]
	}).(VMPodScrapePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMPodScrapePatchInput)(nil)).Elem(), &VMPodScrapePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMPodScrapePatchArrayInput)(nil)).Elem(), VMPodScrapePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMPodScrapePatchMapInput)(nil)).Elem(), VMPodScrapePatchMap{})
	pulumi.RegisterOutputType(VMPodScrapePatchOutput{})
	pulumi.RegisterOutputType(VMPodScrapePatchArrayOutput{})
	pulumi.RegisterOutputType(VMPodScrapePatchMapOutput{})
}
