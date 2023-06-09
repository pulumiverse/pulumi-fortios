// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource to sort firewall security policies by policyid or policy name, in ascending or descending order.
//
// ## Example Usage
// ### Example1
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
//			test, err := fortios.NewFirewallSecurityPolicysort(ctx, "test", &fortios.FirewallSecurityPolicysortArgs{
//				Sortby:        pulumi.String("policyid"),
//				Sortdirection: pulumi.String("descending"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("policylistAfterApply", test.StatePolicyLists)
//			return nil
//		})
//	}
//
// ```
type FirewallSecurityPolicysort struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrOutput `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringOutput `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    pulumi.StringOutput                                  `pulumi:"sortdirection"`
	StatePolicyLists FirewallSecurityPolicysortStatePolicyListArrayOutput `pulumi:"statePolicyLists"`
	Status           pulumi.StringPtrOutput                               `pulumi:"status"`
	Vdomparam        pulumi.StringPtrOutput                               `pulumi:"vdomparam"`
}

// NewFirewallSecurityPolicysort registers a new resource with the given unique name, arguments, and options.
func NewFirewallSecurityPolicysort(ctx *pulumi.Context,
	name string, args *FirewallSecurityPolicysortArgs, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicysort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sortby == nil {
		return nil, errors.New("invalid value for required argument 'Sortby'")
	}
	if args.Sortdirection == nil {
		return nil, errors.New("invalid value for required argument 'Sortdirection'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSecurityPolicysort
	err := ctx.RegisterResource("fortios:index/firewallSecurityPolicysort:FirewallSecurityPolicysort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSecurityPolicysort gets an existing FirewallSecurityPolicysort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSecurityPolicysort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSecurityPolicysortState, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicysort, error) {
	var resource FirewallSecurityPolicysort
	err := ctx.ReadResource("fortios:index/firewallSecurityPolicysort:FirewallSecurityPolicysort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSecurityPolicysort resources.
type firewallSecurityPolicysortState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate *string `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby *string `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    *string                                     `pulumi:"sortdirection"`
	StatePolicyLists []FirewallSecurityPolicysortStatePolicyList `pulumi:"statePolicyLists"`
	Status           *string                                     `pulumi:"status"`
	Vdomparam        *string                                     `pulumi:"vdomparam"`
}

type FirewallSecurityPolicysortState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrInput
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringPtrInput
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    pulumi.StringPtrInput
	StatePolicyLists FirewallSecurityPolicysortStatePolicyListArrayInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (FirewallSecurityPolicysortState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicysortState)(nil)).Elem()
}

type firewallSecurityPolicysortArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate *string `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby string `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection string  `pulumi:"sortdirection"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSecurityPolicysort resource.
type FirewallSecurityPolicysortArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrInput
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringInput
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection pulumi.StringInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
}

func (FirewallSecurityPolicysortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicysortArgs)(nil)).Elem()
}

type FirewallSecurityPolicysortInput interface {
	pulumi.Input

	ToFirewallSecurityPolicysortOutput() FirewallSecurityPolicysortOutput
	ToFirewallSecurityPolicysortOutputWithContext(ctx context.Context) FirewallSecurityPolicysortOutput
}

func (*FirewallSecurityPolicysort) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicysort)(nil)).Elem()
}

func (i *FirewallSecurityPolicysort) ToFirewallSecurityPolicysortOutput() FirewallSecurityPolicysortOutput {
	return i.ToFirewallSecurityPolicysortOutputWithContext(context.Background())
}

func (i *FirewallSecurityPolicysort) ToFirewallSecurityPolicysortOutputWithContext(ctx context.Context) FirewallSecurityPolicysortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicysortOutput)
}

// FirewallSecurityPolicysortArrayInput is an input type that accepts FirewallSecurityPolicysortArray and FirewallSecurityPolicysortArrayOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicysortArrayInput` via:
//
//	FirewallSecurityPolicysortArray{ FirewallSecurityPolicysortArgs{...} }
type FirewallSecurityPolicysortArrayInput interface {
	pulumi.Input

	ToFirewallSecurityPolicysortArrayOutput() FirewallSecurityPolicysortArrayOutput
	ToFirewallSecurityPolicysortArrayOutputWithContext(context.Context) FirewallSecurityPolicysortArrayOutput
}

type FirewallSecurityPolicysortArray []FirewallSecurityPolicysortInput

func (FirewallSecurityPolicysortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicysort)(nil)).Elem()
}

