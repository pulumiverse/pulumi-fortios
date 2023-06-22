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

    public sealed class FirewallSslsshprofileSslExemptGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address object.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// IPv6 address object.
        /// </summary>
        [Input("address6")]
        public Input<string>? Address6 { get; set; }

        /// <summary>
        /// FortiGuard category ID.
        /// </summary>
        [Input("fortiguardCategory")]
        public Input<int>? FortiguardCategory { get; set; }

        /// <summary>
        /// ID number.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Exempt servers by regular expression.
        /// </summary>
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        /// <summary>
        /// Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category`, `address`, `address6`, `wildcard-fqdn`, `regex`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Exempt servers by wildcard FQDN.
        /// </summary>
        [Input("wildcardFqdn")]
        public Input<string>? WildcardFqdn { get; set; }

        public FirewallSslsshprofileSslExemptGetArgs()
        {
        }
        public static new FirewallSslsshprofileSslExemptGetArgs Empty => new FirewallSslsshprofileSslExemptGetArgs();
    }
}
