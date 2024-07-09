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
    /// Configure WCCP.
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
    ///     var trname = new Fortios.System.Wccp("trname", new()
    ///     {
    ///         AssignmentBucketFormat = "cisco-implementation",
    ///         AssignmentDstaddrMask = "0.0.0.0",
    ///         AssignmentMethod = "HASH",
    ///         AssignmentSrcaddrMask = "0.0.23.65",
    ///         AssignmentWeight = 0,
    ///         Authentication = "disable",
    ///         CacheEngineMethod = "GRE",
    ///         CacheId = "1.1.1.1",
    ///         ForwardMethod = "GRE",
    ///         GroupAddress = "0.0.0.0",
    ///         PrimaryHash = "dst-ip",
    ///         Priority = 0,
    ///         Protocol = 0,
    ///         ReturnMethod = "GRE",
    ///         RouterId = "1.1.1.1",
    ///         RouterList = "\"1.0.0.0\" ",
    ///         ServerType = "forward",
    ///         ServiceId = "1",
    ///         ServiceType = "auto",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Wccp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/wccp:Wccp labelname {{service_id}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/wccp:Wccp labelname {{service_id}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/wccp:Wccp")]
    public partial class Wccp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
        /// </summary>
        [Output("assignmentBucketFormat")]
        public Output<string> AssignmentBucketFormat { get; private set; } = null!;

        /// <summary>
        /// Assignment destination address mask.
        /// </summary>
        [Output("assignmentDstaddrMask")]
        public Output<string> AssignmentDstaddrMask { get; private set; } = null!;

        /// <summary>
        /// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
        /// </summary>
        [Output("assignmentMethod")]
        public Output<string> AssignmentMethod { get; private set; } = null!;

        /// <summary>
        /// Assignment source address mask.
        /// </summary>
        [Output("assignmentSrcaddrMask")]
        public Output<string> AssignmentSrcaddrMask { get; private set; } = null!;

        /// <summary>
        /// Assignment of hash weight/ratio for the WCCP cache engine.
        /// </summary>
        [Output("assignmentWeight")]
        public Output<int> AssignmentWeight { get; private set; } = null!;

        /// <summary>
        /// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authentication")]
        public Output<string> Authentication { get; private set; } = null!;

        /// <summary>
        /// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
        /// </summary>
        [Output("cacheEngineMethod")]
        public Output<string> CacheEngineMethod { get; private set; } = null!;

        /// <summary>
        /// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
        /// </summary>
        [Output("cacheId")]
        public Output<string> CacheId { get; private set; } = null!;

        /// <summary>
        /// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Output("forwardMethod")]
        public Output<string> ForwardMethod { get; private set; } = null!;

        /// <summary>
        /// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
        /// </summary>
        [Output("groupAddress")]
        public Output<string> GroupAddress { get; private set; } = null!;

        /// <summary>
        /// Password for MD5 authentication.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Service ports.
        /// </summary>
        [Output("ports")]
        public Output<string> Ports { get; private set; } = null!;

        /// <summary>
        /// Match method. Valid values: `source`, `destination`.
        /// </summary>
        [Output("portsDefined")]
        public Output<string> PortsDefined { get; private set; } = null!;

        /// <summary>
        /// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
        /// </summary>
        [Output("primaryHash")]
        public Output<string> PrimaryHash { get; private set; } = null!;

        /// <summary>
        /// Service priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Service protocol.
        /// </summary>
        [Output("protocol")]
        public Output<int> Protocol { get; private set; } = null!;

        /// <summary>
        /// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Output("returnMethod")]
        public Output<string> ReturnMethod { get; private set; } = null!;

        /// <summary>
        /// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// IP addresses of one or more WCCP routers.
        /// </summary>
        [Output("routerList")]
        public Output<string> RouterList { get; private set; } = null!;

        /// <summary>
        /// IP addresses and netmasks for up to four cache servers.
        /// </summary>
        [Output("serverList")]
        public Output<string> ServerList { get; private set; } = null!;

        /// <summary>
        /// Cache server type. Valid values: `forward`, `proxy`.
        /// </summary>
        [Output("serverType")]
        public Output<string> ServerType { get; private set; } = null!;

        /// <summary>
        /// Service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
        /// </summary>
        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Wccp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wccp(string name, WccpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/wccp:Wccp", name, args ?? new WccpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wccp(string name, Input<string> id, WccpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/wccp:Wccp", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Wccp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wccp Get(string name, Input<string> id, WccpState? state = null, CustomResourceOptions? options = null)
        {
            return new Wccp(name, id, state, options);
        }
    }

    public sealed class WccpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
        /// </summary>
        [Input("assignmentBucketFormat")]
        public Input<string>? AssignmentBucketFormat { get; set; }

        /// <summary>
        /// Assignment destination address mask.
        /// </summary>
        [Input("assignmentDstaddrMask")]
        public Input<string>? AssignmentDstaddrMask { get; set; }

        /// <summary>
        /// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
        /// </summary>
        [Input("assignmentMethod")]
        public Input<string>? AssignmentMethod { get; set; }

        /// <summary>
        /// Assignment source address mask.
        /// </summary>
        [Input("assignmentSrcaddrMask")]
        public Input<string>? AssignmentSrcaddrMask { get; set; }

        /// <summary>
        /// Assignment of hash weight/ratio for the WCCP cache engine.
        /// </summary>
        [Input("assignmentWeight")]
        public Input<int>? AssignmentWeight { get; set; }

        /// <summary>
        /// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
        /// </summary>
        [Input("cacheEngineMethod")]
        public Input<string>? CacheEngineMethod { get; set; }

        /// <summary>
        /// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
        /// </summary>
        [Input("cacheId")]
        public Input<string>? CacheId { get; set; }

        /// <summary>
        /// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Input("forwardMethod")]
        public Input<string>? ForwardMethod { get; set; }

        /// <summary>
        /// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
        /// </summary>
        [Input("groupAddress")]
        public Input<string>? GroupAddress { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for MD5 authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Service ports.
        /// </summary>
        [Input("ports")]
        public Input<string>? Ports { get; set; }

        /// <summary>
        /// Match method. Valid values: `source`, `destination`.
        /// </summary>
        [Input("portsDefined")]
        public Input<string>? PortsDefined { get; set; }

        /// <summary>
        /// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
        /// </summary>
        [Input("primaryHash")]
        public Input<string>? PrimaryHash { get; set; }

        /// <summary>
        /// Service priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Service protocol.
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Input("returnMethod")]
        public Input<string>? ReturnMethod { get; set; }

        /// <summary>
        /// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// IP addresses of one or more WCCP routers.
        /// </summary>
        [Input("routerList")]
        public Input<string>? RouterList { get; set; }

        /// <summary>
        /// IP addresses and netmasks for up to four cache servers.
        /// </summary>
        [Input("serverList")]
        public Input<string>? ServerList { get; set; }

        /// <summary>
        /// Cache server type. Valid values: `forward`, `proxy`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        /// <summary>
        /// Service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WccpArgs()
        {
        }
        public static new WccpArgs Empty => new WccpArgs();
    }

    public sealed class WccpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
        /// </summary>
        [Input("assignmentBucketFormat")]
        public Input<string>? AssignmentBucketFormat { get; set; }

        /// <summary>
        /// Assignment destination address mask.
        /// </summary>
        [Input("assignmentDstaddrMask")]
        public Input<string>? AssignmentDstaddrMask { get; set; }

        /// <summary>
        /// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
        /// </summary>
        [Input("assignmentMethod")]
        public Input<string>? AssignmentMethod { get; set; }

        /// <summary>
        /// Assignment source address mask.
        /// </summary>
        [Input("assignmentSrcaddrMask")]
        public Input<string>? AssignmentSrcaddrMask { get; set; }

        /// <summary>
        /// Assignment of hash weight/ratio for the WCCP cache engine.
        /// </summary>
        [Input("assignmentWeight")]
        public Input<int>? AssignmentWeight { get; set; }

        /// <summary>
        /// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
        /// </summary>
        [Input("cacheEngineMethod")]
        public Input<string>? CacheEngineMethod { get; set; }

        /// <summary>
        /// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
        /// </summary>
        [Input("cacheId")]
        public Input<string>? CacheId { get; set; }

        /// <summary>
        /// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Input("forwardMethod")]
        public Input<string>? ForwardMethod { get; set; }

        /// <summary>
        /// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
        /// </summary>
        [Input("groupAddress")]
        public Input<string>? GroupAddress { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for MD5 authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Service ports.
        /// </summary>
        [Input("ports")]
        public Input<string>? Ports { get; set; }

        /// <summary>
        /// Match method. Valid values: `source`, `destination`.
        /// </summary>
        [Input("portsDefined")]
        public Input<string>? PortsDefined { get; set; }

        /// <summary>
        /// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
        /// </summary>
        [Input("primaryHash")]
        public Input<string>? PrimaryHash { get; set; }

        /// <summary>
        /// Service priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Service protocol.
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
        /// </summary>
        [Input("returnMethod")]
        public Input<string>? ReturnMethod { get; set; }

        /// <summary>
        /// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// IP addresses of one or more WCCP routers.
        /// </summary>
        [Input("routerList")]
        public Input<string>? RouterList { get; set; }

        /// <summary>
        /// IP addresses and netmasks for up to four cache servers.
        /// </summary>
        [Input("serverList")]
        public Input<string>? ServerList { get; set; }

        /// <summary>
        /// Cache server type. Valid values: `forward`, `proxy`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        /// <summary>
        /// Service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WccpState()
        {
        }
        public static new WccpState Empty => new WccpState();
    }
}
