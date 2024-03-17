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
    'GetAutoscriptResult',
    'AwaitableGetAutoscriptResult',
    'get_autoscript',
    'get_autoscript_output',
]

@pulumi.output_type
class GetAutoscriptResult:
    """
    A collection of values returned by getAutoscript.
    """
    def __init__(__self__, id=None, interval=None, name=None, output_size=None, repeat=None, script=None, start=None, timeout=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interval and not isinstance(interval, int):
            raise TypeError("Expected argument 'interval' to be a int")
        pulumi.set(__self__, "interval", interval)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_size and not isinstance(output_size, int):
            raise TypeError("Expected argument 'output_size' to be a int")
        pulumi.set(__self__, "output_size", output_size)
        if repeat and not isinstance(repeat, int):
            raise TypeError("Expected argument 'repeat' to be a int")
        pulumi.set(__self__, "repeat", repeat)
        if script and not isinstance(script, str):
            raise TypeError("Expected argument 'script' to be a str")
        pulumi.set(__self__, "script", script)
        if start and not isinstance(start, str):
            raise TypeError("Expected argument 'start' to be a str")
        pulumi.set(__self__, "start", start)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> int:
        """
        Repeat interval in seconds.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Auto script name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputSize")
    def output_size(self) -> int:
        """
        Number of megabytes to limit script output to (10 - 1024, default = 10).
        """
        return pulumi.get(self, "output_size")

    @property
    @pulumi.getter
    def repeat(self) -> int:
        """
        Number of times to repeat this script (0 = infinite).
        """
        return pulumi.get(self, "repeat")

    @property
    @pulumi.getter
    def script(self) -> str:
        """
        List of FortiOS CLI commands to repeat.
        """
        return pulumi.get(self, "script")

    @property
    @pulumi.getter
    def start(self) -> str:
        """
        Script starting mode.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def timeout(self) -> int:
        """
        Maximum running time for this script in seconds (0 = no timeout).
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetAutoscriptResult(GetAutoscriptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAutoscriptResult(
            id=self.id,
            interval=self.interval,
            name=self.name,
            output_size=self.output_size,
            repeat=self.repeat,
            script=self.script,
            start=self.start,
            timeout=self.timeout,
            vdomparam=self.vdomparam)


def get_autoscript(name: Optional[str] = None,
                   vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAutoscriptResult:
    """
    Use this data source to get information on an fortios system autoscript


    :param str name: Specify the name of the desired system autoscript.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getAutoscript:getAutoscript', __args__, opts=opts, typ=GetAutoscriptResult).value

    return AwaitableGetAutoscriptResult(
        id=pulumi.get(__ret__, 'id'),
        interval=pulumi.get(__ret__, 'interval'),
        name=pulumi.get(__ret__, 'name'),
        output_size=pulumi.get(__ret__, 'output_size'),
        repeat=pulumi.get(__ret__, 'repeat'),
        script=pulumi.get(__ret__, 'script'),
        start=pulumi.get(__ret__, 'start'),
        timeout=pulumi.get(__ret__, 'timeout'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_autoscript)
def get_autoscript_output(name: Optional[pulumi.Input[str]] = None,
                          vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAutoscriptResult]:
    """
    Use this data source to get information on an fortios system autoscript


    :param str name: Specify the name of the desired system autoscript.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
