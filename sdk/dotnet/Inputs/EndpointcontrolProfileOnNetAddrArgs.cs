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

    public sealed class EndpointcontrolProfileOnNetAddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address object from available options.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EndpointcontrolProfileOnNetAddrArgs()
        {
        }
        public static new EndpointcontrolProfileOnNetAddrArgs Empty => new EndpointcontrolProfileOnNetAddrArgs();
    }
}
