// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web
{
    /// <summary>
    /// Configure FortiGuard Web Filter service.
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
    ///     var trname = new Fortios.Filter.Web.Fortiguard("trname", new()
    ///     {
    ///         CacheMemPercent = 2,
    ///         CacheMode = "ttl",
    ///         CachePrefixMatch = "enable",
    ///         ClosePorts = "disable",
    ///         OvrdAuthHttps = "enable",
    ///         OvrdAuthPort = 8008,
    ///         OvrdAuthPortHttp = 8008,
    ///         OvrdAuthPortHttps = 8010,
    ///         OvrdAuthPortHttpsFlow = 8015,
    ///         OvrdAuthPortWarning = 8020,
    ///         WarnAuthHttps = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter Fortiguard can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/fortiguard:Fortiguard labelname WebfilterFortiguard
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/fortiguard:Fortiguard labelname WebfilterFortiguard
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/web/fortiguard:Fortiguard")]
    public partial class Fortiguard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum percentage of available memory allocated to caching (1 - 15%).
        /// </summary>
        [Output("cacheMemPercent")]
        public Output<int> CacheMemPercent { get; private set; } = null!;

        /// <summary>
        /// Maximum permille of available memory allocated to caching (1 - 150).
        /// </summary>
        [Output("cacheMemPermille")]
        public Output<int> CacheMemPermille { get; private set; } = null!;

        /// <summary>
        /// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        /// </summary>
        [Output("cacheMode")]
        public Output<string> CacheMode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cachePrefixMatch")]
        public Output<string> CachePrefixMatch { get; private set; } = null!;

        /// <summary>
        /// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("closePorts")]
        public Output<string> ClosePorts { get; private set; } = null!;

        /// <summary>
        /// Enable/disable embedding images into replacement messages (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("embedImage")]
        public Output<string> EmbedImage { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ovrdAuthHttps")]
        public Output<string> OvrdAuthHttps { get; private set; } = null!;

        /// <summary>
        /// Port to use for FortiGuard Web Filter override authentication.
        /// </summary>
        [Output("ovrdAuthPort")]
        public Output<int> OvrdAuthPort { get; private set; } = null!;

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTP override authentication
        /// </summary>
        [Output("ovrdAuthPortHttp")]
        public Output<int> OvrdAuthPortHttp { get; private set; } = null!;

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        /// </summary>
        [Output("ovrdAuthPortHttps")]
        public Output<int> OvrdAuthPortHttps { get; private set; } = null!;

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        /// </summary>
        [Output("ovrdAuthPortHttpsFlow")]
        public Output<int> OvrdAuthPortHttpsFlow { get; private set; } = null!;

        /// <summary>
        /// Port to use for FortiGuard Web Filter Warning override authentication.
        /// </summary>
        [Output("ovrdAuthPortWarning")]
        public Output<int> OvrdAuthPortWarning { get; private set; } = null!;

        /// <summary>
        /// Limit size of URL request packets sent to FortiGuard server (0 for default).
        /// </summary>
        [Output("requestPacketSizeLimit")]
        public Output<int> RequestPacketSizeLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("warnAuthHttps")]
        public Output<string> WarnAuthHttps { get; private set; } = null!;


        /// <summary>
        /// Create a Fortiguard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fortiguard(string name, FortiguardArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:filter/web/fortiguard:Fortiguard", name, args ?? new FortiguardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fortiguard(string name, Input<string> id, FortiguardState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/web/fortiguard:Fortiguard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fortiguard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fortiguard Get(string name, Input<string> id, FortiguardState? state = null, CustomResourceOptions? options = null)
        {
            return new Fortiguard(name, id, state, options);
        }
    }

    public sealed class FortiguardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum percentage of available memory allocated to caching (1 - 15%).
        /// </summary>
        [Input("cacheMemPercent")]
        public Input<int>? CacheMemPercent { get; set; }

        /// <summary>
        /// Maximum permille of available memory allocated to caching (1 - 150).
        /// </summary>
        [Input("cacheMemPermille")]
        public Input<int>? CacheMemPermille { get; set; }

        /// <summary>
        /// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        /// </summary>
        [Input("cacheMode")]
        public Input<string>? CacheMode { get; set; }

        /// <summary>
        /// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cachePrefixMatch")]
        public Input<string>? CachePrefixMatch { get; set; }

        /// <summary>
        /// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("closePorts")]
        public Input<string>? ClosePorts { get; set; }

        /// <summary>
        /// Enable/disable embedding images into replacement messages (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("embedImage")]
        public Input<string>? EmbedImage { get; set; }

        /// <summary>
        /// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ovrdAuthHttps")]
        public Input<string>? OvrdAuthHttps { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter override authentication.
        /// </summary>
        [Input("ovrdAuthPort")]
        public Input<int>? OvrdAuthPort { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTP override authentication
        /// </summary>
        [Input("ovrdAuthPortHttp")]
        public Input<int>? OvrdAuthPortHttp { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        /// </summary>
        [Input("ovrdAuthPortHttps")]
        public Input<int>? OvrdAuthPortHttps { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        /// </summary>
        [Input("ovrdAuthPortHttpsFlow")]
        public Input<int>? OvrdAuthPortHttpsFlow { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter Warning override authentication.
        /// </summary>
        [Input("ovrdAuthPortWarning")]
        public Input<int>? OvrdAuthPortWarning { get; set; }

        /// <summary>
        /// Limit size of URL request packets sent to FortiGuard server (0 for default).
        /// </summary>
        [Input("requestPacketSizeLimit")]
        public Input<int>? RequestPacketSizeLimit { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("warnAuthHttps")]
        public Input<string>? WarnAuthHttps { get; set; }

        public FortiguardArgs()
        {
        }
        public static new FortiguardArgs Empty => new FortiguardArgs();
    }

    public sealed class FortiguardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum percentage of available memory allocated to caching (1 - 15%).
        /// </summary>
        [Input("cacheMemPercent")]
        public Input<int>? CacheMemPercent { get; set; }

        /// <summary>
        /// Maximum permille of available memory allocated to caching (1 - 150).
        /// </summary>
        [Input("cacheMemPermille")]
        public Input<int>? CacheMemPermille { get; set; }

        /// <summary>
        /// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        /// </summary>
        [Input("cacheMode")]
        public Input<string>? CacheMode { get; set; }

        /// <summary>
        /// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cachePrefixMatch")]
        public Input<string>? CachePrefixMatch { get; set; }

        /// <summary>
        /// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("closePorts")]
        public Input<string>? ClosePorts { get; set; }

        /// <summary>
        /// Enable/disable embedding images into replacement messages (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("embedImage")]
        public Input<string>? EmbedImage { get; set; }

        /// <summary>
        /// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ovrdAuthHttps")]
        public Input<string>? OvrdAuthHttps { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter override authentication.
        /// </summary>
        [Input("ovrdAuthPort")]
        public Input<int>? OvrdAuthPort { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTP override authentication
        /// </summary>
        [Input("ovrdAuthPortHttp")]
        public Input<int>? OvrdAuthPortHttp { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        /// </summary>
        [Input("ovrdAuthPortHttps")]
        public Input<int>? OvrdAuthPortHttps { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        /// </summary>
        [Input("ovrdAuthPortHttpsFlow")]
        public Input<int>? OvrdAuthPortHttpsFlow { get; set; }

        /// <summary>
        /// Port to use for FortiGuard Web Filter Warning override authentication.
        /// </summary>
        [Input("ovrdAuthPortWarning")]
        public Input<int>? OvrdAuthPortWarning { get; set; }

        /// <summary>
        /// Limit size of URL request packets sent to FortiGuard server (0 for default).
        /// </summary>
        [Input("requestPacketSizeLimit")]
        public Input<int>? RequestPacketSizeLimit { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("warnAuthHttps")]
        public Input<string>? WarnAuthHttps { get; set; }

        public FortiguardState()
        {
        }
        public static new FortiguardState Empty => new FortiguardState();
    }
}
