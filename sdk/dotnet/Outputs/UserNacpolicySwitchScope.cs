// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class UserNacpolicySwitchScope
    {
        /// <summary>
        /// Managed FortiSwitch name from available options.
        /// </summary>
        public readonly string? SwitchId;

        [OutputConstructor]
        private UserNacpolicySwitchScope(string? switchId)
        {
            SwitchId = switchId;
        }
    }
}
