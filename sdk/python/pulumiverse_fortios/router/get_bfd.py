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
    'GetBfdResult',
    'AwaitableGetBfdResult',
    'get_bfd',
    'get_bfd_output',
]

@pulumi.output_type
class GetBfdResult:
    """
    A collection of values returned by getBfd.
    """
    def __init__(__self__, id=None, multihop_templates=None, neighbors=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if multihop_templates and not isinstance(multihop_templates, list):
            raise TypeError("Expected argument 'multihop_templates' to be a list")
        pulumi.set(__self__, "multihop_templates", multihop_templates)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
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
    @pulumi.getter(name="multihopTemplates")
    def multihop_templates(self) -> Sequence['outputs.GetBfdMultihopTemplateResult']:
        """
        BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        """
        return pulumi.get(self, "multihop_templates")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetBfdNeighborResult']:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetBfdResult(GetBfdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBfdResult(
            id=self.id,
            multihop_templates=self.multihop_templates,
            neighbors=self.neighbors,
            vdomparam=self.vdomparam)


def get_bfd(vdomparam: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBfdResult:
    """
    Use this data source to get information on fortios router bfd


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:router/getBfd:getBfd', __args__, opts=opts, typ=GetBfdResult).value

    return AwaitableGetBfdResult(
        id=pulumi.get(__ret__, 'id'),
        multihop_templates=pulumi.get(__ret__, 'multihop_templates'),
        neighbors=pulumi.get(__ret__, 'neighbors'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_bfd)
def get_bfd_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBfdResult]:
    """
    Use this data source to get information on fortios router bfd


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
