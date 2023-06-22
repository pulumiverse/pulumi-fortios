// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure dedicated management.
    /// 
    /// ## Import
    /// 
    /// System DedicatedMgmt can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt")]
    public partial class SystemDedicatedmgmt : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Default gateway for dedicated management interface.
        /// </summary>
        [Output("defaultGateway")]
        public Output<string> DefaultGateway { get; private set; } = null!;

        /// <summary>
        /// DHCP end IP for dedicated management.
        /// </summary>
        [Output("dhcpEndIp")]
        public Output<string> DhcpEndIp { get; private set; } = null!;

        /// <summary>
        /// DHCP netmask.
        /// </summary>
        [Output("dhcpNetmask")]
        public Output<string> DhcpNetmask { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dhcpServer")]
        public Output<string> DhcpServer { get; private set; } = null!;

        /// <summary>
        /// DHCP start IP for dedicated management.
        /// </summary>
        [Output("dhcpStartIp")]
        public Output<string> DhcpStartIp { get; private set; } = null!;

        /// <summary>
        /// Dedicated management interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Enable/disable dedicated management. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemDedicatedmgmt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemDedicatedmgmt(string name, SystemDedicatedmgmtArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt", name, args ?? new SystemDedicatedmgmtArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemDedicatedmgmt(string name, Input<string> id, SystemDedicatedmgmtState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemDedicatedmgmt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemDedicatedmgmt Get(string name, Input<string> id, SystemDedicatedmgmtState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemDedicatedmgmt(name, id, state, options);
        }
    }

    public sealed class SystemDedicatedmgmtArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default gateway for dedicated management interface.
        /// </summary>
        [Input("defaultGateway")]
        public Input<string>? DefaultGateway { get; set; }

        /// <summary>
        /// DHCP end IP for dedicated management.
        /// </summary>
        [Input("dhcpEndIp")]
        public Input<string>? DhcpEndIp { get; set; }

        /// <summary>
        /// DHCP netmask.
        /// </summary>
        [Input("dhcpNetmask")]
        public Input<string>? DhcpNetmask { get; set; }

        /// <summary>
        /// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcpServer")]
        public Input<string>? DhcpServer { get; set; }

        /// <summary>
        /// DHCP start IP for dedicated management.
        /// </summary>
        [Input("dhcpStartIp")]
        public Input<string>? DhcpStartIp { get; set; }

        /// <summary>
        /// Dedicated management interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Enable/disable dedicated management. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDedicatedmgmtArgs()
        {
        }
        public static new SystemDedicatedmgmtArgs Empty => new SystemDedicatedmgmtArgs();
    }

    public sealed class SystemDedicatedmgmtState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default gateway for dedicated management interface.
        /// </summary>
        [Input("defaultGateway")]
        public Input<string>? DefaultGateway { get; set; }

        /// <summary>
        /// DHCP end IP for dedicated management.
        /// </summary>
        [Input("dhcpEndIp")]
        public Input<string>? DhcpEndIp { get; set; }

        /// <summary>
        /// DHCP netmask.
        /// </summary>
        [Input("dhcpNetmask")]
        public Input<string>? DhcpNetmask { get; set; }

        /// <summary>
        /// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dhcpServer")]
        public Input<string>? DhcpServer { get; set; }

        /// <summary>
        /// DHCP start IP for dedicated management.
        /// </summary>
        [Input("dhcpStartIp")]
        public Input<string>? DhcpStartIp { get; set; }

        /// <summary>
        /// Dedicated management interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Enable/disable dedicated management. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemDedicatedmgmtState()
        {
        }
        public static new SystemDedicatedmgmtState Empty => new SystemDedicatedmgmtState();
    }
}
