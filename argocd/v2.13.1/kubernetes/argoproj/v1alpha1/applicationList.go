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

// ApplicationList is a list of Application
type ApplicationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of applications. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ApplicationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewApplicationList registers a new resource with the given unique name, arguments, and options.
func NewApplicationList(ctx *pulumi.Context,
	name string, args *ApplicationListArgs, opts ...pulumi.ResourceOption) (*ApplicationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("argoproj.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ApplicationList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationList
	err := ctx.RegisterResource("kubernetes:argoproj.io/v1alpha1:ApplicationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationList gets an existing ApplicationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationListState, opts ...pulumi.ResourceOption) (*ApplicationList, error) {
	var resource ApplicationList
	err := ctx.ReadResource("kubernetes:argoproj.io/v1alpha1:ApplicationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationList resources.
type applicationListState struct {
}

type ApplicationListState struct {
}

func (ApplicationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationListState)(nil)).Elem()
}

type applicationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of applications. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []ApplicationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ApplicationList resource.
type ApplicationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of applications. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items ApplicationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ApplicationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationListArgs)(nil)).Elem()
}

type ApplicationListInput interface {
	pulumi.Input

	ToApplicationListOutput() ApplicationListOutput
	ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput
}

func (*ApplicationList) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationList)(nil)).Elem()
}

func (i *ApplicationList) ToApplicationListOutput() ApplicationListOutput {
	return i.ToApplicationListOutputWithContext(context.Background())
}

func (i *ApplicationList) ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListOutput)
}

// ApplicationListArrayInput is an input type that accepts ApplicationListArray and ApplicationListArrayOutput values.
// You can construct a concrete instance of `ApplicationListArrayInput` via:
//
//	ApplicationListArray{ ApplicationListArgs{...} }
type ApplicationListArrayInput interface {
	pulumi.Input

	ToApplicationListArrayOutput() ApplicationListArrayOutput
	ToApplicationListArrayOutputWithContext(context.Context) ApplicationListArrayOutput
}

type ApplicationListArray []ApplicationListInput

func (ApplicationListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationList)(nil)).Elem()
}

func (i ApplicationListArray) ToApplicationListArrayOutput() ApplicationListArrayOutput {
	return i.ToApplicationListArrayOutputWithContext(context.Background())
}

func (i ApplicationListArray) ToApplicationListArrayOutputWithContext(ctx context.Context) ApplicationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListArrayOutput)
}

// ApplicationListMapInput is an input type that accepts ApplicationListMap and ApplicationListMapOutput values.
// You can construct a concrete instance of `ApplicationListMapInput` via:
//
//	ApplicationListMap{ "key": ApplicationListArgs{...} }
type ApplicationListMapInput interface {
	pulumi.Input

	ToApplicationListMapOutput() ApplicationListMapOutput
	ToApplicationListMapOutputWithContext(context.Context) ApplicationListMapOutput
}

type ApplicationListMap map[string]ApplicationListInput

func (ApplicationListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationList)(nil)).Elem()
}

func (i ApplicationListMap) ToApplicationListMapOutput() ApplicationListMapOutput {
	return i.ToApplicationListMapOutputWithContext(context.Background())
}

func (i ApplicationListMap) ToApplicationListMapOutputWithContext(ctx context.Context) ApplicationListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationListMapOutput)
}

type ApplicationListOutput struct{ *pulumi.OutputState }

func (ApplicationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationList)(nil)).Elem()
}

func (o ApplicationListOutput) ToApplicationListOutput() ApplicationListOutput {
	return o
}

func (o ApplicationListOutput) ToApplicationListOutputWithContext(ctx context.Context) ApplicationListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ApplicationListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of applications. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
func (o ApplicationListOutput) Items() ApplicationTypeArrayOutput {
	return o.ApplyT(func(v *ApplicationList) ApplicationTypeArrayOutput { return v.Items }).(ApplicationTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApplicationListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ApplicationListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ApplicationList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ApplicationListArrayOutput struct{ *pulumi.OutputState }

func (ApplicationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationList)(nil)).Elem()
}

func (o ApplicationListArrayOutput) ToApplicationListArrayOutput() ApplicationListArrayOutput {
	return o
}

func (o ApplicationListArrayOutput) ToApplicationListArrayOutputWithContext(ctx context.Context) ApplicationListArrayOutput {
	return o
}

func (o ApplicationListArrayOutput) Index(i pulumi.IntInput) ApplicationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationList {
		return vs[0].([]*ApplicationList)[vs[1].(int)]
	}).(ApplicationListOutput)
}

type ApplicationListMapOutput struct{ *pulumi.OutputState }

func (ApplicationListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationList)(nil)).Elem()
}

func (o ApplicationListMapOutput) ToApplicationListMapOutput() ApplicationListMapOutput {
	return o
}

func (o ApplicationListMapOutput) ToApplicationListMapOutputWithContext(ctx context.Context) ApplicationListMapOutput {
	return o
}

func (o ApplicationListMapOutput) MapIndex(k pulumi.StringInput) ApplicationListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationList {
		return vs[0].(map[string]*ApplicationList)[vs[1].(string)]
	}).(ApplicationListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationListInput)(nil)).Elem(), &ApplicationList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationListArrayInput)(nil)).Elem(), ApplicationListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationListMapInput)(nil)).Elem(), ApplicationListMap{})
	pulumi.RegisterOutputType(ApplicationListOutput{})
	pulumi.RegisterOutputType(ApplicationListArrayOutput{})
	pulumi.RegisterOutputType(ApplicationListMapOutput{})
}