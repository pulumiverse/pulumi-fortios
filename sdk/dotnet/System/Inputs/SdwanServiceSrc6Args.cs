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

    public sealed class SdwanServiceSrc6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address or address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanServiceSrc6Args()
        {
        }
        public static new SdwanServiceSrc6Args Empty => new SdwanServiceSrc6Args();
    }
}
