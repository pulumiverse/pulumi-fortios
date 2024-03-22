# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ProfileSaasApplicationArgs',
    'ProfileSaasApplicationAccessRuleArgs',
    'ProfileSaasApplicationCustomControlArgs',
    'ProfileSaasApplicationCustomControlOptionArgs',
    'ProfileSaasApplicationCustomControlOptionUserInputArgs',
    'ProfileSaasApplicationDomainControlDomainArgs',
    'ProfileSaasApplicationSafeSearchControlArgs',
    'ProfileSaasApplicationTenantControlTenantArgs',
    'SaasapplicationDomainArgs',
    'UseractivityControlOptionArgs',
    'UseractivityControlOptionOperationArgs',
    'UseractivityControlOptionOperationValueArgs',
    'UseractivityMatchArgs',
    'UseractivityMatchRuleArgs',
    'UseractivityMatchRuleDomainArgs',
    'UseractivityMatchRuleMethodArgs',
]

@pulumi.input_type
class ProfileSaasApplicationArgs:
    def __init__(__self__, *,
                 access_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationAccessRuleArgs']]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlArgs']]]] = None,
                 domain_control: Optional[pulumi.Input[str]] = None,
                 domain_control_domains: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationDomainControlDomainArgs']]]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 safe_search: Optional[pulumi.Input[str]] = None,
                 safe_search_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationSafeSearchControlArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tenant_control: Optional[pulumi.Input[str]] = None,
                 tenant_control_tenants: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationTenantControlTenantArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationAccessRuleArgs']]] access_rules: CASB profile access rule. The structure of `access_rule` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlArgs']]] custom_controls: CASB profile custom control. The structure of `custom_control` block is documented below.
        :param pulumi.Input[str] domain_control: Enable/disable domain control. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationDomainControlDomainArgs']]] domain_control_domains: CASB profile domain control domains. The structure of `domain_control_domains` block is documented below.
        :param pulumi.Input[str] log: Enable/disable log settings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: CASB profile SaaS application name.
        :param pulumi.Input[str] safe_search: Enable/disable safe search. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationSafeSearchControlArgs']]] safe_search_controls: CASB profile safe search control. The structure of `safe_search_control` block is documented below.
        :param pulumi.Input[str] status: Enable/disable setting. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] tenant_control: Enable/disable tenant control. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationTenantControlTenantArgs']]] tenant_control_tenants: CASB profile tenant control tenants. The structure of `tenant_control_tenants` block is documented below.
        """
        if access_rules is not None:
            pulumi.set(__self__, "access_rules", access_rules)
        if custom_controls is not None:
            pulumi.set(__self__, "custom_controls", custom_controls)
        if domain_control is not None:
            pulumi.set(__self__, "domain_control", domain_control)
        if domain_control_domains is not None:
            pulumi.set(__self__, "domain_control_domains", domain_control_domains)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if safe_search is not None:
            pulumi.set(__self__, "safe_search", safe_search)
        if safe_search_controls is not None:
            pulumi.set(__self__, "safe_search_controls", safe_search_controls)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tenant_control is not None:
            pulumi.set(__self__, "tenant_control", tenant_control)
        if tenant_control_tenants is not None:
            pulumi.set(__self__, "tenant_control_tenants", tenant_control_tenants)

    @property
    @pulumi.getter(name="accessRules")
    def access_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationAccessRuleArgs']]]]:
        """
        CASB profile access rule. The structure of `access_rule` block is documented below.
        """
        return pulumi.get(self, "access_rules")

    @access_rules.setter
    def access_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationAccessRuleArgs']]]]):
        pulumi.set(self, "access_rules", value)

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlArgs']]]]:
        """
        CASB profile custom control. The structure of `custom_control` block is documented below.
        """
        return pulumi.get(self, "custom_controls")

    @custom_controls.setter
    def custom_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlArgs']]]]):
        pulumi.set(self, "custom_controls", value)

    @property
    @pulumi.getter(name="domainControl")
    def domain_control(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable domain control. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "domain_control")

    @domain_control.setter
    def domain_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_control", value)

    @property
    @pulumi.getter(name="domainControlDomains")
    def domain_control_domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationDomainControlDomainArgs']]]]:
        """
        CASB profile domain control domains. The structure of `domain_control_domains` block is documented below.
        """
        return pulumi.get(self, "domain_control_domains")

    @domain_control_domains.setter
    def domain_control_domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationDomainControlDomainArgs']]]]):
        pulumi.set(self, "domain_control_domains", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable log settings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB profile SaaS application name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="safeSearch")
    def safe_search(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable safe search. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "safe_search")

    @safe_search.setter
    def safe_search(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "safe_search", value)

    @property
    @pulumi.getter(name="safeSearchControls")
    def safe_search_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationSafeSearchControlArgs']]]]:
        """
        CASB profile safe search control. The structure of `safe_search_control` block is documented below.
        """
        return pulumi.get(self, "safe_search_controls")

    @safe_search_controls.setter
    def safe_search_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationSafeSearchControlArgs']]]]):
        pulumi.set(self, "safe_search_controls", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable setting. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="tenantControl")
    def tenant_control(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable tenant control. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "tenant_control")

    @tenant_control.setter
    def tenant_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_control", value)

    @property
    @pulumi.getter(name="tenantControlTenants")
    def tenant_control_tenants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationTenantControlTenantArgs']]]]:
        """
        CASB profile tenant control tenants. The structure of `tenant_control_tenants` block is documented below.
        """
        return pulumi.get(self, "tenant_control_tenants")

    @tenant_control_tenants.setter
    def tenant_control_tenants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationTenantControlTenantArgs']]]]):
        pulumi.set(self, "tenant_control_tenants", value)


@pulumi.input_type
class ProfileSaasApplicationAccessRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 bypass: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: CASB access rule action. Valid values: `bypass`, `block`, `monitor`.
        :param pulumi.Input[str] bypass: CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.
        :param pulumi.Input[str] name: CASB access rule activity name.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if bypass is not None:
            pulumi.set(__self__, "bypass", bypass)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        CASB access rule action. Valid values: `bypass`, `block`, `monitor`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def bypass(self) -> Optional[pulumi.Input[str]]:
        """
        CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.
        """
        return pulumi.get(self, "bypass")

    @bypass.setter
    def bypass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bypass", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB access rule activity name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProfileSaasApplicationCustomControlArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionArgs']]]] = None):
        """
        :param pulumi.Input[str] name: CASB custom control user activity name.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionArgs']]] options: CASB custom control option. The structure of `option` block is documented below.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB custom control user activity name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionArgs']]]]:
        """
        CASB custom control option. The structure of `option` block is documented below.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionArgs']]]]):
        pulumi.set(self, "options", value)


@pulumi.input_type
class ProfileSaasApplicationCustomControlOptionArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 user_inputs: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionUserInputArgs']]]] = None):
        """
        :param pulumi.Input[str] name: CASB custom control option name.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionUserInputArgs']]] user_inputs: CASB custom control user input. The structure of `user_input` block is documented below.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if user_inputs is not None:
            pulumi.set(__self__, "user_inputs", user_inputs)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB custom control option name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="userInputs")
    def user_inputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionUserInputArgs']]]]:
        """
        CASB custom control user input. The structure of `user_input` block is documented below.
        """
        return pulumi.get(self, "user_inputs")

    @user_inputs.setter
    def user_inputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileSaasApplicationCustomControlOptionUserInputArgs']]]]):
        pulumi.set(self, "user_inputs", value)


