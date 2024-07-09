// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Snmp
{
    /// <summary>
    /// SNMP Access Control MIB View configuration. Applies to FortiOS Version `&gt;= 7.2.0`.
    /// 
    /// ## Import
    /// 
    /// SystemSnmp MibView can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/snmp/mibview:Mibview labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/snmp/mibview:Mibview labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/snmp/mibview:Mibview")]
    public partial class Mibview : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The OID subtrees to be excluded in the view. Maximum 64 allowed.
        /// </summary>
        [Output("exclude")]
        public Output<string> Exclude { get; private set; } = null!;

        /// <summary>
        /// The OID subtrees to be included in the view. Maximum 16 allowed.
        /// </summary>
        [Output("include")]
        public Output<string> Include { get; private set; } = null!;

        /// <summary>
        /// MIB view name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Mibview resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Mibview(string name, MibviewArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/snmp/mibview:Mibview", name, args ?? new MibviewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Mibview(string name, Input<string> id, MibviewState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/snmp/mibview:Mibview", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Mibview resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Mibview Get(string name, Input<string> id, MibviewState? state = null, CustomResourceOptions? options = null)
        {
            return new Mibview(name, id, state, options);
        }
    }

    public sealed class MibviewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OID subtrees to be excluded in the view. Maximum 64 allowed.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        /// <summary>
        /// The OID subtrees to be included in the view. Maximum 16 allowed.
        /// </summary>
        [Input("include")]
        public Input<string>? Include { get; set; }

        /// <summary>
        /// MIB view name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MibviewArgs()
        {
        }
        public static new MibviewArgs Empty => new MibviewArgs();
    }

    public sealed class MibviewState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OID subtrees to be excluded in the view. Maximum 64 allowed.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        /// <summary>
        /// The OID subtrees to be included in the view. Maximum 16 allowed.
        /// </summary>
        [Input("include")]
        public Input<string>? Include { get; set; }

        /// <summary>
        /// MIB view name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public MibviewState()
        {
        }
        public static new MibviewState Empty => new MibviewState();
    }
}
