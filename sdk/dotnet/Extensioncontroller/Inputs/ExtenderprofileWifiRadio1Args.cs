// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller.Inputs
{

    public sealed class ExtenderprofileWifiRadio1Args : global::Pulumi.ResourceArgs
    {
        [Input("band")]
        public Input<string>? Band { get; set; }

        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("beaconInterval")]
        public Input<int>? BeaconInterval { get; set; }

        [Input("bssColor")]
        public Input<int>? BssColor { get; set; }

        [Input("bssColorMode")]
        public Input<string>? BssColorMode { get; set; }

        [Input("channel")]
        public Input<string>? Channel { get; set; }

        [Input("extensionChannel")]
        public Input<string>? ExtensionChannel { get; set; }

        [Input("guardInterval")]
        public Input<string>? GuardInterval { get; set; }

        [Input("lanExtVap")]
        public Input<string>? LanExtVap { get; set; }

        [Input("localVaps")]
        private InputList<Inputs.ExtenderprofileWifiRadio1LocalVapArgs>? _localVaps;
        public InputList<Inputs.ExtenderprofileWifiRadio1LocalVapArgs> LocalVaps
        {
            get => _localVaps ?? (_localVaps = new InputList<Inputs.ExtenderprofileWifiRadio1LocalVapArgs>());
            set => _localVaps = value;
        }

        [Input("maxClients")]
        public Input<int>? MaxClients { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("n80211d")]
        public Input<string>? N80211d { get; set; }

        [Input("operatingStandard")]
        public Input<string>? OperatingStandard { get; set; }

        [Input("powerLevel")]
        public Input<int>? PowerLevel { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExtenderprofileWifiRadio1Args()
        {
        }
        public static new ExtenderprofileWifiRadio1Args Empty => new ExtenderprofileWifiRadio1Args();
    }
}
