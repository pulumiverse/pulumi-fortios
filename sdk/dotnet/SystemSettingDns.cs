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
    /// Provides a resource to configure DNS of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.SystemDns`, we recommend that you use the new resource.
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
    ///     var test1 = new Fortios.SystemSettingDns("test1", new()
    ///     {
    ///         DnsOverTls = "disable",
    ///         Primary = "208.91.112.53",
    ///         Secondary = "208.91.112.22",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/systemSettingDns:SystemSettingDns")]
    public partial class SystemSettingDns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        /// </summary>
        [Output("dnsOverTls")]
        public Output<string> DnsOverTls { get; private set; } = null!;

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Output("primary")]
        public Output<string> Primary { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Output("secondary")]
        public Output<string> Secondary { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSettingDns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSettingDns(string name, SystemSettingDnsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSettingDns:SystemSettingDns", name, args ?? new SystemSettingDnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSettingDns(string name, Input<string> id, SystemSettingDnsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSettingDns:SystemSettingDns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSettingDns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSettingDns Get(string name, Input<string> id, SystemSettingDnsState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSettingDns(name, id, state, options);
        }
    }

    public sealed class SystemSettingDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        public SystemSettingDnsArgs()
        {
        }
        public static new SystemSettingDnsArgs Empty => new SystemSettingDnsArgs();
    }

    public sealed class SystemSettingDnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        /// </summary>
        [Input("dnsOverTls")]
        public Input<string>? DnsOverTls { get; set; }

        /// <summary>
        /// Primary DNS server IP address.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// Secondary DNS server IP address.
        /// </summary>
        [Input("secondary")]
        public Input<string>? Secondary { get; set; }

        public SystemSettingDnsState()
        {
        }
        public static new SystemSettingDnsState Empty => new SystemSettingDnsState();
    }
}
