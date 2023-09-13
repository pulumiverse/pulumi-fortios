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

    public sealed class Vip6SslCipherSuiteGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cipher suite name.
        /// </summary>
        [Input("cipher")]
        public Input<string>? Cipher { get; set; }

        /// <summary>
        /// SSL/TLS cipher suites priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// SSL/TLS versions that the cipher suite can be used with.
        /// </summary>
        [Input("versions")]
        public Input<string>? Versions { get; set; }

        public Vip6SslCipherSuiteGetArgs()
        {
        }
        public static new Vip6SslCipherSuiteGetArgs Empty => new Vip6SslCipherSuiteGetArgs();
    }
}
