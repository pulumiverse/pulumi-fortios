// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wanopt
{
    /// <summary>
    /// Configure global Web cache settings.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Wanopt.Webcache("trname", new()
    ///     {
    ///         AlwaysRevalidate = "disable",
    ///         CacheByDefault = "disable",
    ///         CacheCookie = "disable",
    ///         CacheExpired = "disable",
    ///         DefaultTtl = 1440,
    ///         External = "disable",
    ///         FreshFactor = 100,
    ///         HostValidate = "disable",
    ///         IgnoreConditional = "disable",
    ///         IgnoreIeReload = "enable",
    ///         IgnoreIms = "disable",
    ///         IgnorePnc = "disable",
    ///         MaxObjectSize = 512000,
    ///         MaxTtl = 7200,
    ///         MinTtl = 5,
    ///         NegRespTime = 0,
    ///         RevalPnc = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Wanopt Webcache can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/webcache:Webcache labelname WanoptWebcache
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/webcache:Webcache labelname WanoptWebcache
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wanopt/webcache:Webcache")]
    public partial class Webcache : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("alwaysRevalidate")]
        public Output<string> AlwaysRevalidate { get; private set; } = null!;

        /// <summary>
        /// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cacheByDefault")]
        public Output<string> CacheByDefault { get; private set; } = null!;

        /// <summary>
        /// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cacheCookie")]
        public Output<string> CacheCookie { get; private set; } = null!;

        /// <summary>
        /// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cacheExpired")]
        public Output<string> CacheExpired { get; private set; } = null!;

        /// <summary>
        /// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
        /// </summary>
        [Output("defaultTtl")]
        public Output<int> DefaultTtl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable external Web caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("external")]
        public Output<string> External { get; private set; } = null!;

        /// <summary>
        /// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
        /// </summary>
        [Output("freshFactor")]
        public Output<int> FreshFactor { get; private set; } = null!;

        /// <summary>
        /// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("hostValidate")]
        public Output<string> HostValidate { get; private set; } = null!;

        /// <summary>
        /// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ignoreConditional")]
        public Output<string> IgnoreConditional { get; private set; } = null!;

        /// <summary>
        /// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ignoreIeReload")]
        public Output<string> IgnoreIeReload { get; private set; } = null!;

        /// <summary>
        /// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ignoreIms")]
        public Output<string> IgnoreIms { get; private set; } = null!;

        /// <summary>
        /// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ignorePnc")]
        public Output<string> IgnorePnc { get; private set; } = null!;

        /// <summary>
        /// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
        /// </summary>
        [Output("maxObjectSize")]
        public Output<int> MaxObjectSize { get; private set; } = null!;

        /// <summary>
        /// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
        /// </summary>
        [Output("maxTtl")]
        public Output<int> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
        /// </summary>
        [Output("minTtl")]
        public Output<int> MinTtl { get; private set; } = null!;

        /// <summary>
        /// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
        /// </summary>
        [Output("negRespTime")]
        public Output<int> NegRespTime { get; private set; } = null!;

        /// <summary>
        /// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("revalPnc")]
        public Output<string> RevalPnc { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Webcache resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webcache(string name, WebcacheArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/webcache:Webcache", name, args ?? new WebcacheArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webcache(string name, Input<string> id, WebcacheState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/webcache:Webcache", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Webcache resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webcache Get(string name, Input<string> id, WebcacheState? state = null, CustomResourceOptions? options = null)
        {
            return new Webcache(name, id, state, options);
        }
    }

    public sealed class WebcacheArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysRevalidate")]
        public Input<string>? AlwaysRevalidate { get; set; }

        /// <summary>
        /// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheByDefault")]
        public Input<string>? CacheByDefault { get; set; }

        /// <summary>
        /// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheCookie")]
        public Input<string>? CacheCookie { get; set; }

        /// <summary>
        /// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheExpired")]
        public Input<string>? CacheExpired { get; set; }

        /// <summary>
        /// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// Enable/disable external Web caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("external")]
        public Input<string>? External { get; set; }

        /// <summary>
        /// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
        /// </summary>
        [Input("freshFactor")]
        public Input<int>? FreshFactor { get; set; }

        /// <summary>
        /// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hostValidate")]
        public Input<string>? HostValidate { get; set; }

        /// <summary>
        /// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreConditional")]
        public Input<string>? IgnoreConditional { get; set; }

        /// <summary>
        /// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreIeReload")]
        public Input<string>? IgnoreIeReload { get; set; }

        /// <summary>
        /// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreIms")]
        public Input<string>? IgnoreIms { get; set; }

        /// <summary>
        /// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignorePnc")]
        public Input<string>? IgnorePnc { get; set; }

        /// <summary>
        /// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
        /// </summary>
        [Input("maxObjectSize")]
        public Input<int>? MaxObjectSize { get; set; }

        /// <summary>
        /// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
        /// </summary>
        [Input("minTtl")]
        public Input<int>? MinTtl { get; set; }

        /// <summary>
        /// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
        /// </summary>
        [Input("negRespTime")]
        public Input<int>? NegRespTime { get; set; }

        /// <summary>
        /// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revalPnc")]
        public Input<string>? RevalPnc { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebcacheArgs()
        {
        }
        public static new WebcacheArgs Empty => new WebcacheArgs();
    }

    public sealed class WebcacheState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysRevalidate")]
        public Input<string>? AlwaysRevalidate { get; set; }

        /// <summary>
        /// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheByDefault")]
        public Input<string>? CacheByDefault { get; set; }

        /// <summary>
        /// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheCookie")]
        public Input<string>? CacheCookie { get; set; }

        /// <summary>
        /// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheExpired")]
        public Input<string>? CacheExpired { get; set; }

        /// <summary>
        /// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
        /// </summary>
        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        /// <summary>
        /// Enable/disable external Web caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("external")]
        public Input<string>? External { get; set; }

        /// <summary>
        /// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
        /// </summary>
        [Input("freshFactor")]
        public Input<int>? FreshFactor { get; set; }

        /// <summary>
        /// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hostValidate")]
        public Input<string>? HostValidate { get; set; }

        /// <summary>
        /// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreConditional")]
        public Input<string>? IgnoreConditional { get; set; }

        /// <summary>
        /// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreIeReload")]
        public Input<string>? IgnoreIeReload { get; set; }

        /// <summary>
        /// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreIms")]
        public Input<string>? IgnoreIms { get; set; }

        /// <summary>
        /// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignorePnc")]
        public Input<string>? IgnorePnc { get; set; }

        /// <summary>
        /// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
        /// </summary>
        [Input("maxObjectSize")]
        public Input<int>? MaxObjectSize { get; set; }

        /// <summary>
        /// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
        /// </summary>
        [Input("minTtl")]
        public Input<int>? MinTtl { get; set; }

        /// <summary>
        /// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
        /// </summary>
        [Input("negRespTime")]
        public Input<int>? NegRespTime { get; set; }

        /// <summary>
        /// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("revalPnc")]
        public Input<string>? RevalPnc { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebcacheState()
        {
        }
        public static new WebcacheState Empty => new WebcacheState();
    }
}