// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Consolidated.Inputs
{

    public sealed class PolicyInternetServiceIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public PolicyInternetServiceIdArgs()
        {
        }
        public static new PolicyInternetServiceIdArgs Empty => new PolicyInternetServiceIdArgs();
    }
}
