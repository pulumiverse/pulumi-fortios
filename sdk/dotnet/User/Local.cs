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
    /// Configure local users.
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
    ///     var trname3 = new Fortios.User.Ldap("trname3", new()
    ///     {
    ///         AccountKeyFilter = "(&amp;(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
    ///         AccountKeyProcessing = "same",
    ///         Cnid = "cn",
    ///         Dn = "EIWNCIEW",
    ///         GroupMemberCheck = "user-attr",
    ///         GroupObjectFilter = "(&amp;(objectcategory=group)(member=*))",
    ///         MemberAttr = "memberOf",
    ///         PasswordExpiryWarning = "disable",
    ///         PasswordRenewal = "disable",
    ///         Port = 389,
    ///         Secure = "disable",
    ///         Server = "1.1.1.1",
    ///         ServerIdentityCheck = "disable",
    ///         SourceIp = "0.0.0.0",
    ///         SslMinProtoVersion = "default",
    ///         Type = "simple",
    ///     });
    /// 
    ///     var trname = new Fortios.User.Local("trname", new()
    ///     {
    ///         AuthConcurrentOverride = "disable",
    ///         AuthConcurrentValue = 0,
    ///         Authtimeout = 0,
    ///         LdapServer = trname3.Name,
    ///         PasswdTime = "0000-00-00 00:00:00",
    ///         SmsServer = "fortiguard",
    ///         Status = "enable",
    ///         TwoFactor = "disable",
    ///         Type = "ldap",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Local can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/local:Local labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/local:Local labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/local:Local")]
    public partial class Local : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authConcurrentOverride")]
        public Output<string> AuthConcurrentOverride { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrent logins permitted from the same user.
        /// </summary>
        [Output("authConcurrentValue")]
        public Output<int> AuthConcurrentValue { get; private set; } = null!;

        /// <summary>
        /// Time in minutes before the authentication timeout for a user is reached.
        /// </summary>
        [Output("authtimeout")]
        public Output<int> Authtimeout { get; private set; } = null!;

        /// <summary>
        /// Two-factor recipient's email address.
        /// </summary>
        [Output("emailTo")]
        public Output<string> EmailTo { get; private set; } = null!;

        /// <summary>
        /// Two-factor recipient's FortiToken serial number.
        /// </summary>
        [Output("fortitoken")]
        public Output<string> Fortitoken { get; private set; } = null!;

        /// <summary>
        /// User ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Name of LDAP server with which the user must authenticate.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// User name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User's password.
        /// </summary>
        [Output("passwd")]
        public Output<string?> Passwd { get; private set; } = null!;

        /// <summary>
        /// Password policy to apply to this user, as defined in config user password-policy.
        /// </summary>
        [Output("passwdPolicy")]
        public Output<string> PasswdPolicy { get; private set; } = null!;

        /// <summary>
        /// Time of the last password update.
        /// </summary>
        [Output("passwdTime")]
        public Output<string> PasswdTime { get; private set; } = null!;

        /// <summary>
        /// IKEv2 Postquantum Preshared Key Identity.
        /// </summary>
        [Output("ppkIdentity")]
        public Output<string> PpkIdentity { get; private set; } = null!;

        /// <summary>
        /// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
        /// </summary>
        [Output("ppkSecret")]
        public Output<string?> PpkSecret { get; private set; } = null!;

        /// <summary>
        /// Quantum Key Distribution (QKD) profile.
        /// </summary>
        [Output("qkdProfile")]
        public Output<string> QkdProfile { get; private set; } = null!;

        /// <summary>
        /// Name of RADIUS server with which the user must authenticate.
        /// </summary>
        [Output("radiusServer")]
        public Output<string> RadiusServer { get; private set; } = null!;

        /// <summary>
        /// Two-factor recipient's SMS server.
        /// </summary>
        [Output("smsCustomServer")]
        public Output<string> SmsCustomServer { get; private set; } = null!;

        /// <summary>
        /// Two-factor recipient's mobile phone number.
        /// </summary>
        [Output("smsPhone")]
        public Output<string> SmsPhone { get; private set; } = null!;

        /// <summary>
        /// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Output("smsServer")]
        public Output<string> SmsServer { get; private set; } = null!;

        /// <summary>
        /// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Name of TACACS+ server with which the user must authenticate.
        /// </summary>
        [Output("tacacsServer")]
        public Output<string> TacacsServer { get; private set; } = null!;

        /// <summary>
        /// Enable/disable two-factor authentication.
        /// </summary>
        [Output("twoFactor")]
        public Output<string> TwoFactor { get; private set; } = null!;

        /// <summary>
        /// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
        /// </summary>
        [Output("twoFactorAuthentication")]
        public Output<string> TwoFactorAuthentication { get; private set; } = null!;

        /// <summary>
        /// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
        /// </summary>
        [Output("twoFactorNotification")]
        public Output<string> TwoFactorNotification { get; private set; } = null!;

        /// <summary>
        /// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("usernameCaseInsensitivity")]
        public Output<string> UsernameCaseInsensitivity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("usernameCaseSensitivity")]
        public Output<string> UsernameCaseSensitivity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("usernameSensitivity")]
        public Output<string> UsernameSensitivity { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
        /// </summary>
        [Output("workstation")]
        public Output<string> Workstation { get; private set; } = null!;


        /// <summary>
        /// Create a Local resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Local(string name, LocalArgs args, CustomResourceOptions? options = null)
            : base("fortios:user/local:Local", name, args ?? new LocalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Local(string name, Input<string> id, LocalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/local:Local", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "passwd",
                    "ppkSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Local resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Local Get(string name, Input<string> id, LocalState? state = null, CustomResourceOptions? options = null)
        {
            return new Local(name, id, state, options);
        }
    }

    public sealed class LocalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authConcurrentOverride")]
        public Input<string>? AuthConcurrentOverride { get; set; }

        /// <summary>
        /// Maximum number of concurrent logins permitted from the same user.
        /// </summary>
        [Input("authConcurrentValue")]
        public Input<int>? AuthConcurrentValue { get; set; }

        /// <summary>
        /// Time in minutes before the authentication timeout for a user is reached.
        /// </summary>
        [Input("authtimeout")]
        public Input<int>? Authtimeout { get; set; }

        /// <summary>
        /// Two-factor recipient's email address.
        /// </summary>
        [Input("emailTo")]
        public Input<string>? EmailTo { get; set; }

        /// <summary>
        /// Two-factor recipient's FortiToken serial number.
        /// </summary>
        [Input("fortitoken")]
        public Input<string>? Fortitoken { get; set; }

        /// <summary>
        /// User ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name of LDAP server with which the user must authenticate.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passwd")]
        private Input<string>? _passwd;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Passwd
        {
            get => _passwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Password policy to apply to this user, as defined in config user password-policy.
        /// </summary>
        [Input("passwdPolicy")]
        public Input<string>? PasswdPolicy { get; set; }

        /// <summary>
        /// Time of the last password update.
        /// </summary>
        [Input("passwdTime")]
        public Input<string>? PasswdTime { get; set; }

        /// <summary>
        /// IKEv2 Postquantum Preshared Key Identity.
        /// </summary>
        [Input("ppkIdentity")]
        public Input<string>? PpkIdentity { get; set; }

        [Input("ppkSecret")]
        private Input<string>? _ppkSecret;

        /// <summary>
        /// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
        /// </summary>
        public Input<string>? PpkSecret
        {
            get => _ppkSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ppkSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Quantum Key Distribution (QKD) profile.
        /// </summary>
        [Input("qkdProfile")]
        public Input<string>? QkdProfile { get; set; }

        /// <summary>
        /// Name of RADIUS server with which the user must authenticate.
        /// </summary>
        [Input("radiusServer")]
        public Input<string>? RadiusServer { get; set; }

        /// <summary>
        /// Two-factor recipient's SMS server.
        /// </summary>
        [Input("smsCustomServer")]
        public Input<string>? SmsCustomServer { get; set; }

        /// <summary>
        /// Two-factor recipient's mobile phone number.
        /// </summary>
        [Input("smsPhone")]
        public Input<string>? SmsPhone { get; set; }

        /// <summary>
        /// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Input("smsServer")]
        public Input<string>? SmsServer { get; set; }

        /// <summary>
        /// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// Name of TACACS+ server with which the user must authenticate.
        /// </summary>
        [Input("tacacsServer")]
        public Input<string>? TacacsServer { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication.
        /// </summary>
        [Input("twoFactor")]
        public Input<string>? TwoFactor { get; set; }

        /// <summary>
        /// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
        /// </summary>
        [Input("twoFactorAuthentication")]
        public Input<string>? TwoFactorAuthentication { get; set; }

        /// <summary>
        /// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
        /// </summary>
        [Input("twoFactorNotification")]
        public Input<string>? TwoFactorNotification { get; set; }

        /// <summary>
        /// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("usernameCaseInsensitivity")]
        public Input<string>? UsernameCaseInsensitivity { get; set; }

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("usernameCaseSensitivity")]
        public Input<string>? UsernameCaseSensitivity { get; set; }

        /// <summary>
        /// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("usernameSensitivity")]
        public Input<string>? UsernameSensitivity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
        /// </summary>
        [Input("workstation")]
        public Input<string>? Workstation { get; set; }

        public LocalArgs()
        {
        }
        public static new LocalArgs Empty => new LocalArgs();
    }

    public sealed class LocalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authConcurrentOverride")]
        public Input<string>? AuthConcurrentOverride { get; set; }

        /// <summary>
        /// Maximum number of concurrent logins permitted from the same user.
        /// </summary>
        [Input("authConcurrentValue")]
        public Input<int>? AuthConcurrentValue { get; set; }

        /// <summary>
        /// Time in minutes before the authentication timeout for a user is reached.
        /// </summary>
        [Input("authtimeout")]
        public Input<int>? Authtimeout { get; set; }

        /// <summary>
        /// Two-factor recipient's email address.
        /// </summary>
        [Input("emailTo")]
        public Input<string>? EmailTo { get; set; }

        /// <summary>
        /// Two-factor recipient's FortiToken serial number.
        /// </summary>
        [Input("fortitoken")]
        public Input<string>? Fortitoken { get; set; }

        /// <summary>
        /// User ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name of LDAP server with which the user must authenticate.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passwd")]
        private Input<string>? _passwd;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Passwd
        {
            get => _passwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Password policy to apply to this user, as defined in config user password-policy.
        /// </summary>
        [Input("passwdPolicy")]
        public Input<string>? PasswdPolicy { get; set; }

        /// <summary>
        /// Time of the last password update.
        /// </summary>
        [Input("passwdTime")]
        public Input<string>? PasswdTime { get; set; }

        /// <summary>
        /// IKEv2 Postquantum Preshared Key Identity.
        /// </summary>
        [Input("ppkIdentity")]
        public Input<string>? PpkIdentity { get; set; }

        [Input("ppkSecret")]
        private Input<string>? _ppkSecret;

        /// <summary>
        /// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
        /// </summary>
        public Input<string>? PpkSecret
        {
            get => _ppkSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ppkSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Quantum Key Distribution (QKD) profile.
        /// </summary>
        [Input("qkdProfile")]
        public Input<string>? QkdProfile { get; set; }

        /// <summary>
        /// Name of RADIUS server with which the user must authenticate.
        /// </summary>
        [Input("radiusServer")]
        public Input<string>? RadiusServer { get; set; }

        /// <summary>
        /// Two-factor recipient's SMS server.
        /// </summary>
        [Input("smsCustomServer")]
        public Input<string>? SmsCustomServer { get; set; }

        /// <summary>
        /// Two-factor recipient's mobile phone number.
        /// </summary>
        [Input("smsPhone")]
        public Input<string>? SmsPhone { get; set; }

        /// <summary>
        /// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Input("smsServer")]
        public Input<string>? SmsServer { get; set; }

        /// <summary>
        /// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Name of TACACS+ server with which the user must authenticate.
        /// </summary>
        [Input("tacacsServer")]
        public Input<string>? TacacsServer { get; set; }

        /// <summary>
        /// Enable/disable two-factor authentication.
        /// </summary>
        [Input("twoFactor")]
        public Input<string>? TwoFactor { get; set; }

        /// <summary>
        /// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
        /// </summary>
        [Input("twoFactorAuthentication")]
        public Input<string>? TwoFactorAuthentication { get; set; }

        /// <summary>
        /// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
        /// </summary>
        [Input("twoFactorNotification")]
        public Input<string>? TwoFactorNotification { get; set; }

        /// <summary>
        /// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("usernameCaseInsensitivity")]
        public Input<string>? UsernameCaseInsensitivity { get; set; }

        /// <summary>
        /// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("usernameCaseSensitivity")]
        public Input<string>? UsernameCaseSensitivity { get; set; }

        /// <summary>
        /// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("usernameSensitivity")]
        public Input<string>? UsernameSensitivity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
        /// </summary>
        [Input("workstation")]
        public Input<string>? Workstation { get; set; }

        public LocalState()
        {
        }
        public static new LocalState Empty => new LocalState();
    }
}
