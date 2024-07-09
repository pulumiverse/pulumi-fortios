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
    /// <summary>
    /// Provides a resource to configure firewall address group used in firewall policies of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.firewall.Addrgrp`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s1 = new Fortios.Firewall.ObjectAddressgroup("s1", new()
    ///     {
    ///         Comment = "dfdsad",
    ///         Members = new[]
    ///         {
    ///             "google-play",
    ///             "swscan.apple.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:firewall/objectAddressgroup:ObjectAddressgroup")]
    public partial class ObjectAddressgroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Address objects contained within the group.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Address group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectAddressgroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectAddressgroup(string name, ObjectAddressgroupArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/objectAddressgroup:ObjectAddressgroup", name, args ?? new ObjectAddressgroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectAddressgroup(string name, Input<string> id, ObjectAddressgroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/objectAddressgroup:ObjectAddressgroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectAddressgroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectAddressgroup Get(string name, Input<string> id, ObjectAddressgroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectAddressgroup(name, id, state, options);
        }
    }

    public sealed class ObjectAddressgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// Address objects contained within the group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ObjectAddressgroupArgs()
        {
        }
        public static new ObjectAddressgroupArgs Empty => new ObjectAddressgroupArgs();
    }

    public sealed class ObjectAddressgroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Address objects contained within the group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ObjectAddressgroupState()
        {
        }
        public static new ObjectAddressgroupState Empty => new ObjectAddressgroupState();
    }
}
