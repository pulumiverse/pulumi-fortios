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

    public sealed class RouterMulticast6InterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time before old neighbour information expires (1 - 65535 sec, default = 105).
        /// </summary>
        [Input("helloHoldtime")]
        public Input<int>? HelloHoldtime { get; set; }

        /// <summary>
        /// Interval between sending PIM hello messages  (1 - 65535 sec, default = 30)..
        /// </summary>
        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RouterMulticast6InterfaceArgs()
        {
        }
        public static new RouterMulticast6InterfaceArgs Empty => new RouterMulticast6InterfaceArgs();
    }
}
