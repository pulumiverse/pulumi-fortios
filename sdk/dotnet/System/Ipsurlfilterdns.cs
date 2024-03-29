// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure IPS URL filter DNS servers.
    /// 
    /// ## Import
    /// 
    /// System IpsUrlfilterDns can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipsurlfilterdns:Ipsurlfilterdns labelname {{address}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipsurlfilterdns:Ipsurlfilterdns labelname {{address}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/ipsurlfilterdns:Ipsurlfilterdns")]
    public partial class Ipsurlfilterdns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS server IP address.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipv6Capability")]
        public Output<string> Ipv6Capability { get; private set; } = null!;

        /// <summary>
        /// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipsurlfilterdns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipsurlfilterdns(string name, IpsurlfilterdnsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/ipsurlfilterdns:Ipsurlfilterdns", name, args ?? new IpsurlfilterdnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipsurlfilterdns(string name, Input<string> id, IpsurlfilterdnsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/ipsurlfilterdns:Ipsurlfilterdns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipsurlfilterdns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipsurlfilterdns Get(string name, Input<string> id, IpsurlfilterdnsState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipsurlfilterdns(name, id, state, options);
        }
    }

    public sealed class IpsurlfilterdnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6Capability")]
        public Input<string>? Ipv6Capability { get; set; }

        /// <summary>
        /// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfilterdnsArgs()
        {
        }
        public static new IpsurlfilterdnsArgs Empty => new IpsurlfilterdnsArgs();
    }

    public sealed class IpsurlfilterdnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipv6Capability")]
        public Input<string>? Ipv6Capability { get; set; }

        /// <summary>
        /// Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfilterdnsState()
        {
        }
        public static new IpsurlfilterdnsState Empty => new IpsurlfilterdnsState();
    }
}
