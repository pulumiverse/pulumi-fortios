// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dpdk.Inputs
{

    public sealed class GlobalInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Physical interface name.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        public GlobalInterfaceArgs()
        {
        }
        public static new GlobalInterfaceArgs Empty => new GlobalInterfaceArgs();
    }
}
