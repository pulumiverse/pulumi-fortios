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
    'GetMobiletunnellistResult',
    'AwaitableGetMobiletunnellistResult',
    'get_mobiletunnellist',
    'get_mobiletunnellist_output',
]

@pulumi.output_type
class GetMobiletunnellistResult:
    """
    A collection of values returned by getMobiletunnellist.
    """
    def __init__(__self__, filter=None, id=None, namelists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namelists and not isinstance(namelists, list):
            raise TypeError("Expected argument 'namelists' to be a list")
        pulumi.set(__self__, "namelists", namelists)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namelists(self) -> Sequence[str]:
        """
        A list of the `system.Mobiletunnel`.
        """
        return pulumi.get(self, "namelists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetMobiletunnellistResult(GetMobiletunnellistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMobiletunnellistResult(
            filter=self.filter,
            id=self.id,
            namelists=self.namelists,
            vdomparam=self.vdomparam)


def get_mobiletunnellist(filter: Optional[str] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMobiletunnellistResult:
    """
    Provides a list of `system.Mobiletunnel`.


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getMobiletunnellist:getMobiletunnellist', __args__, opts=opts, typ=GetMobiletunnellistResult).value

    return AwaitableGetMobiletunnellistResult(
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        namelists=pulumi.get(__ret__, 'namelists'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_mobiletunnellist)
def get_mobiletunnellist_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMobiletunnellistResult]:
    """
    Provides a list of `system.Mobiletunnel`.


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
