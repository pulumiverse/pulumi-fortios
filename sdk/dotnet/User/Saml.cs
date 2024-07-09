// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// SAML server entry configuration. Applies to FortiOS Version `&gt;= 6.2.4`.
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
    ///     var tr3 = new Fortios.User.Saml("tr3", new()
    ///     {
    ///         Cert = "Fortinet_Factory",
    ///         EntityId = "https://1.1.1.1",
    ///         IdpCert = "cer11",
    ///         IdpEntityId = "https://1.1.1.1/acc",
    ///         IdpSingleLogoutUrl = "https://1.1.1.1/lo",
    ///         IdpSingleSignOnUrl = "https://1.1.1.1/sou",
    ///         SingleLogoutUrl = "https://1.1.1.1/logout",
    ///         SingleSignOnUrl = "https://1.1.1.1/sign",
    ///         UserName = "ad111",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Saml can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/saml:Saml labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/saml:Saml labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/saml:Saml")]
    public partial class Saml : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("adfsClaim")]
        public Output<string> AdfsClaim { get; private set; } = null!;

        /// <summary>
        /// URL to verify authentication.
        /// </summary>
        [Output("authUrl")]
        public Output<string?> AuthUrl { get; private set; } = null!;

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
        /// </summary>
        [Output("clockTolerance")]
        public Output<int> ClockTolerance { get; private set; } = null!;

        /// <summary>
        /// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
        /// </summary>
        [Output("digestMethod")]
        public Output<string> DigestMethod { get; private set; } = null!;

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Output("groupClaimType")]
        public Output<string> GroupClaimType { get; private set; } = null!;

        /// <summary>
        /// Group name in assertion statement.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// IDP Certificate name.
        /// </summary>
        [Output("idpCert")]
        public Output<string> IdpCert { get; private set; } = null!;

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Output("idpEntityId")]
        public Output<string> IdpEntityId { get; private set; } = null!;

        /// <summary>
        /// IDP single logout url.
        /// </summary>
        [Output("idpSingleLogoutUrl")]
        public Output<string> IdpSingleLogoutUrl { get; private set; } = null!;

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Output("idpSingleSignOnUrl")]
        public Output<string> IdpSingleSignOnUrl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("limitRelaystate")]
        public Output<string> LimitRelaystate { get; private set; } = null!;

        /// <summary>
        /// SAML server entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("reauth")]
        public Output<string> Reauth { get; private set; } = null!;

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
        /// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Output("userClaimType")]
        public Output<string> UserClaimType { get; private set; } = null!;

        /// <summary>
        /// User name in assertion statement.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Saml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Saml(string name, SamlArgs args, CustomResourceOptions? options = null)
            : base("fortios:user/saml:Saml", name, args ?? new SamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Saml(string name, Input<string> id, SamlState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/saml:Saml", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Saml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Saml Get(string name, Input<string> id, SamlState? state = null, CustomResourceOptions? options = null)
        {
            return new Saml(name, id, state, options);
        }
    }

    public sealed class SamlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("adfsClaim")]
        public Input<string>? AdfsClaim { get; set; }

        /// <summary>
        /// URL to verify authentication.
        /// </summary>
        [Input("authUrl")]
        public Input<string>? AuthUrl { get; set; }

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
        /// </summary>
        [Input("clockTolerance")]
        public Input<int>? ClockTolerance { get; set; }

        /// <summary>
        /// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
        /// </summary>
        [Input("digestMethod")]
        public Input<string>? DigestMethod { get; set; }

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Input("groupClaimType")]
        public Input<string>? GroupClaimType { get; set; }

        /// <summary>
        /// Group name in assertion statement.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// IDP Certificate name.
        /// </summary>
        [Input("idpCert", required: true)]
        public Input<string> IdpCert { get; set; } = null!;

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Input("idpEntityId", required: true)]
        public Input<string> IdpEntityId { get; set; } = null!;

        /// <summary>
        /// IDP single logout url.
        /// </summary>
        [Input("idpSingleLogoutUrl")]
        public Input<string>? IdpSingleLogoutUrl { get; set; }

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Input("idpSingleSignOnUrl", required: true)]
        public Input<string> IdpSingleSignOnUrl { get; set; } = null!;

        /// <summary>
        /// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("limitRelaystate")]
        public Input<string>? LimitRelaystate { get; set; }

        /// <summary>
        /// SAML server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reauth")]
        public Input<string>? Reauth { get; set; }

        /// <summary>
        /// SP single logout URL.
        /// </summary>
        [Input("singleLogoutUrl")]
        public Input<string>? SingleLogoutUrl { get; set; }

        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        [Input("singleSignOnUrl", required: true)]
        public Input<string> SingleSignOnUrl { get; set; } = null!;

        /// <summary>
        /// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Input("userClaimType")]
        public Input<string>? UserClaimType { get; set; }

        /// <summary>
        /// User name in assertion statement.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SamlArgs()
        {
        }
        public static new SamlArgs Empty => new SamlArgs();
    }

    public sealed class SamlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("adfsClaim")]
        public Input<string>? AdfsClaim { get; set; }

        /// <summary>
        /// URL to verify authentication.
        /// </summary>
        [Input("authUrl")]
        public Input<string>? AuthUrl { get; set; }

        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
        /// </summary>
        [Input("clockTolerance")]
        public Input<int>? ClockTolerance { get; set; }

        /// <summary>
        /// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
        /// </summary>
        [Input("digestMethod")]
        public Input<string>? DigestMethod { get; set; }

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Input("groupClaimType")]
        public Input<string>? GroupClaimType { get; set; }

        /// <summary>
        /// Group name in assertion statement.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// IDP Certificate name.
        /// </summary>
        [Input("idpCert")]
        public Input<string>? IdpCert { get; set; }

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// IDP single logout url.
        /// </summary>
        [Input("idpSingleLogoutUrl")]
        public Input<string>? IdpSingleLogoutUrl { get; set; }

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Input("idpSingleSignOnUrl")]
        public Input<string>? IdpSingleSignOnUrl { get; set; }

        /// <summary>
        /// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("limitRelaystate")]
        public Input<string>? LimitRelaystate { get; set; }

        /// <summary>
        /// SAML server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reauth")]
        public Input<string>? Reauth { get; set; }

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
        /// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
        /// </summary>
        [Input("userClaimType")]
        public Input<string>? UserClaimType { get; set; }

        /// <summary>
        /// User name in assertion statement.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SamlState()
        {
        }
        public static new SamlState Empty => new SamlState();
    }
}
