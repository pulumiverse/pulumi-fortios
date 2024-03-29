// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb
{
    /// <summary>
    /// Configure CASB SaaS application. Applies to FortiOS Version `&gt;= 7.4.1`.
    /// 
    /// ## Import
    /// 
    /// Casb SaasApplication can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:casb/saasapplication:Saasapplication labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:casb/saasapplication:Saasapplication labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:casb/saasapplication:Saasapplication")]
    public partial class Saasapplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SaaS application signature name.
        /// </summary>
        [Output("casbName")]
        public Output<string> CasbName { get; private set; } = null!;

        /// <summary>
        /// SaaS application description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// SaaS application domain list. The structure of `domains` block is documented below.
        /// </summary>
        [Output("domains")]
        public Output<ImmutableArray<Outputs.SaasapplicationDomain>> Domains { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// SaaS application name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// SaaS application type. Valid values: `built-in`, `customized`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Saasapplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Saasapplication(string name, SaasapplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:casb/saasapplication:Saasapplication", name, args ?? new SaasapplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Saasapplication(string name, Input<string> id, SaasapplicationState? state = null, CustomResourceOptions? options = null)
            : base("fortios:casb/saasapplication:Saasapplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Saasapplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Saasapplication Get(string name, Input<string> id, SaasapplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Saasapplication(name, id, state, options);
        }
    }

    public sealed class SaasapplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SaaS application signature name.
        /// </summary>
        [Input("casbName")]
        public Input<string>? CasbName { get; set; }

        /// <summary>
        /// SaaS application description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domains")]
        private InputList<Inputs.SaasapplicationDomainArgs>? _domains;

        /// <summary>
        /// SaaS application domain list. The structure of `domains` block is documented below.
        /// </summary>
        public InputList<Inputs.SaasapplicationDomainArgs> Domains
        {
            get => _domains ?? (_domains = new InputList<Inputs.SaasapplicationDomainArgs>());
            set => _domains = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// SaaS application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SaaS application type. Valid values: `built-in`, `customized`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SaasapplicationArgs()
        {
        }
        public static new SaasapplicationArgs Empty => new SaasapplicationArgs();
    }

    public sealed class SaasapplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SaaS application signature name.
        /// </summary>
        [Input("casbName")]
        public Input<string>? CasbName { get; set; }

        /// <summary>
        /// SaaS application description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domains")]
        private InputList<Inputs.SaasapplicationDomainGetArgs>? _domains;

        /// <summary>
        /// SaaS application domain list. The structure of `domains` block is documented below.
        /// </summary>
        public InputList<Inputs.SaasapplicationDomainGetArgs> Domains
        {
            get => _domains ?? (_domains = new InputList<Inputs.SaasapplicationDomainGetArgs>());
            set => _domains = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// SaaS application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SaaS application type. Valid values: `built-in`, `customized`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SaasapplicationState()
        {
        }
        public static new SaasapplicationState Empty => new SaasapplicationState();
    }
}
