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
    'GetWccplistResult',
    'AwaitableGetWccplistResult',
    'get_wccplist',
    'get_wccplist_output',
]

@pulumi.output_type
class GetWccplistResult:
    """
    A collection of values returned by getWccplist.
    """
    def __init__(__self__, filter=None, id=None, service_idlists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_idlists and not isinstance(service_idlists, list):
            raise TypeError("Expected argument 'service_idlists' to be a list")
        pulumi.set(__self__, "service_idlists", service_idlists)
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
    @pulumi.getter(name="serviceIdlists")
    def service_idlists(self) -> Sequence[str]:
        """
        A list of the `system.Wccp`.
        """
        return pulumi.get(self, "service_idlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetWccplistResult(GetWccplistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWccplistResult(
            filter=self.filter,
            id=self.id,
            service_idlists=self.service_idlists,
            vdomparam=self.vdomparam)


def get_wccplist(filter: Optional[str] = None,
                 vdomparam: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWccplistResult:
    """
    Provides a list of `system.Wccp`.


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getWccplist:getWccplist', __args__, opts=opts, typ=GetWccplistResult).value

    return AwaitableGetWccplistResult(
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        service_idlists=pulumi.get(__ret__, 'service_idlists'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_wccplist)
def get_wccplist_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                        vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWccplistResult]:
    """
    Provides a list of `system.Wccp`.


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
