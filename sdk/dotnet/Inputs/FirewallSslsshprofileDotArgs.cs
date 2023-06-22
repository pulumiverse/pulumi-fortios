// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class FirewallSslsshprofileDotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("certValidationFailure")]
        public Input<string>? CertValidationFailure { get; set; }

        /// <summary>
        /// Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("certValidationTimeout")]
        public Input<string>? CertValidationTimeout { get; set; }

        /// <summary>
        /// Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("expiredServerCert")]
        public Input<string>? ExpiredServerCert { get; set; }

        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyAfterTcpHandshake")]
        public Input<string>? ProxyAfterTcpHandshake { get; set; }

        /// <summary>
        /// Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("revokedServerCert")]
        public Input<string>? RevokedServerCert { get; set; }

        /// <summary>
        /// Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.
        /// </summary>
        [Input("sniServerCertCheck")]
        public Input<string>? SniServerCertCheck { get; set; }

        /// <summary>
        /// Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
        /// </summary>
        [Input("unsupportedSslCipher")]
        public Input<string>? UnsupportedSslCipher { get; set; }

        /// <summary>
        /// Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
        /// </summary>
        [Input("unsupportedSslNegotiation")]
        public Input<string>? UnsupportedSslNegotiation { get; set; }

        /// <summary>
        /// Action based on the SSL version used being unsupported.
        /// </summary>
        [Input("unsupportedSslVersion")]
        public Input<string>? UnsupportedSslVersion { get; set; }

        /// <summary>
        /// Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("untrustedServerCert")]
        public Input<string>? UntrustedServerCert { get; set; }

        public FirewallSslsshprofileDotArgs()
        {
        }
        public static new FirewallSslsshprofileDotArgs Empty => new FirewallSslsshprofileDotArgs();
    }
}
