// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// This resource supports updating system network interface for FortiManager.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fmg.NewSystemNetworkInterface(ctx, "test1", &fmg.SystemNetworkInterfaceArgs{
//				AllowAccesses: pulumi.StringArray{
//					pulumi.String("ping"),
//					pulumi.String("ssh"),
//					pulumi.String("https"),
//					pulumi.String("http"),
//				},
//				Description: pulumi.String(""),
//				Ip:          pulumi.String("1.1.1.3 255.255.255.0"),
//				ServiceAccesses: pulumi.StringArray{
//					pulumi.String("webfilter"),
//					pulumi.String("fgtupdates"),
//				},
//				Status: pulumi.String("up"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type SystemNetworkInterface struct {
	pulumi.CustomResourceState

	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayOutput `pulumi:"allowAccesses"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Interface Ipaddress.
	Ip pulumi.StringPtrOutput `pulumi:"ip"`
	// Interface port.
	Name            pulumi.StringOutput      `pulumi:"name"`
	ServiceAccesses pulumi.StringArrayOutput `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewSystemNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewSystemNetworkInterface(ctx *pulumi.Context,
	name string, args *SystemNetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*SystemNetworkInterface, error) {
	if args == nil {
		args = &SystemNetworkInterfaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemNetworkInterface
	err := ctx.RegisterResource("fortios:fmg/systemNetworkInterface:SystemNetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNetworkInterface gets an existing SystemNetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNetworkInterfaceState, opts ...pulumi.ResourceOption) (*SystemNetworkInterface, error) {
	var resource SystemNetworkInterface
	err := ctx.ReadResource("fortios:fmg/systemNetworkInterface:SystemNetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNetworkInterface resources.
type systemNetworkInterfaceState struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses []string `pulumi:"allowAccesses"`
	// Description.
	Description *string `pulumi:"description"`
	// Interface Ipaddress.
	Ip *string `pulumi:"ip"`
	// Interface port.
	Name            *string  `pulumi:"name"`
	ServiceAccesses []string `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status *string `pulumi:"status"`
}

type SystemNetworkInterfaceState struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Interface Ipaddress.
	Ip pulumi.StringPtrInput
	// Interface port.
	Name            pulumi.StringPtrInput
	ServiceAccesses pulumi.StringArrayInput
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrInput
}

func (SystemNetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkInterfaceState)(nil)).Elem()
}

type systemNetworkInterfaceArgs struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses []string `pulumi:"allowAccesses"`
	// Description.
	Description *string `pulumi:"description"`
	// Interface Ipaddress.
	Ip *string `pulumi:"ip"`
	// Interface port.
	Name            *string  `pulumi:"name"`
	ServiceAccesses []string `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a SystemNetworkInterface resource.
type SystemNetworkInterfaceArgs struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Interface Ipaddress.
	Ip pulumi.StringPtrInput
	// Interface port.
	Name            pulumi.StringPtrInput
	ServiceAccesses pulumi.StringArrayInput
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrInput
}

func (SystemNetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkInterfaceArgs)(nil)).Elem()
}

type SystemNetworkInterfaceInput interface {
	pulumi.Input

	ToSystemNetworkInterfaceOutput() SystemNetworkInterfaceOutput
	ToSystemNetworkInterfaceOutputWithContext(ctx context.Context) SystemNetworkInterfaceOutput
}

func (*SystemNetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkInterface)(nil)).Elem()
}

func (i *SystemNetworkInterface) ToSystemNetworkInterfaceOutput() SystemNetworkInterfaceOutput {
	return i.ToSystemNetworkInterfaceOutputWithContext(context.Background())
}

func (i *SystemNetworkInterface) ToSystemNetworkInterfaceOutputWithContext(ctx context.Context) SystemNetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkInterfaceOutput)
}

// SystemNetworkInterfaceArrayInput is an input type that accepts SystemNetworkInterfaceArray and SystemNetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `SystemNetworkInterfaceArrayInput` via:
//
//	SystemNetworkInterfaceArray{ SystemNetworkInterfaceArgs{...} }
type SystemNetworkInterfaceArrayInput interface {
	pulumi.Input

	ToSystemNetworkInterfaceArrayOutput() SystemNetworkInterfaceArrayOutput
	ToSystemNetworkInterfaceArrayOutputWithContext(context.Context) SystemNetworkInterfaceArrayOutput
}

