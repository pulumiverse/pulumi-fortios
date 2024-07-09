// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class SslsshprofilePop3sArgs : global::Pulumi.ResourceArgs
    {
        [Input("certValidationFailure")]
        public Input<string>? CertValidationFailure { get; set; }

        [Input("certValidationTimeout")]
        public Input<string>? CertValidationTimeout { get; set; }

        [Input("clientCertRequest")]
        public Input<string>? ClientCertRequest { get; set; }

        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        [Input("expiredServerCert")]
        public Input<string>? ExpiredServerCert { get; set; }

        [Input("invalidServerCert")]
        public Input<string>? InvalidServerCert { get; set; }

        [Input("ports")]
        public Input<string>? Ports { get; set; }

        [Input("proxyAfterTcpHandshake")]
        public Input<string>? ProxyAfterTcpHandshake { get; set; }

        [Input("revokedServerCert")]
        public Input<string>? RevokedServerCert { get; set; }

        [Input("sniServerCertCheck")]
        public Input<string>? SniServerCertCheck { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("unsupportedSsl")]
        public Input<string>? UnsupportedSsl { get; set; }

        [Input("unsupportedSslCipher")]
        public Input<string>? UnsupportedSslCipher { get; set; }

        [Input("unsupportedSslNegotiation")]
        public Input<string>? UnsupportedSslNegotiation { get; set; }

        [Input("unsupportedSslVersion")]
        public Input<string>? UnsupportedSslVersion { get; set; }

        [Input("untrustedServerCert")]
        public Input<string>? UntrustedServerCert { get; set; }

        public SslsshprofilePop3sArgs()
        {
        }
        public static new SslsshprofilePop3sArgs Empty => new SslsshprofilePop3sArgs();
    }
}
