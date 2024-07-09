// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Tacacsaccounting.V2
{
    /// <summary>
    /// Settings for TACACS+ accounting events filter. Applies to FortiOS Version `&gt;= 7.0.2`.
    /// 
    /// ## Import
    /// 
    /// LogTacacsAccounting2 Filter can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/tacacsaccounting/v2/filter:Filter labelname LogTacacsAccounting2Filter
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/tacacsaccounting/v2/filter:Filter labelname LogTacacsAccounting2Filter
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/tacacsaccounting/v2/filter:Filter")]
    public partial class Filter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cliCmdAudit")]
        public Output<string> CliCmdAudit { get; private set; } = null!;

        /// <summary>
        /// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("configChangeAudit")]
        public Output<string> ConfigChangeAudit { get; private set; } = null!;

        /// <summary>
        /// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("loginAudit")]
        public Output<string> LoginAudit { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Filter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Filter(string name, FilterArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/tacacsaccounting/v2/filter:Filter", name, args ?? new FilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Filter(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/tacacsaccounting/v2/filter:Filter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Filter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Filter Get(string name, Input<string> id, FilterState? state = null, CustomResourceOptions? options = null)
        {
            return new Filter(name, id, state, options);
        }
    }

    public sealed class FilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cliCmdAudit")]
        public Input<string>? CliCmdAudit { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("configChangeAudit")]
        public Input<string>? ConfigChangeAudit { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("loginAudit")]
        public Input<string>? LoginAudit { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FilterArgs()
        {
        }
        public static new FilterArgs Empty => new FilterArgs();
    }

    public sealed class FilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cliCmdAudit")]
        public Input<string>? CliCmdAudit { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("configChangeAudit")]
        public Input<string>? ConfigChangeAudit { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("loginAudit")]
        public Input<string>? LoginAudit { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FilterState()
        {
        }
        public static new FilterState Empty => new FilterState();
    }
}
