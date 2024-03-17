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
    'GetTosbasedpriorityResult',
    'AwaitableGetTosbasedpriorityResult',
    'get_tosbasedpriority',
    'get_tosbasedpriority_output',
]

@pulumi.output_type
class GetTosbasedpriorityResult:
    """
    A collection of values returned by getTosbasedpriority.
    """
    def __init__(__self__, fosid=None, id=None, priority=None, tos=None, vdomparam=None):
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if tos and not isinstance(tos, int):
            raise TypeError("Expected argument 'tos' to be a int")
        pulumi.set(__self__, "tos", tos)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Item ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def priority(self) -> str:
        """
        ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium).
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def tos(self) -> int:
        """
        Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
        """
        return pulumi.get(self, "tos")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetTosbasedpriorityResult(GetTosbasedpriorityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTosbasedpriorityResult(
            fosid=self.fosid,
            id=self.id,
            priority=self.priority,
            tos=self.tos,
            vdomparam=self.vdomparam)


def get_tosbasedpriority(fosid: Optional[int] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTosbasedpriorityResult:
    """
    Use this data source to get information on an fortios system tosbasedpriority


    :param int fosid: Specify the fosid of the desired system tosbasedpriority.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getTosbasedpriority:getTosbasedpriority', __args__, opts=opts, typ=GetTosbasedpriorityResult).value

    return AwaitableGetTosbasedpriorityResult(
        fosid=pulumi.get(__ret__, 'fosid'),
        id=pulumi.get(__ret__, 'id'),
        priority=pulumi.get(__ret__, 'priority'),
        tos=pulumi.get(__ret__, 'tos'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_tosbasedpriority)
def get_tosbasedpriority_output(fosid: Optional[pulumi.Input[int]] = None,
                                vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTosbasedpriorityResult]:
    """
    Use this data source to get information on an fortios system tosbasedpriority


    :param int fosid: Specify the fosid of the desired system tosbasedpriority.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
