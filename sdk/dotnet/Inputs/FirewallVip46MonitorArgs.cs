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

    public sealed class FirewallVip46MonitorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Health monitor name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallVip46MonitorArgs()
        {
        }
        public static new FirewallVip46MonitorArgs Empty => new FirewallVip46MonitorArgs();
    }
}