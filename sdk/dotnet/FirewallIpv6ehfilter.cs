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
    /// Configure IPv6 extension header filter.
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
    ///     var trname = new Fortios.FirewallIpv6ehfilter("trname", new()
    ///     {
    ///         Auth = "disable",
    ///         DestOpt = "disable",
    ///         Fragment = "disable",
    ///         HopOpt = "disable",
    ///         NoNext = "disable",
    ///         Routing = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Ipv6EhFilter can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallIpv6ehfilter:FirewallIpv6ehfilter labelname FirewallIpv6EhFilter
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallIpv6ehfilter:FirewallIpv6ehfilter labelname FirewallIpv6EhFilter
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallIpv6ehfilter:FirewallIpv6ehfilter")]
    public partial class FirewallIpv6ehfilter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("auth")]
        public Output<string> Auth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("destOpt")]
        public Output<string> DestOpt { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fragment")]
        public Output<string> Fragment { get; private set; } = null!;

        /// <summary>
        /// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        /// </summary>
        [Output("hdoptType")]
        public Output<int> HdoptType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("hopOpt")]
        public Output<string> HopOpt { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        /// </summary>
        [Output("noNext")]
        public Output<string> NoNext { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("routing")]
        public Output<string> Routing { get; private set; } = null!;

        /// <summary>
        /// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        /// </summary>
        [Output("routingType")]
        public Output<int> RoutingType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallIpv6ehfilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallIpv6ehfilter(string name, FirewallIpv6ehfilterArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallIpv6ehfilter:FirewallIpv6ehfilter", name, args ?? new FirewallIpv6ehfilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallIpv6ehfilter(string name, Input<string> id, FirewallIpv6ehfilterState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallIpv6ehfilter:FirewallIpv6ehfilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallIpv6ehfilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallIpv6ehfilter Get(string name, Input<string> id, FirewallIpv6ehfilterState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallIpv6ehfilter(name, id, state, options);
        }
    }

    public sealed class FirewallIpv6ehfilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("destOpt")]
        public Input<string>? DestOpt { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fragment")]
        public Input<string>? Fragment { get; set; }

        /// <summary>
        /// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        /// </summary>
        [Input("hdoptType")]
        public Input<int>? HdoptType { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hopOpt")]
        public Input<string>? HopOpt { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        /// </summary>
        [Input("noNext")]
        public Input<string>? NoNext { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routing")]
        public Input<string>? Routing { get; set; }

        /// <summary>
        /// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        /// </summary>
        [Input("routingType")]
        public Input<int>? RoutingType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallIpv6ehfilterArgs()
        {
        }
        public static new FirewallIpv6ehfilterArgs Empty => new FirewallIpv6ehfilterArgs();
    }

    public sealed class FirewallIpv6ehfilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("destOpt")]
        public Input<string>? DestOpt { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fragment")]
        public Input<string>? Fragment { get; set; }

        /// <summary>
        /// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        /// </summary>
        [Input("hdoptType")]
        public Input<int>? HdoptType { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hopOpt")]
        public Input<string>? HopOpt { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        /// </summary>
        [Input("noNext")]
        public Input<string>? NoNext { get; set; }

        /// <summary>
        /// Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routing")]
        public Input<string>? Routing { get; set; }

        /// <summary>
        /// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        /// </summary>
        [Input("routingType")]
        public Input<int>? RoutingType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallIpv6ehfilterState()
        {
        }
        public static new FirewallIpv6ehfilterState Empty => new FirewallIpv6ehfilterState();
    }
}
