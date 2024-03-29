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
    'GetInternetserviceextensionResult',
    'AwaitableGetInternetserviceextensionResult',
    'get_internetserviceextension',
    'get_internetserviceextension_output',
]

@pulumi.output_type
class GetInternetserviceextensionResult:
    """
    A collection of values returned by getInternetserviceextension.
    """
    def __init__(__self__, comment=None, disable_entries=None, entries=None, fosid=None, id=None, vdomparam=None):
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if disable_entries and not isinstance(disable_entries, list):
            raise TypeError("Expected argument 'disable_entries' to be a list")
        pulumi.set(__self__, "disable_entries", disable_entries)
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
    @pulumi.getter(name="disableEntries")
    def disable_entries(self) -> Sequence['outputs.GetInternetserviceextensionDisableEntryResult']:
        """
        Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
        """
        return pulumi.get(self, "disable_entries")

    @property
    @pulumi.getter
    def entries(self) -> Sequence['outputs.GetInternetserviceextensionEntryResult']:
        """
        Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Internet Service ID in the Internet Service database.
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
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetInternetserviceextensionResult(GetInternetserviceextensionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInternetserviceextensionResult(
            comment=self.comment,
            disable_entries=self.disable_entries,
            entries=self.entries,
            fosid=self.fosid,
            id=self.id,
            vdomparam=self.vdomparam)


def get_internetserviceextension(fosid: Optional[int] = None,
                                 vdomparam: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInternetserviceextensionResult:
    """
    Use this data source to get information on an fortios firewall internetserviceextension


    :param int fosid: Specify the fosid of the desired firewall internetserviceextension.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getInternetserviceextension:getInternetserviceextension', __args__, opts=opts, typ=GetInternetserviceextensionResult).value

    return AwaitableGetInternetserviceextensionResult(
        comment=pulumi.get(__ret__, 'comment'),
        disable_entries=pulumi.get(__ret__, 'disable_entries'),
        entries=pulumi.get(__ret__, 'entries'),
        fosid=pulumi.get(__ret__, 'fosid'),
        id=pulumi.get(__ret__, 'id'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_internetserviceextension)
def get_internetserviceextension_output(fosid: Optional[pulumi.Input[int]] = None,
                                        vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInternetserviceextensionResult]:
    """
    Use this data source to get information on an fortios firewall internetserviceextension


    :param int fosid: Specify the fosid of the desired firewall internetserviceextension.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
