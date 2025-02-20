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
// AlertmanagerConfig configures the Prometheus Alertmanager,
// specifying how alerts should be grouped, inhibited and notified to external systems.
type AlertmanagerConfigPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput      `pulumi:"metadata"`
	Spec     AlertmanagerConfigSpecPatchPtrOutput `pulumi:"spec"`
}

// NewAlertmanagerConfigPatch registers a new resource with the given unique name, arguments, and options.
func NewAlertmanagerConfigPatch(ctx *pulumi.Context,
	name string, args *AlertmanagerConfigPatchArgs, opts ...pulumi.ResourceOption) (*AlertmanagerConfigPatch, error) {
	if args == nil {
		args = &AlertmanagerConfigPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("monitoring.coreos.com/v1alpha1")
	args.Kind = pulumi.StringPtr("AlertmanagerConfig")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AlertmanagerConfigPatch
	err := ctx.RegisterResource("kubernetes:monitoring.coreos.com/v1alpha1:AlertmanagerConfigPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertmanagerConfigPatch gets an existing AlertmanagerConfigPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertmanagerConfigPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertmanagerConfigPatchState, opts ...pulumi.ResourceOption) (*AlertmanagerConfigPatch, error) {
	var resource AlertmanagerConfigPatch
	err := ctx.ReadResource("kubernetes:monitoring.coreos.com/v1alpha1:AlertmanagerConfigPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertmanagerConfigPatch resources.
type alertmanagerConfigPatchState struct {
}

type AlertmanagerConfigPatchState struct {
}

func (AlertmanagerConfigPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertmanagerConfigPatchState)(nil)).Elem()
}

type alertmanagerConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch      `pulumi:"metadata"`
	Spec     *AlertmanagerConfigSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a AlertmanagerConfigPatch resource.
type AlertmanagerConfigPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	Spec     AlertmanagerConfigSpecPatchPtrInput
}

func (AlertmanagerConfigPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertmanagerConfigPatchArgs)(nil)).Elem()
}

type AlertmanagerConfigPatchInput interface {
	pulumi.Input

	ToAlertmanagerConfigPatchOutput() AlertmanagerConfigPatchOutput
	ToAlertmanagerConfigPatchOutputWithContext(ctx context.Context) AlertmanagerConfigPatchOutput
}

func (*AlertmanagerConfigPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertmanagerConfigPatch)(nil)).Elem()
}

func (i *AlertmanagerConfigPatch) ToAlertmanagerConfigPatchOutput() AlertmanagerConfigPatchOutput {
	return i.ToAlertmanagerConfigPatchOutputWithContext(context.Background())
}

func (i *AlertmanagerConfigPatch) ToAlertmanagerConfigPatchOutputWithContext(ctx context.Context) AlertmanagerConfigPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertmanagerConfigPatchOutput)
}

// AlertmanagerConfigPatchArrayInput is an input type that accepts AlertmanagerConfigPatchArray and AlertmanagerConfigPatchArrayOutput values.
// You can construct a concrete instance of `AlertmanagerConfigPatchArrayInput` via:
//
//	AlertmanagerConfigPatchArray{ AlertmanagerConfigPatchArgs{...} }
type AlertmanagerConfigPatchArrayInput interface {
	pulumi.Input

	ToAlertmanagerConfigPatchArrayOutput() AlertmanagerConfigPatchArrayOutput
	ToAlertmanagerConfigPatchArrayOutputWithContext(context.Context) AlertmanagerConfigPatchArrayOutput
}

type AlertmanagerConfigPatchArray []AlertmanagerConfigPatchInput

func (AlertmanagerConfigPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertmanagerConfigPatch)(nil)).Elem()
}

func (i AlertmanagerConfigPatchArray) ToAlertmanagerConfigPatchArrayOutput() AlertmanagerConfigPatchArrayOutput {
	return i.ToAlertmanagerConfigPatchArrayOutputWithContext(context.Background())
}

func (i AlertmanagerConfigPatchArray) ToAlertmanagerConfigPatchArrayOutputWithContext(ctx context.Context) AlertmanagerConfigPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertmanagerConfigPatchArrayOutput)
}

