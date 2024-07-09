// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web
{
    /// <summary>
    /// Configure SSL VPN user group bookmark.
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
    ///     var trname = new Fortios.Vpn.Ssl.Web.Usergroupbookmark("trname", new()
    ///     {
    ///         Bookmarks = new[]
    ///         {
    ///             new Fortios.Vpn.Ssl.Web.Inputs.UsergroupbookmarkBookmarkArgs
    ///             {
    ///                 Apptype = "citrix",
    ///                 ListeningPort = 0,
    ///                 Name = "b1",
    ///                 Port = 0,
    ///                 RemotePort = 0,
    ///                 Security = "rdp",
    ///                 ServerLayout = "en-us-qwerty",
    ///                 ShowStatusWindow = "disable",
    ///                 Sso = "disable",
    ///                 SsoCredential = "sslvpn-login",
    ///                 SsoCredentialSentOnce = "disable",
    ///                 Url = "www.aaa.com",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpnSslWeb UserGroupBookmark can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/ssl/web/usergroupbookmark:Usergroupbookmark labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/ssl/web/usergroupbookmark:Usergroupbookmark labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpn/ssl/web/usergroupbookmark:Usergroupbookmark")]
    public partial class Usergroupbookmark : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        [Output("bookmarks")]
        public Output<ImmutableArray<Outputs.UsergroupbookmarkBookmark>> Bookmarks { get; private set; } = null!;

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
        /// Group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Usergroupbookmark resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Usergroupbookmark(string name, UsergroupbookmarkArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/ssl/web/usergroupbookmark:Usergroupbookmark", name, args ?? new UsergroupbookmarkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Usergroupbookmark(string name, Input<string> id, UsergroupbookmarkState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/ssl/web/usergroupbookmark:Usergroupbookmark", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Usergroupbookmark resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Usergroupbookmark Get(string name, Input<string> id, UsergroupbookmarkState? state = null, CustomResourceOptions? options = null)
        {
            return new Usergroupbookmark(name, id, state, options);
        }
    }

    public sealed class UsergroupbookmarkArgs : global::Pulumi.ResourceArgs
    {
        [Input("bookmarks")]
        private InputList<Inputs.UsergroupbookmarkBookmarkArgs>? _bookmarks;

        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        public InputList<Inputs.UsergroupbookmarkBookmarkArgs> Bookmarks
        {
            get => _bookmarks ?? (_bookmarks = new InputList<Inputs.UsergroupbookmarkBookmarkArgs>());
            set => _bookmarks = value;
        }

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

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UsergroupbookmarkArgs()
        {
        }
        public static new UsergroupbookmarkArgs Empty => new UsergroupbookmarkArgs();
    }

    public sealed class UsergroupbookmarkState : global::Pulumi.ResourceArgs
    {
        [Input("bookmarks")]
        private InputList<Inputs.UsergroupbookmarkBookmarkGetArgs>? _bookmarks;

        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        public InputList<Inputs.UsergroupbookmarkBookmarkGetArgs> Bookmarks
        {
            get => _bookmarks ?? (_bookmarks = new InputList<Inputs.UsergroupbookmarkBookmarkGetArgs>());
            set => _bookmarks = value;
        }

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

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UsergroupbookmarkState()
        {
        }
        public static new UsergroupbookmarkState Empty => new UsergroupbookmarkState();
    }
}
