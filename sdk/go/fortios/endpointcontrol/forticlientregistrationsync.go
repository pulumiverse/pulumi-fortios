// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiClient registration synchronization settings. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/endpointcontrol"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := endpointcontrol.NewForticlientregistrationsync(ctx, "trname", &endpointcontrol.ForticlientregistrationsyncArgs{
//				PeerIp:   pulumi.String("1.1.1.1"),
//				PeerName: pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:endpointcontrol/forticlientregistrationsync:Forticlientregistrationsync labelname {{peer_name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:endpointcontrol/forticlientregistrationsync:Forticlientregistrationsync labelname {{peer_name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Forticlientregistrationsync struct {
	pulumi.CustomResourceState

	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// Peer name.
	PeerName pulumi.StringOutput `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewForticlientregistrationsync registers a new resource with the given unique name, arguments, and options.
func NewForticlientregistrationsync(ctx *pulumi.Context,
	name string, args *ForticlientregistrationsyncArgs, opts ...pulumi.ResourceOption) (*Forticlientregistrationsync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Forticlientregistrationsync
	err := ctx.RegisterResource("fortios:endpointcontrol/forticlientregistrationsync:Forticlientregistrationsync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForticlientregistrationsync gets an existing Forticlientregistrationsync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForticlientregistrationsync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForticlientregistrationsyncState, opts ...pulumi.ResourceOption) (*Forticlientregistrationsync, error) {
	var resource Forticlientregistrationsync
	err := ctx.ReadResource("fortios:endpointcontrol/forticlientregistrationsync:Forticlientregistrationsync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Forticlientregistrationsync resources.
type forticlientregistrationsyncState struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp *string `pulumi:"peerIp"`
	// Peer name.
	PeerName *string `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ForticlientregistrationsyncState struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringPtrInput
	// Peer name.
	PeerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ForticlientregistrationsyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*forticlientregistrationsyncState)(nil)).Elem()
}

