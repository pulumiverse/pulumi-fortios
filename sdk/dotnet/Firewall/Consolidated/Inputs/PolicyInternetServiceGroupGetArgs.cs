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

    public sealed class PolicyInternetServiceGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyInternetServiceGroupGetArgs()
        {
        }
        public static new PolicyInternetServiceGroupGetArgs Empty => new PolicyInternetServiceGroupGetArgs();
    }
}
