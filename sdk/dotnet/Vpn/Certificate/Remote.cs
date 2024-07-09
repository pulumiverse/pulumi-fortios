// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Certificate
{
    /// <summary>
    /// Remote certificate as a PEM file.
    /// 
    /// ## Import
    /// 
    /// VpnCertificate Remote can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/remote:Remote labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/certificate/remote:Remote labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpn/certificate/remote:Remote")]
    public partial class Remote : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Output("range")]
        public Output<string> Range { get; private set; } = null!;

        /// <summary>
        /// Remote certificate.
        /// </summary>
        [Output("remote")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// Remote certificate source type.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Remote resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Remote(string name, RemoteArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/remote:Remote", name, args ?? new RemoteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Remote(string name, Input<string> id, RemoteState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/certificate/remote:Remote", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Remote resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Remote Get(string name, Input<string> id, RemoteState? state = null, CustomResourceOptions? options = null)
        {
            return new Remote(name, id, state, options);
        }
    }

    public sealed class RemoteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        /// <summary>
        /// Remote certificate.
        /// </summary>
        [Input("remote")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Remote certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RemoteArgs()
        {
        }
        public static new RemoteArgs Empty => new RemoteArgs();
    }

    public sealed class RemoteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        /// <summary>
        /// Remote certificate.
        /// </summary>
        [Input("remote")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Remote certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RemoteState()
        {
        }
        public static new RemoteState Empty => new RemoteState();
    }
}
