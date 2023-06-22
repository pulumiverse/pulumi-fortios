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
    /// <summary>
    /// Global settings for SAML authentication. Applies to FortiOS Version `&gt;= 6.2.4`.
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
    ///     var trname = new Fortios.SystemSaml("trname", new()
    ///     {
    ///         DefaultLoginPage = "normal",
    ///         DefaultProfile = "admin_no_access",
    ///         Life = 30,
    ///         Role = "service-provider",
    ///         Status = "disable",
    ///         Tolerance = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Saml can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemSaml:SystemSaml")]
    public partial class SystemSaml : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IdP Binding protocol. Valid values: `post`, `redirect`.
        /// </summary>
        [Output("bindingProtocol")]
        public Output<string> BindingProtocol { get; private set; } = null!;

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Choose default login page. Valid values: `normal`, `sso`.
        /// </summary>
        [Output("defaultLoginPage")]
        public Output<string> DefaultLoginPage { get; private set; } = null!;

        /// <summary>
        /// Default profile for new SSO admin.
        /// </summary>
        [Output("defaultProfile")]
        public Output<string> DefaultProfile { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// IDP certificate name.
        /// </summary>
        [Output("idpCert")]
        public Output<string> IdpCert { get; private set; } = null!;

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Output("idpEntityId")]
        public Output<string> IdpEntityId { get; private set; } = null!;

        /// <summary>
        /// IDP single logout URL.
        /// </summary>
        [Output("idpSingleLogoutUrl")]
        public Output<string> IdpSingleLogoutUrl { get; private set; } = null!;

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Output("idpSingleSignOnUrl")]
        public Output<string> IdpSingleSignOnUrl { get; private set; } = null!;

        /// <summary>
        /// Length of the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Output("life")]
        public Output<int> Life { get; private set; } = null!;

        /// <summary>
        /// SP portal URL.
        /// </summary>
        [Output("portalUrl")]
        public Output<string> PortalUrl { get; private set; } = null!;

        /// <summary>
        /// SAML role. Valid values: `identity-provider`, `service-provider`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Server address.
        /// </summary>
        [Output("serverAddress")]
        public Output<string> ServerAddress { get; private set; } = null!;

        /// <summary>
        /// Authorized service providers. The structure of `service_providers` block is documented below.
        /// </summary>
        [Output("serviceProviders")]
        public Output<ImmutableArray<Outputs.SystemSamlServiceProvider>> ServiceProviders { get; private set; } = null!;

        /// <summary>
        /// SP single logout URL.
        /// </summary>
        [Output("singleLogoutUrl")]
        public Output<string> SingleLogoutUrl { get; private set; } = null!;

        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        [Output("singleSignOnUrl")]
        public Output<string> SingleSignOnUrl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tolerance to the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Output("tolerance")]
        public Output<int> Tolerance { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSaml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSaml(string name, SystemSamlArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSaml:SystemSaml", name, args ?? new SystemSamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSaml(string name, Input<string> id, SystemSamlState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSaml:SystemSaml", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSaml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSaml Get(string name, Input<string> id, SystemSamlState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSaml(name, id, state, options);
        }
    }

    public sealed class SystemSamlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IdP Binding protocol. Valid values: `post`, `redirect`.
        /// </summary>
        [Input("bindingProtocol")]
        public Input<string>? BindingProtocol { get; set; }

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Choose default login page. Valid values: `normal`, `sso`.
        /// </summary>
        [Input("defaultLoginPage")]
        public Input<string>? DefaultLoginPage { get; set; }

        /// <summary>
        /// Default profile for new SSO admin.
        /// </summary>
        [Input("defaultProfile")]
        public Input<string>? DefaultProfile { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// IDP certificate name.
        /// </summary>
        [Input("idpCert")]
        public Input<string>? IdpCert { get; set; }

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// IDP single logout URL.
        /// </summary>
        [Input("idpSingleLogoutUrl")]
        public Input<string>? IdpSingleLogoutUrl { get; set; }

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Input("idpSingleSignOnUrl")]
        public Input<string>? IdpSingleSignOnUrl { get; set; }

        /// <summary>
        /// Length of the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Input("life")]
        public Input<int>? Life { get; set; }

        /// <summary>
        /// SP portal URL.
        /// </summary>
        [Input("portalUrl")]
        public Input<string>? PortalUrl { get; set; }

        /// <summary>
        /// SAML role. Valid values: `identity-provider`, `service-provider`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Server address.
        /// </summary>
        [Input("serverAddress")]
        public Input<string>? ServerAddress { get; set; }

        [Input("serviceProviders")]
        private InputList<Inputs.SystemSamlServiceProviderArgs>? _serviceProviders;

        /// <summary>
        /// Authorized service providers. The structure of `service_providers` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSamlServiceProviderArgs> ServiceProviders
        {
            get => _serviceProviders ?? (_serviceProviders = new InputList<Inputs.SystemSamlServiceProviderArgs>());
            set => _serviceProviders = value;
        }

        /// <summary>
        /// SP single logout URL.
        /// </summary>
        [Input("singleLogoutUrl")]
        public Input<string>? SingleLogoutUrl { get; set; }

        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        [Input("singleSignOnUrl")]
        public Input<string>? SingleSignOnUrl { get; set; }

        /// <summary>
        /// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Tolerance to the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Input("tolerance")]
        public Input<int>? Tolerance { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSamlArgs()
        {
        }
        public static new SystemSamlArgs Empty => new SystemSamlArgs();
    }

    public sealed class SystemSamlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IdP Binding protocol. Valid values: `post`, `redirect`.
        /// </summary>
        [Input("bindingProtocol")]
        public Input<string>? BindingProtocol { get; set; }

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Choose default login page. Valid values: `normal`, `sso`.
        /// </summary>
        [Input("defaultLoginPage")]
        public Input<string>? DefaultLoginPage { get; set; }

        /// <summary>
        /// Default profile for new SSO admin.
        /// </summary>
        [Input("defaultProfile")]
        public Input<string>? DefaultProfile { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// IDP certificate name.
        /// </summary>
        [Input("idpCert")]
        public Input<string>? IdpCert { get; set; }

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// IDP single logout URL.
        /// </summary>
        [Input("idpSingleLogoutUrl")]
        public Input<string>? IdpSingleLogoutUrl { get; set; }

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Input("idpSingleSignOnUrl")]
        public Input<string>? IdpSingleSignOnUrl { get; set; }

        /// <summary>
        /// Length of the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Input("life")]
        public Input<int>? Life { get; set; }

        /// <summary>
        /// SP portal URL.
        /// </summary>
        [Input("portalUrl")]
        public Input<string>? PortalUrl { get; set; }

        /// <summary>
        /// SAML role. Valid values: `identity-provider`, `service-provider`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Server address.
        /// </summary>
        [Input("serverAddress")]
        public Input<string>? ServerAddress { get; set; }

        [Input("serviceProviders")]
        private InputList<Inputs.SystemSamlServiceProviderGetArgs>? _serviceProviders;

        /// <summary>
        /// Authorized service providers. The structure of `service_providers` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSamlServiceProviderGetArgs> ServiceProviders
        {
            get => _serviceProviders ?? (_serviceProviders = new InputList<Inputs.SystemSamlServiceProviderGetArgs>());
            set => _serviceProviders = value;
        }

        /// <summary>
        /// SP single logout URL.
        /// </summary>
        [Input("singleLogoutUrl")]
        public Input<string>? SingleLogoutUrl { get; set; }

        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        [Input("singleSignOnUrl")]
        public Input<string>? SingleSignOnUrl { get; set; }

        /// <summary>
        /// Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Tolerance to the range of time when the assertion is valid (in minutes).
        /// </summary>
        [Input("tolerance")]
        public Input<int>? Tolerance { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSamlState()
        {
        }
        public static new SystemSamlState Empty => new SystemSamlState();
    }
}
