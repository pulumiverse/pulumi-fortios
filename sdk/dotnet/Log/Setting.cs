// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log
{
    /// <summary>
    /// Configure general log settings.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Log.Setting("trname", new()
    ///     {
    ///         BriefTrafficFormat = "disable",
    ///         DaemonLog = "disable",
    ///         ExpolicyImplicitLog = "disable",
    ///         FazOverride = "disable",
    ///         Fwpolicy6ImplicitLog = "disable",
    ///         FwpolicyImplicitLog = "disable",
    ///         LocalInAllow = "disable",
    ///         LocalInDenyBroadcast = "disable",
    ///         LocalInDenyUnicast = "disable",
    ///         LocalOut = "disable",
    ///         LogInvalidPacket = "disable",
    ///         LogPolicyComment = "disable",
    ///         LogPolicyName = "disable",
    ///         LogUserInUpper = "disable",
    ///         NeighborEvent = "disable",
    ///         ResolveIp = "disable",
    ///         ResolvePort = "enable",
    ///         SyslogOverride = "disable",
    ///         UserAnonymize = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/setting:Setting labelname LogSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/setting:Setting labelname LogSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User name anonymization hash salt.
        /// </summary>
        [Output("anonymizationHash")]
        public Output<string> AnonymizationHash { get; private set; } = null!;

        /// <summary>
        /// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("briefTrafficFormat")]
        public Output<string> BriefTrafficFormat { get; private set; } = null!;

        /// <summary>
        /// Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.
        /// </summary>
        [Output("customLogFields")]
        public Output<ImmutableArray<Outputs.SettingCustomLogField>> CustomLogFields { get; private set; } = null!;

        /// <summary>
        /// Enable/disable daemon logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("daemonLog")]
        public Output<string> DaemonLog { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("expolicyImplicitLog")]
        public Output<string> ExpolicyImplicitLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fazOverride")]
        public Output<string> FazOverride { get; private set; } = null!;

        /// <summary>
        /// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fwpolicy6ImplicitLog")]
        public Output<string> Fwpolicy6ImplicitLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fwpolicyImplicitLog")]
        public Output<string> FwpolicyImplicitLog { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localInAllow")]
        public Output<string> LocalInAllow { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localInDenyBroadcast")]
        public Output<string> LocalInDenyBroadcast { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localInDenyUnicast")]
        public Output<string> LocalInDenyUnicast { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local-out logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localOut")]
        public Output<string> LocalOut { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("localOutIocDetection")]
        public Output<string> LocalOutIocDetection { get; private set; } = null!;

        /// <summary>
        /// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logInvalidPacket")]
        public Output<string> LogInvalidPacket { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logPolicyComment")]
        public Output<string> LogPolicyComment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logPolicyName")]
        public Output<string> LogPolicyName { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logUserInUpper")]
        public Output<string> LogUserInUpper { get; private set; } = null!;

        /// <summary>
        /// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("longLiveSessionStat")]
        public Output<string> LongLiveSessionStat { get; private set; } = null!;

        /// <summary>
        /// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("neighborEvent")]
        public Output<string> NeighborEvent { get; private set; } = null!;

        /// <summary>
        /// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("resolveIp")]
        public Output<string> ResolveIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("resolvePort")]
        public Output<string> ResolvePort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("restApiGet")]
        public Output<string> RestApiGet { get; private set; } = null!;

        /// <summary>
        /// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("restApiSet")]
        public Output<string> RestApiSet { get; private set; } = null!;

        /// <summary>
        /// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("syslogOverride")]
        public Output<string> SyslogOverride { get; private set; } = null!;

        /// <summary>
        /// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("userAnonymize")]
        public Output<string> UserAnonymize { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User name anonymization hash salt.
        /// </summary>
        [Input("anonymizationHash")]
        public Input<string>? AnonymizationHash { get; set; }

        /// <summary>
        /// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("briefTrafficFormat")]
        public Input<string>? BriefTrafficFormat { get; set; }

        [Input("customLogFields")]
        private InputList<Inputs.SettingCustomLogFieldArgs>? _customLogFields;

        /// <summary>
        /// Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingCustomLogFieldArgs> CustomLogFields
        {
            get => _customLogFields ?? (_customLogFields = new InputList<Inputs.SettingCustomLogFieldArgs>());
            set => _customLogFields = value;
        }

        /// <summary>
        /// Enable/disable daemon logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("daemonLog")]
        public Input<string>? DaemonLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("expolicyImplicitLog")]
        public Input<string>? ExpolicyImplicitLog { get; set; }

        /// <summary>
        /// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fazOverride")]
        public Input<string>? FazOverride { get; set; }

        /// <summary>
        /// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fwpolicy6ImplicitLog")]
        public Input<string>? Fwpolicy6ImplicitLog { get; set; }

        /// <summary>
        /// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fwpolicyImplicitLog")]
        public Input<string>? FwpolicyImplicitLog { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInAllow")]
        public Input<string>? LocalInAllow { get; set; }

        /// <summary>
        /// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInDenyBroadcast")]
        public Input<string>? LocalInDenyBroadcast { get; set; }

        /// <summary>
        /// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInDenyUnicast")]
        public Input<string>? LocalInDenyUnicast { get; set; }

        /// <summary>
        /// Enable/disable local-out logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOut")]
        public Input<string>? LocalOut { get; set; }

        /// <summary>
        /// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOutIocDetection")]
        public Input<string>? LocalOutIocDetection { get; set; }

        /// <summary>
        /// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logInvalidPacket")]
        public Input<string>? LogInvalidPacket { get; set; }

        /// <summary>
        /// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logPolicyComment")]
        public Input<string>? LogPolicyComment { get; set; }

        /// <summary>
        /// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logPolicyName")]
        public Input<string>? LogPolicyName { get; set; }

        /// <summary>
        /// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logUserInUpper")]
        public Input<string>? LogUserInUpper { get; set; }

        /// <summary>
        /// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("longLiveSessionStat")]
        public Input<string>? LongLiveSessionStat { get; set; }

        /// <summary>
        /// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("neighborEvent")]
        public Input<string>? NeighborEvent { get; set; }

        /// <summary>
        /// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("resolveIp")]
        public Input<string>? ResolveIp { get; set; }

        /// <summary>
        /// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("resolvePort")]
        public Input<string>? ResolvePort { get; set; }

        /// <summary>
        /// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restApiGet")]
        public Input<string>? RestApiGet { get; set; }

        /// <summary>
        /// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restApiSet")]
        public Input<string>? RestApiSet { get; set; }

        /// <summary>
        /// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("syslogOverride")]
        public Input<string>? SyslogOverride { get; set; }

        /// <summary>
        /// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("userAnonymize")]
        public Input<string>? UserAnonymize { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingArgs()
        {
        }
        public static new SettingArgs Empty => new SettingArgs();
    }

    public sealed class SettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User name anonymization hash salt.
        /// </summary>
        [Input("anonymizationHash")]
        public Input<string>? AnonymizationHash { get; set; }

        /// <summary>
        /// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("briefTrafficFormat")]
        public Input<string>? BriefTrafficFormat { get; set; }

        [Input("customLogFields")]
        private InputList<Inputs.SettingCustomLogFieldGetArgs>? _customLogFields;

        /// <summary>
        /// Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.
        /// </summary>
        public InputList<Inputs.SettingCustomLogFieldGetArgs> CustomLogFields
        {
            get => _customLogFields ?? (_customLogFields = new InputList<Inputs.SettingCustomLogFieldGetArgs>());
            set => _customLogFields = value;
        }

        /// <summary>
        /// Enable/disable daemon logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("daemonLog")]
        public Input<string>? DaemonLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("expolicyImplicitLog")]
        public Input<string>? ExpolicyImplicitLog { get; set; }

        /// <summary>
        /// Enable/disable extended traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fazOverride")]
        public Input<string>? FazOverride { get; set; }

        /// <summary>
        /// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fwpolicy6ImplicitLog")]
        public Input<string>? Fwpolicy6ImplicitLog { get; set; }

        /// <summary>
        /// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fwpolicyImplicitLog")]
        public Input<string>? FwpolicyImplicitLog { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInAllow")]
        public Input<string>? LocalInAllow { get; set; }

        /// <summary>
        /// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInDenyBroadcast")]
        public Input<string>? LocalInDenyBroadcast { get; set; }

        /// <summary>
        /// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localInDenyUnicast")]
        public Input<string>? LocalInDenyUnicast { get; set; }

        /// <summary>
        /// Enable/disable local-out logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOut")]
        public Input<string>? LocalOut { get; set; }

        /// <summary>
        /// Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOutIocDetection")]
        public Input<string>? LocalOutIocDetection { get; set; }

        /// <summary>
        /// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logInvalidPacket")]
        public Input<string>? LogInvalidPacket { get; set; }

        /// <summary>
        /// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logPolicyComment")]
        public Input<string>? LogPolicyComment { get; set; }

        /// <summary>
        /// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logPolicyName")]
        public Input<string>? LogPolicyName { get; set; }

        /// <summary>
        /// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logUserInUpper")]
        public Input<string>? LogUserInUpper { get; set; }

        /// <summary>
        /// Enable/disable long-live-session statistics logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("longLiveSessionStat")]
        public Input<string>? LongLiveSessionStat { get; set; }

        /// <summary>
        /// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("neighborEvent")]
        public Input<string>? NeighborEvent { get; set; }

        /// <summary>
        /// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("resolveIp")]
        public Input<string>? ResolveIp { get; set; }

        /// <summary>
        /// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("resolvePort")]
        public Input<string>? ResolvePort { get; set; }

        /// <summary>
        /// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restApiGet")]
        public Input<string>? RestApiGet { get; set; }

        /// <summary>
        /// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restApiSet")]
        public Input<string>? RestApiSet { get; set; }

        /// <summary>
        /// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("syslogOverride")]
        public Input<string>? SyslogOverride { get; set; }

        /// <summary>
        /// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("userAnonymize")]
        public Input<string>? UserAnonymize { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingState()
        {
        }
        public static new SettingState Empty => new SettingState();
    }
}
