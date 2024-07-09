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

    public sealed class BgpNeighborRange6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// an identifier for the resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("maxNeighborNum")]
        public Input<int>? MaxNeighborNum { get; set; }

        /// <summary>
        /// BGP neighbor group table. The structure of `neighbor_group` block is documented below.
        /// </summary>
        [Input("neighborGroup")]
        public Input<string>? NeighborGroup { get; set; }

        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public BgpNeighborRange6Args()
        {
        }
        public static new BgpNeighborRange6Args Empty => new BgpNeighborRange6Args();
    }
}
