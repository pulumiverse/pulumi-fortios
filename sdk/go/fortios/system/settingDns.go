// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a resource to configure DNS of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `system.Dns`, we recommend that you use the new resource.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewSettingDns(ctx, "test1", &system.SettingDnsArgs{
//				DnsOverTls: pulumi.String("disable"),
//				Primary:    pulumi.String("208.91.112.53"),
//				Secondary:  pulumi.String("208.91.112.22"),
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
type SettingDns struct {
	pulumi.CustomResourceState

	// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
	DnsOverTls pulumi.StringOutput `pulumi:"dnsOverTls"`
	// Primary DNS server IP address.
	Primary pulumi.StringOutput `pulumi:"primary"`
	// Secondary DNS server IP address.
	Secondary pulumi.StringOutput `pulumi:"secondary"`
}

// NewSettingDns registers a new resource with the given unique name, arguments, and options.
func NewSettingDns(ctx *pulumi.Context,
	name string, args *SettingDnsArgs, opts ...pulumi.ResourceOption) (*SettingDns, error) {
	if args == nil {
		args = &SettingDnsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SettingDns
	err := ctx.RegisterResource("fortios:system/settingDns:SettingDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSettingDns gets an existing SettingDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSettingDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingDnsState, opts ...pulumi.ResourceOption) (*SettingDns, error) {
	var resource SettingDns
	err := ctx.ReadResource("fortios:system/settingDns:SettingDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SettingDns resources.
type settingDnsState struct {
	// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
	DnsOverTls *string `pulumi:"dnsOverTls"`
	// Primary DNS server IP address.
	Primary *string `pulumi:"primary"`
	// Secondary DNS server IP address.
	Secondary *string `pulumi:"secondary"`
}

type SettingDnsState struct {
	// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
	DnsOverTls pulumi.StringPtrInput
	// Primary DNS server IP address.
	Primary pulumi.StringPtrInput
	// Secondary DNS server IP address.
	Secondary pulumi.StringPtrInput
}

func (SettingDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingDnsState)(nil)).Elem()
}

type settingDnsArgs struct {
	// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
	DnsOverTls *string `pulumi:"dnsOverTls"`
	// Primary DNS server IP address.
	Primary *string `pulumi:"primary"`
	// Secondary DNS server IP address.
	Secondary *string `pulumi:"secondary"`
}

// The set of arguments for constructing a SettingDns resource.
type SettingDnsArgs struct {
	// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
	DnsOverTls pulumi.StringPtrInput
	// Primary DNS server IP address.
	Primary pulumi.StringPtrInput
	// Secondary DNS server IP address.
	Secondary pulumi.StringPtrInput
}

func (SettingDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingDnsArgs)(nil)).Elem()
}

type SettingDnsInput interface {
	pulumi.Input

	ToSettingDnsOutput() SettingDnsOutput
	ToSettingDnsOutputWithContext(ctx context.Context) SettingDnsOutput
}

func (*SettingDns) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingDns)(nil)).Elem()
}

func (i *SettingDns) ToSettingDnsOutput() SettingDnsOutput {
	return i.ToSettingDnsOutputWithContext(context.Background())
}

func (i *SettingDns) ToSettingDnsOutputWithContext(ctx context.Context) SettingDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingDnsOutput)
}

// SettingDnsArrayInput is an input type that accepts SettingDnsArray and SettingDnsArrayOutput values.
// You can construct a concrete instance of `SettingDnsArrayInput` via:
//
//	SettingDnsArray{ SettingDnsArgs{...} }
type SettingDnsArrayInput interface {
	pulumi.Input

	ToSettingDnsArrayOutput() SettingDnsArrayOutput
	ToSettingDnsArrayOutputWithContext(context.Context) SettingDnsArrayOutput
}

type SettingDnsArray []SettingDnsInput

func (SettingDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SettingDns)(nil)).Elem()
}

func (i SettingDnsArray) ToSettingDnsArrayOutput() SettingDnsArrayOutput {
	return i.ToSettingDnsArrayOutputWithContext(context.Background())
}

func (i SettingDnsArray) ToSettingDnsArrayOutputWithContext(ctx context.Context) SettingDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingDnsArrayOutput)
}

// SettingDnsMapInput is an input type that accepts SettingDnsMap and SettingDnsMapOutput values.
// You can construct a concrete instance of `SettingDnsMapInput` via:
//
//	SettingDnsMap{ "key": SettingDnsArgs{...} }
type SettingDnsMapInput interface {
	pulumi.Input

	ToSettingDnsMapOutput() SettingDnsMapOutput
	ToSettingDnsMapOutputWithContext(context.Context) SettingDnsMapOutput
}

type SettingDnsMap map[string]SettingDnsInput

func (SettingDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SettingDns)(nil)).Elem()
}

func (i SettingDnsMap) ToSettingDnsMapOutput() SettingDnsMapOutput {
	return i.ToSettingDnsMapOutputWithContext(context.Background())
}

func (i SettingDnsMap) ToSettingDnsMapOutputWithContext(ctx context.Context) SettingDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingDnsMapOutput)
}

type SettingDnsOutput struct{ *pulumi.OutputState }

func (SettingDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingDns)(nil)).Elem()
}

func (o SettingDnsOutput) ToSettingDnsOutput() SettingDnsOutput {
	return o
}

func (o SettingDnsOutput) ToSettingDnsOutputWithContext(ctx context.Context) SettingDnsOutput {
	return o
}

// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
func (o SettingDnsOutput) DnsOverTls() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingDns) pulumi.StringOutput { return v.DnsOverTls }).(pulumi.StringOutput)
}

// Primary DNS server IP address.
func (o SettingDnsOutput) Primary() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingDns) pulumi.StringOutput { return v.Primary }).(pulumi.StringOutput)
}

// Secondary DNS server IP address.
func (o SettingDnsOutput) Secondary() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingDns) pulumi.StringOutput { return v.Secondary }).(pulumi.StringOutput)
}

type SettingDnsArrayOutput struct{ *pulumi.OutputState }

func (SettingDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SettingDns)(nil)).Elem()
}

func (o SettingDnsArrayOutput) ToSettingDnsArrayOutput() SettingDnsArrayOutput {
	return o
}

func (o SettingDnsArrayOutput) ToSettingDnsArrayOutputWithContext(ctx context.Context) SettingDnsArrayOutput {
	return o
}

func (o SettingDnsArrayOutput) Index(i pulumi.IntInput) SettingDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SettingDns {
		return vs[0].([]*SettingDns)[vs[1].(int)]
	}).(SettingDnsOutput)
}

type SettingDnsMapOutput struct{ *pulumi.OutputState }

func (SettingDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SettingDns)(nil)).Elem()
}

func (o SettingDnsMapOutput) ToSettingDnsMapOutput() SettingDnsMapOutput {
	return o
}

func (o SettingDnsMapOutput) ToSettingDnsMapOutputWithContext(ctx context.Context) SettingDnsMapOutput {
	return o
}

func (o SettingDnsMapOutput) MapIndex(k pulumi.StringInput) SettingDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SettingDns {
		return vs[0].(map[string]*SettingDns)[vs[1].(string)]
	}).(SettingDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingDnsInput)(nil)).Elem(), &SettingDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingDnsArrayInput)(nil)).Elem(), SettingDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingDnsMapInput)(nil)).Elem(), SettingDnsMap{})
	pulumi.RegisterOutputType(SettingDnsOutput{})
	pulumi.RegisterOutputType(SettingDnsArrayOutput{})
	pulumi.RegisterOutputType(SettingDnsMapOutput{})
}