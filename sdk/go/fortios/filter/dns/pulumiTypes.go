// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

var _ = internal.GetEnvOrDefault

type DomainfilterEntry struct {
	// Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
	Action *string `pulumi:"action"`
	// Domain entries to be filtered.
	Domain *string `pulumi:"domain"`
	// Id.
	Id *int `pulumi:"id"`
	// Enable/disable this domain filter. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
	Type *string `pulumi:"type"`
}

// DomainfilterEntryInput is an input type that accepts DomainfilterEntryArgs and DomainfilterEntryOutput values.
// You can construct a concrete instance of `DomainfilterEntryInput` via:
//
//	DomainfilterEntryArgs{...}
type DomainfilterEntryInput interface {
	pulumi.Input

	ToDomainfilterEntryOutput() DomainfilterEntryOutput
	ToDomainfilterEntryOutputWithContext(context.Context) DomainfilterEntryOutput
}

type DomainfilterEntryArgs struct {
	// Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// Domain entries to be filtered.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// Id.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Enable/disable this domain filter. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (DomainfilterEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainfilterEntry)(nil)).Elem()
}

func (i DomainfilterEntryArgs) ToDomainfilterEntryOutput() DomainfilterEntryOutput {
	return i.ToDomainfilterEntryOutputWithContext(context.Background())
}

func (i DomainfilterEntryArgs) ToDomainfilterEntryOutputWithContext(ctx context.Context) DomainfilterEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainfilterEntryOutput)
}

// DomainfilterEntryArrayInput is an input type that accepts DomainfilterEntryArray and DomainfilterEntryArrayOutput values.
// You can construct a concrete instance of `DomainfilterEntryArrayInput` via:
//
//	DomainfilterEntryArray{ DomainfilterEntryArgs{...} }
type DomainfilterEntryArrayInput interface {
	pulumi.Input

	ToDomainfilterEntryArrayOutput() DomainfilterEntryArrayOutput
	ToDomainfilterEntryArrayOutputWithContext(context.Context) DomainfilterEntryArrayOutput
}

type DomainfilterEntryArray []DomainfilterEntryInput

func (DomainfilterEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainfilterEntry)(nil)).Elem()
}

func (i DomainfilterEntryArray) ToDomainfilterEntryArrayOutput() DomainfilterEntryArrayOutput {
	return i.ToDomainfilterEntryArrayOutputWithContext(context.Background())
}

func (i DomainfilterEntryArray) ToDomainfilterEntryArrayOutputWithContext(ctx context.Context) DomainfilterEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainfilterEntryArrayOutput)
}

type DomainfilterEntryOutput struct{ *pulumi.OutputState }

func (DomainfilterEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainfilterEntry)(nil)).Elem()
}

func (o DomainfilterEntryOutput) ToDomainfilterEntryOutput() DomainfilterEntryOutput {
	return o
}

func (o DomainfilterEntryOutput) ToDomainfilterEntryOutputWithContext(ctx context.Context) DomainfilterEntryOutput {
	return o
}

// Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
func (o DomainfilterEntryOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainfilterEntry) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// Domain entries to be filtered.
func (o DomainfilterEntryOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainfilterEntry) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// Id.
func (o DomainfilterEntryOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainfilterEntry) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Enable/disable this domain filter. Valid values: `enable`, `disable`.
func (o DomainfilterEntryOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainfilterEntry) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
func (o DomainfilterEntryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainfilterEntry) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DomainfilterEntryArrayOutput struct{ *pulumi.OutputState }

func (DomainfilterEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainfilterEntry)(nil)).Elem()
}

func (o DomainfilterEntryArrayOutput) ToDomainfilterEntryArrayOutput() DomainfilterEntryArrayOutput {
	return o
}

func (o DomainfilterEntryArrayOutput) ToDomainfilterEntryArrayOutputWithContext(ctx context.Context) DomainfilterEntryArrayOutput {
	return o
}

func (o DomainfilterEntryArrayOutput) Index(i pulumi.IntInput) DomainfilterEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainfilterEntry {
		return vs[0].([]DomainfilterEntry)[vs[1].(int)]
	}).(DomainfilterEntryOutput)
}

