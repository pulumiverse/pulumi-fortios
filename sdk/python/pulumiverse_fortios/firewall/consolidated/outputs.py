# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'PolicyAppCategory',
    'PolicyAppGroup',
    'PolicyApplication',
    'PolicyDstaddr4',
    'PolicyDstaddr6',
    'PolicyDstintf',
    'PolicyFssoGroup',
    'PolicyGroup',
    'PolicyInternetServiceCustom',
    'PolicyInternetServiceCustomGroup',
    'PolicyInternetServiceGroup',
    'PolicyInternetServiceId',
    'PolicyInternetServiceName',
    'PolicyInternetServiceSrcCustom',
    'PolicyInternetServiceSrcCustomGroup',
    'PolicyInternetServiceSrcGroup',
    'PolicyInternetServiceSrcId',
    'PolicyInternetServiceSrcName',
    'PolicyPoolname4',
    'PolicyPoolname6',
    'PolicyService',
    'PolicySrcaddr4',
    'PolicySrcaddr6',
    'PolicySrcintf',
    'PolicyUrlCategory',
    'PolicyUser',
    'GetPolicyAppCategoryResult',
    'GetPolicyAppGroupResult',
    'GetPolicyApplicationResult',
    'GetPolicyDstaddr4Result',
    'GetPolicyDstaddr6Result',
    'GetPolicyDstintfResult',
    'GetPolicyFssoGroupResult',
    'GetPolicyGroupResult',
    'GetPolicyInternetServiceCustomResult',
    'GetPolicyInternetServiceCustomGroupResult',
    'GetPolicyInternetServiceGroupResult',
    'GetPolicyInternetServiceIdResult',
    'GetPolicyInternetServiceNameResult',
    'GetPolicyInternetServiceSrcCustomResult',
    'GetPolicyInternetServiceSrcCustomGroupResult',
    'GetPolicyInternetServiceSrcGroupResult',
    'GetPolicyInternetServiceSrcIdResult',
    'GetPolicyInternetServiceSrcNameResult',
    'GetPolicyPoolname4Result',
    'GetPolicyPoolname6Result',
    'GetPolicyServiceResult',
    'GetPolicySrcaddr4Result',
    'GetPolicySrcaddr6Result',
    'GetPolicySrcintfResult',
    'GetPolicyUrlCategoryResult',
    'GetPolicyUserResult',
]

@pulumi.output_type
class PolicyAppCategory(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None):
        """
        :param int id: Category IDs.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Category IDs.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PolicyAppGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Application group names.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyApplication(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None):
        """
        :param int id: Application IDs.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Application IDs.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PolicyDstaddr4(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyDstaddr6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyDstintf(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyFssoGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Names of FSSO groups.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Names of FSSO groups.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Group name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceCustom(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Custom Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Custom Internet Service name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceCustomGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Custom Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Custom Internet Service group name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Internet Service group name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceId(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None):
        """
        :param int id: Internet Service ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PolicyInternetServiceName(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceSrcCustom(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Custom Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Custom Internet Service name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceSrcCustomGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Custom Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Custom Internet Service group name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceSrcGroup(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Internet Service group name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyInternetServiceSrcId(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None):
        """
        :param int id: Internet Service ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PolicyInternetServiceSrcName(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyPoolname4(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyPoolname6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyService(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Service name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicySrcaddr4(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicySrcaddr6(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicySrcintf(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Interface name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Interface name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class PolicyUrlCategory(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None):
        """
        :param int id: URL category ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PolicyUser(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: IPv6 pool name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        IPv6 pool name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyAppCategoryResult(dict):
    def __init__(__self__, *,
                 id: int):
        """
        :param int id: URL category ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPolicyAppGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyApplicationResult(dict):
    def __init__(__self__, *,
                 id: int):
        """
        :param int id: URL category ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPolicyDstaddr4Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyDstaddr6Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyDstintfResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyFssoGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceCustomResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceCustomGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceIdResult(dict):
    def __init__(__self__, *,
                 id: int):
        """
        :param int id: URL category ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPolicyInternetServiceNameResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceSrcCustomResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceSrcCustomGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceSrcGroupResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyInternetServiceSrcIdResult(dict):
    def __init__(__self__, *,
                 id: int):
        """
        :param int id: URL category ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPolicyInternetServiceSrcNameResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyPoolname4Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyPoolname6Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyServiceResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicySrcaddr4Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicySrcaddr6Result(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicySrcintfResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPolicyUrlCategoryResult(dict):
    def __init__(__self__, *,
                 id: int):
        """
        :param int id: URL category ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPolicyUserResult(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Application group names.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Application group names.
        """
        return pulumi.get(self, "name")


