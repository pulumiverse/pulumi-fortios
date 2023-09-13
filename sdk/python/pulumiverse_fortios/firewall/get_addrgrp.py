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
    'GetAddrgrpResult',
    'AwaitableGetAddrgrpResult',
    'get_addrgrp',
    'get_addrgrp_output',
]

@pulumi.output_type
class GetAddrgrpResult:
    """
    A collection of values returned by getAddrgrp.
    """
    def __init__(__self__, allow_routing=None, category=None, color=None, comment=None, exclude=None, exclude_members=None, fabric_object=None, id=None, members=None, name=None, taggings=None, type=None, uuid=None, vdomparam=None, visibility=None):
        if allow_routing and not isinstance(allow_routing, str):
            raise TypeError("Expected argument 'allow_routing' to be a str")
        pulumi.set(__self__, "allow_routing", allow_routing)
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if exclude and not isinstance(exclude, str):
            raise TypeError("Expected argument 'exclude' to be a str")
        pulumi.set(__self__, "exclude", exclude)
        if exclude_members and not isinstance(exclude_members, list):
            raise TypeError("Expected argument 'exclude_members' to be a list")
        pulumi.set(__self__, "exclude_members", exclude_members)
        if fabric_object and not isinstance(fabric_object, str):
            raise TypeError("Expected argument 'fabric_object' to be a str")
        pulumi.set(__self__, "fabric_object", fabric_object)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="allowRouting")
    def allow_routing(self) -> str:
        """
        Enable/disable use of this group in the static route configuration.
        """
        return pulumi.get(self, "allow_routing")

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        Tag category.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def exclude(self) -> str:
        """
        Enable/disable address exclusion.
        """
        return pulumi.get(self, "exclude")

    @property
    @pulumi.getter(name="excludeMembers")
    def exclude_members(self) -> Sequence['outputs.GetAddrgrpExcludeMemberResult']:
        """
        Address exclusion member. The structure of `exclude_member` block is documented below.
        """
        return pulumi.get(self, "exclude_members")

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
    def members(self) -> Sequence['outputs.GetAddrgrpMemberResult']:
        """
        Address objects contained within the group. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetAddrgrpTaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Address group type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable address visibility in the GUI.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetAddrgrpResult(GetAddrgrpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddrgrpResult(
            allow_routing=self.allow_routing,
            category=self.category,
            color=self.color,
            comment=self.comment,
            exclude=self.exclude,
            exclude_members=self.exclude_members,
            fabric_object=self.fabric_object,
            id=self.id,
            members=self.members,
            name=self.name,
            taggings=self.taggings,
            type=self.type,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_addrgrp(name: Optional[str] = None,
                vdomparam: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddrgrpResult:
    """
    Use this data source to get information on an fortios firewall addrgrp


    :param str name: Specify the name of the desired firewall addrgrp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getAddrgrp:getAddrgrp', __args__, opts=opts, typ=GetAddrgrpResult).value

    return AwaitableGetAddrgrpResult(
        allow_routing=__ret__.allow_routing,
        category=__ret__.category,
        color=__ret__.color,
        comment=__ret__.comment,
        exclude=__ret__.exclude,
        exclude_members=__ret__.exclude_members,
        fabric_object=__ret__.fabric_object,
        id=__ret__.id,
        members=__ret__.members,
        name=__ret__.name,
        taggings=__ret__.taggings,
        type=__ret__.type,
        uuid=__ret__.uuid,
        vdomparam=__ret__.vdomparam,
        visibility=__ret__.visibility)


@_utilities.lift_output_func(get_addrgrp)
def get_addrgrp_output(name: Optional[pulumi.Input[str]] = None,
                       vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddrgrpResult]:
    """
    Use this data source to get information on an fortios firewall addrgrp


    :param str name: Specify the name of the desired firewall addrgrp.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
