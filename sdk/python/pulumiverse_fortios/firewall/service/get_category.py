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
    'GetCategoryResult',
    'AwaitableGetCategoryResult',
    'get_category',
    'get_category_output',
]

@pulumi.output_type
class GetCategoryResult:
    """
    A collection of values returned by getCategory.
    """
    def __init__(__self__, comment=None, fabric_object=None, id=None, name=None, vdomparam=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if fabric_object and not isinstance(fabric_object, str):
            raise TypeError("Expected argument 'fabric_object' to be a str")
        pulumi.set(__self__, "fabric_object", fabric_object)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

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
        Service category name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetCategoryResult(GetCategoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCategoryResult(
            comment=self.comment,
            fabric_object=self.fabric_object,
            id=self.id,
            name=self.name,
            vdomparam=self.vdomparam)


def get_category(name: Optional[str] = None,
                 vdomparam: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCategoryResult:
    """
    Use this data source to get information on an fortios firewallservice category


    :param str name: Specify the name of the desired firewallservice category.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/service/getCategory:getCategory', __args__, opts=opts, typ=GetCategoryResult).value

    return AwaitableGetCategoryResult(
        comment=pulumi.get(__ret__, 'comment'),
        fabric_object=pulumi.get(__ret__, 'fabric_object'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_category)
def get_category_output(name: Optional[pulumi.Input[str]] = None,
                        vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCategoryResult]:
    """
    Use this data source to get information on an fortios firewallservice category


    :param str name: Specify the name of the desired firewallservice category.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