type forticlientregistrationsyncArgs struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp string `pulumi:"peerIp"`
	// Peer name.
	PeerName *string `pulumi:"peerName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Forticlientregistrationsync resource.
type ForticlientregistrationsyncArgs struct {
	// IP address of the peer FortiGate for endpoint license synchronization.
	PeerIp pulumi.StringInput
	// Peer name.
	PeerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ForticlientregistrationsyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forticlientregistrationsyncArgs)(nil)).Elem()
}

type ForticlientregistrationsyncInput interface {
	pulumi.Input

	ToForticlientregistrationsyncOutput() ForticlientregistrationsyncOutput
	ToForticlientregistrationsyncOutputWithContext(ctx context.Context) ForticlientregistrationsyncOutput
}

func (*Forticlientregistrationsync) ElementType() reflect.Type {
	return reflect.TypeOf((**Forticlientregistrationsync)(nil)).Elem()
}

func (i *Forticlientregistrationsync) ToForticlientregistrationsyncOutput() ForticlientregistrationsyncOutput {
	return i.ToForticlientregistrationsyncOutputWithContext(context.Background())
}

func (i *Forticlientregistrationsync) ToForticlientregistrationsyncOutputWithContext(ctx context.Context) ForticlientregistrationsyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForticlientregistrationsyncOutput)
}

// ForticlientregistrationsyncArrayInput is an input type that accepts ForticlientregistrationsyncArray and ForticlientregistrationsyncArrayOutput values.
// You can construct a concrete instance of `ForticlientregistrationsyncArrayInput` via:
//
//	ForticlientregistrationsyncArray{ ForticlientregistrationsyncArgs{...} }
type ForticlientregistrationsyncArrayInput interface {
	pulumi.Input

	ToForticlientregistrationsyncArrayOutput() ForticlientregistrationsyncArrayOutput
	ToForticlientregistrationsyncArrayOutputWithContext(context.Context) ForticlientregistrationsyncArrayOutput
}

type ForticlientregistrationsyncArray []ForticlientregistrationsyncInput

func (ForticlientregistrationsyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forticlientregistrationsync)(nil)).Elem()
}

func (i ForticlientregistrationsyncArray) ToForticlientregistrationsyncArrayOutput() ForticlientregistrationsyncArrayOutput {
	return i.ToForticlientregistrationsyncArrayOutputWithContext(context.Background())
}

func (i ForticlientregistrationsyncArray) ToForticlientregistrationsyncArrayOutputWithContext(ctx context.Context) ForticlientregistrationsyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForticlientregistrationsyncArrayOutput)
}

// ForticlientregistrationsyncMapInput is an input type that accepts ForticlientregistrationsyncMap and ForticlientregistrationsyncMapOutput values.
// You can construct a concrete instance of `ForticlientregistrationsyncMapInput` via:
//
//	ForticlientregistrationsyncMap{ "key": ForticlientregistrationsyncArgs{...} }
type ForticlientregistrationsyncMapInput interface {
	pulumi.Input

	ToForticlientregistrationsyncMapOutput() ForticlientregistrationsyncMapOutput
	ToForticlientregistrationsyncMapOutputWithContext(context.Context) ForticlientregistrationsyncMapOutput
}

type ForticlientregistrationsyncMap map[string]ForticlientregistrationsyncInput

func (ForticlientregistrationsyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forticlientregistrationsync)(nil)).Elem()
}

func (i ForticlientregistrationsyncMap) ToForticlientregistrationsyncMapOutput() ForticlientregistrationsyncMapOutput {
	return i.ToForticlientregistrationsyncMapOutputWithContext(context.Background())
}

func (i ForticlientregistrationsyncMap) ToForticlientregistrationsyncMapOutputWithContext(ctx context.Context) ForticlientregistrationsyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForticlientregistrationsyncMapOutput)
}

type ForticlientregistrationsyncOutput struct{ *pulumi.OutputState }

func (ForticlientregistrationsyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Forticlientregistrationsync)(nil)).Elem()
}

func (o ForticlientregistrationsyncOutput) ToForticlientregistrationsyncOutput() ForticlientregistrationsyncOutput {
	return o
}

func (o ForticlientregistrationsyncOutput) ToForticlientregistrationsyncOutputWithContext(ctx context.Context) ForticlientregistrationsyncOutput {
	return o
}

// IP address of the peer FortiGate for endpoint license synchronization.
func (o ForticlientregistrationsyncOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Forticlientregistrationsync) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

// Peer name.
func (o ForticlientregistrationsyncOutput) PeerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Forticlientregistrationsync) pulumi.StringOutput { return v.PeerName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ForticlientregistrationsyncOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forticlientregistrationsync) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ForticlientregistrationsyncArrayOutput struct{ *pulumi.OutputState }

func (ForticlientregistrationsyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forticlientregistrationsync)(nil)).Elem()
}

func (o ForticlientregistrationsyncArrayOutput) ToForticlientregistrationsyncArrayOutput() ForticlientregistrationsyncArrayOutput {
	return o
}

func (o ForticlientregistrationsyncArrayOutput) ToForticlientregistrationsyncArrayOutputWithContext(ctx context.Context) ForticlientregistrationsyncArrayOutput {
	return o
}

func (o ForticlientregistrationsyncArrayOutput) Index(i pulumi.IntInput) ForticlientregistrationsyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Forticlientregistrationsync {
		return vs[0].([]*Forticlientregistrationsync)[vs[1].(int)]
	}).(ForticlientregistrationsyncOutput)
}

type ForticlientregistrationsyncMapOutput struct{ *pulumi.OutputState }

func (ForticlientregistrationsyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forticlientregistrationsync)(nil)).Elem()
}

func (o ForticlientregistrationsyncMapOutput) ToForticlientregistrationsyncMapOutput() ForticlientregistrationsyncMapOutput {
	return o
}

func (o ForticlientregistrationsyncMapOutput) ToForticlientregistrationsyncMapOutputWithContext(ctx context.Context) ForticlientregistrationsyncMapOutput {
	return o
}

func (o ForticlientregistrationsyncMapOutput) MapIndex(k pulumi.StringInput) ForticlientregistrationsyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Forticlientregistrationsync {
		return vs[0].(map[string]*Forticlientregistrationsync)[vs[1].(string)]
	}).(ForticlientregistrationsyncOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForticlientregistrationsyncInput)(nil)).Elem(), &Forticlientregistrationsync{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForticlientregistrationsyncArrayInput)(nil)).Elem(), ForticlientregistrationsyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForticlientregistrationsyncMapInput)(nil)).Elem(), ForticlientregistrationsyncMap{})
	pulumi.RegisterOutputType(ForticlientregistrationsyncOutput{})
	pulumi.RegisterOutputType(ForticlientregistrationsyncArrayOutput{})
	pulumi.RegisterOutputType(ForticlientregistrationsyncMapOutput{})
}
