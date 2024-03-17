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

    public sealed class InterfaceIpv6Vrrp6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable accept mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("acceptMode")]
        public Input<string>? AcceptMode { get; set; }

        /// <summary>
        /// Advertisement interval (1 - 255 seconds).
        /// </summary>
        [Input("advInterval")]
        public Input<int>? AdvInterval { get; set; }

        /// <summary>
        /// Enable/disable preempt mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("preempt")]
        public Input<string>? Preempt { get; set; }

        /// <summary>
        /// Priority of learned routes.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Startup time (1 - 255 seconds).
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// Bring the interface up or shut the interface down. Valid values: `up`, `down`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Monitor the route to this destination.
        /// </summary>
        [Input("vrdst6")]
        public Input<string>? Vrdst6 { get; set; }

        /// <summary>
        /// VRRP group ID (1 - 65535).
        /// </summary>
        [Input("vrgrp")]
        public Input<int>? Vrgrp { get; set; }

        /// <summary>
        /// Virtual router identifier (1 - 255).
        /// </summary>
        [Input("vrid")]
        public Input<int>? Vrid { get; set; }

        /// <summary>
        /// IPv6 address of the virtual router.
        /// </summary>
        [Input("vrip6")]
        public Input<string>? Vrip6 { get; set; }

        public InterfaceIpv6Vrrp6Args()
        {
        }
        public static new InterfaceIpv6Vrrp6Args Empty => new InterfaceIpv6Vrrp6Args();
    }
}
