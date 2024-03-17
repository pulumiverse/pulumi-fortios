// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure IPv6 DoS policies.
    /// 
    /// ## Import
    /// 
    /// Firewall DosPolicy6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/doSpolicy6:DoSpolicy6 labelname {{policyid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/doSpolicy6:DoSpolicy6 labelname {{policyid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/doSpolicy6:DoSpolicy6")]
    public partial class DoSpolicy6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Anomaly name. The structure of `anomaly` block is documented below.
        /// </summary>
        [Output("anomalies")]
        public Output<ImmutableArray<Outputs.DoSpolicy6Anomaly>> Anomalies { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
        /// </summary>
        [Output("dstaddrs")]
        public Output<ImmutableArray<Outputs.DoSpolicy6Dstaddr>> Dstaddrs { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Incoming interface name from available interfaces.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("policyid")]
        public Output<int> Policyid { get; private set; } = null!;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<Outputs.DoSpolicy6Service>> Services { get; private set; } = null!;

        /// <summary>
        /// Source address name from available addresses. The structure of `srcaddr` block is documented below.
        /// </summary>
        [Output("srcaddrs")]
        public Output<ImmutableArray<Outputs.DoSpolicy6Srcaddr>> Srcaddrs { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a DoSpolicy6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DoSpolicy6(string name, DoSpolicy6Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/doSpolicy6:DoSpolicy6", name, args ?? new DoSpolicy6Args(), MakeResourceOptions(options, ""))
        {
        }

        private DoSpolicy6(string name, Input<string> id, DoSpolicy6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/doSpolicy6:DoSpolicy6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DoSpolicy6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DoSpolicy6 Get(string name, Input<string> id, DoSpolicy6State? state = null, CustomResourceOptions? options = null)
        {
            return new DoSpolicy6(name, id, state, options);
        }
    }

    public sealed class DoSpolicy6Args : global::Pulumi.ResourceArgs
    {
        [Input("anomalies")]
        private InputList<Inputs.DoSpolicy6AnomalyArgs>? _anomalies;

        /// <summary>
        /// Anomaly name. The structure of `anomaly` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6AnomalyArgs> Anomalies
        {
            get => _anomalies ?? (_anomalies = new InputList<Inputs.DoSpolicy6AnomalyArgs>());
            set => _anomalies = value;
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("dstaddrs", required: true)]
        private InputList<Inputs.DoSpolicy6DstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6DstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.DoSpolicy6DstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Incoming interface name from available interfaces.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        [Input("services")]
        private InputList<Inputs.DoSpolicy6ServiceArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6ServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.DoSpolicy6ServiceArgs>());
            set => _services = value;
        }

        [Input("srcaddrs", required: true)]
        private InputList<Inputs.DoSpolicy6SrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source address name from available addresses. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6SrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.DoSpolicy6SrcaddrArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DoSpolicy6Args()
        {
        }
        public static new DoSpolicy6Args Empty => new DoSpolicy6Args();
    }

    public sealed class DoSpolicy6State : global::Pulumi.ResourceArgs
    {
        [Input("anomalies")]
        private InputList<Inputs.DoSpolicy6AnomalyGetArgs>? _anomalies;

        /// <summary>
        /// Anomaly name. The structure of `anomaly` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6AnomalyGetArgs> Anomalies
        {
            get => _anomalies ?? (_anomalies = new InputList<Inputs.DoSpolicy6AnomalyGetArgs>());
            set => _anomalies = value;
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("dstaddrs")]
        private InputList<Inputs.DoSpolicy6DstaddrGetArgs>? _dstaddrs;

        /// <summary>
        /// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6DstaddrGetArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.DoSpolicy6DstaddrGetArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Incoming interface name from available interfaces.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<int>? Policyid { get; set; }

        [Input("services")]
        private InputList<Inputs.DoSpolicy6ServiceGetArgs>? _services;

        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6ServiceGetArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.DoSpolicy6ServiceGetArgs>());
            set => _services = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.DoSpolicy6SrcaddrGetArgs>? _srcaddrs;

        /// <summary>
        /// Source address name from available addresses. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.DoSpolicy6SrcaddrGetArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.DoSpolicy6SrcaddrGetArgs>());
            set => _srcaddrs = value;
        }

        /// <summary>
        /// Enable/disable this policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DoSpolicy6State()
        {
        }
        public static new DoSpolicy6State Empty => new DoSpolicy6State();
    }
}
