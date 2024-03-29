# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ObjectAddressgroupArgs', 'ObjectAddressgroup']

@pulumi.input_type
class ObjectAddressgroupArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectAddressgroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Address objects contained within the group.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] name: Address group name.
        """
        pulumi.set(__self__, "members", members)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Address objects contained within the group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ObjectAddressgroupState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectAddressgroup resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Address objects contained within the group.
        :param pulumi.Input[str] name: Address group name.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Address objects contained within the group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ObjectAddressgroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to configure firewall address group used in firewall policies of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `firewall.Addrgrp`, we recommend that you use the new resource.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        s1 = fortios.firewall.ObjectAddressgroup("s1",
            comment="dfdsad",
            members=[
                "google-play",
                "swscan.apple.com",
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Address objects contained within the group.
        :param pulumi.Input[str] name: Address group name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObjectAddressgroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure firewall address group used in firewall policies of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `firewall.Addrgrp`, we recommend that you use the new resource.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        s1 = fortios.firewall.ObjectAddressgroup("s1",
            comment="dfdsad",
            members=[
                "google-play",
                "swscan.apple.com",
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ObjectAddressgroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectAddressgroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectAddressgroupArgs.__new__(ObjectAddressgroupArgs)

            __props__.__dict__["comment"] = comment
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
        super(ObjectAddressgroup, __self__).__init__(
            'fortios:firewall/objectAddressgroup:ObjectAddressgroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ObjectAddressgroup':
        """
        Get an existing ObjectAddressgroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Address objects contained within the group.
        :param pulumi.Input[str] name: Address group name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectAddressgroupState.__new__(_ObjectAddressgroupState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        return ObjectAddressgroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        Address objects contained within the group.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Address group name.
        """
        return pulumi.get(self, "name")