type SystemNetworkInterfaceArray []SystemNetworkInterfaceInput

func (SystemNetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkInterface)(nil)).Elem()
}

func (i SystemNetworkInterfaceArray) ToSystemNetworkInterfaceArrayOutput() SystemNetworkInterfaceArrayOutput {
	return i.ToSystemNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i SystemNetworkInterfaceArray) ToSystemNetworkInterfaceArrayOutputWithContext(ctx context.Context) SystemNetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkInterfaceArrayOutput)
}

// SystemNetworkInterfaceMapInput is an input type that accepts SystemNetworkInterfaceMap and SystemNetworkInterfaceMapOutput values.
// You can construct a concrete instance of `SystemNetworkInterfaceMapInput` via:
//
//	SystemNetworkInterfaceMap{ "key": SystemNetworkInterfaceArgs{...} }
type SystemNetworkInterfaceMapInput interface {
	pulumi.Input

	ToSystemNetworkInterfaceMapOutput() SystemNetworkInterfaceMapOutput
	ToSystemNetworkInterfaceMapOutputWithContext(context.Context) SystemNetworkInterfaceMapOutput
}

type SystemNetworkInterfaceMap map[string]SystemNetworkInterfaceInput

func (SystemNetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkInterface)(nil)).Elem()
}

func (i SystemNetworkInterfaceMap) ToSystemNetworkInterfaceMapOutput() SystemNetworkInterfaceMapOutput {
	return i.ToSystemNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i SystemNetworkInterfaceMap) ToSystemNetworkInterfaceMapOutputWithContext(ctx context.Context) SystemNetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkInterfaceMapOutput)
}

type SystemNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (SystemNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkInterface)(nil)).Elem()
}

func (o SystemNetworkInterfaceOutput) ToSystemNetworkInterfaceOutput() SystemNetworkInterfaceOutput {
	return o
}

func (o SystemNetworkInterfaceOutput) ToSystemNetworkInterfaceOutputWithContext(ctx context.Context) SystemNetworkInterfaceOutput {
	return o
}

// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
func (o SystemNetworkInterfaceOutput) AllowAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringArrayOutput { return v.AllowAccesses }).(pulumi.StringArrayOutput)
}

// Description.
func (o SystemNetworkInterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Interface Ipaddress.
func (o SystemNetworkInterfaceOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringPtrOutput { return v.Ip }).(pulumi.StringPtrOutput)
}

// Interface port.
func (o SystemNetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemNetworkInterfaceOutput) ServiceAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringArrayOutput { return v.ServiceAccesses }).(pulumi.StringArrayOutput)
}

// Interface status, Enum: ["down", "up"]
func (o SystemNetworkInterfaceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNetworkInterface) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

type SystemNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (SystemNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkInterface)(nil)).Elem()
}

func (o SystemNetworkInterfaceArrayOutput) ToSystemNetworkInterfaceArrayOutput() SystemNetworkInterfaceArrayOutput {
	return o
}

func (o SystemNetworkInterfaceArrayOutput) ToSystemNetworkInterfaceArrayOutputWithContext(ctx context.Context) SystemNetworkInterfaceArrayOutput {
	return o
}

func (o SystemNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) SystemNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemNetworkInterface {
		return vs[0].([]*SystemNetworkInterface)[vs[1].(int)]
	}).(SystemNetworkInterfaceOutput)
}

type SystemNetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (SystemNetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkInterface)(nil)).Elem()
}

func (o SystemNetworkInterfaceMapOutput) ToSystemNetworkInterfaceMapOutput() SystemNetworkInterfaceMapOutput {
	return o
}

func (o SystemNetworkInterfaceMapOutput) ToSystemNetworkInterfaceMapOutputWithContext(ctx context.Context) SystemNetworkInterfaceMapOutput {
	return o
}

func (o SystemNetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) SystemNetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemNetworkInterface {
		return vs[0].(map[string]*SystemNetworkInterface)[vs[1].(string)]
	}).(SystemNetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkInterfaceInput)(nil)).Elem(), &SystemNetworkInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkInterfaceArrayInput)(nil)).Elem(), SystemNetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkInterfaceMapInput)(nil)).Elem(), SystemNetworkInterfaceMap{})
	pulumi.RegisterOutputType(SystemNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(SystemNetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(SystemNetworkInterfaceMapOutput{})
}
