// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class VirtualwirepairMemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        public VirtualwirepairMemberGetArgs()
        {
        }
        public static new VirtualwirepairMemberGetArgs Empty => new VirtualwirepairMemberGetArgs();
    }
}
