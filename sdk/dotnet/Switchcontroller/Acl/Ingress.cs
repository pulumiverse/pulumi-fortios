// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Acl
{
    /// <summary>
    /// Configure ingress ACL policies to be applied on managed FortiSwitch ports. Applies to FortiOS Version `&gt;= 7.4.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerAcl Ingress can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/acl/ingress:Ingress labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/acl/ingress:Ingress labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/acl/ingress:Ingress")]
    public partial class Ingress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ACL actions. The structure of `action` block is documented below.
        /// </summary>
        [Output("action")]
        public Output<Outputs.IngressAction> Action { get; private set; } = null!;

        /// <summary>
        /// ACL classifiers. The structure of `classifier` block is documented below.
        /// </summary>
        [Output("classifier")]
        public Output<Outputs.IngressClassifier> Classifier { get; private set; } = null!;

        /// <summary>
        /// Description for the ACL policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// ACL ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ingress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ingress(string name, IngressArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/acl/ingress:Ingress", name, args ?? new IngressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ingress(string name, Input<string> id, IngressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/acl/ingress:Ingress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ingress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ingress Get(string name, Input<string> id, IngressState? state = null, CustomResourceOptions? options = null)
        {
            return new Ingress(name, id, state, options);
        }
    }

    public sealed class IngressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ACL actions. The structure of `action` block is documented below.
        /// </summary>
        [Input("action")]
        public Input<Inputs.IngressActionArgs>? Action { get; set; }

        /// <summary>
        /// ACL classifiers. The structure of `classifier` block is documented below.
        /// </summary>
        [Input("classifier")]
        public Input<Inputs.IngressClassifierArgs>? Classifier { get; set; }

        /// <summary>
        /// Description for the ACL policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ACL ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IngressArgs()
        {
        }
        public static new IngressArgs Empty => new IngressArgs();
    }

    public sealed class IngressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ACL actions. The structure of `action` block is documented below.
        /// </summary>
        [Input("action")]
        public Input<Inputs.IngressActionGetArgs>? Action { get; set; }

        /// <summary>
        /// ACL classifiers. The structure of `classifier` block is documented below.
        /// </summary>
        [Input("classifier")]
        public Input<Inputs.IngressClassifierGetArgs>? Classifier { get; set; }

        /// <summary>
        /// Description for the ACL policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ACL ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IngressState()
        {
        }
        public static new IngressState Empty => new IngressState();
    }
}
