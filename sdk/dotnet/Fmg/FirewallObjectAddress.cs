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
    /// This resource supports Create/Read/Update/Delete firewall object address for FortiManager.
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
    ///     var test1 = new Fortios.Fmg.FirewallObjectAddress("test1", new()
    ///     {
    ///         AssociatedIntf = "any",
    ///         Comment = "test obj address",
    ///         Fqdn = "fqdn.google.com",
    ///         Type = "fqdn",
    ///     });
    /// 
    ///     var test2 = new Fortios.Fmg.FirewallObjectAddress("test2", new()
    ///     {
    ///         AllowRouting = "disable",
    ///         AssociatedIntf = "any",
    ///         Comment = "test obj address",
    ///         Subnet = "2.2.2.0 255.255.255.0",
    ///         Type = "ipmask",
    ///     });
    /// 
    ///     var test3 = new Fortios.Fmg.FirewallObjectAddress("test3", new()
    ///     {
    ///         AssociatedIntf = "any",
    ///         Comment = "test obj address",
    ///         EndIp = "2.2.2.100",
    ///         StartIp = "2.2.2.1",
    ///         Type = "iprange",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/firewallObjectAddress:FirewallObjectAddress")]
    public partial class FirewallObjectAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Output("adom")]
        public Output<string?> Adom { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Output("allowRouting")]
        public Output<string?> AllowRouting { get; private set; } = null!;

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Output("associatedIntf")]
        public Output<string?> AssociatedIntf { get; private set; } = null!;

        /// <summary>
        /// Comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("endIp")]
        public Output<string?> EndIp { get; private set; } = null!;

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Output("fqdn")]
        public Output<string?> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("startIp")]
        public Output<string?> StartIp { get; private set; } = null!;

        /// <summary>
        /// IPv4 address/mask
        /// </summary>
        [Output("subnet")]
        public Output<string?> Subnet { get; private set; } = null!;

        /// <summary>
        /// Type of address, Enum: ["ipmask", "iprange", "fqdn"].
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallObjectAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallObjectAddress(string name, FirewallObjectAddressArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/firewallObjectAddress:FirewallObjectAddress", name, args ?? new FirewallObjectAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallObjectAddress(string name, Input<string> id, FirewallObjectAddressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/firewallObjectAddress:FirewallObjectAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallObjectAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallObjectAddress Get(string name, Input<string> id, FirewallObjectAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallObjectAddress(name, id, state, options);
        }
    }

    public sealed class FirewallObjectAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedIntf")]
        public Input<string>? AssociatedIntf { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// IPv4 address/mask
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Type of address, Enum: ["ipmask", "iprange", "fqdn"].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallObjectAddressArgs()
        {
        }
        public static new FirewallObjectAddressArgs Empty => new FirewallObjectAddressArgs();
    }

    public sealed class FirewallObjectAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedIntf")]
        public Input<string>? AssociatedIntf { get; set; }

        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// IPv4 address/mask
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Type of address, Enum: ["ipmask", "iprange", "fqdn"].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallObjectAddressState()
        {
        }
        public static new FirewallObjectAddressState Empty => new FirewallObjectAddressState();
    }
}
