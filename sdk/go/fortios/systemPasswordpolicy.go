// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
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
//			_, err := fortios.NewSystemPasswordpolicy(ctx, "trname", &fortios.SystemPasswordpolicyArgs{
//				ApplyTo:            pulumi.String("admin-password"),
//				Change4Characters:  pulumi.String("disable"),
//				ExpireDay:          pulumi.Int(90),
//				ExpireStatus:       pulumi.String("disable"),
//				MinLowerCaseLetter: pulumi.Int(0),
//				MinNonAlphanumeric: pulumi.Int(0),
//				MinNumber:          pulumi.Int(0),
//				MinUpperCaseLetter: pulumi.Int(0),
//				MinimumLength:      pulumi.Int(8),
//				ReusePassword:      pulumi.String("enable"),
//				Status:             pulumi.String("disable"),
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
// # System PasswordPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemPasswordpolicy:SystemPasswordpolicy labelname SystemPasswordPolicy
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemPasswordpolicy:SystemPasswordpolicy labelname SystemPasswordPolicy
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemPasswordpolicy struct {
	pulumi.CustomResourceState

	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringOutput `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringOutput `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntOutput `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringOutput `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntOutput `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntOutput `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntOutput `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntOutput `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntOutput `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntOutput `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringOutput `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemPasswordpolicy registers a new resource with the given unique name, arguments, and options.
func NewSystemPasswordpolicy(ctx *pulumi.Context,
	name string, args *SystemPasswordpolicyArgs, opts ...pulumi.ResourceOption) (*SystemPasswordpolicy, error) {
	if args == nil {
		args = &SystemPasswordpolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemPasswordpolicy
	err := ctx.RegisterResource("fortios:index/systemPasswordpolicy:SystemPasswordpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemPasswordpolicy gets an existing SystemPasswordpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemPasswordpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemPasswordpolicyState, opts ...pulumi.ResourceOption) (*SystemPasswordpolicy, error) {
	var resource SystemPasswordpolicy
	err := ctx.ReadResource("fortios:index/systemPasswordpolicy:SystemPasswordpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemPasswordpolicy resources.
type systemPasswordpolicyState struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo *string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters *string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay *int `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus *string `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters *int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter *int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric *int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber *int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter *int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength *int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword *string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemPasswordpolicyState struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringPtrInput
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringPtrInput
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntPtrInput
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringPtrInput
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntPtrInput
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntPtrInput
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntPtrInput
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntPtrInput
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntPtrInput
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntPtrInput
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringPtrInput
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemPasswordpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordpolicyState)(nil)).Elem()
}

type systemPasswordpolicyArgs struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo *string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters *string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay *int `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus *string `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters *int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter *int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric *int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber *int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter *int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength *int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword *string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemPasswordpolicy resource.
type SystemPasswordpolicyArgs struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringPtrInput
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringPtrInput
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntPtrInput
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringPtrInput
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntPtrInput
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntPtrInput
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntPtrInput
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntPtrInput
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntPtrInput
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntPtrInput
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringPtrInput
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemPasswordpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordpolicyArgs)(nil)).Elem()
}

type SystemPasswordpolicyInput interface {
	pulumi.Input

	ToSystemPasswordpolicyOutput() SystemPasswordpolicyOutput
	ToSystemPasswordpolicyOutputWithContext(ctx context.Context) SystemPasswordpolicyOutput
}

func (*SystemPasswordpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordpolicy)(nil)).Elem()
}

func (i *SystemPasswordpolicy) ToSystemPasswordpolicyOutput() SystemPasswordpolicyOutput {
	return i.ToSystemPasswordpolicyOutputWithContext(context.Background())
}

func (i *SystemPasswordpolicy) ToSystemPasswordpolicyOutputWithContext(ctx context.Context) SystemPasswordpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordpolicyOutput)
}

// SystemPasswordpolicyArrayInput is an input type that accepts SystemPasswordpolicyArray and SystemPasswordpolicyArrayOutput values.
// You can construct a concrete instance of `SystemPasswordpolicyArrayInput` via:
//
//	SystemPasswordpolicyArray{ SystemPasswordpolicyArgs{...} }
type SystemPasswordpolicyArrayInput interface {
	pulumi.Input

	ToSystemPasswordpolicyArrayOutput() SystemPasswordpolicyArrayOutput
	ToSystemPasswordpolicyArrayOutputWithContext(context.Context) SystemPasswordpolicyArrayOutput
}

type SystemPasswordpolicyArray []SystemPasswordpolicyInput

func (SystemPasswordpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordpolicy)(nil)).Elem()
}

func (i SystemPasswordpolicyArray) ToSystemPasswordpolicyArrayOutput() SystemPasswordpolicyArrayOutput {
	return i.ToSystemPasswordpolicyArrayOutputWithContext(context.Background())
}

func (i SystemPasswordpolicyArray) ToSystemPasswordpolicyArrayOutputWithContext(ctx context.Context) SystemPasswordpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordpolicyArrayOutput)
}

// SystemPasswordpolicyMapInput is an input type that accepts SystemPasswordpolicyMap and SystemPasswordpolicyMapOutput values.
// You can construct a concrete instance of `SystemPasswordpolicyMapInput` via:
//
//	SystemPasswordpolicyMap{ "key": SystemPasswordpolicyArgs{...} }
type SystemPasswordpolicyMapInput interface {
	pulumi.Input

	ToSystemPasswordpolicyMapOutput() SystemPasswordpolicyMapOutput
	ToSystemPasswordpolicyMapOutputWithContext(context.Context) SystemPasswordpolicyMapOutput
}

type SystemPasswordpolicyMap map[string]SystemPasswordpolicyInput

func (SystemPasswordpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordpolicy)(nil)).Elem()
}

func (i SystemPasswordpolicyMap) ToSystemPasswordpolicyMapOutput() SystemPasswordpolicyMapOutput {
	return i.ToSystemPasswordpolicyMapOutputWithContext(context.Background())
}

func (i SystemPasswordpolicyMap) ToSystemPasswordpolicyMapOutputWithContext(ctx context.Context) SystemPasswordpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordpolicyMapOutput)
}

type SystemPasswordpolicyOutput struct{ *pulumi.OutputState }

func (SystemPasswordpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordpolicy)(nil)).Elem()
}

func (o SystemPasswordpolicyOutput) ToSystemPasswordpolicyOutput() SystemPasswordpolicyOutput {
	return o
}

func (o SystemPasswordpolicyOutput) ToSystemPasswordpolicyOutputWithContext(ctx context.Context) SystemPasswordpolicyOutput {
	return o
}

// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
func (o SystemPasswordpolicyOutput) ApplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringOutput { return v.ApplyTo }).(pulumi.StringOutput)
}

// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
func (o SystemPasswordpolicyOutput) Change4Characters() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringOutput { return v.Change4Characters }).(pulumi.StringOutput)
}

// Number of days after which passwords expire (1 - 999 days, default = 90).
func (o SystemPasswordpolicyOutput) ExpireDay() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.ExpireDay }).(pulumi.IntOutput)
}

// Enable/disable password expiration. Valid values: `enable`, `disable`.
func (o SystemPasswordpolicyOutput) ExpireStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringOutput { return v.ExpireStatus }).(pulumi.StringOutput)
}

// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
func (o SystemPasswordpolicyOutput) MinChangeCharacters() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinChangeCharacters }).(pulumi.IntOutput)
}

// Minimum number of lowercase characters in password (0 - 128, default = 0).
func (o SystemPasswordpolicyOutput) MinLowerCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinLowerCaseLetter }).(pulumi.IntOutput)
}

// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
func (o SystemPasswordpolicyOutput) MinNonAlphanumeric() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinNonAlphanumeric }).(pulumi.IntOutput)
}

// Minimum number of numeric characters in password (0 - 128, default = 0).
func (o SystemPasswordpolicyOutput) MinNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinNumber }).(pulumi.IntOutput)
}

// Minimum number of uppercase characters in password (0 - 128, default = 0).
func (o SystemPasswordpolicyOutput) MinUpperCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinUpperCaseLetter }).(pulumi.IntOutput)
}

// Minimum password length (8 - 128, default = 8).
func (o SystemPasswordpolicyOutput) MinimumLength() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.IntOutput { return v.MinimumLength }).(pulumi.IntOutput)
}

// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
func (o SystemPasswordpolicyOutput) ReusePassword() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringOutput { return v.ReusePassword }).(pulumi.StringOutput)
}

// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
func (o SystemPasswordpolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemPasswordpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemPasswordpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemPasswordpolicyArrayOutput struct{ *pulumi.OutputState }

func (SystemPasswordpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordpolicy)(nil)).Elem()
}

func (o SystemPasswordpolicyArrayOutput) ToSystemPasswordpolicyArrayOutput() SystemPasswordpolicyArrayOutput {
	return o
}

func (o SystemPasswordpolicyArrayOutput) ToSystemPasswordpolicyArrayOutputWithContext(ctx context.Context) SystemPasswordpolicyArrayOutput {
	return o
}

func (o SystemPasswordpolicyArrayOutput) Index(i pulumi.IntInput) SystemPasswordpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemPasswordpolicy {
		return vs[0].([]*SystemPasswordpolicy)[vs[1].(int)]
	}).(SystemPasswordpolicyOutput)
}

type SystemPasswordpolicyMapOutput struct{ *pulumi.OutputState }

func (SystemPasswordpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordpolicy)(nil)).Elem()
}

func (o SystemPasswordpolicyMapOutput) ToSystemPasswordpolicyMapOutput() SystemPasswordpolicyMapOutput {
	return o
}

func (o SystemPasswordpolicyMapOutput) ToSystemPasswordpolicyMapOutputWithContext(ctx context.Context) SystemPasswordpolicyMapOutput {
	return o
}

func (o SystemPasswordpolicyMapOutput) MapIndex(k pulumi.StringInput) SystemPasswordpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemPasswordpolicy {
		return vs[0].(map[string]*SystemPasswordpolicy)[vs[1].(string)]
	}).(SystemPasswordpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordpolicyInput)(nil)).Elem(), &SystemPasswordpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordpolicyArrayInput)(nil)).Elem(), SystemPasswordpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordpolicyMapInput)(nil)).Elem(), SystemPasswordpolicyMap{})
	pulumi.RegisterOutputType(SystemPasswordpolicyOutput{})
	pulumi.RegisterOutputType(SystemPasswordpolicyArrayOutput{})
	pulumi.RegisterOutputType(SystemPasswordpolicyMapOutput{})
}