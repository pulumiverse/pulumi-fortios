// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Ospf.Inputs
{

    public sealed class OspfinterfaceMd5KeyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key ID (1 - 255).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Password for the key.
        /// </summary>
        [Input("keyString")]
        public Input<string>? KeyString { get; set; }

        public OspfinterfaceMd5KeyGetArgs()
        {
        }
        public static new OspfinterfaceMd5KeyGetArgs Empty => new OspfinterfaceMd5KeyGetArgs();
    }
}