type ProfileDnsTranslation struct {
	// DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
	AddrType *string `pulumi:"addrType"`
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst *string `pulumi:"dst"`
	// IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
	Dst6 *string `pulumi:"dst6"`
	// ID.
	Id *int `pulumi:"id"`
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask *string `pulumi:"netmask"`
	// If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
	Prefix *int `pulumi:"prefix"`
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src *string `pulumi:"src"`
	// IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
	Src6 *string `pulumi:"src6"`
	// Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
}

// ProfileDnsTranslationInput is an input type that accepts ProfileDnsTranslationArgs and ProfileDnsTranslationOutput values.
// You can construct a concrete instance of `ProfileDnsTranslationInput` via:
//
//	ProfileDnsTranslationArgs{...}
type ProfileDnsTranslationInput interface {
	pulumi.Input

	ToProfileDnsTranslationOutput() ProfileDnsTranslationOutput
	ToProfileDnsTranslationOutputWithContext(context.Context) ProfileDnsTranslationOutput
}

type ProfileDnsTranslationArgs struct {
	// DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
	AddrType pulumi.StringPtrInput `pulumi:"addrType"`
	// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
	Dst pulumi.StringPtrInput `pulumi:"dst"`
	// IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
	Dst6 pulumi.StringPtrInput `pulumi:"dst6"`
	// ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
	Netmask pulumi.StringPtrInput `pulumi:"netmask"`
	// If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
	Prefix pulumi.IntPtrInput `pulumi:"prefix"`
	// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
	Src pulumi.StringPtrInput `pulumi:"src"`
	// IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
	Src6 pulumi.StringPtrInput `pulumi:"src6"`
	// Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ProfileDnsTranslationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDnsTranslation)(nil)).Elem()
}

func (i ProfileDnsTranslationArgs) ToProfileDnsTranslationOutput() ProfileDnsTranslationOutput {
	return i.ToProfileDnsTranslationOutputWithContext(context.Background())
}

func (i ProfileDnsTranslationArgs) ToProfileDnsTranslationOutputWithContext(ctx context.Context) ProfileDnsTranslationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDnsTranslationOutput)
}

// ProfileDnsTranslationArrayInput is an input type that accepts ProfileDnsTranslationArray and ProfileDnsTranslationArrayOutput values.
// You can construct a concrete instance of `ProfileDnsTranslationArrayInput` via:
//
//	ProfileDnsTranslationArray{ ProfileDnsTranslationArgs{...} }
type ProfileDnsTranslationArrayInput interface {
	pulumi.Input

	ToProfileDnsTranslationArrayOutput() ProfileDnsTranslationArrayOutput
	ToProfileDnsTranslationArrayOutputWithContext(context.Context) ProfileDnsTranslationArrayOutput
}

type ProfileDnsTranslationArray []ProfileDnsTranslationInput

func (ProfileDnsTranslationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileDnsTranslation)(nil)).Elem()
}

func (i ProfileDnsTranslationArray) ToProfileDnsTranslationArrayOutput() ProfileDnsTranslationArrayOutput {
	return i.ToProfileDnsTranslationArrayOutputWithContext(context.Background())
}

func (i ProfileDnsTranslationArray) ToProfileDnsTranslationArrayOutputWithContext(ctx context.Context) ProfileDnsTranslationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDnsTranslationArrayOutput)
}

type ProfileDnsTranslationOutput struct{ *pulumi.OutputState }

func (ProfileDnsTranslationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDnsTranslation)(nil)).Elem()
}

func (o ProfileDnsTranslationOutput) ToProfileDnsTranslationOutput() ProfileDnsTranslationOutput {
	return o
}

func (o ProfileDnsTranslationOutput) ToProfileDnsTranslationOutputWithContext(ctx context.Context) ProfileDnsTranslationOutput {
	return o
}

// DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
func (o ProfileDnsTranslationOutput) AddrType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.AddrType }).(pulumi.StringPtrOutput)
}

// IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
func (o ProfileDnsTranslationOutput) Dst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Dst }).(pulumi.StringPtrOutput)
}

// IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
func (o ProfileDnsTranslationOutput) Dst6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Dst6 }).(pulumi.StringPtrOutput)
}

// ID.
func (o ProfileDnsTranslationOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
func (o ProfileDnsTranslationOutput) Netmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Netmask }).(pulumi.StringPtrOutput)
}

// If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
func (o ProfileDnsTranslationOutput) Prefix() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *int { return v.Prefix }).(pulumi.IntPtrOutput)
}

// IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
func (o ProfileDnsTranslationOutput) Src() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Src }).(pulumi.StringPtrOutput)
}

// IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
func (o ProfileDnsTranslationOutput) Src6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Src6 }).(pulumi.StringPtrOutput)
}

// Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
func (o ProfileDnsTranslationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileDnsTranslation) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ProfileDnsTranslationArrayOutput struct{ *pulumi.OutputState }

func (ProfileDnsTranslationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileDnsTranslation)(nil)).Elem()
}

func (o ProfileDnsTranslationArrayOutput) ToProfileDnsTranslationArrayOutput() ProfileDnsTranslationArrayOutput {
	return o
}

func (o ProfileDnsTranslationArrayOutput) ToProfileDnsTranslationArrayOutputWithContext(ctx context.Context) ProfileDnsTranslationArrayOutput {
	return o
}

func (o ProfileDnsTranslationArrayOutput) Index(i pulumi.IntInput) ProfileDnsTranslationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileDnsTranslation {
		return vs[0].([]ProfileDnsTranslation)[vs[1].(int)]
	}).(ProfileDnsTranslationOutput)
}

type ProfileDomainFilter struct {
	// DNS domain filter table ID.
	DomainFilterTable *int `pulumi:"domainFilterTable"`
}

// ProfileDomainFilterInput is an input type that accepts ProfileDomainFilterArgs and ProfileDomainFilterOutput values.
// You can construct a concrete instance of `ProfileDomainFilterInput` via:
//
//	ProfileDomainFilterArgs{...}
type ProfileDomainFilterInput interface {
	pulumi.Input

	ToProfileDomainFilterOutput() ProfileDomainFilterOutput
	ToProfileDomainFilterOutputWithContext(context.Context) ProfileDomainFilterOutput
}

type ProfileDomainFilterArgs struct {
	// DNS domain filter table ID.
	DomainFilterTable pulumi.IntPtrInput `pulumi:"domainFilterTable"`
}

func (ProfileDomainFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDomainFilter)(nil)).Elem()
}

func (i ProfileDomainFilterArgs) ToProfileDomainFilterOutput() ProfileDomainFilterOutput {
	return i.ToProfileDomainFilterOutputWithContext(context.Background())
}

func (i ProfileDomainFilterArgs) ToProfileDomainFilterOutputWithContext(ctx context.Context) ProfileDomainFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDomainFilterOutput)
}

func (i ProfileDomainFilterArgs) ToProfileDomainFilterPtrOutput() ProfileDomainFilterPtrOutput {
	return i.ToProfileDomainFilterPtrOutputWithContext(context.Background())
}

func (i ProfileDomainFilterArgs) ToProfileDomainFilterPtrOutputWithContext(ctx context.Context) ProfileDomainFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDomainFilterOutput).ToProfileDomainFilterPtrOutputWithContext(ctx)
}

// ProfileDomainFilterPtrInput is an input type that accepts ProfileDomainFilterArgs, ProfileDomainFilterPtr and ProfileDomainFilterPtrOutput values.
// You can construct a concrete instance of `ProfileDomainFilterPtrInput` via:
//
//	        ProfileDomainFilterArgs{...}
//
//	or:
//
//	        nil
type ProfileDomainFilterPtrInput interface {
	pulumi.Input

	ToProfileDomainFilterPtrOutput() ProfileDomainFilterPtrOutput
	ToProfileDomainFilterPtrOutputWithContext(context.Context) ProfileDomainFilterPtrOutput
}

type profileDomainFilterPtrType ProfileDomainFilterArgs

func ProfileDomainFilterPtr(v *ProfileDomainFilterArgs) ProfileDomainFilterPtrInput {
	return (*profileDomainFilterPtrType)(v)
}

