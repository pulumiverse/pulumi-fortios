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

    public sealed class HaHaMgmtInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("gateway6")]
        public Input<string>? Gateway6 { get; set; }

        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("interface")]
        public Input<string>? Interface { get; set; }

        public HaHaMgmtInterfaceGetArgs()
        {
        }
        public static new HaHaMgmtInterfaceGetArgs Empty => new HaHaMgmtInterfaceGetArgs();
    }
}
