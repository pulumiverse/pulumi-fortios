// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Outputs
{

    [OutputType]
    public sealed class ProfileSaasApplication
    {
        /// <summary>
        /// CASB profile access rule. The structure of `access_rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileSaasApplicationAccessRule> AccessRules;
        /// <summary>
        /// CASB profile custom control. The structure of `custom_control` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileSaasApplicationCustomControl> CustomControls;
        /// <summary>
        /// Enable/disable domain control. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? DomainControl;
        /// <summary>
        /// CASB profile domain control domains. The structure of `domain_control_domains` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileSaasApplicationDomainControlDomain> DomainControlDomains;
        /// <summary>
        /// Enable/disable log settings. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// CASB profile SaaS application name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Enable/disable safe search. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SafeSearch;
        /// <summary>
        /// CASB profile safe search control. The structure of `safe_search_control` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileSaasApplicationSafeSearchControl> SafeSearchControls;
        /// <summary>
        /// Enable/disable setting. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Enable/disable tenant control. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? TenantControl;
        /// <summary>
        /// CASB profile tenant control tenants. The structure of `tenant_control_tenants` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileSaasApplicationTenantControlTenant> TenantControlTenants;

        [OutputConstructor]
        private ProfileSaasApplication(
            ImmutableArray<Outputs.ProfileSaasApplicationAccessRule> accessRules,

            ImmutableArray<Outputs.ProfileSaasApplicationCustomControl> customControls,

            string? domainControl,

            ImmutableArray<Outputs.ProfileSaasApplicationDomainControlDomain> domainControlDomains,

            string? log,

            string? name,

            string? safeSearch,

            ImmutableArray<Outputs.ProfileSaasApplicationSafeSearchControl> safeSearchControls,

            string? status,

            string? tenantControl,

            ImmutableArray<Outputs.ProfileSaasApplicationTenantControlTenant> tenantControlTenants)
        {
            AccessRules = accessRules;
            CustomControls = customControls;
            DomainControl = domainControl;
            DomainControlDomains = domainControlDomains;
            Log = log;
            Name = name;
            SafeSearch = safeSearch;
            SafeSearchControls = safeSearchControls;
            Status = status;
            TenantControl = tenantControl;
            TenantControlTenants = tenantControlTenants;
        }
    }
}
