// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class InterfaceIpv6Ip6ExtraAddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv6 prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public InterfaceIpv6Ip6ExtraAddrArgs()
        {
        }
        public static new InterfaceIpv6Ip6ExtraAddrArgs Empty => new InterfaceIpv6Ip6ExtraAddrArgs();
    }
}
