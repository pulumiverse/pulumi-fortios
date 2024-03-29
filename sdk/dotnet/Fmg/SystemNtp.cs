// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Fmg
{
    /// <summary>
    /// This resource supports modifying system ntp setting for FortiManager.
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
    ///     var test1 = new Fortios.Fmg.SystemNtp("test1", new()
    ///     {
    ///         Server = "ntp1.fortinet.com",
    ///         Status = "enable",
    ///         SyncInterval = 30,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/systemNtp:SystemNtp")]
    public partial class SystemNtp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Output("syncInterval")]
        public Output<int?> SyncInterval { get; private set; } = null!;


        /// <summary>
        /// Create a SystemNtp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemNtp(string name, SystemNtpArgs args, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemNtp:SystemNtp", name, args ?? new SystemNtpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemNtp(string name, Input<string> id, SystemNtpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemNtp:SystemNtp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemNtp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemNtp Get(string name, Input<string> id, SystemNtpState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemNtp(name, id, state, options);
        }
    }

    public sealed class SystemNtpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Input("syncInterval")]
        public Input<int>? SyncInterval { get; set; }

        public SystemNtpArgs()
        {
        }
        public static new SystemNtpArgs Empty => new SystemNtpArgs();
    }

    public sealed class SystemNtpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Input("syncInterval")]
        public Input<int>? SyncInterval { get; set; }

        public SystemNtpState()
        {
        }
        public static new SystemNtpState Empty => new SystemNtpState();
    }
}
