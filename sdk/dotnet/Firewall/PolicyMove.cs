// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    [FortiosResourceType("fortios:firewall/policyMove:PolicyMove")]
    public partial class PolicyMove : global::Pulumi.CustomResource
    {
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("move")]
        public Output<string> Move { get; private set; } = null!;

        [Output("policyidDst")]
        public Output<int> PolicyidDst { get; private set; } = null!;

        [Output("policyidSrc")]
        public Output<int> PolicyidSrc { get; private set; } = null!;

        [Output("statePolicySrcdstPos")]
        public Output<string?> StatePolicySrcdstPos { get; private set; } = null!;

        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyMove resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyMove(string name, PolicyMoveArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/policyMove:PolicyMove", name, args ?? new PolicyMoveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyMove(string name, Input<string> id, PolicyMoveState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/policyMove:PolicyMove", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyMove resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyMove Get(string name, Input<string> id, PolicyMoveState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyMove(name, id, state, options);
        }
    }

    public sealed class PolicyMoveArgs : global::Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("move", required: true)]
        public Input<string> Move { get; set; } = null!;

        [Input("policyidDst", required: true)]
        public Input<int> PolicyidDst { get; set; } = null!;

        [Input("policyidSrc", required: true)]
        public Input<int> PolicyidSrc { get; set; } = null!;

        [Input("statePolicySrcdstPos")]
        public Input<string>? StatePolicySrcdstPos { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PolicyMoveArgs()
        {
        }
        public static new PolicyMoveArgs Empty => new PolicyMoveArgs();
    }

    public sealed class PolicyMoveState : global::Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("move")]
        public Input<string>? Move { get; set; }

        [Input("policyidDst")]
        public Input<int>? PolicyidDst { get; set; }

        [Input("policyidSrc")]
        public Input<int>? PolicyidSrc { get; set; }

        [Input("statePolicySrcdstPos")]
        public Input<string>? StatePolicySrcdstPos { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PolicyMoveState()
        {
        }
        public static new PolicyMoveState Empty => new PolicyMoveState();
    }
}
