// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class SdnconnectorRouteTableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Route table name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource group of Azure route table.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        [Input("routes")]
        private InputList<Inputs.SdnconnectorRouteTableRouteGetArgs>? _routes;

        /// <summary>
        /// Configure Azure route. The structure of `route` block is documented below.
        /// </summary>
        public InputList<Inputs.SdnconnectorRouteTableRouteGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.SdnconnectorRouteTableRouteGetArgs>());
            set => _routes = value;
        }

        /// <summary>
        /// Subscription ID of Azure route table.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public SdnconnectorRouteTableGetArgs()
        {
        }
        public static new SdnconnectorRouteTableGetArgs Empty => new SdnconnectorRouteTableGetArgs();
    }
}
