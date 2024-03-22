// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Inputs
{

    public sealed class ProfileSaasApplicationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.ProfileSaasApplicationAccessRuleArgs>? _accessRules;

        /// <summary>
        /// CASB profile access rule. The structure of `access_rule` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationAccessRuleArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.ProfileSaasApplicationAccessRuleArgs>());
            set => _accessRules = value;
        }

        [Input("customControls")]
        private InputList<Inputs.ProfileSaasApplicationCustomControlArgs>? _customControls;

        /// <summary>
        /// CASB profile custom control. The structure of `custom_control` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationCustomControlArgs> CustomControls
        {
            get => _customControls ?? (_customControls = new InputList<Inputs.ProfileSaasApplicationCustomControlArgs>());
            set => _customControls = value;
        }

        /// <summary>
        /// Enable/disable domain control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("domainControl")]
        public Input<string>? DomainControl { get; set; }

        [Input("domainControlDomains")]
        private InputList<Inputs.ProfileSaasApplicationDomainControlDomainArgs>? _domainControlDomains;

        /// <summary>
        /// CASB profile domain control domains. The structure of `domain_control_domains` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationDomainControlDomainArgs> DomainControlDomains
        {
            get => _domainControlDomains ?? (_domainControlDomains = new InputList<Inputs.ProfileSaasApplicationDomainControlDomainArgs>());
            set => _domainControlDomains = value;
        }

        /// <summary>
        /// Enable/disable log settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// CASB profile SaaS application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable safe search. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("safeSearch")]
        public Input<string>? SafeSearch { get; set; }

        [Input("safeSearchControls")]
        private InputList<Inputs.ProfileSaasApplicationSafeSearchControlArgs>? _safeSearchControls;

        /// <summary>
        /// CASB profile safe search control. The structure of `safe_search_control` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationSafeSearchControlArgs> SafeSearchControls
        {
            get => _safeSearchControls ?? (_safeSearchControls = new InputList<Inputs.ProfileSaasApplicationSafeSearchControlArgs>());
            set => _safeSearchControls = value;
        }

        /// <summary>
        /// Enable/disable setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable tenant control. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("tenantControl")]
        public Input<string>? TenantControl { get; set; }

        [Input("tenantControlTenants")]
        private InputList<Inputs.ProfileSaasApplicationTenantControlTenantArgs>? _tenantControlTenants;

        /// <summary>
        /// CASB profile tenant control tenants. The structure of `tenant_control_tenants` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationTenantControlTenantArgs> TenantControlTenants
        {
            get => _tenantControlTenants ?? (_tenantControlTenants = new InputList<Inputs.ProfileSaasApplicationTenantControlTenantArgs>());
            set => _tenantControlTenants = value;
        }

        public ProfileSaasApplicationArgs()
        {
        }
        public static new ProfileSaasApplicationArgs Empty => new ProfileSaasApplicationArgs();
    }
}
