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

    public sealed class Addrgrp6MemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address6/addrgrp6 name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Addrgrp6MemberArgs()
        {
        }
        public static new Addrgrp6MemberArgs Empty => new Addrgrp6MemberArgs();
    }
}
