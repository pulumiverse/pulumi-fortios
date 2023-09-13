// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20
{
    /// <summary>
    /// Configure hotspot profile.
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 HsProfile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile")]
    public partial class Hsprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable additional step required for access (ASRA). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessNetworkAsra")]
        public Output<string> AccessNetworkAsra { get; private set; } = null!;

        /// <summary>
        /// Enable/disable emergency services reachable (ESR). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessNetworkEsr")]
        public Output<string> AccessNetworkEsr { get; private set; } = null!;

        /// <summary>
        /// Enable/disable connectivity to the Internet. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessNetworkInternet")]
        public Output<string> AccessNetworkInternet { get; private set; } = null!;

        /// <summary>
        /// Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.
        /// </summary>
        [Output("accessNetworkType")]
        public Output<string> AccessNetworkType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessNetworkUesa")]
        public Output<string> AccessNetworkUesa { get; private set; } = null!;

        /// <summary>
        /// Advice of charge.
        /// </summary>
        [Output("adviceOfCharge")]
        public Output<string> AdviceOfCharge { get; private set; } = null!;

        /// <summary>
        /// ANQP Domain ID (0-65535).
        /// </summary>
        [Output("anqpDomainId")]
        public Output<int> AnqpDomainId { get; private set; } = null!;

        /// <summary>
        /// Enable/disable basic service set (BSS) transition Support. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bssTransition")]
        public Output<string> BssTransition { get; private set; } = null!;

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Output("connCap")]
        public Output<string> ConnCap { get; private set; } = null!;

        /// <summary>
        /// Deauthentication request timeout (in seconds).
        /// </summary>
        [Output("deauthRequestTimeout")]
        public Output<int> DeauthRequestTimeout { get; private set; } = null!;

        /// <summary>
        /// Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dgaf")]
        public Output<string> Dgaf { get; private set; } = null!;

        /// <summary>
        /// Domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
        /// </summary>
        [Output("gasComebackDelay")]
        public Output<int> GasComebackDelay { get; private set; } = null!;

        /// <summary>
        /// GAS fragmentation limit (512 - 4096, default = 1024).
        /// </summary>
        [Output("gasFragmentationLimit")]
        public Output<int> GasFragmentationLimit { get; private set; } = null!;

        /// <summary>
        /// Homogeneous extended service set identifier (HESSID).
        /// </summary>
        [Output("hessid")]
        public Output<string> Hessid { get; private set; } = null!;

        /// <summary>
        /// IP address type name.
        /// </summary>
        [Output("ipAddrType")]
        public Output<string> IpAddrType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Layer 2 traffic inspection and filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("l2tif")]
        public Output<string> L2tif { get; private set; } = null!;

        /// <summary>
        /// 3GPP PLMN name.
        /// </summary>
        [Output("n3gppPlmn")]
        public Output<string> N3gppPlmn { get; private set; } = null!;

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Output("naiRealm")]
        public Output<string> NaiRealm { get; private set; } = null!;

        /// <summary>
        /// Hotspot profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network authentication name.
        /// </summary>
        [Output("networkAuth")]
        public Output<string> NetworkAuth { get; private set; } = null!;

        /// <summary>
        /// Operator friendly name.
        /// </summary>
        [Output("operFriendlyName")]
        public Output<string> OperFriendlyName { get; private set; } = null!;

        /// <summary>
        /// Operator icon.
        /// </summary>
        [Output("operIcon")]
        public Output<string> OperIcon { get; private set; } = null!;

        /// <summary>
        /// OSU Provider NAI.
        /// </summary>
        [Output("osuProviderNai")]
        public Output<string> OsuProviderNai { get; private set; } = null!;

        /// <summary>
        /// Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.
        /// </summary>
        [Output("osuProviders")]
        public Output<ImmutableArray<Outputs.HsprofileOsuProvider>> OsuProviders { get; private set; } = null!;

        /// <summary>
        /// Online sign up (OSU) SSID.
        /// </summary>
        [Output("osuSsid")]
        public Output<string> OsuSsid { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("pameBi")]
        public Output<string> PameBi { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Proxy ARP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("proxyArp")]
        public Output<string> ProxyArp { get; private set; } = null!;

        /// <summary>
        /// QoS MAP set ID.
        /// </summary>
        [Output("qosMap")]
        public Output<string> QosMap { get; private set; } = null!;

        /// <summary>
        /// Hotspot 2.0 Release number (1, 2, 3, default = 2).
        /// </summary>
        [Output("release")]
        public Output<int> Release { get; private set; } = null!;

        /// <summary>
        /// Roaming consortium list name.
        /// </summary>
        [Output("roamingConsortium")]
        public Output<string> RoamingConsortium { get; private set; } = null!;

        /// <summary>
        /// Terms and conditions.
        /// </summary>
        [Output("termsAndConditions")]
        public Output<string> TermsAndConditions { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.
        /// </summary>
        [Output("venueGroup")]
        public Output<string> VenueGroup { get; private set; } = null!;

        /// <summary>
        /// Venue name.
        /// </summary>
        [Output("venueName")]
        public Output<string> VenueName { get; private set; } = null!;

        /// <summary>
        /// Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.
        /// </summary>
        [Output("venueType")]
        public Output<string> VenueType { get; private set; } = null!;

        /// <summary>
        /// Venue name.
        /// </summary>
        [Output("venueUrl")]
        public Output<string> VenueUrl { get; private set; } = null!;

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Output("wanMetrics")]
        public Output<string> WanMetrics { get; private set; } = null!;

        /// <summary>
        /// Enable/disable wireless network management (WNM) sleep mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wnmSleepMode")]
        public Output<string> WnmSleepMode { get; private set; } = null!;


        /// <summary>
        /// Create a Hsprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hsprofile(string name, HsprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile", name, args ?? new HsprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hsprofile(string name, Input<string> id, HsprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Hsprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hsprofile Get(string name, Input<string> id, HsprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Hsprofile(name, id, state, options);
        }
    }

    public sealed class HsprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable additional step required for access (ASRA). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkAsra")]
        public Input<string>? AccessNetworkAsra { get; set; }

        /// <summary>
        /// Enable/disable emergency services reachable (ESR). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkEsr")]
        public Input<string>? AccessNetworkEsr { get; set; }

        /// <summary>
        /// Enable/disable connectivity to the Internet. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkInternet")]
        public Input<string>? AccessNetworkInternet { get; set; }

        /// <summary>
        /// Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.
        /// </summary>
        [Input("accessNetworkType")]
        public Input<string>? AccessNetworkType { get; set; }

        /// <summary>
        /// Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkUesa")]
        public Input<string>? AccessNetworkUesa { get; set; }

        /// <summary>
        /// Advice of charge.
        /// </summary>
        [Input("adviceOfCharge")]
        public Input<string>? AdviceOfCharge { get; set; }

        /// <summary>
        /// ANQP Domain ID (0-65535).
        /// </summary>
        [Input("anqpDomainId")]
        public Input<int>? AnqpDomainId { get; set; }

        /// <summary>
        /// Enable/disable basic service set (BSS) transition Support. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bssTransition")]
        public Input<string>? BssTransition { get; set; }

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Input("connCap")]
        public Input<string>? ConnCap { get; set; }

        /// <summary>
        /// Deauthentication request timeout (in seconds).
        /// </summary>
        [Input("deauthRequestTimeout")]
        public Input<int>? DeauthRequestTimeout { get; set; }

        /// <summary>
        /// Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dgaf")]
        public Input<string>? Dgaf { get; set; }

        /// <summary>
        /// Domain name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
        /// </summary>
        [Input("gasComebackDelay")]
        public Input<int>? GasComebackDelay { get; set; }

        /// <summary>
        /// GAS fragmentation limit (512 - 4096, default = 1024).
        /// </summary>
        [Input("gasFragmentationLimit")]
        public Input<int>? GasFragmentationLimit { get; set; }

        /// <summary>
        /// Homogeneous extended service set identifier (HESSID).
        /// </summary>
        [Input("hessid")]
        public Input<string>? Hessid { get; set; }

        /// <summary>
        /// IP address type name.
        /// </summary>
        [Input("ipAddrType")]
        public Input<string>? IpAddrType { get; set; }

        /// <summary>
        /// Enable/disable Layer 2 traffic inspection and filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("l2tif")]
        public Input<string>? L2tif { get; set; }

        /// <summary>
        /// 3GPP PLMN name.
        /// </summary>
        [Input("n3gppPlmn")]
        public Input<string>? N3gppPlmn { get; set; }

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Input("naiRealm")]
        public Input<string>? NaiRealm { get; set; }

        /// <summary>
        /// Hotspot profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network authentication name.
        /// </summary>
        [Input("networkAuth")]
        public Input<string>? NetworkAuth { get; set; }

        /// <summary>
        /// Operator friendly name.
        /// </summary>
        [Input("operFriendlyName")]
        public Input<string>? OperFriendlyName { get; set; }

        /// <summary>
        /// Operator icon.
        /// </summary>
        [Input("operIcon")]
        public Input<string>? OperIcon { get; set; }

        /// <summary>
        /// OSU Provider NAI.
        /// </summary>
        [Input("osuProviderNai")]
        public Input<string>? OsuProviderNai { get; set; }

        [Input("osuProviders")]
        private InputList<Inputs.HsprofileOsuProviderArgs>? _osuProviders;

        /// <summary>
        /// Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.
        /// </summary>
        public InputList<Inputs.HsprofileOsuProviderArgs> OsuProviders
        {
            get => _osuProviders ?? (_osuProviders = new InputList<Inputs.HsprofileOsuProviderArgs>());
            set => _osuProviders = value;
        }

        /// <summary>
        /// Online sign up (OSU) SSID.
        /// </summary>
        [Input("osuSsid")]
        public Input<string>? OsuSsid { get; set; }

        /// <summary>
        /// Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pameBi")]
        public Input<string>? PameBi { get; set; }

        /// <summary>
        /// Enable/disable Proxy ARP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyArp")]
        public Input<string>? ProxyArp { get; set; }

        /// <summary>
        /// QoS MAP set ID.
        /// </summary>
        [Input("qosMap")]
        public Input<string>? QosMap { get; set; }

        /// <summary>
        /// Hotspot 2.0 Release number (1, 2, 3, default = 2).
        /// </summary>
        [Input("release")]
        public Input<int>? Release { get; set; }

        /// <summary>
        /// Roaming consortium list name.
        /// </summary>
        [Input("roamingConsortium")]
        public Input<string>? RoamingConsortium { get; set; }

        /// <summary>
        /// Terms and conditions.
        /// </summary>
        [Input("termsAndConditions")]
        public Input<string>? TermsAndConditions { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.
        /// </summary>
        [Input("venueGroup")]
        public Input<string>? VenueGroup { get; set; }

        /// <summary>
        /// Venue name.
        /// </summary>
        [Input("venueName")]
        public Input<string>? VenueName { get; set; }

        /// <summary>
        /// Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.
        /// </summary>
        [Input("venueType")]
        public Input<string>? VenueType { get; set; }

        /// <summary>
        /// Venue name.
        /// </summary>
        [Input("venueUrl")]
        public Input<string>? VenueUrl { get; set; }

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Input("wanMetrics")]
        public Input<string>? WanMetrics { get; set; }

        /// <summary>
        /// Enable/disable wireless network management (WNM) sleep mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wnmSleepMode")]
        public Input<string>? WnmSleepMode { get; set; }

        public HsprofileArgs()
        {
        }
        public static new HsprofileArgs Empty => new HsprofileArgs();
    }

    public sealed class HsprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable additional step required for access (ASRA). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkAsra")]
        public Input<string>? AccessNetworkAsra { get; set; }

        /// <summary>
        /// Enable/disable emergency services reachable (ESR). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkEsr")]
        public Input<string>? AccessNetworkEsr { get; set; }

        /// <summary>
        /// Enable/disable connectivity to the Internet. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkInternet")]
        public Input<string>? AccessNetworkInternet { get; set; }

        /// <summary>
        /// Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.
        /// </summary>
        [Input("accessNetworkType")]
        public Input<string>? AccessNetworkType { get; set; }

        /// <summary>
        /// Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessNetworkUesa")]
        public Input<string>? AccessNetworkUesa { get; set; }

        /// <summary>
        /// Advice of charge.
        /// </summary>
        [Input("adviceOfCharge")]
        public Input<string>? AdviceOfCharge { get; set; }

        /// <summary>
        /// ANQP Domain ID (0-65535).
        /// </summary>
        [Input("anqpDomainId")]
        public Input<int>? AnqpDomainId { get; set; }

        /// <summary>
        /// Enable/disable basic service set (BSS) transition Support. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bssTransition")]
        public Input<string>? BssTransition { get; set; }

        /// <summary>
        /// Connection capability name.
        /// </summary>
        [Input("connCap")]
        public Input<string>? ConnCap { get; set; }

        /// <summary>
        /// Deauthentication request timeout (in seconds).
        /// </summary>
        [Input("deauthRequestTimeout")]
        public Input<int>? DeauthRequestTimeout { get; set; }

        /// <summary>
        /// Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dgaf")]
        public Input<string>? Dgaf { get; set; }

        /// <summary>
        /// Domain name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
        /// </summary>
        [Input("gasComebackDelay")]
        public Input<int>? GasComebackDelay { get; set; }

        /// <summary>
        /// GAS fragmentation limit (512 - 4096, default = 1024).
        /// </summary>
        [Input("gasFragmentationLimit")]
        public Input<int>? GasFragmentationLimit { get; set; }

        /// <summary>
        /// Homogeneous extended service set identifier (HESSID).
        /// </summary>
        [Input("hessid")]
        public Input<string>? Hessid { get; set; }

        /// <summary>
        /// IP address type name.
        /// </summary>
        [Input("ipAddrType")]
        public Input<string>? IpAddrType { get; set; }

        /// <summary>
        /// Enable/disable Layer 2 traffic inspection and filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("l2tif")]
        public Input<string>? L2tif { get; set; }

        /// <summary>
        /// 3GPP PLMN name.
        /// </summary>
        [Input("n3gppPlmn")]
        public Input<string>? N3gppPlmn { get; set; }

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Input("naiRealm")]
        public Input<string>? NaiRealm { get; set; }

        /// <summary>
        /// Hotspot profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network authentication name.
        /// </summary>
        [Input("networkAuth")]
        public Input<string>? NetworkAuth { get; set; }

        /// <summary>
        /// Operator friendly name.
        /// </summary>
        [Input("operFriendlyName")]
        public Input<string>? OperFriendlyName { get; set; }

        /// <summary>
        /// Operator icon.
        /// </summary>
        [Input("operIcon")]
        public Input<string>? OperIcon { get; set; }

        /// <summary>
        /// OSU Provider NAI.
        /// </summary>
        [Input("osuProviderNai")]
        public Input<string>? OsuProviderNai { get; set; }

        [Input("osuProviders")]
        private InputList<Inputs.HsprofileOsuProviderGetArgs>? _osuProviders;

        /// <summary>
        /// Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.
        /// </summary>
        public InputList<Inputs.HsprofileOsuProviderGetArgs> OsuProviders
        {
            get => _osuProviders ?? (_osuProviders = new InputList<Inputs.HsprofileOsuProviderGetArgs>());
            set => _osuProviders = value;
        }

        /// <summary>
        /// Online sign up (OSU) SSID.
        /// </summary>
        [Input("osuSsid")]
        public Input<string>? OsuSsid { get; set; }

        /// <summary>
        /// Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pameBi")]
        public Input<string>? PameBi { get; set; }

        /// <summary>
        /// Enable/disable Proxy ARP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyArp")]
        public Input<string>? ProxyArp { get; set; }

        /// <summary>
        /// QoS MAP set ID.
        /// </summary>
        [Input("qosMap")]
        public Input<string>? QosMap { get; set; }

        /// <summary>
        /// Hotspot 2.0 Release number (1, 2, 3, default = 2).
        /// </summary>
        [Input("release")]
        public Input<int>? Release { get; set; }

        /// <summary>
        /// Roaming consortium list name.
        /// </summary>
        [Input("roamingConsortium")]
        public Input<string>? RoamingConsortium { get; set; }

        /// <summary>
        /// Terms and conditions.
        /// </summary>
        [Input("termsAndConditions")]
        public Input<string>? TermsAndConditions { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.
        /// </summary>
        [Input("venueGroup")]
        public Input<string>? VenueGroup { get; set; }

        /// <summary>
        /// Venue name.
        /// </summary>
        [Input("venueName")]
        public Input<string>? VenueName { get; set; }

        /// <summary>
        /// Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.
        /// </summary>
        [Input("venueType")]
        public Input<string>? VenueType { get; set; }

        /// <summary>
        /// Venue name.
        /// </summary>
        [Input("venueUrl")]
        public Input<string>? VenueUrl { get; set; }

        /// <summary>
        /// WAN metric name.
        /// </summary>
        [Input("wanMetrics")]
        public Input<string>? WanMetrics { get; set; }

        /// <summary>
        /// Enable/disable wireless network management (WNM) sleep mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wnmSleepMode")]
        public Input<string>? WnmSleepMode { get; set; }

        public HsprofileState()
        {
        }
        public static new HsprofileState Empty => new HsprofileState();
    }
}