func (*profileDomainFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileDomainFilter)(nil)).Elem()
}

func (i *profileDomainFilterPtrType) ToProfileDomainFilterPtrOutput() ProfileDomainFilterPtrOutput {
	return i.ToProfileDomainFilterPtrOutputWithContext(context.Background())
}

func (i *profileDomainFilterPtrType) ToProfileDomainFilterPtrOutputWithContext(ctx context.Context) ProfileDomainFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileDomainFilterPtrOutput)
}

type ProfileDomainFilterOutput struct{ *pulumi.OutputState }

func (ProfileDomainFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileDomainFilter)(nil)).Elem()
}

func (o ProfileDomainFilterOutput) ToProfileDomainFilterOutput() ProfileDomainFilterOutput {
	return o
}

func (o ProfileDomainFilterOutput) ToProfileDomainFilterOutputWithContext(ctx context.Context) ProfileDomainFilterOutput {
	return o
}

func (o ProfileDomainFilterOutput) ToProfileDomainFilterPtrOutput() ProfileDomainFilterPtrOutput {
	return o.ToProfileDomainFilterPtrOutputWithContext(context.Background())
}

func (o ProfileDomainFilterOutput) ToProfileDomainFilterPtrOutputWithContext(ctx context.Context) ProfileDomainFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileDomainFilter) *ProfileDomainFilter {
		return &v
	}).(ProfileDomainFilterPtrOutput)
}

// DNS domain filter table ID.
func (o ProfileDomainFilterOutput) DomainFilterTable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileDomainFilter) *int { return v.DomainFilterTable }).(pulumi.IntPtrOutput)
}

type ProfileDomainFilterPtrOutput struct{ *pulumi.OutputState }

func (ProfileDomainFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileDomainFilter)(nil)).Elem()
}

func (o ProfileDomainFilterPtrOutput) ToProfileDomainFilterPtrOutput() ProfileDomainFilterPtrOutput {
	return o
}

func (o ProfileDomainFilterPtrOutput) ToProfileDomainFilterPtrOutputWithContext(ctx context.Context) ProfileDomainFilterPtrOutput {
	return o
}

func (o ProfileDomainFilterPtrOutput) Elem() ProfileDomainFilterOutput {
	return o.ApplyT(func(v *ProfileDomainFilter) ProfileDomainFilter {
		if v != nil {
			return *v
		}
		var ret ProfileDomainFilter
		return ret
	}).(ProfileDomainFilterOutput)
}

// DNS domain filter table ID.
func (o ProfileDomainFilterPtrOutput) DomainFilterTable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProfileDomainFilter) *int {
		if v == nil {
			return nil
		}
		return v.DomainFilterTable
	}).(pulumi.IntPtrOutput)
}

type ProfileExternalIpBlocklist struct {
	// External domain block list name.
	Name *string `pulumi:"name"`
}

// ProfileExternalIpBlocklistInput is an input type that accepts ProfileExternalIpBlocklistArgs and ProfileExternalIpBlocklistOutput values.
// You can construct a concrete instance of `ProfileExternalIpBlocklistInput` via:
//
//	ProfileExternalIpBlocklistArgs{...}
type ProfileExternalIpBlocklistInput interface {
	pulumi.Input

	ToProfileExternalIpBlocklistOutput() ProfileExternalIpBlocklistOutput
	ToProfileExternalIpBlocklistOutputWithContext(context.Context) ProfileExternalIpBlocklistOutput
}

type ProfileExternalIpBlocklistArgs struct {
	// External domain block list name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ProfileExternalIpBlocklistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileExternalIpBlocklist)(nil)).Elem()
}

func (i ProfileExternalIpBlocklistArgs) ToProfileExternalIpBlocklistOutput() ProfileExternalIpBlocklistOutput {
	return i.ToProfileExternalIpBlocklistOutputWithContext(context.Background())
}

func (i ProfileExternalIpBlocklistArgs) ToProfileExternalIpBlocklistOutputWithContext(ctx context.Context) ProfileExternalIpBlocklistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileExternalIpBlocklistOutput)
}

