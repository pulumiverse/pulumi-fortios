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
    'GetZoneResult',
    'AwaitableGetZoneResult',
    'get_zone',
    'get_zone_output',
]

@pulumi.output_type
class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, description=None, id=None, interfaces=None, intrazone=None, name=None, taggings=None, vdomparam=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if intrazone and not isinstance(intrazone, str):
            raise TypeError("Expected argument 'intrazone' to be a str")
        pulumi.set(__self__, "intrazone", intrazone)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetZoneInterfaceResult']:
        """
        Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def intrazone(self) -> str:
        """
        Allow or deny traffic routing between different interfaces in the same zone (default = deny).
        """
        return pulumi.get(self, "intrazone")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetZoneTaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            description=self.description,
            id=self.id,
            interfaces=self.interfaces,
            intrazone=self.intrazone,
            name=self.name,
            taggings=self.taggings,
            vdomparam=self.vdomparam)


def get_zone(name: Optional[str] = None,
             vdomparam: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneResult:
    """
    Use this data source to get information on an fortios system zone


    :param str name: Specify the name of the desired system zone.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getZone:getZone', __args__, opts=opts, typ=GetZoneResult).value

    return AwaitableGetZoneResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        interfaces=pulumi.get(__ret__, 'interfaces'),
        intrazone=pulumi.get(__ret__, 'intrazone'),
        name=pulumi.get(__ret__, 'name'),
        taggings=pulumi.get(__ret__, 'taggings'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_zone)
def get_zone_output(name: Optional[pulumi.Input[str]] = None,
                    vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneResult]:
    """
    Use this data source to get information on an fortios system zone


    :param str name: Specify the name of the desired system zone.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
