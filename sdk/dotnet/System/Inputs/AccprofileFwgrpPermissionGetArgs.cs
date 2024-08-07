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

    public sealed class AccprofileFwgrpPermissionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Other Firewall Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("others")]
        public Input<string>? Others { get; set; }

        /// <summary>
        /// Policy Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Schedule Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Service Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public AccprofileFwgrpPermissionGetArgs()
        {
        }
        public static new AccprofileFwgrpPermissionGetArgs Empty => new AccprofileFwgrpPermissionGetArgs();
    }
}
