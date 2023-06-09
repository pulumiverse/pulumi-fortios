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
    /// Settings for memory buffer.
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
    ///     var trname = new Fortios.LogmemorySetting("trname", new()
    ///     {
    ///         Diskfull = "overwrite",
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LogMemory Setting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/logmemorySetting:LogmemorySetting labelname LogMemorySetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/logmemorySetting:LogmemorySetting labelname LogMemorySetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/logmemorySetting:LogmemorySetting")]
    public partial class LogmemorySetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to take when memory is full. Valid values: `overwrite`.
        /// </summary>
        [Output("diskfull")]
        public Output<string> Diskfull { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a LogmemorySetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogmemorySetting(string name, LogmemorySettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/logmemorySetting:LogmemorySetting", name, args ?? new LogmemorySettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogmemorySetting(string name, Input<string> id, LogmemorySettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/logmemorySetting:LogmemorySetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogmemorySetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogmemorySetting Get(string name, Input<string> id, LogmemorySettingState? state = null, CustomResourceOptions? options = null)
        {
            return new LogmemorySetting(name, id, state, options);
        }
    }

    public sealed class LogmemorySettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when memory is full. Valid values: `overwrite`.
        /// </summary>
        [Input("diskfull")]
        public Input<string>? Diskfull { get; set; }

        /// <summary>
        /// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LogmemorySettingArgs()
        {
        }
        public static new LogmemorySettingArgs Empty => new LogmemorySettingArgs();
    }

    public sealed class LogmemorySettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when memory is full. Valid values: `overwrite`.
        /// </summary>
        [Input("diskfull")]
        public Input<string>? Diskfull { get; set; }

        /// <summary>
        /// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LogmemorySettingState()
        {
        }
        public static new LogmemorySettingState Empty => new LogmemorySettingState();
    }
}
