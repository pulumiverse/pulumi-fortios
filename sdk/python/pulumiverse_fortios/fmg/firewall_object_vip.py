# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FirewallObjectVipArgs', 'FirewallObjectVip']

@pulumi.input_type
class FirewallObjectVipArgs:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 config_default: Optional[pulumi.Input[str]] = None,
                 ext_intf: Optional[pulumi.Input[str]] = None,
                 ext_ip: Optional[pulumi.Input[str]] = None,
                 mapped_addr: Optional[pulumi.Input[str]] = None,
                 mapped_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallObjectVip resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_reply: Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] config_default: Enable/Disable config default value. enabled by default.
        :param pulumi.Input[str] ext_intf: Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        :param pulumi.Input[str] ext_ip: IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        :param pulumi.Input[str] mapped_addr: Mapped FQDN address name.
        :param pulumi.Input[str] mapped_ip: Mapped Ip address.
        :param pulumi.Input[str] name: Virtual IP name.
        :param pulumi.Input[str] type: Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if config_default is not None:
            pulumi.set(__self__, "config_default", config_default)
        if ext_intf is not None:
            pulumi.set(__self__, "ext_intf", ext_intf)
        if ext_ip is not None:
            pulumi.set(__self__, "ext_ip", ext_ip)
        if mapped_addr is not None:
            pulumi.set(__self__, "mapped_addr", mapped_addr)
        if mapped_ip is not None:
            pulumi.set(__self__, "mapped_ip", mapped_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        """
        Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        """
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="configDefault")
    def config_default(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/Disable config default value. enabled by default.
        """
        return pulumi.get(self, "config_default")

    @config_default.setter
    def config_default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_default", value)

    @property
    @pulumi.getter(name="extIntf")
    def ext_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        """
        return pulumi.get(self, "ext_intf")

    @ext_intf.setter
    def ext_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_intf", value)

    @property
    @pulumi.getter(name="extIp")
    def ext_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        """
        return pulumi.get(self, "ext_ip")

    @ext_ip.setter
    def ext_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_ip", value)

    @property
    @pulumi.getter(name="mappedAddr")
    def mapped_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Mapped FQDN address name.
        """
        return pulumi.get(self, "mapped_addr")

    @mapped_addr.setter
    def mapped_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapped_addr", value)

    @property
    @pulumi.getter(name="mappedIp")
    def mapped_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Mapped Ip address.
        """
        return pulumi.get(self, "mapped_ip")

    @mapped_ip.setter
    def mapped_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapped_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FirewallObjectVipState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 config_default: Optional[pulumi.Input[str]] = None,
                 ext_intf: Optional[pulumi.Input[str]] = None,
                 ext_ip: Optional[pulumi.Input[str]] = None,
                 mapped_addr: Optional[pulumi.Input[str]] = None,
                 mapped_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallObjectVip resources.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_reply: Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] config_default: Enable/Disable config default value. enabled by default.
        :param pulumi.Input[str] ext_intf: Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        :param pulumi.Input[str] ext_ip: IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        :param pulumi.Input[str] mapped_addr: Mapped FQDN address name.
        :param pulumi.Input[str] mapped_ip: Mapped Ip address.
        :param pulumi.Input[str] name: Virtual IP name.
        :param pulumi.Input[str] type: Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if config_default is not None:
            pulumi.set(__self__, "config_default", config_default)
        if ext_intf is not None:
            pulumi.set(__self__, "ext_intf", ext_intf)
        if ext_ip is not None:
            pulumi.set(__self__, "ext_ip", ext_ip)
        if mapped_addr is not None:
            pulumi.set(__self__, "mapped_addr", mapped_addr)
        if mapped_ip is not None:
            pulumi.set(__self__, "mapped_ip", mapped_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        """
        Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        """
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="configDefault")
    def config_default(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/Disable config default value. enabled by default.
        """
        return pulumi.get(self, "config_default")

    @config_default.setter
    def config_default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_default", value)

    @property
    @pulumi.getter(name="extIntf")
    def ext_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        """
        return pulumi.get(self, "ext_intf")

    @ext_intf.setter
    def ext_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_intf", value)

    @property
    @pulumi.getter(name="extIp")
    def ext_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        """
        return pulumi.get(self, "ext_ip")

    @ext_ip.setter
    def ext_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_ip", value)

    @property
    @pulumi.getter(name="mappedAddr")
    def mapped_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Mapped FQDN address name.
        """
        return pulumi.get(self, "mapped_addr")

    @mapped_addr.setter
    def mapped_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapped_addr", value)

    @property
    @pulumi.getter(name="mappedIp")
    def mapped_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Mapped Ip address.
        """
        return pulumi.get(self, "mapped_ip")

    @mapped_ip.setter
    def mapped_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapped_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FirewallObjectVip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 config_default: Optional[pulumi.Input[str]] = None,
                 ext_intf: Optional[pulumi.Input[str]] = None,
                 ext_ip: Optional[pulumi.Input[str]] = None,
                 mapped_addr: Optional[pulumi.Input[str]] = None,
                 mapped_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.FirewallObjectVip("test1",
            arp_reply="enable",
            comment="test obj vip",
            config_default="enable",
            ext_intf="any",
            ext_ip="2.2.2.2",
            mapped_ip="1.1.1.1",
            type="static-nat")
        test2 = fortios.fmg.FirewallObjectVip("test2",
            arp_reply="disable",
            comment="test obj vip",
            config_default="enable",
            ext_ip="2.2.2.2-2.2.2.100",
            mapped_addr="update.microsoft.com",
            type="fqdn")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_reply: Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] config_default: Enable/Disable config default value. enabled by default.
        :param pulumi.Input[str] ext_intf: Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        :param pulumi.Input[str] ext_ip: IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        :param pulumi.Input[str] mapped_addr: Mapped FQDN address name.
        :param pulumi.Input[str] mapped_ip: Mapped Ip address.
        :param pulumi.Input[str] name: Virtual IP name.
        :param pulumi.Input[str] type: Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallObjectVipArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.FirewallObjectVip("test1",
            arp_reply="enable",
            comment="test obj vip",
            config_default="enable",
            ext_intf="any",
            ext_ip="2.2.2.2",
            mapped_ip="1.1.1.1",
            type="static-nat")
        test2 = fortios.fmg.FirewallObjectVip("test2",
            arp_reply="disable",
            comment="test obj vip",
            config_default="enable",
            ext_ip="2.2.2.2-2.2.2.100",
            mapped_addr="update.microsoft.com",
            type="fqdn")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param FirewallObjectVipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallObjectVipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 config_default: Optional[pulumi.Input[str]] = None,
                 ext_intf: Optional[pulumi.Input[str]] = None,
                 ext_ip: Optional[pulumi.Input[str]] = None,
                 mapped_addr: Optional[pulumi.Input[str]] = None,
                 mapped_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallObjectVipArgs.__new__(FirewallObjectVipArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["arp_reply"] = arp_reply
            __props__.__dict__["comment"] = comment
            __props__.__dict__["config_default"] = config_default
            __props__.__dict__["ext_intf"] = ext_intf
            __props__.__dict__["ext_ip"] = ext_ip
            __props__.__dict__["mapped_addr"] = mapped_addr
            __props__.__dict__["mapped_ip"] = mapped_ip
            __props__.__dict__["name"] = name
            __props__.__dict__["type"] = type
        super(FirewallObjectVip, __self__).__init__(
            'fortios:fmg/firewallObjectVip:FirewallObjectVip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            arp_reply: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            config_default: Optional[pulumi.Input[str]] = None,
            ext_intf: Optional[pulumi.Input[str]] = None,
            ext_ip: Optional[pulumi.Input[str]] = None,
            mapped_addr: Optional[pulumi.Input[str]] = None,
            mapped_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FirewallObjectVip':
        """
        Get an existing FirewallObjectVip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_reply: Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] config_default: Enable/Disable config default value. enabled by default.
        :param pulumi.Input[str] ext_intf: Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        :param pulumi.Input[str] ext_ip: IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        :param pulumi.Input[str] mapped_addr: Mapped FQDN address name.
        :param pulumi.Input[str] mapped_ip: Mapped Ip address.
        :param pulumi.Input[str] name: Virtual IP name.
        :param pulumi.Input[str] type: Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallObjectVipState.__new__(_FirewallObjectVipState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["arp_reply"] = arp_reply
        __props__.__dict__["comment"] = comment
        __props__.__dict__["config_default"] = config_default
        __props__.__dict__["ext_intf"] = ext_intf
        __props__.__dict__["ext_ip"] = ext_ip
        __props__.__dict__["mapped_addr"] = mapped_addr
        __props__.__dict__["mapped_ip"] = mapped_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        return FirewallObjectVip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> pulumi.Output[Optional[str]]:
        """
        Enable to respond to ARP requests for this virtual IP address. Enabled by default.
        """
        return pulumi.get(self, "arp_reply")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="configDefault")
    def config_default(self) -> pulumi.Output[Optional[str]]:
        """
        Enable/Disable config default value. enabled by default.
        """
        return pulumi.get(self, "config_default")

    @property
    @pulumi.getter(name="extIntf")
    def ext_intf(self) -> pulumi.Output[Optional[str]]:
        """
        Interface connected to the source network that receives the packets that will be forwarded to the destination network.
        """
        return pulumi.get(self, "ext_intf")

    @property
    @pulumi.getter(name="extIp")
    def ext_ip(self) -> pulumi.Output[Optional[str]]:
        """
        IP address or address range on the external interface that you want to map to an address or address range on the destination network.
        """
        return pulumi.get(self, "ext_ip")

    @property
    @pulumi.getter(name="mappedAddr")
    def mapped_addr(self) -> pulumi.Output[Optional[str]]:
        """
        Mapped FQDN address name.
        """
        return pulumi.get(self, "mapped_addr")

    @property
    @pulumi.getter(name="mappedIp")
    def mapped_ip(self) -> pulumi.Output[Optional[str]]:
        """
        Mapped Ip address.
        """
        return pulumi.get(self, "mapped_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
        """
        return pulumi.get(self, "type")

