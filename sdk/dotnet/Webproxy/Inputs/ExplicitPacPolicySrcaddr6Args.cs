// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy.Inputs
{

    public sealed class ExplicitPacPolicySrcaddr6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExplicitPacPolicySrcaddr6Args()
        {
        }
        public static new ExplicitPacPolicySrcaddr6Args Empty => new ExplicitPacPolicySrcaddr6Args();
    }
}
