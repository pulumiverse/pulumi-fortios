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
    /// Configure IPv6 virtual IP groups.
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
    ///     var trname1 = new Fortios.Firewall.Vip6("trname1", new()
    ///     {
    ///         ArpReply = "enable",
    ///         Color = 0,
    ///         Extip = "2001:1:1:2::100",
    ///         Extport = "0-65535",
    ///         Fosid = 0,
    ///         HttpCookieAge = 60,
    ///         HttpCookieDomainFromHost = "disable",
    ///         HttpCookieGeneration = 0,
    ///         HttpCookieShare = "same-ip",
    ///         HttpIpHeader = "disable",
    ///         HttpMultiplex = "disable",
    ///         HttpsCookieSecure = "disable",
    ///         LdbMethod = "static",
    ///         Mappedip = "2001:1:1:2::200",
    ///         Mappedport = "0-65535",
    ///         MaxEmbryonicConnections = 1000,
    ///         OutlookWebAccess = "disable",
    ///         Persistence = "none",
    ///         Portforward = "disable",
    ///         Protocol = "tcp",
    ///         SslAlgorithm = "high",
    ///         SslClientFallback = "enable",
    ///         SslClientRenegotiation = "secure",
    ///         SslClientSessionStateMax = 1000,
    ///         SslClientSessionStateTimeout = 30,
    ///         SslClientSessionStateType = "both",
    ///         SslDhBits = "2048",
    ///         SslHpkp = "disable",
    ///         SslHpkpAge = 5184000,
    ///         SslHpkpIncludeSubdomains = "disable",
    ///         SslHsts = "disable",
    ///         SslHstsAge = 5184000,
    ///         SslHstsIncludeSubdomains = "disable",
    ///         SslHttpLocationConversion = "disable",
    ///         SslHttpMatchHost = "enable",
    ///         SslMaxVersion = "tls-1.2",
    ///         SslMinVersion = "tls-1.1",
    ///         SslMode = "half",
    ///         SslPfs = "require",
    ///         SslSendEmptyFrags = "enable",
    ///         SslServerAlgorithm = "client",
    ///         SslServerMaxVersion = "client",
    ///         SslServerMinVersion = "client",
    ///         SslServerSessionStateMax = 100,
    ///         SslServerSessionStateTimeout = 60,
    ///         SslServerSessionStateType = "both",
    ///         Type = "static-nat",
    ///         WeblogicServer = "disable",
    ///         WebsphereServer = "disable",
    ///     });
    /// 
    ///     var trname = new Fortios.Firewall.Vipgrp6("trname", new()
    ///     {
    ///         Color = 0,
    ///         Members = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Vipgrp6MemberArgs
    ///             {
    ///                 Name = trname1.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Vipgrp6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/vipgrp6:Vipgrp6 labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/vipgrp6:Vipgrp6 labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/vipgrp6:Vipgrp6")]
    public partial class Vipgrp6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.Vipgrp6Member>> Members { get; private set; } = null!;

        /// <summary>
        /// IPv6 VIP group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Vipgrp6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vipgrp6(string name, Vipgrp6Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp6:Vipgrp6", name, args ?? new Vipgrp6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Vipgrp6(string name, Input<string> id, Vipgrp6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp6:Vipgrp6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vipgrp6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vipgrp6 Get(string name, Input<string> id, Vipgrp6State? state = null, CustomResourceOptions? options = null)
        {
            return new Vipgrp6(name, id, state, options);
        }
    }

    public sealed class Vipgrp6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.Vipgrp6MemberArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Vipgrp6MemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Vipgrp6MemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPv6 VIP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Vipgrp6Args()
        {
        }
        public static new Vipgrp6Args Empty => new Vipgrp6Args();
    }

    public sealed class Vipgrp6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members")]
        private InputList<Inputs.Vipgrp6MemberGetArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Vipgrp6MemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Vipgrp6MemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPv6 VIP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Vipgrp6State()
        {
        }
        public static new Vipgrp6State Empty => new Vipgrp6State();
    }
}
