// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnsslweb.Inputs
{

    public sealed class PortalMacAddrCheckRuleMacAddrListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client MAC address.
        /// </summary>
        [Input("addr")]
        public Input<string>? Addr { get; set; }

        public PortalMacAddrCheckRuleMacAddrListGetArgs()
        {
        }
        public static new PortalMacAddrCheckRuleMacAddrListGetArgs Empty => new PortalMacAddrCheckRuleMacAddrListGetArgs();
    }
}
