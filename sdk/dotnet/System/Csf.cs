// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
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
    ///     var trname = new Fortios.System.Csf("trname", new()
    ///     {
    ///         ConfigurationSync = "default",
    ///         GroupPassword = "tmp",
    ///         ManagementIp = "0.0.0.0",
    ///         ManagementPort = 33,
    ///         Status = "disable",
    ///         UpstreamIp = "0.0.0.0",
    ///         UpstreamPort = 8013,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Csf can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/csf:Csf labelname SystemCsf
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/csf:Csf labelname SystemCsf
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/csf:Csf")]
    public partial class Csf : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("acceptAuthByCert")]
        public Output<string> AcceptAuthByCert { get; private set; } = null!;

        /// <summary>
        /// Authorization request type. Valid values: `serial`, `certificate`.
        /// </summary>
        [Output("authorizationRequestType")]
        public Output<string> AuthorizationRequestType { get; private set; } = null!;

        /// <summary>
        /// Certificate.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Configuration sync mode. Valid values: `default`, `local`.
        /// </summary>
        [Output("configurationSync")]
        public Output<string> ConfigurationSync { get; private set; } = null!;

        /// <summary>
        /// Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("downstreamAccess")]
        public Output<string> DownstreamAccess { get; private set; } = null!;

        /// <summary>
        /// Default access profile for requests from downstream devices.
        /// </summary>
        [Output("downstreamAccprofile")]
        public Output<string> DownstreamAccprofile { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Fabric connector configuration. The structure of `fabric_connector` block is documented below.
        /// </summary>
        [Output("fabricConnectors")]
        public Output<ImmutableArray<Outputs.CsfFabricConnector>> FabricConnectors { get; private set; } = null!;

        /// <summary>
        /// Fabric device configuration. The structure of `fabric_device` block is documented below.
        /// </summary>
        [Output("fabricDevices")]
        public Output<ImmutableArray<Outputs.CsfFabricDevice>> FabricDevices { get; private set; } = null!;

        /// <summary>
        /// Fabric CMDB Object Unification Valid values: `default`, `local`.
        /// </summary>
        [Output("fabricObjectUnification")]
        public Output<string> FabricObjectUnification { get; private set; } = null!;

        /// <summary>
        /// Number of worker processes for Security Fabric daemon.
        /// </summary>
        [Output("fabricWorkers")]
        public Output<int> FabricWorkers { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fileMgmt")]
        public Output<string> FileMgmt { get; private set; } = null!;

        /// <summary>
        /// Maximum amount of memory that can be used by the daemon files (in bytes).
        /// </summary>
        [Output("fileQuota")]
        public Output<int> FileQuota { get; private set; } = null!;

        /// <summary>
        /// Warn when the set percentage of quota has been used.
        /// </summary>
        [Output("fileQuotaWarning")]
        public Output<int> FileQuotaWarning { get; private set; } = null!;

        /// <summary>
        /// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
        /// </summary>
        [Output("fixedKey")]
        public Output<string?> FixedKey { get; private set; } = null!;

        /// <summary>
        /// Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticloudAccountEnforcement")]
        public Output<string> ForticloudAccountEnforcement { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
        /// </summary>
        [Output("groupPassword")]
        public Output<string?> GroupPassword { get; private set; } = null!;

        /// <summary>
        /// Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("logUnification")]
        public Output<string> LogUnification { get; private set; } = null!;

        /// <summary>
        /// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
        /// </summary>
        [Output("managementIp")]
        public Output<string> ManagementIp { get; private set; } = null!;

        /// <summary>
        /// Overriding port for management connection (Overrides admin port).
        /// </summary>
        [Output("managementPort")]
        public Output<int> ManagementPort { get; private set; } = null!;

        /// <summary>
        /// SAML setting configuration synchronization. Valid values: `default`, `local`.
        /// </summary>
        [Output("samlConfigurationSync")]
        public Output<string> SamlConfigurationSync { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Security Fabric. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
        /// </summary>
        [Output("trustedLists")]
        public Output<ImmutableArray<Outputs.CsfTrustedList>> TrustedLists { get; private set; } = null!;

        /// <summary>
        /// Unique ID of the current CSF node
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Output("upstream")]
        public Output<string> Upstream { get; private set; } = null!;

        /// <summary>
        /// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Output("upstreamIp")]
        public Output<string> UpstreamIp { get; private set; } = null!;

        /// <summary>
        /// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
        /// </summary>
        [Output("upstreamPort")]
        public Output<int> UpstreamPort { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Csf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Csf(string name, CsfArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/csf:Csf", name, args ?? new CsfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Csf(string name, Input<string> id, CsfState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/csf:Csf", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "fixedKey",
                    "groupPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Csf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Csf Get(string name, Input<string> id, CsfState? state = null, CustomResourceOptions? options = null)
        {
            return new Csf(name, id, state, options);
        }
    }

    public sealed class CsfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("acceptAuthByCert")]
        public Input<string>? AcceptAuthByCert { get; set; }

        /// <summary>
        /// Authorization request type. Valid values: `serial`, `certificate`.
        /// </summary>
        [Input("authorizationRequestType")]
        public Input<string>? AuthorizationRequestType { get; set; }

        /// <summary>
        /// Certificate.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Configuration sync mode. Valid values: `default`, `local`.
        /// </summary>
        [Input("configurationSync")]
        public Input<string>? ConfigurationSync { get; set; }

        /// <summary>
        /// Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("downstreamAccess")]
        public Input<string>? DownstreamAccess { get; set; }

        /// <summary>
        /// Default access profile for requests from downstream devices.
        /// </summary>
        [Input("downstreamAccprofile")]
        public Input<string>? DownstreamAccprofile { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("fabricConnectors")]
        private InputList<Inputs.CsfFabricConnectorArgs>? _fabricConnectors;

        /// <summary>
        /// Fabric connector configuration. The structure of `fabric_connector` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfFabricConnectorArgs> FabricConnectors
        {
            get => _fabricConnectors ?? (_fabricConnectors = new InputList<Inputs.CsfFabricConnectorArgs>());
            set => _fabricConnectors = value;
        }

        [Input("fabricDevices")]
        private InputList<Inputs.CsfFabricDeviceArgs>? _fabricDevices;

        /// <summary>
        /// Fabric device configuration. The structure of `fabric_device` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfFabricDeviceArgs> FabricDevices
        {
            get => _fabricDevices ?? (_fabricDevices = new InputList<Inputs.CsfFabricDeviceArgs>());
            set => _fabricDevices = value;
        }

        /// <summary>
        /// Fabric CMDB Object Unification Valid values: `default`, `local`.
        /// </summary>
        [Input("fabricObjectUnification")]
        public Input<string>? FabricObjectUnification { get; set; }

        /// <summary>
        /// Number of worker processes for Security Fabric daemon.
        /// </summary>
        [Input("fabricWorkers")]
        public Input<int>? FabricWorkers { get; set; }

        /// <summary>
        /// Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fileMgmt")]
        public Input<string>? FileMgmt { get; set; }

        /// <summary>
        /// Maximum amount of memory that can be used by the daemon files (in bytes).
        /// </summary>
        [Input("fileQuota")]
        public Input<int>? FileQuota { get; set; }

        /// <summary>
        /// Warn when the set percentage of quota has been used.
        /// </summary>
        [Input("fileQuotaWarning")]
        public Input<int>? FileQuotaWarning { get; set; }

        [Input("fixedKey")]
        private Input<string>? _fixedKey;

        /// <summary>
        /// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
        /// </summary>
        public Input<string>? FixedKey
        {
            get => _fixedKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _fixedKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticloudAccountEnforcement")]
        public Input<string>? ForticloudAccountEnforcement { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("groupPassword")]
        private Input<string>? _groupPassword;

        /// <summary>
        /// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
        /// </summary>
        public Input<string>? GroupPassword
        {
            get => _groupPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _groupPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logUnification")]
        public Input<string>? LogUnification { get; set; }

        /// <summary>
        /// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
        /// </summary>
        [Input("managementIp")]
        public Input<string>? ManagementIp { get; set; }

        /// <summary>
        /// Overriding port for management connection (Overrides admin port).
        /// </summary>
        [Input("managementPort")]
        public Input<int>? ManagementPort { get; set; }

        /// <summary>
        /// SAML setting configuration synchronization. Valid values: `default`, `local`.
        /// </summary>
        [Input("samlConfigurationSync")]
        public Input<string>? SamlConfigurationSync { get; set; }

        /// <summary>
        /// Enable/disable Security Fabric. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        [Input("trustedLists")]
        private InputList<Inputs.CsfTrustedListArgs>? _trustedLists;

        /// <summary>
        /// Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfTrustedListArgs> TrustedLists
        {
            get => _trustedLists ?? (_trustedLists = new InputList<Inputs.CsfTrustedListArgs>());
            set => _trustedLists = value;
        }

        /// <summary>
        /// Unique ID of the current CSF node
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Input("upstream")]
        public Input<string>? Upstream { get; set; }

        /// <summary>
        /// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Input("upstreamIp")]
        public Input<string>? UpstreamIp { get; set; }

        /// <summary>
        /// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
        /// </summary>
        [Input("upstreamPort")]
        public Input<int>? UpstreamPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CsfArgs()
        {
        }
        public static new CsfArgs Empty => new CsfArgs();
    }

    public sealed class CsfState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("acceptAuthByCert")]
        public Input<string>? AcceptAuthByCert { get; set; }

        /// <summary>
        /// Authorization request type. Valid values: `serial`, `certificate`.
        /// </summary>
        [Input("authorizationRequestType")]
        public Input<string>? AuthorizationRequestType { get; set; }

        /// <summary>
        /// Certificate.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Configuration sync mode. Valid values: `default`, `local`.
        /// </summary>
        [Input("configurationSync")]
        public Input<string>? ConfigurationSync { get; set; }

        /// <summary>
        /// Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("downstreamAccess")]
        public Input<string>? DownstreamAccess { get; set; }

        /// <summary>
        /// Default access profile for requests from downstream devices.
        /// </summary>
        [Input("downstreamAccprofile")]
        public Input<string>? DownstreamAccprofile { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("fabricConnectors")]
        private InputList<Inputs.CsfFabricConnectorGetArgs>? _fabricConnectors;

        /// <summary>
        /// Fabric connector configuration. The structure of `fabric_connector` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfFabricConnectorGetArgs> FabricConnectors
        {
            get => _fabricConnectors ?? (_fabricConnectors = new InputList<Inputs.CsfFabricConnectorGetArgs>());
            set => _fabricConnectors = value;
        }

        [Input("fabricDevices")]
        private InputList<Inputs.CsfFabricDeviceGetArgs>? _fabricDevices;

        /// <summary>
        /// Fabric device configuration. The structure of `fabric_device` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfFabricDeviceGetArgs> FabricDevices
        {
            get => _fabricDevices ?? (_fabricDevices = new InputList<Inputs.CsfFabricDeviceGetArgs>());
            set => _fabricDevices = value;
        }

        /// <summary>
        /// Fabric CMDB Object Unification Valid values: `default`, `local`.
        /// </summary>
        [Input("fabricObjectUnification")]
        public Input<string>? FabricObjectUnification { get; set; }

        /// <summary>
        /// Number of worker processes for Security Fabric daemon.
        /// </summary>
        [Input("fabricWorkers")]
        public Input<int>? FabricWorkers { get; set; }

        /// <summary>
        /// Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fileMgmt")]
        public Input<string>? FileMgmt { get; set; }

        /// <summary>
        /// Maximum amount of memory that can be used by the daemon files (in bytes).
        /// </summary>
        [Input("fileQuota")]
        public Input<int>? FileQuota { get; set; }

        /// <summary>
        /// Warn when the set percentage of quota has been used.
        /// </summary>
        [Input("fileQuotaWarning")]
        public Input<int>? FileQuotaWarning { get; set; }

        [Input("fixedKey")]
        private Input<string>? _fixedKey;

        /// <summary>
        /// Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
        /// </summary>
        public Input<string>? FixedKey
        {
            get => _fixedKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _fixedKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticloudAccountEnforcement")]
        public Input<string>? ForticloudAccountEnforcement { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("groupPassword")]
        private Input<string>? _groupPassword;

        /// <summary>
        /// Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
        /// </summary>
        public Input<string>? GroupPassword
        {
            get => _groupPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _groupPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logUnification")]
        public Input<string>? LogUnification { get; set; }

        /// <summary>
        /// Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
        /// </summary>
        [Input("managementIp")]
        public Input<string>? ManagementIp { get; set; }

        /// <summary>
        /// Overriding port for management connection (Overrides admin port).
        /// </summary>
        [Input("managementPort")]
        public Input<int>? ManagementPort { get; set; }

        /// <summary>
        /// SAML setting configuration synchronization. Valid values: `default`, `local`.
        /// </summary>
        [Input("samlConfigurationSync")]
        public Input<string>? SamlConfigurationSync { get; set; }

        /// <summary>
        /// Enable/disable Security Fabric. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("trustedLists")]
        private InputList<Inputs.CsfTrustedListGetArgs>? _trustedLists;

        /// <summary>
        /// Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
        /// </summary>
        public InputList<Inputs.CsfTrustedListGetArgs> TrustedLists
        {
            get => _trustedLists ?? (_trustedLists = new InputList<Inputs.CsfTrustedListGetArgs>());
            set => _trustedLists = value;
        }

        /// <summary>
        /// Unique ID of the current CSF node
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Input("upstream")]
        public Input<string>? Upstream { get; set; }

        /// <summary>
        /// IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
        /// </summary>
        [Input("upstreamIp")]
        public Input<string>? UpstreamIp { get; set; }

        /// <summary>
        /// The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
        /// </summary>
        [Input("upstreamPort")]
        public Input<int>? UpstreamPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CsfState()
        {
        }
        public static new CsfState Empty => new CsfState();
    }
}