// ProfileExternalIpBlocklistArrayInput is an input type that accepts ProfileExternalIpBlocklistArray and ProfileExternalIpBlocklistArrayOutput values.
// You can construct a concrete instance of `ProfileExternalIpBlocklistArrayInput` via:
//
//	ProfileExternalIpBlocklistArray{ ProfileExternalIpBlocklistArgs{...} }
type ProfileExternalIpBlocklistArrayInput interface {
	pulumi.Input

	ToProfileExternalIpBlocklistArrayOutput() ProfileExternalIpBlocklistArrayOutput
	ToProfileExternalIpBlocklistArrayOutputWithContext(context.Context) ProfileExternalIpBlocklistArrayOutput
}

type ProfileExternalIpBlocklistArray []ProfileExternalIpBlocklistInput

func (ProfileExternalIpBlocklistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileExternalIpBlocklist)(nil)).Elem()
}

func (i ProfileExternalIpBlocklistArray) ToProfileExternalIpBlocklistArrayOutput() ProfileExternalIpBlocklistArrayOutput {
	return i.ToProfileExternalIpBlocklistArrayOutputWithContext(context.Background())
}

func (i ProfileExternalIpBlocklistArray) ToProfileExternalIpBlocklistArrayOutputWithContext(ctx context.Context) ProfileExternalIpBlocklistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileExternalIpBlocklistArrayOutput)
}

type ProfileExternalIpBlocklistOutput struct{ *pulumi.OutputState }

func (ProfileExternalIpBlocklistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileExternalIpBlocklist)(nil)).Elem()
}

func (o ProfileExternalIpBlocklistOutput) ToProfileExternalIpBlocklistOutput() ProfileExternalIpBlocklistOutput {
	return o
}

func (o ProfileExternalIpBlocklistOutput) ToProfileExternalIpBlocklistOutputWithContext(ctx context.Context) ProfileExternalIpBlocklistOutput {
	return o
}

// External domain block list name.
func (o ProfileExternalIpBlocklistOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileExternalIpBlocklist) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ProfileExternalIpBlocklistArrayOutput struct{ *pulumi.OutputState }

func (ProfileExternalIpBlocklistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileExternalIpBlocklist)(nil)).Elem()
}

func (o ProfileExternalIpBlocklistArrayOutput) ToProfileExternalIpBlocklistArrayOutput() ProfileExternalIpBlocklistArrayOutput {
	return o
}

func (o ProfileExternalIpBlocklistArrayOutput) ToProfileExternalIpBlocklistArrayOutputWithContext(ctx context.Context) ProfileExternalIpBlocklistArrayOutput {
	return o
}

func (o ProfileExternalIpBlocklistArrayOutput) Index(i pulumi.IntInput) ProfileExternalIpBlocklistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileExternalIpBlocklist {
		return vs[0].([]ProfileExternalIpBlocklist)[vs[1].(int)]
	}).(ProfileExternalIpBlocklistOutput)
}

type ProfileFtgdDns struct {
	// FortiGuard DNS domain filters. The structure of `filters` block is documented below.
	Filters []ProfileFtgdDnsFilter `pulumi:"filters"`
	// FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
	Options *string `pulumi:"options"`
}

// ProfileFtgdDnsInput is an input type that accepts ProfileFtgdDnsArgs and ProfileFtgdDnsOutput values.
// You can construct a concrete instance of `ProfileFtgdDnsInput` via:
//
//	ProfileFtgdDnsArgs{...}
type ProfileFtgdDnsInput interface {
	pulumi.Input

	ToProfileFtgdDnsOutput() ProfileFtgdDnsOutput
	ToProfileFtgdDnsOutputWithContext(context.Context) ProfileFtgdDnsOutput
}

type ProfileFtgdDnsArgs struct {
	// FortiGuard DNS domain filters. The structure of `filters` block is documented below.
	Filters ProfileFtgdDnsFilterArrayInput `pulumi:"filters"`
	// FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
	Options pulumi.StringPtrInput `pulumi:"options"`
}

func (ProfileFtgdDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFtgdDns)(nil)).Elem()
}

