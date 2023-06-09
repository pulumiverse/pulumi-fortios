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
    /// Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.SystemVdom`, we recommend that you use the new resource.
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
    ///     var test2 = new Fortios.SystemVdomSetting("test2", new()
    ///     {
    ///         ShortName = "aa1122",
    ///         Temporary = "0",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/systemVdomSetting:SystemVdomSetting")]
    public partial class SystemVdomSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VDOM name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;

        /// <summary>
        /// Temporary.
        /// </summary>
        [Output("temporary")]
        public Output<string> Temporary { get; private set; } = null!;


        /// <summary>
        /// Create a SystemVdomSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemVdomSetting(string name, SystemVdomSettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVdomSetting:SystemVdomSetting", name, args ?? new SystemVdomSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemVdomSetting(string name, Input<string> id, SystemVdomSettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVdomSetting:SystemVdomSetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemVdomSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemVdomSetting Get(string name, Input<string> id, SystemVdomSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemVdomSetting(name, id, state, options);
        }
    }

    public sealed class SystemVdomSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// Temporary.
        /// </summary>
        [Input("temporary")]
        public Input<string>? Temporary { get; set; }

        public SystemVdomSettingArgs()
        {
        }
        public static new SystemVdomSettingArgs Empty => new SystemVdomSettingArgs();
    }

    public sealed class SystemVdomSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// Temporary.
        /// </summary>
        [Input("temporary")]
        public Input<string>? Temporary { get; set; }

        public SystemVdomSettingState()
        {
        }
        public static new SystemVdomSettingState Empty => new SystemVdomSettingState();
    }
}
