// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20.Inputs
{

    public sealed class Anqp3gppcellularMccMncListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Mobile country code.
        /// </summary>
        [Input("mcc")]
        public Input<string>? Mcc { get; set; }

        /// <summary>
        /// Mobile network code.
        /// </summary>
        [Input("mnc")]
        public Input<string>? Mnc { get; set; }

        public Anqp3gppcellularMccMncListGetArgs()
        {
        }
        public static new Anqp3gppcellularMccMncListGetArgs Empty => new Anqp3gppcellularMccMncListGetArgs();
    }
}
