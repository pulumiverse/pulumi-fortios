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

    public sealed class Policy64PoolnameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Policy64PoolnameArgs()
        {
        }
        public static new Policy64PoolnameArgs Empty => new Policy64PoolnameArgs();
    }
}
