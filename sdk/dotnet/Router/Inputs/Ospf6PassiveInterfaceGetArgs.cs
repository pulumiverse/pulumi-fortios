// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class Ospf6PassiveInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Passive interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Ospf6PassiveInterfaceGetArgs()
        {
        }
        public static new Ospf6PassiveInterfaceGetArgs Empty => new Ospf6PassiveInterfaceGetArgs();
    }
}