@pulumi.input_type
class ProfileSaasApplicationCustomControlOptionUserInputArgs:
    def __init__(__self__, *,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] value: user input value.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        user input value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ProfileSaasApplicationDomainControlDomainArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Domain control domain name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain control domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProfileSaasApplicationSafeSearchControlArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Safe search control name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Safe search control name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProfileSaasApplicationTenantControlTenantArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Tenant control tenants name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Tenant control tenants name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SaasapplicationDomainArgs:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] domain: Domain list separated by space.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain list separated by space.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)


@pulumi.input_type
class UseractivityControlOptionArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 operations: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: CASB control option name.
        :param pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationArgs']]] operations: CASB control option operations. The structure of `operations` block is documented below.
        :param pulumi.Input[str] status: CASB control option status. Valid values: `enable`, `disable`.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operations is not None:
            pulumi.set(__self__, "operations", operations)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB control option name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationArgs']]]]:
        """
        CASB control option operations. The structure of `operations` block is documented below.
        """
        return pulumi.get(self, "operations")

    @operations.setter
    def operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationArgs']]]]):
        pulumi.set(self, "operations", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        CASB control option status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class UseractivityControlOptionOperationArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 case_sensitive: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 header_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_key: Optional[pulumi.Input[str]] = None,
                 search_pattern: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 value_from_input: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationValueArgs']]]] = None):
        """
        :param pulumi.Input[str] action: CASB operation action. Valid values: `append`, `prepend`, `replace`, `new`, `new-on-not-found`, `delete`.
        :param pulumi.Input[str] case_sensitive: CASB operation search case sensitive. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] direction: CASB operation direction. Valid values: `request`.
        :param pulumi.Input[str] header_name: CASB operation header name to search.
        :param pulumi.Input[str] name: CASB control option operation name.
        :param pulumi.Input[str] search_key: CASB operation key to search.
        :param pulumi.Input[str] search_pattern: CASB operation search pattern. Valid values: `simple`, `substr`, `regexp`.
        :param pulumi.Input[str] target: CASB operation target. Valid values: `header`, `path`.
        :param pulumi.Input[str] value_from_input: Enable/disable value from user input. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationValueArgs']]] values: CASB operation new values. The structure of `values` block is documented below.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if case_sensitive is not None:
            pulumi.set(__self__, "case_sensitive", case_sensitive)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if header_name is not None:
            pulumi.set(__self__, "header_name", header_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if search_key is not None:
            pulumi.set(__self__, "search_key", search_key)
        if search_pattern is not None:
            pulumi.set(__self__, "search_pattern", search_pattern)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if value_from_input is not None:
            pulumi.set(__self__, "value_from_input", value_from_input)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation action. Valid values: `append`, `prepend`, `replace`, `new`, `new-on-not-found`, `delete`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="caseSensitive")
    def case_sensitive(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation search case sensitive. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "case_sensitive")

    @case_sensitive.setter
    def case_sensitive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "case_sensitive", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation direction. Valid values: `request`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation header name to search.
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB control option operation name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="searchKey")
    def search_key(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation key to search.
        """
        return pulumi.get(self, "search_key")

    @search_key.setter
    def search_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_key", value)

    @property
    @pulumi.getter(name="searchPattern")
    def search_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation search pattern. Valid values: `simple`, `substr`, `regexp`.
        """
        return pulumi.get(self, "search_pattern")

    @search_pattern.setter
    def search_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_pattern", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        CASB operation target. Valid values: `header`, `path`.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter(name="valueFromInput")
    def value_from_input(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable value from user input. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "value_from_input")

    @value_from_input.setter
    def value_from_input(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_from_input", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationValueArgs']]]]:
        """
        CASB operation new values. The structure of `values` block is documented below.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityControlOptionOperationValueArgs']]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class UseractivityControlOptionOperationValueArgs:
    def __init__(__self__, *,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] value: Operation value.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Operation value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class UseractivityMatchArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleArgs']]]] = None,
                 strategy: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] id: CASB user activity match rules ID.
        :param pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleArgs']]] rules: CASB user activity rules. The structure of `rules` block is documented below.
        :param pulumi.Input[str] strategy: CASB user activity rules strategy. Valid values: `and`, `or`.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if strategy is not None:
            pulumi.set(__self__, "strategy", strategy)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        CASB user activity match rules ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleArgs']]]]:
        """
        CASB user activity rules. The structure of `rules` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def strategy(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity rules strategy. Valid values: `and`, `or`.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strategy", value)


@pulumi.input_type
class UseractivityMatchRuleArgs:
    def __init__(__self__, *,
                 case_sensitive: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleDomainArgs']]]] = None,
                 header_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 match_pattern: Optional[pulumi.Input[str]] = None,
                 match_value: Optional[pulumi.Input[str]] = None,
                 methods: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleMethodArgs']]]] = None,
                 negate: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] case_sensitive: CASB user activity match case sensitive. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleDomainArgs']]] domains: CASB user activity domain list. The structure of `domains` block is documented below.
        :param pulumi.Input[str] header_name: CASB user activity rule header name.
        :param pulumi.Input[int] id: CASB user activity rule ID.
        :param pulumi.Input[str] match_pattern: CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.
        :param pulumi.Input[str] match_value: CASB user activity rule match value.
        :param pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleMethodArgs']]] methods: CASB user activity method list. The structure of `methods` block is documented below.
        :param pulumi.Input[str] negate: Enable/disable what the matching strategy must not be. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`.
        """
        if case_sensitive is not None:
            pulumi.set(__self__, "case_sensitive", case_sensitive)
        if domains is not None:
            pulumi.set(__self__, "domains", domains)
        if header_name is not None:
            pulumi.set(__self__, "header_name", header_name)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if match_pattern is not None:
            pulumi.set(__self__, "match_pattern", match_pattern)
        if match_value is not None:
            pulumi.set(__self__, "match_value", match_value)
        if methods is not None:
            pulumi.set(__self__, "methods", methods)
        if negate is not None:
            pulumi.set(__self__, "negate", negate)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="caseSensitive")
    def case_sensitive(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity match case sensitive. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "case_sensitive")

    @case_sensitive.setter
    def case_sensitive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "case_sensitive", value)

    @property
    @pulumi.getter
    def domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleDomainArgs']]]]:
        """
        CASB user activity domain list. The structure of `domains` block is documented below.
        """
        return pulumi.get(self, "domains")

    @domains.setter
    def domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleDomainArgs']]]]):
        pulumi.set(self, "domains", value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity rule header name.
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        CASB user activity rule ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="matchPattern")
    def match_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.
        """
        return pulumi.get(self, "match_pattern")

    @match_pattern.setter
    def match_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_pattern", value)

    @property
    @pulumi.getter(name="matchValue")
    def match_value(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity rule match value.
        """
        return pulumi.get(self, "match_value")

    @match_value.setter
    def match_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_value", value)

    @property
    @pulumi.getter
    def methods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleMethodArgs']]]]:
        """
        CASB user activity method list. The structure of `methods` block is documented below.
        """
        return pulumi.get(self, "methods")

    @methods.setter
    def methods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UseractivityMatchRuleMethodArgs']]]]):
        pulumi.set(self, "methods", value)

    @property
    @pulumi.getter
    def negate(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable what the matching strategy must not be. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "negate")

    @negate.setter
    def negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "negate", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class UseractivityMatchRuleDomainArgs:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] domain: Domain list separated by space.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain list separated by space.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)


@pulumi.input_type
class UseractivityMatchRuleMethodArgs:
    def __init__(__self__, *,
                 method: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] method: User activity method.
        """
        if method is not None:
            pulumi.set(__self__, "method", method)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        User activity method.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)


