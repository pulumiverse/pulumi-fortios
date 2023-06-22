// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class FirewallMulticastaddress6TaggingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Tagging entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.FirewallMulticastaddress6TaggingTagGetArgs>? _tags;

        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallMulticastaddress6TaggingTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FirewallMulticastaddress6TaggingTagGetArgs>());
            set => _tags = value;
        }

        public FirewallMulticastaddress6TaggingGetArgs()
        {
        }
        public static new FirewallMulticastaddress6TaggingGetArgs Empty => new FirewallMulticastaddress6TaggingGetArgs();
    }
}
