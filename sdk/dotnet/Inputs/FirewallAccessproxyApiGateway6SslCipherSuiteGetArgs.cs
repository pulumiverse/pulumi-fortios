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

    public sealed class FirewallAccessproxyApiGateway6SslCipherSuiteGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`.
        /// </summary>
        [Input("cipher")]
        public Input<string>? Cipher { get; set; }

        /// <summary>
        /// SSL/TLS cipher suites priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
        /// </summary>
        [Input("versions")]
        public Input<string>? Versions { get; set; }

        public FirewallAccessproxyApiGateway6SslCipherSuiteGetArgs()
        {
        }
        public static new FirewallAccessproxyApiGateway6SslCipherSuiteGetArgs Empty => new FirewallAccessproxyApiGateway6SslCipherSuiteGetArgs();
    }
}