func (i FirewallSecurityPolicysortArray) ToFirewallSecurityPolicysortArrayOutput() FirewallSecurityPolicysortArrayOutput {
	return i.ToFirewallSecurityPolicysortArrayOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicysortArray) ToFirewallSecurityPolicysortArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicysortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicysortArrayOutput)
}

// FirewallSecurityPolicysortMapInput is an input type that accepts FirewallSecurityPolicysortMap and FirewallSecurityPolicysortMapOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicysortMapInput` via:
//
//	FirewallSecurityPolicysortMap{ "key": FirewallSecurityPolicysortArgs{...} }
type FirewallSecurityPolicysortMapInput interface {
	pulumi.Input

	ToFirewallSecurityPolicysortMapOutput() FirewallSecurityPolicysortMapOutput
	ToFirewallSecurityPolicysortMapOutputWithContext(context.Context) FirewallSecurityPolicysortMapOutput
}

type FirewallSecurityPolicysortMap map[string]FirewallSecurityPolicysortInput

func (FirewallSecurityPolicysortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicysort)(nil)).Elem()
}

func (i FirewallSecurityPolicysortMap) ToFirewallSecurityPolicysortMapOutput() FirewallSecurityPolicysortMapOutput {
	return i.ToFirewallSecurityPolicysortMapOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicysortMap) ToFirewallSecurityPolicysortMapOutputWithContext(ctx context.Context) FirewallSecurityPolicysortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicysortMapOutput)
}

type FirewallSecurityPolicysortOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicysortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicysort)(nil)).Elem()
}

func (o FirewallSecurityPolicysortOutput) ToFirewallSecurityPolicysortOutput() FirewallSecurityPolicysortOutput {
	return o
}

func (o FirewallSecurityPolicysortOutput) ToFirewallSecurityPolicysortOutputWithContext(ctx context.Context) FirewallSecurityPolicysortOutput {
	return o
}

// Comment.
func (o FirewallSecurityPolicysortOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
func (o FirewallSecurityPolicysortOutput) ForceRecreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringPtrOutput { return v.ForceRecreate }).(pulumi.StringPtrOutput)
}

// Sort security policies by the value, it currently supports "policyid" and "name".
func (o FirewallSecurityPolicysortOutput) Sortby() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringOutput { return v.Sortby }).(pulumi.StringOutput)
}

// Sort dirction, supports "ascending" and "descending".
func (o FirewallSecurityPolicysortOutput) Sortdirection() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringOutput { return v.Sortdirection }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicysortOutput) StatePolicyLists() FirewallSecurityPolicysortStatePolicyListArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) FirewallSecurityPolicysortStatePolicyListArrayOutput {
		return v.StatePolicyLists
	}).(FirewallSecurityPolicysortStatePolicyListArrayOutput)
}

func (o FirewallSecurityPolicysortOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o FirewallSecurityPolicysortOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicysort) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSecurityPolicysortArrayOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicysortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicysort)(nil)).Elem()
}

func (o FirewallSecurityPolicysortArrayOutput) ToFirewallSecurityPolicysortArrayOutput() FirewallSecurityPolicysortArrayOutput {
	return o
}

func (o FirewallSecurityPolicysortArrayOutput) ToFirewallSecurityPolicysortArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicysortArrayOutput {
	return o
}

func (o FirewallSecurityPolicysortArrayOutput) Index(i pulumi.IntInput) FirewallSecurityPolicysortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSecurityPolicysort {
		return vs[0].([]*FirewallSecurityPolicysort)[vs[1].(int)]
	}).(FirewallSecurityPolicysortOutput)
}

type FirewallSecurityPolicysortMapOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicysortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicysort)(nil)).Elem()
}

func (o FirewallSecurityPolicysortMapOutput) ToFirewallSecurityPolicysortMapOutput() FirewallSecurityPolicysortMapOutput {
	return o
}

func (o FirewallSecurityPolicysortMapOutput) ToFirewallSecurityPolicysortMapOutputWithContext(ctx context.Context) FirewallSecurityPolicysortMapOutput {
	return o
}

func (o FirewallSecurityPolicysortMapOutput) MapIndex(k pulumi.StringInput) FirewallSecurityPolicysortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSecurityPolicysort {
		return vs[0].(map[string]*FirewallSecurityPolicysort)[vs[1].(string)]
	}).(FirewallSecurityPolicysortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicysortInput)(nil)).Elem(), &FirewallSecurityPolicysort{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicysortArrayInput)(nil)).Elem(), FirewallSecurityPolicysortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicysortMapInput)(nil)).Elem(), FirewallSecurityPolicysortMap{})
	pulumi.RegisterOutputType(FirewallSecurityPolicysortOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicysortArrayOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicysortMapOutput{})
}
