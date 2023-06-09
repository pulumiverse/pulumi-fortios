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

    public sealed class UserExchangeKdcIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// KDC IPv4 addresses for Kerberos authentication.
        /// </summary>
        [Input("ipv4")]
        public Input<string>? Ipv4 { get; set; }

        public UserExchangeKdcIpArgs()
        {
        }
        public static new UserExchangeKdcIpArgs Empty => new UserExchangeKdcIpArgs();
    }
}
