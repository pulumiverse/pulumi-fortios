# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemDscpbasedpriorityResult',
    'AwaitableGetSystemDscpbasedpriorityResult',
    'get_system_dscpbasedpriority',
    'get_system_dscpbasedpriority_output',
]

@pulumi.output_type
class GetSystemDscpbasedpriorityResult:
    """
    A collection of values returned by getSystemDscpbasedpriority.
    """
    def __init__(__self__, ds=None, fosid=None, id=None, priority=None, vdomparam=None):
        if ds and not isinstance(ds, int):
            raise TypeError("Expected argument 'ds' to be a int")
        pulumi.set(__self__, "ds", ds)
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ds(self) -> int:
        """
        DSCP(DiffServ) DS value (0 - 63).
        """
        return pulumi.get(self, "ds")

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
        DSCP based priority level.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemDscpbasedpriorityResult(GetSystemDscpbasedpriorityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemDscpbasedpriorityResult(
            ds=self.ds,
            fosid=self.fosid,
            id=self.id,
            priority=self.priority,
            vdomparam=self.vdomparam)


def get_system_dscpbasedpriority(fosid: Optional[int] = None,
                                 vdomparam: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemDscpbasedpriorityResult:
    """
    Use this data source to get information on an fortios system dscpbasedpriority


    :param int fosid: Specify the fosid of the desired system dscpbasedpriority.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemDscpbasedpriority:getSystemDscpbasedpriority', __args__, opts=opts, typ=GetSystemDscpbasedpriorityResult).value

    return AwaitableGetSystemDscpbasedpriorityResult(
        ds=__ret__.ds,
        fosid=__ret__.fosid,
        id=__ret__.id,
        priority=__ret__.priority,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_dscpbasedpriority)
def get_system_dscpbasedpriority_output(fosid: Optional[pulumi.Input[int]] = None,
                                        vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemDscpbasedpriorityResult]:
    """
    Use this data source to get information on an fortios system dscpbasedpriority


    :param int fosid: Specify the fosid of the desired system dscpbasedpriority.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
