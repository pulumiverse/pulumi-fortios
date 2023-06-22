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

    public sealed class FirewallSslsshprofileSslGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action based on certificate probe failure. Valid values: `allow`, `block`.
        /// </summary>
        [Input("certProbeFailure")]
        public Input<string>? CertProbeFailure { get; set; }

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
        /// Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        [Input("clientCertRequest")]
        public Input<string>? ClientCertRequest { get; set; }

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
        /// Level of SSL inspection. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.
        /// </summary>
        [Input("inspectAll")]
        public Input<string>? InspectAll { get; set; }

        /// <summary>
        /// Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
        /// </summary>
        [Input("invalidServerCert")]
        public Input<string>? InvalidServerCert { get; set; }

        /// <summary>
        /// Minimum SSL version to be allowed. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
        /// </summary>
        [Input("minAllowedSslVersion")]
        public Input<string>? MinAllowedSslVersion { get; set; }

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
        /// Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        [Input("unsupportedSsl")]
        public Input<string>? UnsupportedSsl { get; set; }

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
        /// Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        [Input("untrustedServerCert")]
        public Input<string>? UntrustedServerCert { get; set; }

        public FirewallSslsshprofileSslGetArgs()
        {
        }
        public static new FirewallSslsshprofileSslGetArgs Empty => new FirewallSslsshprofileSslGetArgs();
    }
}
