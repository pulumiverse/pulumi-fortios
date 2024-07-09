// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wanopt
{
    /// <summary>
    /// Designate cache-service for wan-optimization and webcache.
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
    ///     var trname = new Fortios.Wanopt.Cacheservice("trname", new()
    ///     {
    ///         AcceptableConnections = "any",
    ///         Collaboration = "disable",
    ///         DeviceId = "default_dev_id",
    ///         PreferScenario = "balance",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Wanopt CacheService can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/cacheservice:Cacheservice labelname WanoptCacheService
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/cacheservice:Cacheservice labelname WanoptCacheService
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wanopt/cacheservice:Cacheservice")]
    public partial class Cacheservice : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        /// </summary>
        [Output("acceptableConnections")]
        public Output<string> AcceptableConnections { get; private set; } = null!;

        /// <summary>
        /// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("collaboration")]
        public Output<string> Collaboration { get; private set; } = null!;

        /// <summary>
        /// Set identifier for this cache device.
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        /// </summary>
        [Output("dstPeers")]
        public Output<ImmutableArray<Outputs.CacheserviceDstPeer>> DstPeers { get; private set; } = null!;

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
        /// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        /// </summary>
        [Output("preferScenario")]
        public Output<string> PreferScenario { get; private set; } = null!;

        /// <summary>
        /// Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        /// </summary>
        [Output("srcPeers")]
        public Output<ImmutableArray<Outputs.CacheserviceSrcPeer>> SrcPeers { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Cacheservice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cacheservice(string name, CacheserviceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/cacheservice:Cacheservice", name, args ?? new CacheserviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cacheservice(string name, Input<string> id, CacheserviceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/cacheservice:Cacheservice", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cacheservice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cacheservice Get(string name, Input<string> id, CacheserviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Cacheservice(name, id, state, options);
        }
    }

    public sealed class CacheserviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        /// </summary>
        [Input("acceptableConnections")]
        public Input<string>? AcceptableConnections { get; set; }

        /// <summary>
        /// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("collaboration")]
        public Input<string>? Collaboration { get; set; }

        /// <summary>
        /// Set identifier for this cache device.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        [Input("dstPeers")]
        private InputList<Inputs.CacheserviceDstPeerArgs>? _dstPeers;

        /// <summary>
        /// Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.CacheserviceDstPeerArgs> DstPeers
        {
            get => _dstPeers ?? (_dstPeers = new InputList<Inputs.CacheserviceDstPeerArgs>());
            set => _dstPeers = value;
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
        /// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        /// </summary>
        [Input("preferScenario")]
        public Input<string>? PreferScenario { get; set; }

        [Input("srcPeers")]
        private InputList<Inputs.CacheserviceSrcPeerArgs>? _srcPeers;

        /// <summary>
        /// Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.CacheserviceSrcPeerArgs> SrcPeers
        {
            get => _srcPeers ?? (_srcPeers = new InputList<Inputs.CacheserviceSrcPeerArgs>());
            set => _srcPeers = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CacheserviceArgs()
        {
        }
        public static new CacheserviceArgs Empty => new CacheserviceArgs();
    }

    public sealed class CacheserviceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
        /// </summary>
        [Input("acceptableConnections")]
        public Input<string>? AcceptableConnections { get; set; }

        /// <summary>
        /// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("collaboration")]
        public Input<string>? Collaboration { get; set; }

        /// <summary>
        /// Set identifier for this cache device.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        [Input("dstPeers")]
        private InputList<Inputs.CacheserviceDstPeerGetArgs>? _dstPeers;

        /// <summary>
        /// Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.CacheserviceDstPeerGetArgs> DstPeers
        {
            get => _dstPeers ?? (_dstPeers = new InputList<Inputs.CacheserviceDstPeerGetArgs>());
            set => _dstPeers = value;
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
        /// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
        /// </summary>
        [Input("preferScenario")]
        public Input<string>? PreferScenario { get; set; }

        [Input("srcPeers")]
        private InputList<Inputs.CacheserviceSrcPeerGetArgs>? _srcPeers;

        /// <summary>
        /// Modify cache-service source peer list. The structure of `src_peer` block is documented below.
        /// </summary>
        public InputList<Inputs.CacheserviceSrcPeerGetArgs> SrcPeers
        {
            get => _srcPeers ?? (_srcPeers = new InputList<Inputs.CacheserviceSrcPeerGetArgs>());
            set => _srcPeers = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CacheserviceState()
        {
        }
        public static new CacheserviceState Empty => new CacheserviceState();
    }
}
