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

    public sealed class FirewallCentralsnatmapDstAddr6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address name.
        /// 
        /// The `orig_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `dst_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `nat_ippool6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallCentralsnatmapDstAddr6GetArgs()
        {
        }
        public static new FirewallCentralsnatmapDstAddr6GetArgs Empty => new FirewallCentralsnatmapDstAddr6GetArgs();
    }
}
