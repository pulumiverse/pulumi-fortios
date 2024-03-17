// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Networking
{
    /// <summary>
    /// Provides a resource to configure interface settings of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.system.Interface`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// 
    /// ### Loopback Interface
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var loopback1 = new Fortios.Networking.InterfacePort("loopback1", new()
    ///     {
    ///         Alias = "cc1",
    ///         Allowaccess = "ping http",
    ///         Description = "description",
    ///         Ip = "23.123.33.10 255.255.255.0",
    ///         Mode = "static",
    ///         Role = "lan",
    ///         Status = "up",
    ///         Type = "loopback",
    ///         Vdom = "root",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### VLAN Interface
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vlan1 = new Fortios.Networking.InterfacePort("vlan1", new()
    ///     {
    ///         Allowaccess = "ping",
    ///         Defaultgw = "enable",
    ///         Distance = "33",
    ///         Interface = "port2",
    ///         Ip = "3.123.33.10 255.255.255.0",
    ///         Mode = "static",
    ///         Role = "lan",
    ///         Type = "vlan",
    ///         Vdom = "root",
    ///         Vlanid = "3",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Physical Interface
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test1 = new Fortios.Networking.InterfacePort("test1", new()
    ///     {
    ///         Alias = "dkeeew",
    ///         Allowaccess = "ping https",
    ///         Defaultgw = "enable",
    ///         Description = "description",
    ///         DeviceIdentification = "enable",
    ///         Distance = "33",
    ///         DnsServerOverride = "enable",
    ///         Ip = "93.133.133.110 255.255.255.0",
    ///         Mode = "static",
    ///         Mtu = "2933",
    ///         MtuOverride = "enable",
    ///         Role = "lan",
    ///         Speed = "auto",
    ///         Status = "up",
    ///         TcpMss = "3232",
    ///         Type = "physical",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:networking/interfacePort:InterfacePort")]
    public partial class InterfacePort : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alias will be displayed with the interface name to make it easier to distinguish.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Permitted types of management access to this interface.
        /// </summary>
        [Output("allowaccess")]
        public Output<string> Allowaccess { get; private set; } = null!;

        /// <summary>
        /// Enable to get the gateway IP from the DHCP or PPPoE server.
        /// </summary>
        [Output("defaultgw")]
        public Output<string> Defaultgw { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
        /// </summary>
        [Output("deviceIdentification")]
        public Output<string> DeviceIdentification { get; private set; } = null!;

        /// <summary>
        /// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
        /// </summary>
        [Output("distance")]
        public Output<string> Distance { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use DNS acquired by DHCP or PPPoE.
        /// </summary>
        [Output("dnsServerOverride")]
        public Output<string> DnsServerOverride { get; private set; } = null!;

        /// <summary>
        /// Interface name.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Addressing mode.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// MTU value for this interface.
        /// </summary>
        [Output("mtu")]
        public Output<string> Mtu { get; private set; } = null!;

        /// <summary>
        /// Enable to set a custom MTU for this interface.
        /// </summary>
        [Output("mtuOverride")]
        public Output<string> MtuOverride { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Interface role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Interface speed. The default setting and the options available depend on the interface hardware.
        /// </summary>
        [Output("speed")]
        public Output<string> Speed { get; private set; } = null!;

        /// <summary>
        /// Bring the interface up or shut the interface down.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// TCP maximum segment size. 0 means do not change segment size.
        /// </summary>
        [Output("tcpMss")]
        public Output<string> TcpMss { get; private set; } = null!;

        /// <summary>
        /// Interface type (support physical, vlan, loopback).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Interface is in this virtual domain (VDOM).
        /// </summary>
        [Output("vdom")]
        public Output<string> Vdom { get; private set; } = null!;

        /// <summary>
        /// VLAN ID.
        /// </summary>
        [Output("vlanid")]
        public Output<string> Vlanid { get; private set; } = null!;


        /// <summary>
        /// Create a InterfacePort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterfacePort(string name, InterfacePortArgs args, CustomResourceOptions? options = null)
            : base("fortios:networking/interfacePort:InterfacePort", name, args ?? new InterfacePortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterfacePort(string name, Input<string> id, InterfacePortState? state = null, CustomResourceOptions? options = null)
            : base("fortios:networking/interfacePort:InterfacePort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterfacePort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterfacePort Get(string name, Input<string> id, InterfacePortState? state = null, CustomResourceOptions? options = null)
        {
            return new InterfacePort(name, id, state, options);
        }
    }

    public sealed class InterfacePortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias will be displayed with the interface name to make it easier to distinguish.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Permitted types of management access to this interface.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// Enable to get the gateway IP from the DHCP or PPPoE server.
        /// </summary>
        [Input("defaultgw")]
        public Input<string>? Defaultgw { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
        /// </summary>
        [Input("deviceIdentification")]
        public Input<string>? DeviceIdentification { get; set; }

        /// <summary>
        /// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Enable/disable use DNS acquired by DHCP or PPPoE.
        /// </summary>
        [Input("dnsServerOverride")]
        public Input<string>? DnsServerOverride { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Addressing mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// MTU value for this interface.
        /// </summary>
        [Input("mtu")]
        public Input<string>? Mtu { get; set; }

        /// <summary>
        /// Enable to set a custom MTU for this interface.
        /// </summary>
        [Input("mtuOverride")]
        public Input<string>? MtuOverride { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Interface role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Interface speed. The default setting and the options available depend on the interface hardware.
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Bring the interface up or shut the interface down.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TCP maximum segment size. 0 means do not change segment size.
        /// </summary>
        [Input("tcpMss")]
        public Input<string>? TcpMss { get; set; }

        /// <summary>
        /// Interface type (support physical, vlan, loopback).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Interface is in this virtual domain (VDOM).
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// VLAN ID.
        /// </summary>
        [Input("vlanid")]
        public Input<string>? Vlanid { get; set; }

        public InterfacePortArgs()
        {
        }
        public static new InterfacePortArgs Empty => new InterfacePortArgs();
    }

    public sealed class InterfacePortState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias will be displayed with the interface name to make it easier to distinguish.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Permitted types of management access to this interface.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// Enable to get the gateway IP from the DHCP or PPPoE server.
        /// </summary>
        [Input("defaultgw")]
        public Input<string>? Defaultgw { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
        /// </summary>
        [Input("deviceIdentification")]
        public Input<string>? DeviceIdentification { get; set; }

        /// <summary>
        /// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
        /// </summary>
        [Input("distance")]
        public Input<string>? Distance { get; set; }

        /// <summary>
        /// Enable/disable use DNS acquired by DHCP or PPPoE.
        /// </summary>
        [Input("dnsServerOverride")]
        public Input<string>? DnsServerOverride { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Addressing mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// MTU value for this interface.
        /// </summary>
        [Input("mtu")]
        public Input<string>? Mtu { get; set; }

        /// <summary>
        /// Enable to set a custom MTU for this interface.
        /// </summary>
        [Input("mtuOverride")]
        public Input<string>? MtuOverride { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Interface role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Interface speed. The default setting and the options available depend on the interface hardware.
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Bring the interface up or shut the interface down.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TCP maximum segment size. 0 means do not change segment size.
        /// </summary>
        [Input("tcpMss")]
        public Input<string>? TcpMss { get; set; }

        /// <summary>
        /// Interface type (support physical, vlan, loopback).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Interface is in this virtual domain (VDOM).
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// VLAN ID.
        /// </summary>
        [Input("vlanid")]
        public Input<string>? Vlanid { get; set; }

        public InterfacePortState()
        {
        }
        public static new InterfacePortState Empty => new InterfacePortState();
    }
}