func (i ProfileFtgdDnsArgs) ToProfileFtgdDnsOutput() ProfileFtgdDnsOutput {
	return i.ToProfileFtgdDnsOutputWithContext(context.Background())
}

func (i ProfileFtgdDnsArgs) ToProfileFtgdDnsOutputWithContext(ctx context.Context) ProfileFtgdDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFtgdDnsOutput)
}

func (i ProfileFtgdDnsArgs) ToProfileFtgdDnsPtrOutput() ProfileFtgdDnsPtrOutput {
	return i.ToProfileFtgdDnsPtrOutputWithContext(context.Background())
}

func (i ProfileFtgdDnsArgs) ToProfileFtgdDnsPtrOutputWithContext(ctx context.Context) ProfileFtgdDnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFtgdDnsOutput).ToProfileFtgdDnsPtrOutputWithContext(ctx)
}

// ProfileFtgdDnsPtrInput is an input type that accepts ProfileFtgdDnsArgs, ProfileFtgdDnsPtr and ProfileFtgdDnsPtrOutput values.
// You can construct a concrete instance of `ProfileFtgdDnsPtrInput` via:
//
//	        ProfileFtgdDnsArgs{...}
//
//	or:
//
//	        nil
type ProfileFtgdDnsPtrInput interface {
	pulumi.Input

	ToProfileFtgdDnsPtrOutput() ProfileFtgdDnsPtrOutput
	ToProfileFtgdDnsPtrOutputWithContext(context.Context) ProfileFtgdDnsPtrOutput
}

type profileFtgdDnsPtrType ProfileFtgdDnsArgs

func ProfileFtgdDnsPtr(v *ProfileFtgdDnsArgs) ProfileFtgdDnsPtrInput {
	return (*profileFtgdDnsPtrType)(v)
}

func (*profileFtgdDnsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileFtgdDns)(nil)).Elem()
}

func (i *profileFtgdDnsPtrType) ToProfileFtgdDnsPtrOutput() ProfileFtgdDnsPtrOutput {
	return i.ToProfileFtgdDnsPtrOutputWithContext(context.Background())
}

func (i *profileFtgdDnsPtrType) ToProfileFtgdDnsPtrOutputWithContext(ctx context.Context) ProfileFtgdDnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFtgdDnsPtrOutput)
}

type ProfileFtgdDnsOutput struct{ *pulumi.OutputState }

func (ProfileFtgdDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFtgdDns)(nil)).Elem()
}

func (o ProfileFtgdDnsOutput) ToProfileFtgdDnsOutput() ProfileFtgdDnsOutput {
	return o
}

func (o ProfileFtgdDnsOutput) ToProfileFtgdDnsOutputWithContext(ctx context.Context) ProfileFtgdDnsOutput {
	return o
}

func (o ProfileFtgdDnsOutput) ToProfileFtgdDnsPtrOutput() ProfileFtgdDnsPtrOutput {
	return o.ToProfileFtgdDnsPtrOutputWithContext(context.Background())
}

func (o ProfileFtgdDnsOutput) ToProfileFtgdDnsPtrOutputWithContext(ctx context.Context) ProfileFtgdDnsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileFtgdDns) *ProfileFtgdDns {
		return &v
	}).(ProfileFtgdDnsPtrOutput)
}

// FortiGuard DNS domain filters. The structure of `filters` block is documented below.
func (o ProfileFtgdDnsOutput) Filters() ProfileFtgdDnsFilterArrayOutput {
	return o.ApplyT(func(v ProfileFtgdDns) []ProfileFtgdDnsFilter { return v.Filters }).(ProfileFtgdDnsFilterArrayOutput)
}

// FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
func (o ProfileFtgdDnsOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFtgdDns) *string { return v.Options }).(pulumi.StringPtrOutput)
}

type ProfileFtgdDnsPtrOutput struct{ *pulumi.OutputState }

func (ProfileFtgdDnsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileFtgdDns)(nil)).Elem()
}

func (o ProfileFtgdDnsPtrOutput) ToProfileFtgdDnsPtrOutput() ProfileFtgdDnsPtrOutput {
	return o
}

