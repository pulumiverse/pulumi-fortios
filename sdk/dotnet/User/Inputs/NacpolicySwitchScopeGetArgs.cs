// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Inputs
{

    public sealed class NacpolicySwitchScopeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed FortiSwitch name from available options.
        /// </summary>
        [Input("switchId")]
        public Input<string>? SwitchId { get; set; }

        public NacpolicySwitchScopeGetArgs()
        {
        }
        public static new NacpolicySwitchScopeGetArgs Empty => new NacpolicySwitchScopeGetArgs();
    }
}
