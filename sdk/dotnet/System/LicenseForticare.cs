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
    /// Provides a resource to add a FortiCare license for FortiOS.
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
    ///     var test2 = new Fortios.System.LicenseForticare("test2", new()
    ///     {
    ///         RegistrationCode = "license",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:system/licenseForticare:LicenseForticare")]
    public partial class LicenseForticare : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Output("registrationCode")]
        public Output<string> RegistrationCode { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseForticare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseForticare(string name, LicenseForticareArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/licenseForticare:LicenseForticare", name, args ?? new LicenseForticareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseForticare(string name, Input<string> id, LicenseForticareState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/licenseForticare:LicenseForticare", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseForticare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseForticare Get(string name, Input<string> id, LicenseForticareState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseForticare(name, id, state, options);
        }
    }

    public sealed class LicenseForticareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Input("registrationCode", required: true)]
        public Input<string> RegistrationCode { get; set; } = null!;

        public LicenseForticareArgs()
        {
        }
        public static new LicenseForticareArgs Empty => new LicenseForticareArgs();
    }

    public sealed class LicenseForticareState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registration code.
        /// </summary>
        [Input("registrationCode")]
        public Input<string>? RegistrationCode { get; set; }

        public LicenseForticareState()
        {
        }
        public static new LicenseForticareState Empty => new LicenseForticareState();
    }
}
