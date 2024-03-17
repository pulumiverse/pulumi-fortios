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
    /// Configure DNS64. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// System Dns64 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/dns64:Dns64 labelname SystemDns64
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/dns64:Dns64 labelname SystemDns64
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/dns64:Dns64")]
    public partial class Dns64 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("alwaysSynthesizeAaaaRecord")]
        public Output<string> AlwaysSynthesizeAaaaRecord { get; private set; } = null!;

        /// <summary>
        /// DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Output("dns64Prefix")]
        public Output<string> Dns64Prefix { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DNS64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Dns64 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dns64(string name, Dns64Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/dns64:Dns64", name, args ?? new Dns64Args(), MakeResourceOptions(options, ""))
        {
        }

        private Dns64(string name, Input<string> id, Dns64State? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/dns64:Dns64", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dns64 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dns64 Get(string name, Input<string> id, Dns64State? state = null, CustomResourceOptions? options = null)
        {
            return new Dns64(name, id, state, options);
        }
    }

    public sealed class Dns64Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysSynthesizeAaaaRecord")]
        public Input<string>? AlwaysSynthesizeAaaaRecord { get; set; }

        /// <summary>
        /// DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Input("dns64Prefix")]
        public Input<string>? Dns64Prefix { get; set; }

        /// <summary>
        /// Enable/disable DNS64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Dns64Args()
        {
        }
        public static new Dns64Args Empty => new Dns64Args();
    }

    public sealed class Dns64State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("alwaysSynthesizeAaaaRecord")]
        public Input<string>? AlwaysSynthesizeAaaaRecord { get; set; }

        /// <summary>
        /// DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
        /// </summary>
        [Input("dns64Prefix")]
        public Input<string>? Dns64Prefix { get; set; }

        /// <summary>
        /// Enable/disable DNS64 (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Dns64State()
        {
        }
        public static new Dns64State Empty => new Dns64State();
    }
}
