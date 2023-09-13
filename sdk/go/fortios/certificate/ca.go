// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CA certificate.
//
// ## Import
//
// # Certificate Ca can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:certificate/ca:Ca labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:certificate/ca:Ca labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Ca struct {
	pulumi.CustomResourceState

	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntOutput `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntOutput `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca pulumi.StringOutput `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringOutput `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated pulumi.IntOutput `pulumi:"lastUpdated"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringOutput `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl pulumi.StringOutput `pulumi:"scepUrl"`
	// CA certificate source type.
	Source pulumi.StringOutput `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringOutput `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringOutput `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCa registers a new resource with the given unique name, arguments, and options.
func NewCa(ctx *pulumi.Context,
	name string, args *CaArgs, opts ...pulumi.ResourceOption) (*Ca, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ca == nil {
		return nil, errors.New("invalid value for required argument 'Ca'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Ca
	err := ctx.RegisterResource("fortios:certificate/ca:Ca", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCa gets an existing Ca resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaState, opts ...pulumi.ResourceOption) (*Ca, error) {
	var resource Ca
	err := ctx.ReadResource("fortios:certificate/ca:Ca", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ca resources.
type caState struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays *int `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning *int `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca *string `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier *string `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete *string `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl *string `pulumi:"scepUrl"`
	// CA certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted *string `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted *string `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CaState struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntPtrInput
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntPtrInput
	// CA certificate as a PEM file.
	Ca pulumi.StringPtrInput
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringPtrInput
	// Time at which CA was last updated.
	LastUpdated pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringPtrInput
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// URL of the SCEP server.
	ScepUrl pulumi.StringPtrInput
	// CA certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringPtrInput
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CaState) ElementType() reflect.Type {
	return reflect.TypeOf((*caState)(nil)).Elem()
}

type caArgs struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays *int `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning *int `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca string `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier *string `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete *string `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl *string `pulumi:"scepUrl"`
	// CA certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted *string `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted *string `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ca resource.
type CaArgs struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntPtrInput
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntPtrInput
	// CA certificate as a PEM file.
	Ca pulumi.StringInput
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringPtrInput
	// Time at which CA was last updated.
	LastUpdated pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringPtrInput
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// URL of the SCEP server.
	ScepUrl pulumi.StringPtrInput
	// CA certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringPtrInput
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caArgs)(nil)).Elem()
}

type CaInput interface {
	pulumi.Input

	ToCaOutput() CaOutput
	ToCaOutputWithContext(ctx context.Context) CaOutput
}

func (*Ca) ElementType() reflect.Type {
	return reflect.TypeOf((**Ca)(nil)).Elem()
}

func (i *Ca) ToCaOutput() CaOutput {
	return i.ToCaOutputWithContext(context.Background())
}

func (i *Ca) ToCaOutputWithContext(ctx context.Context) CaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaOutput)
}

// CaArrayInput is an input type that accepts CaArray and CaArrayOutput values.
// You can construct a concrete instance of `CaArrayInput` via:
//
//	CaArray{ CaArgs{...} }
type CaArrayInput interface {
	pulumi.Input

	ToCaArrayOutput() CaArrayOutput
	ToCaArrayOutputWithContext(context.Context) CaArrayOutput
}

type CaArray []CaInput

func (CaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ca)(nil)).Elem()
}

func (i CaArray) ToCaArrayOutput() CaArrayOutput {
	return i.ToCaArrayOutputWithContext(context.Background())
}

func (i CaArray) ToCaArrayOutputWithContext(ctx context.Context) CaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaArrayOutput)
}

// CaMapInput is an input type that accepts CaMap and CaMapOutput values.
// You can construct a concrete instance of `CaMapInput` via:
//
//	CaMap{ "key": CaArgs{...} }
type CaMapInput interface {
	pulumi.Input

	ToCaMapOutput() CaMapOutput
	ToCaMapOutputWithContext(context.Context) CaMapOutput
}

type CaMap map[string]CaInput

func (CaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ca)(nil)).Elem()
}

func (i CaMap) ToCaMapOutput() CaMapOutput {
	return i.ToCaMapOutputWithContext(context.Background())
}

func (i CaMap) ToCaMapOutputWithContext(ctx context.Context) CaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaMapOutput)
}

type CaOutput struct{ *pulumi.OutputState }

func (CaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ca)(nil)).Elem()
}

func (o CaOutput) ToCaOutput() CaOutput {
	return o
}

func (o CaOutput) ToCaOutputWithContext(ctx context.Context) CaOutput {
	return o
}

// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
func (o CaOutput) AutoUpdateDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Ca) pulumi.IntOutput { return v.AutoUpdateDays }).(pulumi.IntOutput)
}

// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
func (o CaOutput) AutoUpdateDaysWarning() pulumi.IntOutput {
	return o.ApplyT(func(v *Ca) pulumi.IntOutput { return v.AutoUpdateDaysWarning }).(pulumi.IntOutput)
}

// CA certificate as a PEM file.
func (o CaOutput) Ca() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Ca }).(pulumi.StringOutput)
}

// CA identifier of the SCEP server.
func (o CaOutput) CaIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.CaIdentifier }).(pulumi.StringOutput)
}

// Time at which CA was last updated.
func (o CaOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *Ca) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

// Name.
func (o CaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
func (o CaOutput) Obsolete() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Obsolete }).(pulumi.StringOutput)
}

// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
func (o CaOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

// URL of the SCEP server.
func (o CaOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

// CA certificate source type.
func (o CaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Source IP address for communications to the SCEP server.
func (o CaOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
func (o CaOutput) SslInspectionTrusted() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.SslInspectionTrusted }).(pulumi.StringOutput)
}

// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
func (o CaOutput) Trusted() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.Trusted }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CaArrayOutput struct{ *pulumi.OutputState }

func (CaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ca)(nil)).Elem()
}

func (o CaArrayOutput) ToCaArrayOutput() CaArrayOutput {
	return o
}

func (o CaArrayOutput) ToCaArrayOutputWithContext(ctx context.Context) CaArrayOutput {
	return o
}

func (o CaArrayOutput) Index(i pulumi.IntInput) CaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ca {
		return vs[0].([]*Ca)[vs[1].(int)]
	}).(CaOutput)
}

type CaMapOutput struct{ *pulumi.OutputState }

func (CaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ca)(nil)).Elem()
}

func (o CaMapOutput) ToCaMapOutput() CaMapOutput {
	return o
}

func (o CaMapOutput) ToCaMapOutputWithContext(ctx context.Context) CaMapOutput {
	return o
}

func (o CaMapOutput) MapIndex(k pulumi.StringInput) CaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ca {
		return vs[0].(map[string]*Ca)[vs[1].(string)]
	}).(CaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaInput)(nil)).Elem(), &Ca{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaArrayInput)(nil)).Elem(), CaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaMapInput)(nil)).Elem(), CaMap{})
	pulumi.RegisterOutputType(CaOutput{})
	pulumi.RegisterOutputType(CaArrayOutput{})
	pulumi.RegisterOutputType(CaMapOutput{})
}
