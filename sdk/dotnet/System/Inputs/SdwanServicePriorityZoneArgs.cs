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

    public sealed class SdwanServicePriorityZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Priority zone name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanServicePriorityZoneArgs()
        {
        }
        public static new SdwanServicePriorityZoneArgs Empty => new SdwanServicePriorityZoneArgs();
    }
}
