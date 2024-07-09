// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antivirus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure quarantine options.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/antivirus"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := antivirus.NewQuarantine(ctx, "trname", &antivirus.QuarantineArgs{
//				Agelimit:        pulumi.Int(0),
//				Destination:     pulumi.String("disk"),
//				Lowspace:        pulumi.String("ovrw-old"),
//				Maxfilesize:     pulumi.Int(0),
//				QuarantineQuota: pulumi.Int(0),
//				StoreBlocked:    pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs"),
//				StoreHeuristic:  pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
//				StoreInfected:   pulumi.String("imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
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
// Antivirus Quarantine can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:antivirus/quarantine:Quarantine labelname AntivirusQuarantine
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:antivirus/quarantine:Quarantine labelname AntivirusQuarantine
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Quarantine struct {
	pulumi.CustomResourceState

	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntOutput `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringOutput `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringOutput `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringOutput `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringOutput `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringOutput `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntOutput `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntOutput `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringOutput `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringOutput `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringOutput `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringOutput `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewQuarantine registers a new resource with the given unique name, arguments, and options.
func NewQuarantine(ctx *pulumi.Context,
	name string, args *QuarantineArgs, opts ...pulumi.ResourceOption) (*Quarantine, error) {
	if args == nil {
		args = &QuarantineArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Quarantine
	err := ctx.RegisterResource("fortios:antivirus/quarantine:Quarantine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuarantine gets an existing Quarantine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuarantine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuarantineState, opts ...pulumi.ResourceOption) (*Quarantine, error) {
	var resource Quarantine
	err := ctx.ReadResource("fortios:antivirus/quarantine:Quarantine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Quarantine resources.
type quarantineState struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit *int `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination *string `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked *string `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic *string `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected *string `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning *string `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace *string `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize *int `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota *int `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked *string `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic *string `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected *string `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning *string `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type QuarantineState struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntPtrInput
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringPtrInput
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringPtrInput
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringPtrInput
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringPtrInput
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringPtrInput
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringPtrInput
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntPtrInput
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntPtrInput
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringPtrInput
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringPtrInput
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringPtrInput
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QuarantineState) ElementType() reflect.Type {
	return reflect.TypeOf((*quarantineState)(nil)).Elem()
}

type quarantineArgs struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit *int `pulumi:"agelimit"`
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination *string `pulumi:"destination"`
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked *string `pulumi:"dropBlocked"`
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic *string `pulumi:"dropHeuristic"`
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected *string `pulumi:"dropInfected"`
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning *string `pulumi:"dropMachineLearning"`
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace *string `pulumi:"lowspace"`
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize *int `pulumi:"maxfilesize"`
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota *int `pulumi:"quarantineQuota"`
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked *string `pulumi:"storeBlocked"`
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic *string `pulumi:"storeHeuristic"`
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected *string `pulumi:"storeInfected"`
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning *string `pulumi:"storeMachineLearning"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Quarantine resource.
type QuarantineArgs struct {
	// Age limit for quarantined files (0 - 479 hours, 0 means forever).
	Agelimit pulumi.IntPtrInput
	// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
	Destination pulumi.StringPtrInput
	// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropBlocked pulumi.StringPtrInput
	// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropHeuristic pulumi.StringPtrInput
	// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
	DropInfected pulumi.StringPtrInput
	// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	DropMachineLearning pulumi.StringPtrInput
	// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
	Lowspace pulumi.StringPtrInput
	// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
	Maxfilesize pulumi.IntPtrInput
	// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
	QuarantineQuota pulumi.IntPtrInput
	// Quarantine blocked files found in sessions using the selected protocols.
	StoreBlocked pulumi.StringPtrInput
	// Quarantine files detected by heuristics found in sessions using the selected protocols.
	StoreHeuristic pulumi.StringPtrInput
	// Quarantine infected files found in sessions using the selected protocols.
	StoreInfected pulumi.StringPtrInput
	// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
	StoreMachineLearning pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QuarantineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quarantineArgs)(nil)).Elem()
}

type QuarantineInput interface {
	pulumi.Input

	ToQuarantineOutput() QuarantineOutput
	ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput
}

func (*Quarantine) ElementType() reflect.Type {
	return reflect.TypeOf((**Quarantine)(nil)).Elem()
}

func (i *Quarantine) ToQuarantineOutput() QuarantineOutput {
	return i.ToQuarantineOutputWithContext(context.Background())
}

func (i *Quarantine) ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineOutput)
}

// QuarantineArrayInput is an input type that accepts QuarantineArray and QuarantineArrayOutput values.
// You can construct a concrete instance of `QuarantineArrayInput` via:
//
//	QuarantineArray{ QuarantineArgs{...} }
type QuarantineArrayInput interface {
	pulumi.Input

	ToQuarantineArrayOutput() QuarantineArrayOutput
	ToQuarantineArrayOutputWithContext(context.Context) QuarantineArrayOutput
}

type QuarantineArray []QuarantineInput

func (QuarantineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quarantine)(nil)).Elem()
}

