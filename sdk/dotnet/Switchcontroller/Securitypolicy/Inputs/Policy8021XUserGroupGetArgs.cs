// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Securitypolicy.Inputs
{

    public sealed class Policy8021XUserGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Policy8021XUserGroupGetArgs()
        {
        }
        public static new Policy8021XUserGroupGetArgs Empty => new Policy8021XUserGroupGetArgs();
    }
}
