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

    public sealed class HsprofileOsuProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OSU provider name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public HsprofileOsuProviderArgs()
        {
        }
        public static new HsprofileOsuProviderArgs Empty => new HsprofileOsuProviderArgs();
    }
}
