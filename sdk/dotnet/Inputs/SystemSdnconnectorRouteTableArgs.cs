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

    public sealed class SystemSdnconnectorRouteTableArgs : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.SystemSdnconnectorRouteTableRouteArgs>? _routes;

        /// <summary>
        /// Configure Azure route. The structure of `route` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdnconnectorRouteTableRouteArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.SystemSdnconnectorRouteTableRouteArgs>());
            set => _routes = value;
        }

        /// <summary>
        /// Subscription ID of Azure route table.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public SystemSdnconnectorRouteTableArgs()
        {
        }
        public static new SystemSdnconnectorRouteTableArgs Empty => new SystemSdnconnectorRouteTableArgs();
    }
}
