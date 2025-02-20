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

// VMAlert  executes a list of given alerting or recording rules against configured address.
type VMAlert struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	Spec     VMAlertSpecOutput       `pulumi:"spec"`
	Status   VMAlertStatusPtrOutput  `pulumi:"status"`
}

// NewVMAlert registers a new resource with the given unique name, arguments, and options.
func NewVMAlert(ctx *pulumi.Context,
	name string, args *VMAlertArgs, opts ...pulumi.ResourceOption) (*VMAlert, error) {
	if args == nil {
		args = &VMAlertArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("operator.victoriametrics.com/v1beta1")
	args.Kind = pulumi.StringPtr("VMAlert")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMAlert
	err := ctx.RegisterResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMAlert gets an existing VMAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMAlertState, opts ...pulumi.ResourceOption) (*VMAlert, error) {
	var resource VMAlert
	err := ctx.ReadResource("kubernetes:operator.victoriametrics.com/v1beta1:VMAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMAlert resources.
type vmalertState struct {
}

type VMAlertState struct {
}

func (VMAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertState)(nil)).Elem()
}

type vmalertArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *VMAlertSpec       `pulumi:"spec"`
}

// The set of arguments for constructing a VMAlert resource.
type VMAlertArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	Spec     VMAlertSpecPtrInput
}

func (VMAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmalertArgs)(nil)).Elem()
}

type VMAlertInput interface {
	pulumi.Input

	ToVMAlertOutput() VMAlertOutput
	ToVMAlertOutputWithContext(ctx context.Context) VMAlertOutput
}

func (*VMAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlert)(nil)).Elem()
}

func (i *VMAlert) ToVMAlertOutput() VMAlertOutput {
	return i.ToVMAlertOutputWithContext(context.Background())
}

func (i *VMAlert) ToVMAlertOutputWithContext(ctx context.Context) VMAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertOutput)
}

// VMAlertArrayInput is an input type that accepts VMAlertArray and VMAlertArrayOutput values.
// You can construct a concrete instance of `VMAlertArrayInput` via:
//
//	VMAlertArray{ VMAlertArgs{...} }
type VMAlertArrayInput interface {
	pulumi.Input

	ToVMAlertArrayOutput() VMAlertArrayOutput
	ToVMAlertArrayOutputWithContext(context.Context) VMAlertArrayOutput
}

type VMAlertArray []VMAlertInput

func (VMAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlert)(nil)).Elem()
}

func (i VMAlertArray) ToVMAlertArrayOutput() VMAlertArrayOutput {
	return i.ToVMAlertArrayOutputWithContext(context.Background())
}

func (i VMAlertArray) ToVMAlertArrayOutputWithContext(ctx context.Context) VMAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertArrayOutput)
}

// VMAlertMapInput is an input type that accepts VMAlertMap and VMAlertMapOutput values.
// You can construct a concrete instance of `VMAlertMapInput` via:
//
//	VMAlertMap{ "key": VMAlertArgs{...} }
type VMAlertMapInput interface {
	pulumi.Input

	ToVMAlertMapOutput() VMAlertMapOutput
	ToVMAlertMapOutputWithContext(context.Context) VMAlertMapOutput
}

type VMAlertMap map[string]VMAlertInput

func (VMAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlert)(nil)).Elem()
}

func (i VMAlertMap) ToVMAlertMapOutput() VMAlertMapOutput {
	return i.ToVMAlertMapOutputWithContext(context.Background())
}

func (i VMAlertMap) ToVMAlertMapOutputWithContext(ctx context.Context) VMAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMAlertMapOutput)
}

type VMAlertOutput struct{ *pulumi.OutputState }

func (VMAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMAlert)(nil)).Elem()
}

func (o VMAlertOutput) ToVMAlertOutput() VMAlertOutput {
	return o
}

func (o VMAlertOutput) ToVMAlertOutputWithContext(ctx context.Context) VMAlertOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VMAlertOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAlert) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VMAlertOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VMAlert) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VMAlertOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *VMAlert) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

func (o VMAlertOutput) Spec() VMAlertSpecOutput {
	return o.ApplyT(func(v *VMAlert) VMAlertSpecOutput { return v.Spec }).(VMAlertSpecOutput)
}

func (o VMAlertOutput) Status() VMAlertStatusPtrOutput {
	return o.ApplyT(func(v *VMAlert) VMAlertStatusPtrOutput { return v.Status }).(VMAlertStatusPtrOutput)
}

type VMAlertArrayOutput struct{ *pulumi.OutputState }

func (VMAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMAlert)(nil)).Elem()
}

func (o VMAlertArrayOutput) ToVMAlertArrayOutput() VMAlertArrayOutput {
	return o
}

func (o VMAlertArrayOutput) ToVMAlertArrayOutputWithContext(ctx context.Context) VMAlertArrayOutput {
	return o
}

func (o VMAlertArrayOutput) Index(i pulumi.IntInput) VMAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMAlert {
		return vs[0].([]*VMAlert)[vs[1].(int)]
	}).(VMAlertOutput)
}

type VMAlertMapOutput struct{ *pulumi.OutputState }

func (VMAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMAlert)(nil)).Elem()
}

func (o VMAlertMapOutput) ToVMAlertMapOutput() VMAlertMapOutput {
	return o
}

func (o VMAlertMapOutput) ToVMAlertMapOutputWithContext(ctx context.Context) VMAlertMapOutput {
	return o
}

func (o VMAlertMapOutput) MapIndex(k pulumi.StringInput) VMAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMAlert {
		return vs[0].(map[string]*VMAlert)[vs[1].(string)]
	}).(VMAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertInput)(nil)).Elem(), &VMAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertArrayInput)(nil)).Elem(), VMAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMAlertMapInput)(nil)).Elem(), VMAlertMap{})
	pulumi.RegisterOutputType(VMAlertOutput{})
	pulumi.RegisterOutputType(VMAlertArrayOutput{})
	pulumi.RegisterOutputType(VMAlertMapOutput{})
}
