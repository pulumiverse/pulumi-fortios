// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Inputs
{

    public sealed class SettingsTunnelIpv6PoolGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SettingsTunnelIpv6PoolGetArgs()
        {
        }
        public static new SettingsTunnelIpv6PoolGetArgs Empty => new SettingsTunnelIpv6PoolGetArgs();
    }
}