func (i QuarantineArray) ToQuarantineArrayOutput() QuarantineArrayOutput {
	return i.ToQuarantineArrayOutputWithContext(context.Background())
}

func (i QuarantineArray) ToQuarantineArrayOutputWithContext(ctx context.Context) QuarantineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineArrayOutput)
}

// QuarantineMapInput is an input type that accepts QuarantineMap and QuarantineMapOutput values.
// You can construct a concrete instance of `QuarantineMapInput` via:
//
//	QuarantineMap{ "key": QuarantineArgs{...} }
type QuarantineMapInput interface {
	pulumi.Input

	ToQuarantineMapOutput() QuarantineMapOutput
	ToQuarantineMapOutputWithContext(context.Context) QuarantineMapOutput
}

type QuarantineMap map[string]QuarantineInput

func (QuarantineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quarantine)(nil)).Elem()
}

func (i QuarantineMap) ToQuarantineMapOutput() QuarantineMapOutput {
	return i.ToQuarantineMapOutputWithContext(context.Background())
}

func (i QuarantineMap) ToQuarantineMapOutputWithContext(ctx context.Context) QuarantineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantineMapOutput)
}

type QuarantineOutput struct{ *pulumi.OutputState }

func (QuarantineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Quarantine)(nil)).Elem()
}

func (o QuarantineOutput) ToQuarantineOutput() QuarantineOutput {
	return o
}

func (o QuarantineOutput) ToQuarantineOutputWithContext(ctx context.Context) QuarantineOutput {
	return o
}

// Age limit for quarantined files (0 - 479 hours, 0 means forever).
func (o QuarantineOutput) Agelimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.IntOutput { return v.Agelimit }).(pulumi.IntOutput)
}

// Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL`, `disk`, `FortiAnalyzer`.
func (o QuarantineOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o QuarantineOutput) DropBlocked() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.DropBlocked }).(pulumi.StringOutput)
}

// Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o QuarantineOutput) DropHeuristic() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.DropHeuristic }).(pulumi.StringOutput)
}

// Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
func (o QuarantineOutput) DropInfected() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.DropInfected }).(pulumi.StringOutput)
}

// Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
func (o QuarantineOutput) DropMachineLearning() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.DropMachineLearning }).(pulumi.StringOutput)
}

// Select the method for handling additional files when running low on disk space. Valid values: `drop-new`, `ovrw-old`.
func (o QuarantineOutput) Lowspace() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.Lowspace }).(pulumi.StringOutput)
}

// Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
func (o QuarantineOutput) Maxfilesize() pulumi.IntOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.IntOutput { return v.Maxfilesize }).(pulumi.IntOutput)
}

// The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
func (o QuarantineOutput) QuarantineQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.IntOutput { return v.QuarantineQuota }).(pulumi.IntOutput)
}

// Quarantine blocked files found in sessions using the selected protocols.
func (o QuarantineOutput) StoreBlocked() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.StoreBlocked }).(pulumi.StringOutput)
}

// Quarantine files detected by heuristics found in sessions using the selected protocols.
func (o QuarantineOutput) StoreHeuristic() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.StoreHeuristic }).(pulumi.StringOutput)
}

// Quarantine infected files found in sessions using the selected protocols.
func (o QuarantineOutput) StoreInfected() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.StoreInfected }).(pulumi.StringOutput)
}

// Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `mapi`, `cifs`, `ssh`.
func (o QuarantineOutput) StoreMachineLearning() pulumi.StringOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringOutput { return v.StoreMachineLearning }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o QuarantineOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Quarantine) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type QuarantineArrayOutput struct{ *pulumi.OutputState }

func (QuarantineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Quarantine)(nil)).Elem()
}

func (o QuarantineArrayOutput) ToQuarantineArrayOutput() QuarantineArrayOutput {
	return o
}

func (o QuarantineArrayOutput) ToQuarantineArrayOutputWithContext(ctx context.Context) QuarantineArrayOutput {
	return o
}

func (o QuarantineArrayOutput) Index(i pulumi.IntInput) QuarantineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Quarantine {
		return vs[0].([]*Quarantine)[vs[1].(int)]
	}).(QuarantineOutput)
}

type QuarantineMapOutput struct{ *pulumi.OutputState }

func (QuarantineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Quarantine)(nil)).Elem()
}

func (o QuarantineMapOutput) ToQuarantineMapOutput() QuarantineMapOutput {
	return o
}

func (o QuarantineMapOutput) ToQuarantineMapOutputWithContext(ctx context.Context) QuarantineMapOutput {
	return o
}

func (o QuarantineMapOutput) MapIndex(k pulumi.StringInput) QuarantineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Quarantine {
		return vs[0].(map[string]*Quarantine)[vs[1].(string)]
	}).(QuarantineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineInput)(nil)).Elem(), &Quarantine{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineArrayInput)(nil)).Elem(), QuarantineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuarantineMapInput)(nil)).Elem(), QuarantineMap{})
	pulumi.RegisterOutputType(QuarantineOutput{})
	pulumi.RegisterOutputType(QuarantineArrayOutput{})
	pulumi.RegisterOutputType(QuarantineMapOutput{})
}
