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
    /// Configure GENEVE devices. Applies to FortiOS Version `&gt;= 6.2.4`.
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
    ///     var trname = new Fortios.SystemGeneve("trname", new()
    ///     {
    ///         Dstport = 22,
    ///         Interface = "port2",
    ///         IpVersion = "ipv4-unicast",
    ///         RemoteIp = "1.1.1.1",
    ///         RemoteIp6 = "::",
    ///         Vni = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Geneve can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemGeneve:SystemGeneve")]
    public partial class SystemGeneve : global::Pulumi.CustomResource
    {
        /// <summary>
        /// GENEVE destination port (1 - 65535, default = 6081).
        /// </summary>
        [Output("dstport")]
        public Output<int> Dstport { get; private set; } = null!;

        /// <summary>
        /// Outgoing interface for GENEVE encapsulated traffic.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        /// </summary>
        [Output("ipVersion")]
        public Output<string> IpVersion { get; private set; } = null!;

        /// <summary>
        /// GENEVE device or interface name. Must be an unique interface name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Output("remoteIp")]
        public Output<string> RemoteIp { get; private set; } = null!;

        /// <summary>
        /// IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Output("remoteIp6")]
        public Output<string> RemoteIp6 { get; private set; } = null!;

        /// <summary>
        /// GENEVE type. Valid values: `ethernet`, `ppp`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// GENEVE network ID.
        /// </summary>
        [Output("vni")]
        public Output<int> Vni { get; private set; } = null!;


        /// <summary>
        /// Create a SystemGeneve resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemGeneve(string name, SystemGeneveArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/systemGeneve:SystemGeneve", name, args ?? new SystemGeneveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemGeneve(string name, Input<string> id, SystemGeneveState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemGeneve:SystemGeneve", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemGeneve resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemGeneve Get(string name, Input<string> id, SystemGeneveState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemGeneve(name, id, state, options);
        }
    }

    public sealed class SystemGeneveArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GENEVE destination port (1 - 65535, default = 6081).
        /// </summary>
        [Input("dstport")]
        public Input<int>? Dstport { get; set; }

        /// <summary>
        /// Outgoing interface for GENEVE encapsulated traffic.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        /// </summary>
        [Input("ipVersion", required: true)]
        public Input<string> IpVersion { get; set; } = null!;

        /// <summary>
        /// GENEVE device or interface name. Must be an unique interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Input("remoteIp", required: true)]
        public Input<string> RemoteIp { get; set; } = null!;

        /// <summary>
        /// IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Input("remoteIp6")]
        public Input<string>? RemoteIp6 { get; set; }

        /// <summary>
        /// GENEVE type. Valid values: `ethernet`, `ppp`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// GENEVE network ID.
        /// </summary>
        [Input("vni", required: true)]
        public Input<int> Vni { get; set; } = null!;

        public SystemGeneveArgs()
        {
        }
        public static new SystemGeneveArgs Empty => new SystemGeneveArgs();
    }

    public sealed class SystemGeneveState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GENEVE destination port (1 - 65535, default = 6081).
        /// </summary>
        [Input("dstport")]
        public Input<int>? Dstport { get; set; }

        /// <summary>
        /// Outgoing interface for GENEVE encapsulated traffic.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// GENEVE device or interface name. Must be an unique interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Input("remoteIp")]
        public Input<string>? RemoteIp { get; set; }

        /// <summary>
        /// IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        /// </summary>
        [Input("remoteIp6")]
        public Input<string>? RemoteIp6 { get; set; }

        /// <summary>
        /// GENEVE type. Valid values: `ethernet`, `ppp`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// GENEVE network ID.
        /// </summary>
        [Input("vni")]
        public Input<int>? Vni { get; set; }

        public SystemGeneveState()
        {
        }
        public static new SystemGeneveState Empty => new SystemGeneveState();
    }
}
