// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure FSSO groups.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname1 = new Fortios.User.Fsso("trname1", new()
    ///     {
    ///         Port = 32381,
    ///         Port2 = 8000,
    ///         Port3 = 8000,
    ///         Port4 = 8000,
    ///         Port5 = 8000,
    ///         Server = "1.1.1.1",
    ///         SourceIp = "0.0.0.0",
    ///         SourceIp6 = "::",
    ///     });
    /// 
    ///     var trname = new Fortios.User.Adgrp("trname", new()
    ///     {
    ///         ServerName = trname1.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User Adgrp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/adgrp:Adgrp")]
    public partial class Adgrp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// FSSO connector source.
        /// </summary>
        [Output("connectorSource")]
        public Output<string> ConnectorSource { get; private set; } = null!;

        /// <summary>
        /// Group ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// FSSO agent name.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Adgrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Adgrp(string name, AdgrpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/adgrp:Adgrp", name, args ?? new AdgrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Adgrp(string name, Input<string> id, AdgrpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/adgrp:Adgrp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Adgrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Adgrp Get(string name, Input<string> id, AdgrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Adgrp(name, id, state, options);
        }
    }

    public sealed class AdgrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FSSO connector source.
        /// </summary>
        [Input("connectorSource")]
        public Input<string>? ConnectorSource { get; set; }

        /// <summary>
        /// Group ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FSSO agent name.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AdgrpArgs()
        {
        }
        public static new AdgrpArgs Empty => new AdgrpArgs();
    }

    public sealed class AdgrpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FSSO connector source.
        /// </summary>
        [Input("connectorSource")]
        public Input<string>? ConnectorSource { get; set; }

        /// <summary>
        /// Group ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FSSO agent name.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AdgrpState()
        {
        }
        public static new AdgrpState Empty => new AdgrpState();
    }
}
