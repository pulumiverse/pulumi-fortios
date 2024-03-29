# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetMulticastResult',
    'AwaitableGetMulticastResult',
    'get_multicast',
    'get_multicast_output',
]

@pulumi.output_type
class GetMulticastResult:
    """
    A collection of values returned by getMulticast.
    """
    def __init__(__self__, id=None, interfaces=None, multicast_routing=None, pim_sm_globals=None, route_limit=None, route_threshold=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if multicast_routing and not isinstance(multicast_routing, str):
            raise TypeError("Expected argument 'multicast_routing' to be a str")
        pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_globals and not isinstance(pim_sm_globals, list):
            raise TypeError("Expected argument 'pim_sm_globals' to be a list")
        pulumi.set(__self__, "pim_sm_globals", pim_sm_globals)
        if route_limit and not isinstance(route_limit, int):
            raise TypeError("Expected argument 'route_limit' to be a int")
        pulumi.set(__self__, "route_limit", route_limit)
        if route_threshold and not isinstance(route_threshold, int):
            raise TypeError("Expected argument 'route_threshold' to be a int")
        pulumi.set(__self__, "route_threshold", route_threshold)
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
    def interfaces(self) -> Sequence['outputs.GetMulticastInterfaceResult']:
        """
        PIM interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> str:
        """
        Enable/disable IP multicast routing.
        """
        return pulumi.get(self, "multicast_routing")

    @property
    @pulumi.getter(name="pimSmGlobals")
    def pim_sm_globals(self) -> Sequence['outputs.GetMulticastPimSmGlobalResult']:
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_globals")

    @property
    @pulumi.getter(name="routeLimit")
    def route_limit(self) -> int:
        """
        Maximum number of multicast routes.
        """
        return pulumi.get(self, "route_limit")

    @property
    @pulumi.getter(name="routeThreshold")
    def route_threshold(self) -> int:
        """
        Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
        """
        return pulumi.get(self, "route_threshold")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetMulticastResult(GetMulticastResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMulticastResult(
            id=self.id,
            interfaces=self.interfaces,
            multicast_routing=self.multicast_routing,
            pim_sm_globals=self.pim_sm_globals,
            route_limit=self.route_limit,
            route_threshold=self.route_threshold,
            vdomparam=self.vdomparam)


def get_multicast(vdomparam: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMulticastResult:
    """
    Use this data source to get information on fortios router multicast


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:router/getMulticast:getMulticast', __args__, opts=opts, typ=GetMulticastResult).value

    return AwaitableGetMulticastResult(
        id=pulumi.get(__ret__, 'id'),
        interfaces=pulumi.get(__ret__, 'interfaces'),
        multicast_routing=pulumi.get(__ret__, 'multicast_routing'),
        pim_sm_globals=pulumi.get(__ret__, 'pim_sm_globals'),
        route_limit=pulumi.get(__ret__, 'route_limit'),
        route_threshold=pulumi.get(__ret__, 'route_threshold'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_multicast)
def get_multicast_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMulticastResult]:
    """
    Use this data source to get information on fortios router multicast


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
