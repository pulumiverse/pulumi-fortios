// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure console.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewSystemConsole(ctx, "trname", &fortios.SystemConsoleArgs{
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
//
// ## Import
//
// # System Console can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemConsole:SystemConsole labelname SystemConsole
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemConsole:SystemConsole labelname SystemConsole
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemConsole struct {
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

// NewSystemConsole registers a new resource with the given unique name, arguments, and options.
func NewSystemConsole(ctx *pulumi.Context,
	name string, args *SystemConsoleArgs, opts ...pulumi.ResourceOption) (*SystemConsole, error) {
	if args == nil {
		args = &SystemConsoleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemConsole
	err := ctx.RegisterResource("fortios:index/systemConsole:SystemConsole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemConsole gets an existing SystemConsole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemConsoleState, opts ...pulumi.ResourceOption) (*SystemConsole, error) {
	var resource SystemConsole
	err := ctx.ReadResource("fortios:index/systemConsole:SystemConsole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemConsole resources.
type systemConsoleState struct {
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

type SystemConsoleState struct {
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

func (SystemConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemConsoleState)(nil)).Elem()
}

type systemConsoleArgs struct {
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

// The set of arguments for constructing a SystemConsole resource.
type SystemConsoleArgs struct {
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

func (SystemConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemConsoleArgs)(nil)).Elem()
}

type SystemConsoleInput interface {
	pulumi.Input

	ToSystemConsoleOutput() SystemConsoleOutput
	ToSystemConsoleOutputWithContext(ctx context.Context) SystemConsoleOutput
}

func (*SystemConsole) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemConsole)(nil)).Elem()
}

func (i *SystemConsole) ToSystemConsoleOutput() SystemConsoleOutput {
	return i.ToSystemConsoleOutputWithContext(context.Background())
}

func (i *SystemConsole) ToSystemConsoleOutputWithContext(ctx context.Context) SystemConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemConsoleOutput)
}

// SystemConsoleArrayInput is an input type that accepts SystemConsoleArray and SystemConsoleArrayOutput values.
// You can construct a concrete instance of `SystemConsoleArrayInput` via:
//
//	SystemConsoleArray{ SystemConsoleArgs{...} }
type SystemConsoleArrayInput interface {
	pulumi.Input

	ToSystemConsoleArrayOutput() SystemConsoleArrayOutput
	ToSystemConsoleArrayOutputWithContext(context.Context) SystemConsoleArrayOutput
}

type SystemConsoleArray []SystemConsoleInput

func (SystemConsoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemConsole)(nil)).Elem()
}

func (i SystemConsoleArray) ToSystemConsoleArrayOutput() SystemConsoleArrayOutput {
	return i.ToSystemConsoleArrayOutputWithContext(context.Background())
}

func (i SystemConsoleArray) ToSystemConsoleArrayOutputWithContext(ctx context.Context) SystemConsoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemConsoleArrayOutput)
}

// SystemConsoleMapInput is an input type that accepts SystemConsoleMap and SystemConsoleMapOutput values.
// You can construct a concrete instance of `SystemConsoleMapInput` via:
//
//	SystemConsoleMap{ "key": SystemConsoleArgs{...} }
type SystemConsoleMapInput interface {
	pulumi.Input

	ToSystemConsoleMapOutput() SystemConsoleMapOutput
	ToSystemConsoleMapOutputWithContext(context.Context) SystemConsoleMapOutput
}

type SystemConsoleMap map[string]SystemConsoleInput

func (SystemConsoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemConsole)(nil)).Elem()
}

func (i SystemConsoleMap) ToSystemConsoleMapOutput() SystemConsoleMapOutput {
	return i.ToSystemConsoleMapOutputWithContext(context.Background())
}

func (i SystemConsoleMap) ToSystemConsoleMapOutputWithContext(ctx context.Context) SystemConsoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemConsoleMapOutput)
}

type SystemConsoleOutput struct{ *pulumi.OutputState }

func (SystemConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemConsole)(nil)).Elem()
}

func (o SystemConsoleOutput) ToSystemConsoleOutput() SystemConsoleOutput {
	return o
}

func (o SystemConsoleOutput) ToSystemConsoleOutputWithContext(ctx context.Context) SystemConsoleOutput {
	return o
}

// Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
func (o SystemConsoleOutput) Baudrate() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringOutput { return v.Baudrate }).(pulumi.StringOutput)
}

// Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
func (o SystemConsoleOutput) Fortiexplorer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringOutput { return v.Fortiexplorer }).(pulumi.StringOutput)
}

// Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
func (o SystemConsoleOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// Console mode. Valid values: `batch`, `line`.
func (o SystemConsoleOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Console output mode. Valid values: `standard`, `more`.
func (o SystemConsoleOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringOutput { return v.Output }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemConsoleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemConsole) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemConsoleArrayOutput struct{ *pulumi.OutputState }

func (SystemConsoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemConsole)(nil)).Elem()
}

func (o SystemConsoleArrayOutput) ToSystemConsoleArrayOutput() SystemConsoleArrayOutput {
	return o
}

func (o SystemConsoleArrayOutput) ToSystemConsoleArrayOutputWithContext(ctx context.Context) SystemConsoleArrayOutput {
	return o
}

func (o SystemConsoleArrayOutput) Index(i pulumi.IntInput) SystemConsoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemConsole {
		return vs[0].([]*SystemConsole)[vs[1].(int)]
	}).(SystemConsoleOutput)
}

type SystemConsoleMapOutput struct{ *pulumi.OutputState }

func (SystemConsoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemConsole)(nil)).Elem()
}

func (o SystemConsoleMapOutput) ToSystemConsoleMapOutput() SystemConsoleMapOutput {
	return o
}

func (o SystemConsoleMapOutput) ToSystemConsoleMapOutputWithContext(ctx context.Context) SystemConsoleMapOutput {
	return o
}

func (o SystemConsoleMapOutput) MapIndex(k pulumi.StringInput) SystemConsoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemConsole {
		return vs[0].(map[string]*SystemConsole)[vs[1].(string)]
	}).(SystemConsoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemConsoleInput)(nil)).Elem(), &SystemConsole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemConsoleArrayInput)(nil)).Elem(), SystemConsoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemConsoleMapInput)(nil)).Elem(), SystemConsoleMap{})
	pulumi.RegisterOutputType(SystemConsoleOutput{})
	pulumi.RegisterOutputType(SystemConsoleArrayOutput{})
	pulumi.RegisterOutputType(SystemConsoleMapOutput{})
}
