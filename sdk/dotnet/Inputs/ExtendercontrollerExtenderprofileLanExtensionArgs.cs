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

    public sealed class ExtendercontrollerExtenderprofileLanExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPsec phase1 interface.
        /// </summary>
        [Input("backhaulInterface")]
        public Input<string>? BackhaulInterface { get; set; }

        /// <summary>
        /// IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
        /// </summary>
        [Input("backhaulIp")]
        public Input<string>? BackhaulIp { get; set; }

        [Input("backhauls")]
        private InputList<Inputs.ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs>? _backhauls;

        /// <summary>
        /// LAN extension backhaul tunnel configuration. The structure of `backhaul` block is documented below.
        /// </summary>
        public InputList<Inputs.ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs> Backhauls
        {
            get => _backhauls ?? (_backhauls = new InputList<Inputs.ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs>());
            set => _backhauls = value;
        }

        /// <summary>
        /// IPsec tunnel name.
        /// </summary>
        [Input("ipsecTunnel")]
        public Input<string>? IpsecTunnel { get; set; }

        /// <summary>
        /// LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.
        /// </summary>
        [Input("linkLoadbalance")]
        public Input<string>? LinkLoadbalance { get; set; }

        public ExtendercontrollerExtenderprofileLanExtensionArgs()
        {
        }
        public static new ExtendercontrollerExtenderprofileLanExtensionArgs Empty => new ExtendercontrollerExtenderprofileLanExtensionArgs();
    }
}
