// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol
{
    /// <summary>
    /// Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `&gt;= 7.4.0`.
    /// 
    /// ## Import
    /// 
    /// EndpointControl FctemsOverride can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:endpointcontrol/fctemsoverride:Fctemsoverride labelname {{ems_id}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:endpointcontrol/fctemsoverride:Fctemsoverride labelname {{ems_id}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:endpointcontrol/fctemsoverride:Fctemsoverride")]
    public partial class Fctemsoverride : global::Pulumi.CustomResource
    {
        /// <summary>
        /// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
        /// </summary>
        [Output("callTimeout")]
        public Output<int> CallTimeout { get; private set; } = null!;

        /// <summary>
        /// List of EMS capabilities.
        /// </summary>
        [Output("capabilities")]
        public Output<string> Capabilities { get; private set; } = null!;

        /// <summary>
        /// Cloud server type. Valid values: `production`, `alpha`, `beta`.
        /// </summary>
        [Output("cloudServerType")]
        public Output<string> CloudServerType { get; private set; } = null!;

        /// <summary>
        /// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
        /// </summary>
        [Output("dirtyReason")]
        public Output<string> DirtyReason { get; private set; } = null!;

        /// <summary>
        /// EMS ID in order (1 - 7).
        /// </summary>
        [Output("emsId")]
        public Output<int> EmsId { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fortinetoneCloudAuthentication")]
        public Output<string> FortinetoneCloudAuthentication { get; private set; } = null!;

        /// <summary>
        /// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        /// </summary>
        [Output("httpsPort")]
        public Output<int> HttpsPort { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// FortiClient Enterprise Management Server (EMS) name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Outdated resource threshold in seconds (10 - 3600, default = 180).
        /// </summary>
        [Output("outOfSyncThreshold")]
        public Output<int> OutOfSyncThreshold { get; private set; } = null!;

        /// <summary>
        /// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("preserveSslSession")]
        public Output<string> PreserveSslSession { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pullAvatars")]
        public Output<string> PullAvatars { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pullMalwareHash")]
        public Output<string> PullMalwareHash { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pullSysinfo")]
        public Output<string> PullSysinfo { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pullTags")]
        public Output<string> PullTags { get; private set; } = null!;

        /// <summary>
        /// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pullVulnerabilities")]
        public Output<string> PullVulnerabilities { get; private set; } = null!;

        /// <summary>
        /// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sendTagsToAllVdoms")]
        public Output<string> SendTagsToAllVdoms { get; private set; } = null!;

        /// <summary>
        /// EMS Serial Number.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// FortiClient EMS FQDN or IPv4 address.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// REST API call source IP.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// EMS Tenant ID.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("trustCaCn")]
        public Output<string> TrustCaCn { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Lowest CA cert on Fortigate in verified EMS cert chain.
        /// </summary>
        [Output("verifyingCa")]
        public Output<string> VerifyingCa { get; private set; } = null!;

        /// <summary>
        /// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("websocketOverride")]
        public Output<string> WebsocketOverride { get; private set; } = null!;


        /// <summary>
        /// Create a Fctemsoverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fctemsoverride(string name, FctemsoverrideArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:endpointcontrol/fctemsoverride:Fctemsoverride", name, args ?? new FctemsoverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fctemsoverride(string name, Input<string> id, FctemsoverrideState? state = null, CustomResourceOptions? options = null)
            : base("fortios:endpointcontrol/fctemsoverride:Fctemsoverride", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fctemsoverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fctemsoverride Get(string name, Input<string> id, FctemsoverrideState? state = null, CustomResourceOptions? options = null)
        {
            return new Fctemsoverride(name, id, state, options);
        }
    }

    public sealed class FctemsoverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
        /// </summary>
        [Input("callTimeout")]
        public Input<int>? CallTimeout { get; set; }

        /// <summary>
        /// List of EMS capabilities.
        /// </summary>
        [Input("capabilities")]
        public Input<string>? Capabilities { get; set; }

        /// <summary>
        /// Cloud server type. Valid values: `production`, `alpha`, `beta`.
        /// </summary>
        [Input("cloudServerType")]
        public Input<string>? CloudServerType { get; set; }

        /// <summary>
        /// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
        /// </summary>
        [Input("dirtyReason")]
        public Input<string>? DirtyReason { get; set; }

        /// <summary>
        /// EMS ID in order (1 - 7).
        /// </summary>
        [Input("emsId")]
        public Input<int>? EmsId { get; set; }

        /// <summary>
        /// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortinetoneCloudAuthentication")]
        public Input<string>? FortinetoneCloudAuthentication { get; set; }

        /// <summary>
        /// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// FortiClient Enterprise Management Server (EMS) name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Outdated resource threshold in seconds (10 - 3600, default = 180).
        /// </summary>
        [Input("outOfSyncThreshold")]
        public Input<int>? OutOfSyncThreshold { get; set; }

        /// <summary>
        /// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("preserveSslSession")]
        public Input<string>? PreserveSslSession { get; set; }

        /// <summary>
        /// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullAvatars")]
        public Input<string>? PullAvatars { get; set; }

        /// <summary>
        /// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullMalwareHash")]
        public Input<string>? PullMalwareHash { get; set; }

        /// <summary>
        /// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullSysinfo")]
        public Input<string>? PullSysinfo { get; set; }

        /// <summary>
        /// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullTags")]
        public Input<string>? PullTags { get; set; }

        /// <summary>
        /// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullVulnerabilities")]
        public Input<string>? PullVulnerabilities { get; set; }

        /// <summary>
        /// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sendTagsToAllVdoms")]
        public Input<string>? SendTagsToAllVdoms { get; set; }

        /// <summary>
        /// EMS Serial Number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// FortiClient EMS FQDN or IPv4 address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// REST API call source IP.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// EMS Tenant ID.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("trustCaCn")]
        public Input<string>? TrustCaCn { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Lowest CA cert on Fortigate in verified EMS cert chain.
        /// </summary>
        [Input("verifyingCa")]
        public Input<string>? VerifyingCa { get; set; }

        /// <summary>
        /// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("websocketOverride")]
        public Input<string>? WebsocketOverride { get; set; }

        public FctemsoverrideArgs()
        {
        }
        public static new FctemsoverrideArgs Empty => new FctemsoverrideArgs();
    }

    public sealed class FctemsoverrideState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
        /// </summary>
        [Input("callTimeout")]
        public Input<int>? CallTimeout { get; set; }

        /// <summary>
        /// List of EMS capabilities.
        /// </summary>
        [Input("capabilities")]
        public Input<string>? Capabilities { get; set; }

        /// <summary>
        /// Cloud server type. Valid values: `production`, `alpha`, `beta`.
        /// </summary>
        [Input("cloudServerType")]
        public Input<string>? CloudServerType { get; set; }

        /// <summary>
        /// Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
        /// </summary>
        [Input("dirtyReason")]
        public Input<string>? DirtyReason { get; set; }

        /// <summary>
        /// EMS ID in order (1 - 7).
        /// </summary>
        [Input("emsId")]
        public Input<int>? EmsId { get; set; }

        /// <summary>
        /// Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortinetoneCloudAuthentication")]
        public Input<string>? FortinetoneCloudAuthentication { get; set; }

        /// <summary>
        /// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// FortiClient Enterprise Management Server (EMS) name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Outdated resource threshold in seconds (10 - 3600, default = 180).
        /// </summary>
        [Input("outOfSyncThreshold")]
        public Input<int>? OutOfSyncThreshold { get; set; }

        /// <summary>
        /// Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("preserveSslSession")]
        public Input<string>? PreserveSslSession { get; set; }

        /// <summary>
        /// Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullAvatars")]
        public Input<string>? PullAvatars { get; set; }

        /// <summary>
        /// Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullMalwareHash")]
        public Input<string>? PullMalwareHash { get; set; }

        /// <summary>
        /// Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullSysinfo")]
        public Input<string>? PullSysinfo { get; set; }

        /// <summary>
        /// Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullTags")]
        public Input<string>? PullTags { get; set; }

        /// <summary>
        /// Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pullVulnerabilities")]
        public Input<string>? PullVulnerabilities { get; set; }

        /// <summary>
        /// Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sendTagsToAllVdoms")]
        public Input<string>? SendTagsToAllVdoms { get; set; }

        /// <summary>
        /// EMS Serial Number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// FortiClient EMS FQDN or IPv4 address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// REST API call source IP.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// EMS Tenant ID.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("trustCaCn")]
        public Input<string>? TrustCaCn { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Lowest CA cert on Fortigate in verified EMS cert chain.
        /// </summary>
        [Input("verifyingCa")]
        public Input<string>? VerifyingCa { get; set; }

        /// <summary>
        /// Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("websocketOverride")]
        public Input<string>? WebsocketOverride { get; set; }

        public FctemsoverrideState()
        {
        }
        public static new FctemsoverrideState Empty => new FctemsoverrideState();
    }
}
