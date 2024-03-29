// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Fmg
{
    /// <summary>
    /// This resource supports Create/Read/Update/Delete firewall object ippool for FortiManager.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test1 = new Fortios.Fmg.FirewallObjectIppool("test1", new()
    ///     {
    ///         ArpIntf = "any",
    ///         ArpReply = "enable",
    ///         AssociatedIntf = "any",
    ///         Comment = "test obj ippool",
    ///         Endip = "1.1.10.100",
    ///         Startip = "1.1.10.1",
    ///         Type = "one-to-one",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/firewallObjectIppool:FirewallObjectIppool")]
    public partial class FirewallObjectIppool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Output("adom")]
        public Output<string?> Adom { get; private set; } = null!;

        /// <summary>
        /// Select an interface that will reply to ARP requests.
        /// </summary>
        [Output("arpIntf")]
        public Output<string?> ArpIntf { get; private set; } = null!;

        /// <summary>
        /// Enable/disable replying to ARP request, default is "enable".
        /// </summary>
        [Output("arpReply")]
        public Output<string?> ArpReply { get; private set; } = null!;

        /// <summary>
        /// Associated interface name.
        /// </summary>
        [Output("associatedIntf")]
        public Output<string?> AssociatedIntf { get; private set; } = null!;

        /// <summary>
        /// Comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Output("endip")]
        public Output<string> Endip { get; private set; } = null!;

        /// <summary>
        /// Ippool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Output("startip")]
        public Output<string> Startip { get; private set; } = null!;

        /// <summary>
        /// Ip pool type, Enum: ["overload", "one-to-one"].
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallObjectIppool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallObjectIppool(string name, FirewallObjectIppoolArgs args, CustomResourceOptions? options = null)
            : base("fortios:fmg/firewallObjectIppool:FirewallObjectIppool", name, args ?? new FirewallObjectIppoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallObjectIppool(string name, Input<string> id, FirewallObjectIppoolState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/firewallObjectIppool:FirewallObjectIppool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallObjectIppool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallObjectIppool Get(string name, Input<string> id, FirewallObjectIppoolState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallObjectIppool(name, id, state, options);
        }
    }

    public sealed class FirewallObjectIppoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Select an interface that will reply to ARP requests.
        /// </summary>
        [Input("arpIntf")]
        public Input<string>? ArpIntf { get; set; }

        /// <summary>
        /// Enable/disable replying to ARP request, default is "enable".
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Associated interface name.
        /// </summary>
        [Input("associatedIntf")]
        public Input<string>? AssociatedIntf { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Input("endip", required: true)]
        public Input<string> Endip { get; set; } = null!;

        /// <summary>
        /// Ippool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Input("startip", required: true)]
        public Input<string> Startip { get; set; } = null!;

        /// <summary>
        /// Ip pool type, Enum: ["overload", "one-to-one"].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallObjectIppoolArgs()
        {
        }
        public static new FirewallObjectIppoolArgs Empty => new FirewallObjectIppoolArgs();
    }

    public sealed class FirewallObjectIppoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Select an interface that will reply to ARP requests.
        /// </summary>
        [Input("arpIntf")]
        public Input<string>? ArpIntf { get; set; }

        /// <summary>
        /// Enable/disable replying to ARP request, default is "enable".
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Associated interface name.
        /// </summary>
        [Input("associatedIntf")]
        public Input<string>? AssociatedIntf { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Input("endip")]
        public Input<string>? Endip { get; set; }

        /// <summary>
        /// Ippool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool.
        /// </summary>
        [Input("startip")]
        public Input<string>? Startip { get; set; }

        /// <summary>
        /// Ip pool type, Enum: ["overload", "one-to-one"].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallObjectIppoolState()
        {
        }
        public static new FirewallObjectIppoolState Empty => new FirewallObjectIppoolState();
    }
}
