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
    'GetInterfacelistResult',
    'AwaitableGetInterfacelistResult',
    'get_interfacelist',
    'get_interfacelist_output',
]

@pulumi.output_type
class GetInterfacelistResult:
    """
    A collection of values returned by getInterfacelist.
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
        A list of the `system.Interface`.
        """
        return pulumi.get(self, "namelists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetInterfacelistResult(GetInterfacelistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInterfacelistResult(
            filter=self.filter,
            id=self.id,
            namelists=self.namelists,
            vdomparam=self.vdomparam)


def get_interfacelist(filter: Optional[str] = None,
                      vdomparam: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInterfacelistResult:
    """
    Provides a list of `system.Interface`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.system.get_interfacelist(filter="name!=port1")
    pulumi.export("output1", data["fortios_system_interfacelist"]["sample2"]["namelist"])
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getInterfacelist:getInterfacelist', __args__, opts=opts, typ=GetInterfacelistResult).value

    return AwaitableGetInterfacelistResult(
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        namelists=pulumi.get(__ret__, 'namelists'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_interfacelist)
def get_interfacelist_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                             vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInterfacelistResult]:
    """
    Provides a list of `system.Interface`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.system.get_interfacelist(filter="name!=port1")
    pulumi.export("output1", data["fortios_system_interfacelist"]["sample2"]["namelist"])
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
