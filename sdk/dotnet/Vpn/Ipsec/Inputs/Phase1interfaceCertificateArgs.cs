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

    public sealed class Phase1interfaceCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate name.
        /// 
        /// The `ipv4_exclude_range` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Phase1interfaceCertificateArgs()
        {
        }
        public static new Phase1interfaceCertificateArgs Empty => new Phase1interfaceCertificateArgs();
    }
}
