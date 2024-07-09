// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure NAC policy matching pattern to identify matching NAC devices. Applies to FortiOS Version `&gt;= 6.4.0`.
    /// 
    /// ## Import
    /// 
    /// User NacPolicy can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/nacpolicy:Nacpolicy labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/nacpolicy:Nacpolicy labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/nacpolicy:Nacpolicy")]
    public partial class Nacpolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Category of NAC policy.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Description for the NAC policy matching pattern.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching EMS tag.
        /// </summary>
        [Output("emsTag")]
        public Output<string> EmsTag { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching family.
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// Dynamic firewall address to associate MAC which match this policy.
        /// </summary>
        [Output("firewallAddress")]
        public Output<string> FirewallAddress { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching FortiVoice tag.
        /// </summary>
        [Output("fortivoiceTag")]
        public Output<string> FortivoiceTag { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching host.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching hardware vendor.
        /// </summary>
        [Output("hwVendor")]
        public Output<string> HwVendor { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching hardware version.
        /// </summary>
        [Output("hwVersion")]
        public Output<string> HwVersion { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching MAC address.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// Number of days the matched devices will be retained (0 - always retain)
        /// </summary>
        [Output("matchPeriod")]
        public Output<int> MatchPeriod { get; private set; } = null!;

        /// <summary>
        /// Match and retain the devices based on the type. Valid values: `dynamic`, `override`.
        /// </summary>
        [Output("matchType")]
        public Output<string> MatchType { get; private set; } = null!;

        /// <summary>
        /// NAC policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching operating system.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching devices vulnerability severity lists. The structure of `severity` block is documented below.
        /// </summary>
        [Output("severities")]
        public Output<ImmutableArray<Outputs.NacpolicySeverity>> Severities { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching source.
        /// </summary>
        [Output("src")]
        public Output<string> Src { get; private set; } = null!;

        /// <summary>
        /// SSID policy to be applied on the matched NAC policy.
        /// </summary>
        [Output("ssidPolicy")]
        public Output<string> SsidPolicy { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NAC policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching software version.
        /// </summary>
        [Output("swVersion")]
        public Output<string> SwVersion { get; private set; } = null!;

        /// <summary>
        /// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
        /// </summary>
        [Output("switchAutoAuth")]
        public Output<string> SwitchAutoAuth { get; private set; } = null!;

        /// <summary>
        /// FortiLink interface for which this NAC policy belongs to.
        /// </summary>
        [Output("switchFortilink")]
        public Output<string> SwitchFortilink { get; private set; } = null!;

        /// <summary>
        /// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switch_group` block is documented below.
        /// </summary>
        [Output("switchGroups")]
        public Output<ImmutableArray<Outputs.NacpolicySwitchGroup>> SwitchGroups { get; private set; } = null!;

        /// <summary>
        /// switch-mac-policy to be applied on the matched NAC policy.
        /// </summary>
        [Output("switchMacPolicy")]
        public Output<string> SwitchMacPolicy { get; private set; } = null!;

        /// <summary>
        /// switch-port-policy to be applied on the matched NAC policy.
        /// </summary>
        [Output("switchPortPolicy")]
        public Output<string> SwitchPortPolicy { get; private set; } = null!;

        /// <summary>
        /// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.
        /// </summary>
        [Output("switchScopes")]
        public Output<ImmutableArray<Outputs.NacpolicySwitchScope>> SwitchScopes { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching user.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// NAC policy matching user group.
        /// </summary>
        [Output("userGroup")]
        public Output<string> UserGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Nacpolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Nacpolicy(string name, NacpolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/nacpolicy:Nacpolicy", name, args ?? new NacpolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Nacpolicy(string name, Input<string> id, NacpolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/nacpolicy:Nacpolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Nacpolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Nacpolicy Get(string name, Input<string> id, NacpolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Nacpolicy(name, id, state, options);
        }
    }

    public sealed class NacpolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Category of NAC policy.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Description for the NAC policy matching pattern.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// NAC policy matching EMS tag.
        /// </summary>
        [Input("emsTag")]
        public Input<string>? EmsTag { get; set; }

        /// <summary>
        /// NAC policy matching family.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// Dynamic firewall address to associate MAC which match this policy.
        /// </summary>
        [Input("firewallAddress")]
        public Input<string>? FirewallAddress { get; set; }

        /// <summary>
        /// NAC policy matching FortiVoice tag.
        /// </summary>
        [Input("fortivoiceTag")]
        public Input<string>? FortivoiceTag { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// NAC policy matching host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// NAC policy matching hardware vendor.
        /// </summary>
        [Input("hwVendor")]
        public Input<string>? HwVendor { get; set; }

        /// <summary>
        /// NAC policy matching hardware version.
        /// </summary>
        [Input("hwVersion")]
        public Input<string>? HwVersion { get; set; }

        /// <summary>
        /// NAC policy matching MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Number of days the matched devices will be retained (0 - always retain)
        /// </summary>
        [Input("matchPeriod")]
        public Input<int>? MatchPeriod { get; set; }

        /// <summary>
        /// Match and retain the devices based on the type. Valid values: `dynamic`, `override`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// NAC policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NAC policy matching operating system.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("severities")]
        private InputList<Inputs.NacpolicySeverityArgs>? _severities;

        /// <summary>
        /// NAC policy matching devices vulnerability severity lists. The structure of `severity` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySeverityArgs> Severities
        {
            get => _severities ?? (_severities = new InputList<Inputs.NacpolicySeverityArgs>());
            set => _severities = value;
        }

        /// <summary>
        /// NAC policy matching source.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// SSID policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("ssidPolicy")]
        public Input<string>? SsidPolicy { get; set; }

        /// <summary>
        /// Enable/disable NAC policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NAC policy matching software version.
        /// </summary>
        [Input("swVersion")]
        public Input<string>? SwVersion { get; set; }

        /// <summary>
        /// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
        /// </summary>
        [Input("switchAutoAuth")]
        public Input<string>? SwitchAutoAuth { get; set; }

        /// <summary>
        /// FortiLink interface for which this NAC policy belongs to.
        /// </summary>
        [Input("switchFortilink")]
        public Input<string>? SwitchFortilink { get; set; }

        [Input("switchGroups")]
        private InputList<Inputs.NacpolicySwitchGroupArgs>? _switchGroups;

        /// <summary>
        /// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switch_group` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySwitchGroupArgs> SwitchGroups
        {
            get => _switchGroups ?? (_switchGroups = new InputList<Inputs.NacpolicySwitchGroupArgs>());
            set => _switchGroups = value;
        }

        /// <summary>
        /// switch-mac-policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("switchMacPolicy")]
        public Input<string>? SwitchMacPolicy { get; set; }

        /// <summary>
        /// switch-port-policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("switchPortPolicy")]
        public Input<string>? SwitchPortPolicy { get; set; }

        [Input("switchScopes")]
        private InputList<Inputs.NacpolicySwitchScopeArgs>? _switchScopes;

        /// <summary>
        /// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySwitchScopeArgs> SwitchScopes
        {
            get => _switchScopes ?? (_switchScopes = new InputList<Inputs.NacpolicySwitchScopeArgs>());
            set => _switchScopes = value;
        }

        /// <summary>
        /// NAC policy matching type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// NAC policy matching user.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// NAC policy matching user group.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NacpolicyArgs()
        {
        }
        public static new NacpolicyArgs Empty => new NacpolicyArgs();
    }

    public sealed class NacpolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Category of NAC policy.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Description for the NAC policy matching pattern.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// NAC policy matching EMS tag.
        /// </summary>
        [Input("emsTag")]
        public Input<string>? EmsTag { get; set; }

        /// <summary>
        /// NAC policy matching family.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// Dynamic firewall address to associate MAC which match this policy.
        /// </summary>
        [Input("firewallAddress")]
        public Input<string>? FirewallAddress { get; set; }

        /// <summary>
        /// NAC policy matching FortiVoice tag.
        /// </summary>
        [Input("fortivoiceTag")]
        public Input<string>? FortivoiceTag { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// NAC policy matching host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// NAC policy matching hardware vendor.
        /// </summary>
        [Input("hwVendor")]
        public Input<string>? HwVendor { get; set; }

        /// <summary>
        /// NAC policy matching hardware version.
        /// </summary>
        [Input("hwVersion")]
        public Input<string>? HwVersion { get; set; }

        /// <summary>
        /// NAC policy matching MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Number of days the matched devices will be retained (0 - always retain)
        /// </summary>
        [Input("matchPeriod")]
        public Input<int>? MatchPeriod { get; set; }

        /// <summary>
        /// Match and retain the devices based on the type. Valid values: `dynamic`, `override`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// NAC policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NAC policy matching operating system.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("severities")]
        private InputList<Inputs.NacpolicySeverityGetArgs>? _severities;

        /// <summary>
        /// NAC policy matching devices vulnerability severity lists. The structure of `severity` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySeverityGetArgs> Severities
        {
            get => _severities ?? (_severities = new InputList<Inputs.NacpolicySeverityGetArgs>());
            set => _severities = value;
        }

        /// <summary>
        /// NAC policy matching source.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        /// <summary>
        /// SSID policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("ssidPolicy")]
        public Input<string>? SsidPolicy { get; set; }

        /// <summary>
        /// Enable/disable NAC policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NAC policy matching software version.
        /// </summary>
        [Input("swVersion")]
        public Input<string>? SwVersion { get; set; }

        /// <summary>
        /// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
        /// </summary>
        [Input("switchAutoAuth")]
        public Input<string>? SwitchAutoAuth { get; set; }

        /// <summary>
        /// FortiLink interface for which this NAC policy belongs to.
        /// </summary>
        [Input("switchFortilink")]
        public Input<string>? SwitchFortilink { get; set; }

        [Input("switchGroups")]
        private InputList<Inputs.NacpolicySwitchGroupGetArgs>? _switchGroups;

        /// <summary>
        /// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switch_group` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySwitchGroupGetArgs> SwitchGroups
        {
            get => _switchGroups ?? (_switchGroups = new InputList<Inputs.NacpolicySwitchGroupGetArgs>());
            set => _switchGroups = value;
        }

        /// <summary>
        /// switch-mac-policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("switchMacPolicy")]
        public Input<string>? SwitchMacPolicy { get; set; }

        /// <summary>
        /// switch-port-policy to be applied on the matched NAC policy.
        /// </summary>
        [Input("switchPortPolicy")]
        public Input<string>? SwitchPortPolicy { get; set; }

        [Input("switchScopes")]
        private InputList<Inputs.NacpolicySwitchScopeGetArgs>? _switchScopes;

        /// <summary>
        /// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.
        /// </summary>
        public InputList<Inputs.NacpolicySwitchScopeGetArgs> SwitchScopes
        {
            get => _switchScopes ?? (_switchScopes = new InputList<Inputs.NacpolicySwitchScopeGetArgs>());
            set => _switchScopes = value;
        }

        /// <summary>
        /// NAC policy matching type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// NAC policy matching user.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// NAC policy matching user group.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NacpolicyState()
        {
        }
        public static new NacpolicyState Empty => new NacpolicyState();
    }
}
