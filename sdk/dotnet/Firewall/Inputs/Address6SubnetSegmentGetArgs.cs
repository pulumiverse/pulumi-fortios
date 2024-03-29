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

    public sealed class Address6SubnetSegmentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Subnet segment type. Valid values: `any`, `specific`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Subnet segment value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public Address6SubnetSegmentGetArgs()
        {
        }
        public static new Address6SubnetSegmentGetArgs Empty => new Address6SubnetSegmentGetArgs();
    }
}
