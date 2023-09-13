// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class SdwanDuplicationDstaddr6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Physical interface name.
        /// 
        /// 
        /// 
        /// 
        /// 
        /// The `dst6` block supports:
        /// 
        /// 
        /// The `src6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// The `srcaddr6` block supports:
        /// 
        /// 
        /// The `dstaddr6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanDuplicationDstaddr6GetArgs()
        {
        }
        public static new SdwanDuplicationDstaddr6GetArgs Empty => new SdwanDuplicationDstaddr6GetArgs();
    }
}