func (o ProfileFtgdDnsPtrOutput) ToProfileFtgdDnsPtrOutputWithContext(ctx context.Context) ProfileFtgdDnsPtrOutput {
	return o
}

func (o ProfileFtgdDnsPtrOutput) Elem() ProfileFtgdDnsOutput {
	return o.ApplyT(func(v *ProfileFtgdDns) ProfileFtgdDns {
		if v != nil {
			return *v
		}
		var ret ProfileFtgdDns
		return ret
	}).(ProfileFtgdDnsOutput)
}

// FortiGuard DNS domain filters. The structure of `filters` block is documented below.
func (o ProfileFtgdDnsPtrOutput) Filters() ProfileFtgdDnsFilterArrayOutput {
	return o.ApplyT(func(v *ProfileFtgdDns) []ProfileFtgdDnsFilter {
		if v == nil {
			return nil
		}
		return v.Filters
	}).(ProfileFtgdDnsFilterArrayOutput)
}

// FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
func (o ProfileFtgdDnsPtrOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileFtgdDns) *string {
		if v == nil {
			return nil
		}
		return v.Options
	}).(pulumi.StringPtrOutput)
}

type ProfileFtgdDnsFilter struct {
	// Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
	Action *string `pulumi:"action"`
	// Category number.
	Category *int `pulumi:"category"`
	// ID number.
	Id *int `pulumi:"id"`
	// Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
	Log *string `pulumi:"log"`
}

// ProfileFtgdDnsFilterInput is an input type that accepts ProfileFtgdDnsFilterArgs and ProfileFtgdDnsFilterOutput values.
// You can construct a concrete instance of `ProfileFtgdDnsFilterInput` via:
//
//	ProfileFtgdDnsFilterArgs{...}
type ProfileFtgdDnsFilterInput interface {
	pulumi.Input

	ToProfileFtgdDnsFilterOutput() ProfileFtgdDnsFilterOutput
	ToProfileFtgdDnsFilterOutputWithContext(context.Context) ProfileFtgdDnsFilterOutput
}

type ProfileFtgdDnsFilterArgs struct {
	// Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// Category number.
	Category pulumi.IntPtrInput `pulumi:"category"`
	// ID number.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
	Log pulumi.StringPtrInput `pulumi:"log"`
}

func (ProfileFtgdDnsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFtgdDnsFilter)(nil)).Elem()
}

func (i ProfileFtgdDnsFilterArgs) ToProfileFtgdDnsFilterOutput() ProfileFtgdDnsFilterOutput {
	return i.ToProfileFtgdDnsFilterOutputWithContext(context.Background())
}

func (i ProfileFtgdDnsFilterArgs) ToProfileFtgdDnsFilterOutputWithContext(ctx context.Context) ProfileFtgdDnsFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFtgdDnsFilterOutput)
}

// ProfileFtgdDnsFilterArrayInput is an input type that accepts ProfileFtgdDnsFilterArray and ProfileFtgdDnsFilterArrayOutput values.
// You can construct a concrete instance of `ProfileFtgdDnsFilterArrayInput` via:
//
//	ProfileFtgdDnsFilterArray{ ProfileFtgdDnsFilterArgs{...} }
type ProfileFtgdDnsFilterArrayInput interface {
	pulumi.Input

	ToProfileFtgdDnsFilterArrayOutput() ProfileFtgdDnsFilterArrayOutput
	ToProfileFtgdDnsFilterArrayOutputWithContext(context.Context) ProfileFtgdDnsFilterArrayOutput
}

type ProfileFtgdDnsFilterArray []ProfileFtgdDnsFilterInput

func (ProfileFtgdDnsFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFtgdDnsFilter)(nil)).Elem()
}

func (i ProfileFtgdDnsFilterArray) ToProfileFtgdDnsFilterArrayOutput() ProfileFtgdDnsFilterArrayOutput {
	return i.ToProfileFtgdDnsFilterArrayOutputWithContext(context.Background())
}

func (i ProfileFtgdDnsFilterArray) ToProfileFtgdDnsFilterArrayOutputWithContext(ctx context.Context) ProfileFtgdDnsFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFtgdDnsFilterArrayOutput)
}

