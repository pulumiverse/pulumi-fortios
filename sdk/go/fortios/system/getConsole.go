// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on fortios system console
func LookupConsole(ctx *pulumi.Context, args *LookupConsoleArgs, opts ...pulumi.InvokeOption) (*LookupConsoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConsoleResult
	err := ctx.Invoke("fortios:system/getConsole:getConsole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConsole.
type LookupConsoleArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getConsole.
type LookupConsoleResult struct {
	// Console baud rate.
	Baudrate string `pulumi:"baudrate"`
	// Enable/disable access for FortiExplorer.
	Fortiexplorer string `pulumi:"fortiexplorer"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable serial console and FortiExplorer.
	Login string `pulumi:"login"`
	// Console mode.
	Mode string `pulumi:"mode"`
	// Console output mode.
	Output    string  `pulumi:"output"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupConsoleOutput(ctx *pulumi.Context, args LookupConsoleOutputArgs, opts ...pulumi.InvokeOption) LookupConsoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsoleResult, error) {
			args := v.(LookupConsoleArgs)
			r, err := LookupConsole(ctx, &args, opts...)
			var s LookupConsoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConsoleResultOutput)
}

// A collection of arguments for invoking getConsole.
type LookupConsoleOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupConsoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleArgs)(nil)).Elem()
}

// A collection of values returned by getConsole.
type LookupConsoleResultOutput struct{ *pulumi.OutputState }

func (LookupConsoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleResult)(nil)).Elem()
}

func (o LookupConsoleResultOutput) ToLookupConsoleResultOutput() LookupConsoleResultOutput {
	return o
}

func (o LookupConsoleResultOutput) ToLookupConsoleResultOutputWithContext(ctx context.Context) LookupConsoleResultOutput {
	return o
}

// Console baud rate.
func (o LookupConsoleResultOutput) Baudrate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Baudrate }).(pulumi.StringOutput)
}

// Enable/disable access for FortiExplorer.
func (o LookupConsoleResultOutput) Fortiexplorer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Fortiexplorer }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConsoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable serial console and FortiExplorer.
func (o LookupConsoleResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Login }).(pulumi.StringOutput)
}

// Console mode.
func (o LookupConsoleResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Console output mode.
func (o LookupConsoleResultOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsoleResult) string { return v.Output }).(pulumi.StringOutput)
}

func (o LookupConsoleResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConsoleResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsoleResultOutput{})
}
