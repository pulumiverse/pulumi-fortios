// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ipsec.Inputs
{

    public sealed class Phase1CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Phase1CertificateArgs()
        {
        }
        public static new Phase1CertificateArgs Empty => new Phase1CertificateArgs();
    }
}
