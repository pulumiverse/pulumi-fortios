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

    public sealed class VirtualwanlinkServiceSrc6Args : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VirtualwanlinkServiceSrc6Args()
        {
        }
        public static new VirtualwanlinkServiceSrc6Args Empty => new VirtualwanlinkServiceSrc6Args();
    }
}
