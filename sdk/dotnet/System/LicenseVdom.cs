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
    /// Provides a resource to add a VDOM license for FortiOS.
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
    ///     var test2 = new Fortios.System.LicenseVdom("test2", new()
    ///     {
    ///         License = "license",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:system/licenseVdom:LicenseVdom")]
    public partial class LicenseVdom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Output("license")]
        public Output<string> License { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseVdom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseVdom(string name, LicenseVdomArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/licenseVdom:LicenseVdom", name, args ?? new LicenseVdomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseVdom(string name, Input<string> id, LicenseVdomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/licenseVdom:LicenseVdom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseVdom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseVdom Get(string name, Input<string> id, LicenseVdomState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseVdom(name, id, state, options);
        }
    }

    public sealed class LicenseVdomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Input("license", required: true)]
        public Input<string> License { get; set; } = null!;

        public LicenseVdomArgs()
        {
        }
        public static new LicenseVdomArgs Empty => new LicenseVdomArgs();
    }

    public sealed class LicenseVdomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        public LicenseVdomState()
        {
        }
        public static new LicenseVdomState Empty => new LicenseVdomState();
    }
}
