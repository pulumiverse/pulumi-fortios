// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class Localinpolicy6InternetService6SrcNameArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Localinpolicy6InternetService6SrcNameArgs()
        {
        }
        public static new Localinpolicy6InternetService6SrcNameArgs Empty => new Localinpolicy6InternetService6SrcNameArgs();
    }
}