type ProfileFtgdDnsFilterOutput struct{ *pulumi.OutputState }

func (ProfileFtgdDnsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFtgdDnsFilter)(nil)).Elem()
}

func (o ProfileFtgdDnsFilterOutput) ToProfileFtgdDnsFilterOutput() ProfileFtgdDnsFilterOutput {
	return o
}

func (o ProfileFtgdDnsFilterOutput) ToProfileFtgdDnsFilterOutputWithContext(ctx context.Context) ProfileFtgdDnsFilterOutput {
	return o
}

// Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
func (o ProfileFtgdDnsFilterOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFtgdDnsFilter) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// Category number.
func (o ProfileFtgdDnsFilterOutput) Category() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileFtgdDnsFilter) *int { return v.Category }).(pulumi.IntPtrOutput)
}

// ID number.
func (o ProfileFtgdDnsFilterOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileFtgdDnsFilter) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
func (o ProfileFtgdDnsFilterOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFtgdDnsFilter) *string { return v.Log }).(pulumi.StringPtrOutput)
}

type ProfileFtgdDnsFilterArrayOutput struct{ *pulumi.OutputState }

func (ProfileFtgdDnsFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFtgdDnsFilter)(nil)).Elem()
}

func (o ProfileFtgdDnsFilterArrayOutput) ToProfileFtgdDnsFilterArrayOutput() ProfileFtgdDnsFilterArrayOutput {
	return o
}

func (o ProfileFtgdDnsFilterArrayOutput) ToProfileFtgdDnsFilterArrayOutputWithContext(ctx context.Context) ProfileFtgdDnsFilterArrayOutput {
	return o
}

func (o ProfileFtgdDnsFilterArrayOutput) Index(i pulumi.IntInput) ProfileFtgdDnsFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileFtgdDnsFilter {
		return vs[0].([]ProfileFtgdDnsFilter)[vs[1].(int)]
	}).(ProfileFtgdDnsFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainfilterEntryInput)(nil)).Elem(), DomainfilterEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainfilterEntryArrayInput)(nil)).Elem(), DomainfilterEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileDnsTranslationInput)(nil)).Elem(), ProfileDnsTranslationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileDnsTranslationArrayInput)(nil)).Elem(), ProfileDnsTranslationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileDomainFilterInput)(nil)).Elem(), ProfileDomainFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileDomainFilterPtrInput)(nil)).Elem(), ProfileDomainFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileExternalIpBlocklistInput)(nil)).Elem(), ProfileExternalIpBlocklistArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileExternalIpBlocklistArrayInput)(nil)).Elem(), ProfileExternalIpBlocklistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFtgdDnsInput)(nil)).Elem(), ProfileFtgdDnsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFtgdDnsPtrInput)(nil)).Elem(), ProfileFtgdDnsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFtgdDnsFilterInput)(nil)).Elem(), ProfileFtgdDnsFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFtgdDnsFilterArrayInput)(nil)).Elem(), ProfileFtgdDnsFilterArray{})
	pulumi.RegisterOutputType(DomainfilterEntryOutput{})
	pulumi.RegisterOutputType(DomainfilterEntryArrayOutput{})
	pulumi.RegisterOutputType(ProfileDnsTranslationOutput{})
	pulumi.RegisterOutputType(ProfileDnsTranslationArrayOutput{})
	pulumi.RegisterOutputType(ProfileDomainFilterOutput{})
	pulumi.RegisterOutputType(ProfileDomainFilterPtrOutput{})
	pulumi.RegisterOutputType(ProfileExternalIpBlocklistOutput{})
	pulumi.RegisterOutputType(ProfileExternalIpBlocklistArrayOutput{})
	pulumi.RegisterOutputType(ProfileFtgdDnsOutput{})
	pulumi.RegisterOutputType(ProfileFtgdDnsPtrOutput{})
	pulumi.RegisterOutputType(ProfileFtgdDnsFilterOutput{})
	pulumi.RegisterOutputType(ProfileFtgdDnsFilterArrayOutput{})
}
