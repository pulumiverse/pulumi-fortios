// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class SdwanDuplicationSrcaddr6Args : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanDuplicationSrcaddr6Args()
        {
        }
        public static new SdwanDuplicationSrcaddr6Args Empty => new SdwanDuplicationSrcaddr6Args();
    }
}
