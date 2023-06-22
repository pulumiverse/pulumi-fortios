// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    [FortiosResourceType("fortios:index/firewallProxypolicySort:FirewallProxypolicySort")]
    public partial class FirewallProxypolicySort : global::Pulumi.CustomResource
    {
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("forceRecreate")]
        public Output<string?> ForceRecreate { get; private set; } = null!;

        [Output("sortby")]
        public Output<string> Sortby { get; private set; } = null!;

        [Output("sortdirection")]
        public Output<string> Sortdirection { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallProxypolicySort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallProxypolicySort(string name, FirewallProxypolicySortArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/firewallProxypolicySort:FirewallProxypolicySort", name, args ?? new FirewallProxypolicySortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallProxypolicySort(string name, Input<string> id, FirewallProxypolicySortState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallProxypolicySort:FirewallProxypolicySort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallProxypolicySort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallProxypolicySort Get(string name, Input<string> id, FirewallProxypolicySortState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallProxypolicySort(name, id, state, options);
        }
    }

    public sealed class FirewallProxypolicySortArgs : global::Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("forceRecreate")]
        public Input<string>? ForceRecreate { get; set; }

        [Input("sortby", required: true)]
        public Input<string> Sortby { get; set; } = null!;

        [Input("sortdirection", required: true)]
        public Input<string> Sortdirection { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallProxypolicySortArgs()
        {
        }
        public static new FirewallProxypolicySortArgs Empty => new FirewallProxypolicySortArgs();
    }

    public sealed class FirewallProxypolicySortState : global::Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("forceRecreate")]
        public Input<string>? ForceRecreate { get; set; }

        [Input("sortby")]
        public Input<string>? Sortby { get; set; }

        [Input("sortdirection")]
        public Input<string>? Sortdirection { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallProxypolicySortState()
        {
        }
        public static new FirewallProxypolicySortState Empty => new FirewallProxypolicySortState();
    }
}
