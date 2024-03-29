// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure console.
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
//			_, err := system.NewConsole(ctx, "trname", &system.ConsoleArgs{
//				Baudrate: pulumi.String("9600"),
//				Login:    pulumi.String("enable"),
//				Mode:     pulumi.String("line"),
//				Output:   pulumi.String("more"),
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
//
// ## Import
//
// System Console can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/console:Console labelname SystemConsole
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/console:Console labelname SystemConsole
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Console struct {
	pulumi.CustomResourceState

	// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
	Baudrate pulumi.StringOutput `pulumi:"baudrate"`
	// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
	Fortiexplorer pulumi.StringOutput `pulumi:"fortiexplorer"`
	// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
	Login pulumi.StringOutput `pulumi:"login"`
	// Console mode. Valid values: `batch`, `line`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Console output mode. Valid values: `standard`, `more`.
	Output pulumi.StringOutput `pulumi:"output"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewConsole registers a new resource with the given unique name, arguments, and options.
func NewConsole(ctx *pulumi.Context,
	name string, args *ConsoleArgs, opts ...pulumi.ResourceOption) (*Console, error) {
	if args == nil {
		args = &ConsoleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Console
	err := ctx.RegisterResource("fortios:system/console:Console", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsole gets an existing Console resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsoleState, opts ...pulumi.ResourceOption) (*Console, error) {
	var resource Console
	err := ctx.ReadResource("fortios:system/console:Console", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Console resources.
type consoleState struct {
	// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
	Baudrate *string `pulumi:"baudrate"`
	// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
	Fortiexplorer *string `pulumi:"fortiexplorer"`
	// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
	Login *string `pulumi:"login"`
	// Console mode. Valid values: `batch`, `line`.
	Mode *string `pulumi:"mode"`
	// Console output mode. Valid values: `standard`, `more`.
	Output *string `pulumi:"output"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ConsoleState struct {
	// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
	Baudrate pulumi.StringPtrInput
	// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
	Fortiexplorer pulumi.StringPtrInput
	// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
	Login pulumi.StringPtrInput
	// Console mode. Valid values: `batch`, `line`.
	Mode pulumi.StringPtrInput
	// Console output mode. Valid values: `standard`, `more`.
	Output pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleState)(nil)).Elem()
}

type consoleArgs struct {
	// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
	Baudrate *string `pulumi:"baudrate"`
	// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
	Fortiexplorer *string `pulumi:"fortiexplorer"`
	// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
	Login *string `pulumi:"login"`
	// Console mode. Valid values: `batch`, `line`.
	Mode *string `pulumi:"mode"`
	// Console output mode. Valid values: `standard`, `more`.
	Output *string `pulumi:"output"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Console resource.
type ConsoleArgs struct {
	// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
	Baudrate pulumi.StringPtrInput
	// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
	Fortiexplorer pulumi.StringPtrInput
	// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
	Login pulumi.StringPtrInput
	// Console mode. Valid values: `batch`, `line`.
	Mode pulumi.StringPtrInput
	// Console output mode. Valid values: `standard`, `more`.
	Output pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleArgs)(nil)).Elem()
}

type ConsoleInput interface {
	pulumi.Input

	ToConsoleOutput() ConsoleOutput
	ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput
}

func (*Console) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (i *Console) ToConsoleOutput() ConsoleOutput {
	return i.ToConsoleOutputWithContext(context.Background())
}

func (i *Console) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleOutput)
}

// ConsoleArrayInput is an input type that accepts ConsoleArray and ConsoleArrayOutput values.
// You can construct a concrete instance of `ConsoleArrayInput` via:
//
//	ConsoleArray{ ConsoleArgs{...} }
type ConsoleArrayInput interface {
	pulumi.Input

	ToConsoleArrayOutput() ConsoleArrayOutput
	ToConsoleArrayOutputWithContext(context.Context) ConsoleArrayOutput
}

type ConsoleArray []ConsoleInput

func (ConsoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Console)(nil)).Elem()
}

func (i ConsoleArray) ToConsoleArrayOutput() ConsoleArrayOutput {
	return i.ToConsoleArrayOutputWithContext(context.Background())
}

func (i ConsoleArray) ToConsoleArrayOutputWithContext(ctx context.Context) ConsoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleArrayOutput)
}

// ConsoleMapInput is an input type that accepts ConsoleMap and ConsoleMapOutput values.
// You can construct a concrete instance of `ConsoleMapInput` via:
//
//	ConsoleMap{ "key": ConsoleArgs{...} }
type ConsoleMapInput interface {
	pulumi.Input

	ToConsoleMapOutput() ConsoleMapOutput
	ToConsoleMapOutputWithContext(context.Context) ConsoleMapOutput
}

type ConsoleMap map[string]ConsoleInput

func (ConsoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Console)(nil)).Elem()
}

func (i ConsoleMap) ToConsoleMapOutput() ConsoleMapOutput {
	return i.ToConsoleMapOutputWithContext(context.Background())
}

func (i ConsoleMap) ToConsoleMapOutputWithContext(ctx context.Context) ConsoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleMapOutput)
}

type ConsoleOutput struct{ *pulumi.OutputState }

func (ConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (o ConsoleOutput) ToConsoleOutput() ConsoleOutput {
	return o
}

func (o ConsoleOutput) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return o
}

// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
func (o ConsoleOutput) Baudrate() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Baudrate }).(pulumi.StringOutput)
}

// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
func (o ConsoleOutput) Fortiexplorer() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Fortiexplorer }).(pulumi.StringOutput)
}

// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
func (o ConsoleOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// Console mode. Valid values: `batch`, `line`.
func (o ConsoleOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Console output mode. Valid values: `standard`, `more`.
func (o ConsoleOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Output }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ConsoleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Console) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ConsoleArrayOutput struct{ *pulumi.OutputState }

func (ConsoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Console)(nil)).Elem()
}

func (o ConsoleArrayOutput) ToConsoleArrayOutput() ConsoleArrayOutput {
	return o
}

func (o ConsoleArrayOutput) ToConsoleArrayOutputWithContext(ctx context.Context) ConsoleArrayOutput {
	return o
}

func (o ConsoleArrayOutput) Index(i pulumi.IntInput) ConsoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Console {
		return vs[0].([]*Console)[vs[1].(int)]
	}).(ConsoleOutput)
}

type ConsoleMapOutput struct{ *pulumi.OutputState }

func (ConsoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Console)(nil)).Elem()
}

func (o ConsoleMapOutput) ToConsoleMapOutput() ConsoleMapOutput {
	return o
}

func (o ConsoleMapOutput) ToConsoleMapOutputWithContext(ctx context.Context) ConsoleMapOutput {
	return o
}

func (o ConsoleMapOutput) MapIndex(k pulumi.StringInput) ConsoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Console {
		return vs[0].(map[string]*Console)[vs[1].(string)]
	}).(ConsoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleInput)(nil)).Elem(), &Console{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleArrayInput)(nil)).Elem(), ConsoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleMapInput)(nil)).Elem(), ConsoleMap{})
	pulumi.RegisterOutputType(ConsoleOutput{})
	pulumi.RegisterOutputType(ConsoleArrayOutput{})
	pulumi.RegisterOutputType(ConsoleMapOutput{})
}
