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
    'GetOnetimeResult',
    'AwaitableGetOnetimeResult',
    'get_onetime',
    'get_onetime_output',
]

@pulumi.output_type
class GetOnetimeResult:
    """
    A collection of values returned by getOnetime.
    """
    def __init__(__self__, color=None, end=None, end_utc=None, expiration_days=None, fabric_object=None, id=None, name=None, start=None, start_utc=None, vdomparam=None):
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if end and not isinstance(end, str):
            raise TypeError("Expected argument 'end' to be a str")
        pulumi.set(__self__, "end", end)
        if end_utc and not isinstance(end_utc, str):
            raise TypeError("Expected argument 'end_utc' to be a str")
        pulumi.set(__self__, "end_utc", end_utc)
        if expiration_days and not isinstance(expiration_days, int):
            raise TypeError("Expected argument 'expiration_days' to be a int")
        pulumi.set(__self__, "expiration_days", expiration_days)
        if fabric_object and not isinstance(fabric_object, str):
            raise TypeError("Expected argument 'fabric_object' to be a str")
        pulumi.set(__self__, "fabric_object", fabric_object)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start and not isinstance(start, str):
            raise TypeError("Expected argument 'start' to be a str")
        pulumi.set(__self__, "start", start)
        if start_utc and not isinstance(start_utc, str):
            raise TypeError("Expected argument 'start_utc' to be a str")
        pulumi.set(__self__, "start_utc", start_utc)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def end(self) -> str:
        """
        Schedule end date and time, format hh:mm yyyy/mm/dd.
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter(name="endUtc")
    def end_utc(self) -> str:
        """
        Schedule end date and time, in epoch format.
        """
        return pulumi.get(self, "end_utc")

    @property
    @pulumi.getter(name="expirationDays")
    def expiration_days(self) -> int:
        """
        Write an event log message this many days before the schedule expires.
        """
        return pulumi.get(self, "expiration_days")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> str:
        """
        Security Fabric global object setting.
        """
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Onetime schedule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def start(self) -> str:
        """
        Schedule start date and time, format hh:mm yyyy/mm/dd.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter(name="startUtc")
    def start_utc(self) -> str:
        """
        Schedule start date and time, in epoch format.
        """
        return pulumi.get(self, "start_utc")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetOnetimeResult(GetOnetimeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOnetimeResult(
            color=self.color,
            end=self.end,
            end_utc=self.end_utc,
            expiration_days=self.expiration_days,
            fabric_object=self.fabric_object,
            id=self.id,
            name=self.name,
            start=self.start,
            start_utc=self.start_utc,
            vdomparam=self.vdomparam)


def get_onetime(name: Optional[str] = None,
                vdomparam: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOnetimeResult:
    """
    Use this data source to get information on an fortios firewallschedule onetime


    :param str name: Specify the name of the desired firewallschedule onetime.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/schedule/getOnetime:getOnetime', __args__, opts=opts, typ=GetOnetimeResult).value

    return AwaitableGetOnetimeResult(
        color=pulumi.get(__ret__, 'color'),
        end=pulumi.get(__ret__, 'end'),
        end_utc=pulumi.get(__ret__, 'end_utc'),
        expiration_days=pulumi.get(__ret__, 'expiration_days'),
        fabric_object=pulumi.get(__ret__, 'fabric_object'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        start=pulumi.get(__ret__, 'start'),
        start_utc=pulumi.get(__ret__, 'start_utc'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_onetime)
def get_onetime_output(name: Optional[pulumi.Input[str]] = None,
                       vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOnetimeResult]:
    """
    Use this data source to get information on an fortios firewallschedule onetime


    :param str name: Specify the name of the desired firewallschedule onetime.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
