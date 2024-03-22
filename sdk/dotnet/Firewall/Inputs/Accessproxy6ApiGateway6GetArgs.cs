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

    public sealed class Accessproxy6ApiGateway6GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.Accessproxy6ApiGateway6ApplicationGetArgs>? _applications;

        /// <summary>
        /// SaaS application controlled by this Access Proxy. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.Accessproxy6ApiGateway6ApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.Accessproxy6ApiGateway6ApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// HTTP2 support, default=Enable. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("h2Support")]
        public Input<string>? H2Support { get; set; }

        /// <summary>
        /// HTTP3/QUIC support, default=Disable. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("h3Support")]
        public Input<string>? H3Support { get; set; }

        /// <summary>
        /// Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
        /// </summary>
        [Input("httpCookieAge")]
        public Input<int>? HttpCookieAge { get; set; }

        /// <summary>
        /// Domain that HTTP cookie persistence should apply to.
        /// </summary>
        [Input("httpCookieDomain")]
        public Input<string>? HttpCookieDomain { get; set; }

        /// <summary>
        /// Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("httpCookieDomainFromHost")]
        public Input<string>? HttpCookieDomainFromHost { get; set; }

        /// <summary>
        /// Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
        /// </summary>
        [Input("httpCookieGeneration")]
        public Input<int>? HttpCookieGeneration { get; set; }

        /// <summary>
        /// Limit HTTP cookie persistence to the specified path.
        /// </summary>
        [Input("httpCookiePath")]
        public Input<string>? HttpCookiePath { get; set; }

        /// <summary>
        /// Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.
        /// </summary>
        [Input("httpCookieShare")]
        public Input<string>? HttpCookieShare { get; set; }

        /// <summary>
        /// Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("httpsCookieSecure")]
        public Input<string>? HttpsCookieSecure { get; set; }

        /// <summary>
        /// API Gateway ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `first-alive`, `http-host`.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`.
        /// </summary>
        [Input("persistence")]
        public Input<string>? Persistence { get; set; }

        /// <summary>
        /// QUIC setting. The structure of `quic` block is documented below.
        /// </summary>
        [Input("quic")]
        public Input<Inputs.Accessproxy6ApiGateway6QuicGetArgs>? Quic { get; set; }

        [Input("realservers")]
        private InputList<Inputs.Accessproxy6ApiGateway6RealserverGetArgs>? _realservers;

        /// <summary>
        /// Select the real servers that this Access Proxy will distribute traffic to. The structure of `realservers` block is documented below.
        /// </summary>
        public InputList<Inputs.Accessproxy6ApiGateway6RealserverGetArgs> Realservers
        {
            get => _realservers ?? (_realservers = new InputList<Inputs.Accessproxy6ApiGateway6RealserverGetArgs>());
            set => _realservers = value;
        }

        /// <summary>
        /// Enable/disable SAML redirection after successful authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("samlRedirect")]
        public Input<string>? SamlRedirect { get; set; }

        /// <summary>
        /// SAML service provider configuration for VIP authentication.
        /// </summary>
        [Input("samlServer")]
        public Input<string>? SamlServer { get; set; }

        /// <summary>
        /// Service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("sslAlgorithm")]
        public Input<string>? SslAlgorithm { get; set; }

        [Input("sslCipherSuites")]
        private InputList<Inputs.Accessproxy6ApiGateway6SslCipherSuiteGetArgs>? _sslCipherSuites;

        /// <summary>
        /// SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.
        /// </summary>
        public InputList<Inputs.Accessproxy6ApiGateway6SslCipherSuiteGetArgs> SslCipherSuites
        {
            get => _sslCipherSuites ?? (_sslCipherSuites = new InputList<Inputs.Accessproxy6ApiGateway6SslCipherSuiteGetArgs>());
            set => _sslCipherSuites = value;
        }

        /// <summary>
        /// Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.
        /// </summary>
        [Input("sslDhBits")]
        public Input<string>? SslDhBits { get; set; }

        /// <summary>
        /// Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
        /// </summary>
        [Input("sslMaxVersion")]
        public Input<string>? SslMaxVersion { get; set; }

        /// <summary>
        /// Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
        /// </summary>
        [Input("sslMinVersion")]
        public Input<string>? SslMinVersion { get; set; }

        /// <summary>
        /// Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sslRenegotiation")]
        public Input<string>? SslRenegotiation { get; set; }

        /// <summary>
        /// SSL-VPN web portal.
        /// </summary>
        [Input("sslVpnWebPortal")]
        public Input<string>? SslVpnWebPortal { get; set; }

        /// <summary>
        /// URL pattern to match.
        /// </summary>
        [Input("urlMap")]
        public Input<string>? UrlMap { get; set; }

        /// <summary>
        /// Type of url-map. Valid values: `sub-string`, `wildcard`, `regex`.
        /// </summary>
        [Input("urlMapType")]
        public Input<string>? UrlMapType { get; set; }

        /// <summary>
        /// Virtual host.
        /// </summary>
        [Input("virtualHost")]
        public Input<string>? VirtualHost { get; set; }

        public Accessproxy6ApiGateway6GetArgs()
        {
        }
        public static new Accessproxy6ApiGateway6GetArgs Empty => new Accessproxy6ApiGateway6GetArgs();
    }
}
