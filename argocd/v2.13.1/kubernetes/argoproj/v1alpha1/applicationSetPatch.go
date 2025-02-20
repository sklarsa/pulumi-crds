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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
type ApplicationSetPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput    `pulumi:"metadata"`
	Spec     ApplicationSetSpecPatchPtrOutput   `pulumi:"spec"`
	Status   ApplicationSetStatusPatchPtrOutput `pulumi:"status"`
}

// NewApplicationSetPatch registers a new resource with the given unique name, arguments, and options.
func NewApplicationSetPatch(ctx *pulumi.Context,
	name string, args *ApplicationSetPatchArgs, opts ...pulumi.ResourceOption) (*ApplicationSetPatch, error) {
	if args == nil {
		args = &ApplicationSetPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("argoproj.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ApplicationSet")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationSetPatch
	err := ctx.RegisterResource("kubernetes:argoproj.io/v1alpha1:ApplicationSetPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSetPatch gets an existing ApplicationSetPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSetPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSetPatchState, opts ...pulumi.ResourceOption) (*ApplicationSetPatch, error) {
	var resource ApplicationSetPatch
	err := ctx.ReadResource("kubernetes:argoproj.io/v1alpha1:ApplicationSetPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSetPatch resources.
type applicationSetPatchState struct {
}

type ApplicationSetPatchState struct {
}

func (ApplicationSetPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSetPatchState)(nil)).Elem()
}

type applicationSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch  `pulumi:"metadata"`
	Spec     *ApplicationSetSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ApplicationSetPatch resource.
type ApplicationSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     ApplicationSetSpecPatchPtrInput
}

func (ApplicationSetPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSetPatchArgs)(nil)).Elem()
}

type ApplicationSetPatchInput interface {
	pulumi.Input

	ToApplicationSetPatchOutput() ApplicationSetPatchOutput
	ToApplicationSetPatchOutputWithContext(ctx context.Context) ApplicationSetPatchOutput
}

func (*ApplicationSetPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSetPatch)(nil)).Elem()
}

func (i *ApplicationSetPatch) ToApplicationSetPatchOutput() ApplicationSetPatchOutput {
	return i.ToApplicationSetPatchOutputWithContext(context.Background())
}

func (i *ApplicationSetPatch) ToApplicationSetPatchOutputWithContext(ctx context.Context) ApplicationSetPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSetPatchOutput)
}

// ApplicationSetPatchArrayInput is an input type that accepts ApplicationSetPatchArray and ApplicationSetPatchArrayOutput values.
// You can construct a concrete instance of `ApplicationSetPatchArrayInput` via:
//
//	ApplicationSetPatchArray{ ApplicationSetPatchArgs{...} }
type ApplicationSetPatchArrayInput interface {
	pulumi.Input

	ToApplicationSetPatchArrayOutput() ApplicationSetPatchArrayOutput
	ToApplicationSetPatchArrayOutputWithContext(context.Context) ApplicationSetPatchArrayOutput
}

type ApplicationSetPatchArray []ApplicationSetPatchInput

func (ApplicationSetPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSetPatch)(nil)).Elem()
}

func (i ApplicationSetPatchArray) ToApplicationSetPatchArrayOutput() ApplicationSetPatchArrayOutput {
	return i.ToApplicationSetPatchArrayOutputWithContext(context.Background())
}

func (i ApplicationSetPatchArray) ToApplicationSetPatchArrayOutputWithContext(ctx context.Context) ApplicationSetPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSetPatchArrayOutput)
}

// ApplicationSetPatchMapInput is an input type that accepts ApplicationSetPatchMap and ApplicationSetPatchMapOutput values.
// You can construct a concrete instance of `ApplicationSetPatchMapInput` via:
//
//	ApplicationSetPatchMap{ "key": ApplicationSetPatchArgs{...} }
type ApplicationSetPatchMapInput interface {
	pulumi.Input

	ToApplicationSetPatchMapOutput() ApplicationSetPatchMapOutput
	ToApplicationSetPatchMapOutputWithContext(context.Context) ApplicationSetPatchMapOutput
}

type ApplicationSetPatchMap map[string]ApplicationSetPatchInput

func (ApplicationSetPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSetPatch)(nil)).Elem()
}

func (i ApplicationSetPatchMap) ToApplicationSetPatchMapOutput() ApplicationSetPatchMapOutput {
	return i.ToApplicationSetPatchMapOutputWithContext(context.Background())
}

func (i ApplicationSetPatchMap) ToApplicationSetPatchMapOutputWithContext(ctx context.Context) ApplicationSetPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSetPatchMapOutput)
}

type ApplicationSetPatchOutput struct{ *pulumi.OutputState }

func (ApplicationSetPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSetPatch)(nil)).Elem()
}

func (o ApplicationSetPatchOutput) ToApplicationSetPatchOutput() ApplicationSetPatchOutput {
	return o
}

func (o ApplicationSetPatchOutput) ToApplicationSetPatchOutputWithContext(ctx context.Context) ApplicationSetPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ApplicationSetPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSetPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApplicationSetPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSetPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ApplicationSetPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationSetPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o ApplicationSetPatchOutput) Spec() ApplicationSetSpecPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationSetPatch) ApplicationSetSpecPatchPtrOutput { return v.Spec }).(ApplicationSetSpecPatchPtrOutput)
}

func (o ApplicationSetPatchOutput) Status() ApplicationSetStatusPatchPtrOutput {
	return o.ApplyT(func(v *ApplicationSetPatch) ApplicationSetStatusPatchPtrOutput { return v.Status }).(ApplicationSetStatusPatchPtrOutput)
}

type ApplicationSetPatchArrayOutput struct{ *pulumi.OutputState }

func (ApplicationSetPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationSetPatch)(nil)).Elem()
}

func (o ApplicationSetPatchArrayOutput) ToApplicationSetPatchArrayOutput() ApplicationSetPatchArrayOutput {
	return o
}

func (o ApplicationSetPatchArrayOutput) ToApplicationSetPatchArrayOutputWithContext(ctx context.Context) ApplicationSetPatchArrayOutput {
	return o
}

func (o ApplicationSetPatchArrayOutput) Index(i pulumi.IntInput) ApplicationSetPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationSetPatch {
		return vs[0].([]*ApplicationSetPatch)[vs[1].(int)]
	}).(ApplicationSetPatchOutput)
}

type ApplicationSetPatchMapOutput struct{ *pulumi.OutputState }

func (ApplicationSetPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationSetPatch)(nil)).Elem()
}

func (o ApplicationSetPatchMapOutput) ToApplicationSetPatchMapOutput() ApplicationSetPatchMapOutput {
	return o
}

func (o ApplicationSetPatchMapOutput) ToApplicationSetPatchMapOutputWithContext(ctx context.Context) ApplicationSetPatchMapOutput {
	return o
}

func (o ApplicationSetPatchMapOutput) MapIndex(k pulumi.StringInput) ApplicationSetPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationSetPatch {
		return vs[0].(map[string]*ApplicationSetPatch)[vs[1].(string)]
	}).(ApplicationSetPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSetPatchInput)(nil)).Elem(), &ApplicationSetPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSetPatchArrayInput)(nil)).Elem(), ApplicationSetPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSetPatchMapInput)(nil)).Elem(), ApplicationSetPatchMap{})
	pulumi.RegisterOutputType(ApplicationSetPatchOutput{})
	pulumi.RegisterOutputType(ApplicationSetPatchArrayOutput{})
	pulumi.RegisterOutputType(ApplicationSetPatchMapOutput{})
}
