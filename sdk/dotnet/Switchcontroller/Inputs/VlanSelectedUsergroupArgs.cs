// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class VlanSelectedUsergroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VlanSelectedUsergroupArgs()
        {
        }
        public static new VlanSelectedUsergroupArgs Empty => new VlanSelectedUsergroupArgs();
    }
}
