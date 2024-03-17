// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class AutomationstitchAction
    {
        /// <summary>
        /// Action name.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Delay before execution (in seconds).
        /// </summary>
        public readonly int? Delay;
        /// <summary>
        /// Entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Required in action chain. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Required;

        [OutputConstructor]
        private AutomationstitchAction(
            string? action,

            int? delay,

            int? id,

            string? required)
        {
            Action = action;
            Delay = delay;
            Id = id;
            Required = required;
        }
    }
}
