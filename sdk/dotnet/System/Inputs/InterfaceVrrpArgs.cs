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

    public sealed class InterfaceVrrpArgs : global::Pulumi.ResourceArgs
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
        /// Enable/disable ignoring of default route when checking destination. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ignoreDefaultRoute")]
        public Input<string>? IgnoreDefaultRoute { get; set; }

        /// <summary>
        /// Enable/disable preempt mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("preempt")]
        public Input<string>? Preempt { get; set; }

        /// <summary>
        /// Priority of the virtual router (1 - 255).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("proxyArps")]
        private InputList<Inputs.InterfaceVrrpProxyArpArgs>? _proxyArps;

        /// <summary>
        /// VRRP Proxy ARP configuration. The structure of `proxy_arp` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfaceVrrpProxyArpArgs> ProxyArps
        {
            get => _proxyArps ?? (_proxyArps = new InputList<Inputs.InterfaceVrrpProxyArpArgs>());
            set => _proxyArps = value;
        }

        /// <summary>
        /// Startup time (1 - 255 seconds).
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// Enable/disable this VRRP configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// VRRP version. Valid values: `2`, `3`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Monitor the route to this destination.
        /// </summary>
        [Input("vrdst")]
        public Input<string>? Vrdst { get; set; }

        /// <summary>
        /// Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
        /// </summary>
        [Input("vrdstPriority")]
        public Input<int>? VrdstPriority { get; set; }

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
        /// IP address of the virtual router.
        /// </summary>
        [Input("vrip")]
        public Input<string>? Vrip { get; set; }

        public InterfaceVrrpArgs()
        {
        }
        public static new InterfaceVrrpArgs Empty => new InterfaceVrrpArgs();
    }
}