// AlertmanagerConfigPatchMapInput is an input type that accepts AlertmanagerConfigPatchMap and AlertmanagerConfigPatchMapOutput values.
// You can construct a concrete instance of `AlertmanagerConfigPatchMapInput` via:
//
//	AlertmanagerConfigPatchMap{ "key": AlertmanagerConfigPatchArgs{...} }
type AlertmanagerConfigPatchMapInput interface {
	pulumi.Input

	ToAlertmanagerConfigPatchMapOutput() AlertmanagerConfigPatchMapOutput
	ToAlertmanagerConfigPatchMapOutputWithContext(context.Context) AlertmanagerConfigPatchMapOutput
}

type AlertmanagerConfigPatchMap map[string]AlertmanagerConfigPatchInput

func (AlertmanagerConfigPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertmanagerConfigPatch)(nil)).Elem()
}

func (i AlertmanagerConfigPatchMap) ToAlertmanagerConfigPatchMapOutput() AlertmanagerConfigPatchMapOutput {
	return i.ToAlertmanagerConfigPatchMapOutputWithContext(context.Background())
}

func (i AlertmanagerConfigPatchMap) ToAlertmanagerConfigPatchMapOutputWithContext(ctx context.Context) AlertmanagerConfigPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertmanagerConfigPatchMapOutput)
}

type AlertmanagerConfigPatchOutput struct{ *pulumi.OutputState }

func (AlertmanagerConfigPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertmanagerConfigPatch)(nil)).Elem()
}

func (o AlertmanagerConfigPatchOutput) ToAlertmanagerConfigPatchOutput() AlertmanagerConfigPatchOutput {
	return o
}

func (o AlertmanagerConfigPatchOutput) ToAlertmanagerConfigPatchOutputWithContext(ctx context.Context) AlertmanagerConfigPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o AlertmanagerConfigPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertmanagerConfigPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o AlertmanagerConfigPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertmanagerConfigPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o AlertmanagerConfigPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *AlertmanagerConfigPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

func (o AlertmanagerConfigPatchOutput) Spec() AlertmanagerConfigSpecPatchPtrOutput {
	return o.ApplyT(func(v *AlertmanagerConfigPatch) AlertmanagerConfigSpecPatchPtrOutput { return v.Spec }).(AlertmanagerConfigSpecPatchPtrOutput)
}

type AlertmanagerConfigPatchArrayOutput struct{ *pulumi.OutputState }

func (AlertmanagerConfigPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertmanagerConfigPatch)(nil)).Elem()
}

func (o AlertmanagerConfigPatchArrayOutput) ToAlertmanagerConfigPatchArrayOutput() AlertmanagerConfigPatchArrayOutput {
	return o
}

func (o AlertmanagerConfigPatchArrayOutput) ToAlertmanagerConfigPatchArrayOutputWithContext(ctx context.Context) AlertmanagerConfigPatchArrayOutput {
	return o
}

func (o AlertmanagerConfigPatchArrayOutput) Index(i pulumi.IntInput) AlertmanagerConfigPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlertmanagerConfigPatch {
		return vs[0].([]*AlertmanagerConfigPatch)[vs[1].(int)]
	}).(AlertmanagerConfigPatchOutput)
}

type AlertmanagerConfigPatchMapOutput struct{ *pulumi.OutputState }

func (AlertmanagerConfigPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertmanagerConfigPatch)(nil)).Elem()
}

func (o AlertmanagerConfigPatchMapOutput) ToAlertmanagerConfigPatchMapOutput() AlertmanagerConfigPatchMapOutput {
	return o
}

func (o AlertmanagerConfigPatchMapOutput) ToAlertmanagerConfigPatchMapOutputWithContext(ctx context.Context) AlertmanagerConfigPatchMapOutput {
	return o
}

func (o AlertmanagerConfigPatchMapOutput) MapIndex(k pulumi.StringInput) AlertmanagerConfigPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlertmanagerConfigPatch {
		return vs[0].(map[string]*AlertmanagerConfigPatch)[vs[1].(string)]
	}).(AlertmanagerConfigPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertmanagerConfigPatchInput)(nil)).Elem(), &AlertmanagerConfigPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertmanagerConfigPatchArrayInput)(nil)).Elem(), AlertmanagerConfigPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertmanagerConfigPatchMapInput)(nil)).Elem(), AlertmanagerConfigPatchMap{})
	pulumi.RegisterOutputType(AlertmanagerConfigPatchOutput{})
	pulumi.RegisterOutputType(AlertmanagerConfigPatchArrayOutput{})
	pulumi.RegisterOutputType(AlertmanagerConfigPatchMapOutput{})
}
