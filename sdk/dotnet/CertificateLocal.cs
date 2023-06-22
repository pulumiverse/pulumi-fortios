// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Local keys and certificates.
    /// 
    /// By design considerations, the feature is using the fortios.JsonGenericApi resource as documented below.
    /// 
    /// ## Example
    /// 
    /// ### Delete Certificate:
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname1 = new Fortios.SystemAutoscript("trname1", new()
    ///     {
    ///         Interval = 1,
    ///         OutputSize = 10,
    ///         Repeat = 1,
    ///         Script = @"config vpn certificate local
    /// delete testcer
    /// end
    /// 
    /// ",
    ///         Start = "auto",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/certificateLocal:CertificateLocal")]
    public partial class CertificateLocal : global::Pulumi.CustomResource
    {
        [Output("acmeCaUrl")]
        public Output<string> AcmeCaUrl { get; private set; } = null!;

        [Output("acmeDomain")]
        public Output<string> AcmeDomain { get; private set; } = null!;

        [Output("acmeEmail")]
        public Output<string> AcmeEmail { get; private set; } = null!;

        [Output("acmeRenewWindow")]
        public Output<int> AcmeRenewWindow { get; private set; } = null!;

        [Output("acmeRsaKeySize")]
        public Output<int> AcmeRsaKeySize { get; private set; } = null!;

        [Output("autoRegenerateDays")]
        public Output<int> AutoRegenerateDays { get; private set; } = null!;

        [Output("autoRegenerateDaysWarning")]
        public Output<int> AutoRegenerateDaysWarning { get; private set; } = null!;

        [Output("caIdentifier")]
        public Output<string> CaIdentifier { get; private set; } = null!;

        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        [Output("cmpPath")]
        public Output<string> CmpPath { get; private set; } = null!;

        [Output("cmpRegenerationMethod")]
        public Output<string> CmpRegenerationMethod { get; private set; } = null!;

        [Output("cmpServer")]
        public Output<string> CmpServer { get; private set; } = null!;

        [Output("cmpServerCert")]
        public Output<string> CmpServerCert { get; private set; } = null!;

        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

        [Output("enrollProtocol")]
        public Output<string> EnrollProtocol { get; private set; } = null!;

        [Output("ikeLocalid")]
        public Output<string> IkeLocalid { get; private set; } = null!;

        [Output("ikeLocalidType")]
        public Output<string> IkeLocalidType { get; private set; } = null!;

        [Output("lastUpdated")]
        public Output<int> LastUpdated { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nameEncoding")]
        public Output<string> NameEncoding { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        [Output("privateKeyRetain")]
        public Output<string> PrivateKeyRetain { get; private set; } = null!;

        [Output("range")]
        public Output<string> Range { get; private set; } = null!;

        [Output("scepPassword")]
        public Output<string?> ScepPassword { get; private set; } = null!;

        [Output("scepUrl")]
        public Output<string> ScepUrl { get; private set; } = null!;

        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateLocal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateLocal(string name, CertificateLocalArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/certificateLocal:CertificateLocal", name, args ?? new CertificateLocalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateLocal(string name, Input<string> id, CertificateLocalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/certificateLocal:CertificateLocal", name, state, MakeResourceOptions(options, id))
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
                    "password",
                    "privateKey",
                    "scepPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateLocal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateLocal Get(string name, Input<string> id, CertificateLocalState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateLocal(name, id, state, options);
        }
    }

    public sealed class CertificateLocalArgs : global::Pulumi.ResourceArgs
    {
        [Input("acmeCaUrl")]
        public Input<string>? AcmeCaUrl { get; set; }

        [Input("acmeDomain")]
        public Input<string>? AcmeDomain { get; set; }

        [Input("acmeEmail")]
        public Input<string>? AcmeEmail { get; set; }

        [Input("acmeRenewWindow")]
        public Input<int>? AcmeRenewWindow { get; set; }

        [Input("acmeRsaKeySize")]
        public Input<int>? AcmeRsaKeySize { get; set; }

        [Input("autoRegenerateDays")]
        public Input<int>? AutoRegenerateDays { get; set; }

        [Input("autoRegenerateDaysWarning")]
        public Input<int>? AutoRegenerateDaysWarning { get; set; }

        [Input("caIdentifier")]
        public Input<string>? CaIdentifier { get; set; }

        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("cmpPath")]
        public Input<string>? CmpPath { get; set; }

        [Input("cmpRegenerationMethod")]
        public Input<string>? CmpRegenerationMethod { get; set; }

        [Input("cmpServer")]
        public Input<string>? CmpServer { get; set; }

        [Input("cmpServerCert")]
        public Input<string>? CmpServerCert { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("csr")]
        public Input<string>? Csr { get; set; }

        [Input("enrollProtocol")]
        public Input<string>? EnrollProtocol { get; set; }

        [Input("ikeLocalid")]
        public Input<string>? IkeLocalid { get; set; }

        [Input("ikeLocalidType")]
        public Input<string>? IkeLocalidType { get; set; }

        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameEncoding")]
        public Input<string>? NameEncoding { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyRetain")]
        public Input<string>? PrivateKeyRetain { get; set; }

        [Input("range")]
        public Input<string>? Range { get; set; }

        [Input("scepPassword")]
        private Input<string>? _scepPassword;
        public Input<string>? ScepPassword
        {
            get => _scepPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _scepPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CertificateLocalArgs()
        {
        }
        public static new CertificateLocalArgs Empty => new CertificateLocalArgs();
    }

    public sealed class CertificateLocalState : global::Pulumi.ResourceArgs
    {
        [Input("acmeCaUrl")]
        public Input<string>? AcmeCaUrl { get; set; }

        [Input("acmeDomain")]
        public Input<string>? AcmeDomain { get; set; }

        [Input("acmeEmail")]
        public Input<string>? AcmeEmail { get; set; }

        [Input("acmeRenewWindow")]
        public Input<int>? AcmeRenewWindow { get; set; }

        [Input("acmeRsaKeySize")]
        public Input<int>? AcmeRsaKeySize { get; set; }

        [Input("autoRegenerateDays")]
        public Input<int>? AutoRegenerateDays { get; set; }

        [Input("autoRegenerateDaysWarning")]
        public Input<int>? AutoRegenerateDaysWarning { get; set; }

        [Input("caIdentifier")]
        public Input<string>? CaIdentifier { get; set; }

        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("cmpPath")]
        public Input<string>? CmpPath { get; set; }

        [Input("cmpRegenerationMethod")]
        public Input<string>? CmpRegenerationMethod { get; set; }

        [Input("cmpServer")]
        public Input<string>? CmpServer { get; set; }

        [Input("cmpServerCert")]
        public Input<string>? CmpServerCert { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("csr")]
        public Input<string>? Csr { get; set; }

        [Input("enrollProtocol")]
        public Input<string>? EnrollProtocol { get; set; }

        [Input("ikeLocalid")]
        public Input<string>? IkeLocalid { get; set; }

        [Input("ikeLocalidType")]
        public Input<string>? IkeLocalidType { get; set; }

        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameEncoding")]
        public Input<string>? NameEncoding { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyRetain")]
        public Input<string>? PrivateKeyRetain { get; set; }

        [Input("range")]
        public Input<string>? Range { get; set; }

        [Input("scepPassword")]
        private Input<string>? _scepPassword;
        public Input<string>? ScepPassword
        {
            get => _scepPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _scepPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CertificateLocalState()
        {
        }
        public static new CertificateLocalState Empty => new CertificateLocalState();
    }
}
