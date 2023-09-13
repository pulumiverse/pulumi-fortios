// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpncertificate
{
    /// <summary>
    /// OCSP server configuration.
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
    ///     var trname = new Fortios.Vpncertificate.Ocspserver("trname", new()
    ///     {
    ///         Cert = "ACCVRAIZ1",
    ///         SourceIp = "0.0.0.0",
    ///         UnavailAction = "revoke",
    ///         Url = "www.tetserv.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpnCertificate OcspServer can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:vpncertificate/ocspserver:Ocspserver labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:vpncertificate/ocspserver:Ocspserver labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpncertificate/ocspserver:Ocspserver")]
    public partial class Ocspserver : global::Pulumi.CustomResource
    {
        /// <summary>
        /// OCSP server certificate.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// OCSP server entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Secondary OCSP server certificate.
        /// </summary>
        [Output("secondaryCert")]
        public Output<string> SecondaryCert { get; private set; } = null!;

        /// <summary>
        /// Secondary OCSP server URL.
        /// </summary>
        [Output("secondaryUrl")]
        public Output<string> SecondaryUrl { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communications to the OCSP server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        /// </summary>
        [Output("unavailAction")]
        public Output<string> UnavailAction { get; private set; } = null!;

        /// <summary>
        /// OCSP server URL.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ocspserver resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ocspserver(string name, OcspserverArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:vpncertificate/ocspserver:Ocspserver", name, args ?? new OcspserverArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ocspserver(string name, Input<string> id, OcspserverState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpncertificate/ocspserver:Ocspserver", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ocspserver resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ocspserver Get(string name, Input<string> id, OcspserverState? state = null, CustomResourceOptions? options = null)
        {
            return new Ocspserver(name, id, state, options);
        }
    }

    public sealed class OcspserverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OCSP server certificate.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// OCSP server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Secondary OCSP server certificate.
        /// </summary>
        [Input("secondaryCert")]
        public Input<string>? SecondaryCert { get; set; }

        /// <summary>
        /// Secondary OCSP server URL.
        /// </summary>
        [Input("secondaryUrl")]
        public Input<string>? SecondaryUrl { get; set; }

        /// <summary>
        /// Source IP address for communications to the OCSP server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        /// </summary>
        [Input("unavailAction")]
        public Input<string>? UnavailAction { get; set; }

        /// <summary>
        /// OCSP server URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OcspserverArgs()
        {
        }
        public static new OcspserverArgs Empty => new OcspserverArgs();
    }

    public sealed class OcspserverState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OCSP server certificate.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// OCSP server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Secondary OCSP server certificate.
        /// </summary>
        [Input("secondaryCert")]
        public Input<string>? SecondaryCert { get; set; }

        /// <summary>
        /// Secondary OCSP server URL.
        /// </summary>
        [Input("secondaryUrl")]
        public Input<string>? SecondaryUrl { get; set; }

        /// <summary>
        /// Source IP address for communications to the OCSP server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
        /// </summary>
        [Input("unavailAction")]
        public Input<string>? UnavailAction { get; set; }

        /// <summary>
        /// OCSP server URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OcspserverState()
        {
        }
        public static new OcspserverState Empty => new OcspserverState();
    }
}
