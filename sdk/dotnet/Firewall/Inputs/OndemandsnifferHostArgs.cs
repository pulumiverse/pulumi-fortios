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

    public sealed class OndemandsnifferHostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 or IPv6 host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        public OndemandsnifferHostArgs()
        {
        }
        public static new OndemandsnifferHostArgs Empty => new OndemandsnifferHostArgs();
    }
}
