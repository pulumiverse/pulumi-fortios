// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class LocationCoordinatesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// +/- Floating point no. eg. 117.47.
        /// </summary>
        [Input("altitude")]
        public Input<string>? Altitude { get; set; }

        /// <summary>
        /// Configure the unit for which the altitude is to (m = meters, f = floors of a building). Valid values: `m`, `f`.
        /// </summary>
        [Input("altitudeUnit")]
        public Input<string>? AltitudeUnit { get; set; }

        /// <summary>
        /// WGS84, NAD83, NAD83/MLLW. Valid values: `WGS84`, `NAD83`, `NAD83/MLLW`.
        /// </summary>
        [Input("datum")]
        public Input<string>? Datum { get; set; }

        /// <summary>
        /// Floating point starting with +/- or ending with (N or S). For example, +/-16.67 or 16.67N.
        /// </summary>
        [Input("latitude")]
        public Input<string>? Latitude { get; set; }

        /// <summary>
        /// Floating point starting with +/- or ending with (N or S). For example, +/-26.789 or 26.789E.
        /// </summary>
        [Input("longitude")]
        public Input<string>? Longitude { get; set; }

        /// <summary>
        /// Parent key name.
        /// </summary>
        [Input("parentKey")]
        public Input<string>? ParentKey { get; set; }

        public LocationCoordinatesGetArgs()
        {
        }
        public static new LocationCoordinatesGetArgs Empty => new LocationCoordinatesGetArgs();
    }
}
