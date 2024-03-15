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
    /// ## Import
    /// 
    /// System ApiUser can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/apiuser:Apiuser labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/apiuser:Apiuser labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/apiuser:Apiuser")]
    public partial class Apiuser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Admin user access profile.
        /// </summary>
        [Output("accprofile")]
        public Output<string> Accprofile { get; private set; } = null!;

        /// <summary>
        /// Admin user password.
        /// </summary>
        [Output("apiKey")]
        public Output<string?> ApiKey { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        /// </summary>
        [Output("corsAllowOrigin")]
        public Output<string> CorsAllowOrigin { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// User name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable peer authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("peerAuth")]
        public Output<string> PeerAuth { get; private set; } = null!;

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Output("peerGroup")]
        public Output<string> PeerGroup { get; private set; } = null!;

        /// <summary>
        /// Schedule name.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// Trusthost. The structure of `trusthost` block is documented below.
        /// </summary>
        [Output("trusthosts")]
        public Output<ImmutableArray<Outputs.ApiuserTrusthost>> Trusthosts { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Virtual domains. The structure of `vdom` block is documented below.
        /// </summary>
        [Output("vdoms")]
        public Output<ImmutableArray<Outputs.ApiuserVdom>> Vdoms { get; private set; } = null!;


        /// <summary>
        /// Create a Apiuser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Apiuser(string name, ApiuserArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/apiuser:Apiuser", name, args ?? new ApiuserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Apiuser(string name, Input<string> id, ApiuserState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/apiuser:Apiuser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "apiKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Apiuser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Apiuser Get(string name, Input<string> id, ApiuserState? state = null, CustomResourceOptions? options = null)
        {
            return new Apiuser(name, id, state, options);
        }
    }

    public sealed class ApiuserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin user access profile.
        /// </summary>
        [Input("accprofile", required: true)]
        public Input<string> Accprofile { get; set; } = null!;

        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// Admin user password.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        /// </summary>
        [Input("corsAllowOrigin")]
        public Input<string>? CorsAllowOrigin { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable peer authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("peerAuth")]
        public Input<string>? PeerAuth { get; set; }

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Input("peerGroup")]
        public Input<string>? PeerGroup { get; set; }

        /// <summary>
        /// Schedule name.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("trusthosts")]
        private InputList<Inputs.ApiuserTrusthostArgs>? _trusthosts;

        /// <summary>
        /// Trusthost. The structure of `trusthost` block is documented below.
        /// </summary>
        public InputList<Inputs.ApiuserTrusthostArgs> Trusthosts
        {
            get => _trusthosts ?? (_trusthosts = new InputList<Inputs.ApiuserTrusthostArgs>());
            set => _trusthosts = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.ApiuserVdomArgs>? _vdoms;

        /// <summary>
        /// Virtual domains. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.ApiuserVdomArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.ApiuserVdomArgs>());
            set => _vdoms = value;
        }

        public ApiuserArgs()
        {
        }
        public static new ApiuserArgs Empty => new ApiuserArgs();
    }

    public sealed class ApiuserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin user access profile.
        /// </summary>
        [Input("accprofile")]
        public Input<string>? Accprofile { get; set; }

        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// Admin user password.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
        /// </summary>
        [Input("corsAllowOrigin")]
        public Input<string>? CorsAllowOrigin { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable peer authentication. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("peerAuth")]
        public Input<string>? PeerAuth { get; set; }

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Input("peerGroup")]
        public Input<string>? PeerGroup { get; set; }

        /// <summary>
        /// Schedule name.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("trusthosts")]
        private InputList<Inputs.ApiuserTrusthostGetArgs>? _trusthosts;

        /// <summary>
        /// Trusthost. The structure of `trusthost` block is documented below.
        /// </summary>
        public InputList<Inputs.ApiuserTrusthostGetArgs> Trusthosts
        {
            get => _trusthosts ?? (_trusthosts = new InputList<Inputs.ApiuserTrusthostGetArgs>());
            set => _trusthosts = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.ApiuserVdomGetArgs>? _vdoms;

        /// <summary>
        /// Virtual domains. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.ApiuserVdomGetArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.ApiuserVdomGetArgs>());
            set => _vdoms = value;
        }

        public ApiuserState()
        {
        }
        public static new ApiuserState Empty => new ApiuserState();
    }
}