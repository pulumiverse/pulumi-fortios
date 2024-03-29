// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class AddressMacaddrGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// MAC address ranges &lt;start&gt;[-&lt;end&gt;] separated by space.
        /// </summary>
        [Input("macaddr")]
        public Input<string>? Macaddr { get; set; }

        public AddressMacaddrGetArgs()
        {
        }
        public static new AddressMacaddrGetArgs Empty => new AddressMacaddrGetArgs();
    }
}
