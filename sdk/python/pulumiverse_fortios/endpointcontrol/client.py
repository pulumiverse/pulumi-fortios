# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientArgs', 'Client']

@pulumi.input_type
class ClientArgs:
    def __init__(__self__, *,
                 ad_groups: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ftcl_uid: Optional[pulumi.Input[str]] = None,
                 info: Optional[pulumi.Input[str]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 src_mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Client resource.
        :param pulumi.Input[str] ad_groups: Endpoint client AD logon groups.
        :param pulumi.Input[int] fosid: Endpoint client ID.
        :param pulumi.Input[str] ftcl_uid: Endpoint FortiClient UID.
        :param pulumi.Input[str] info: Endpoint client information.
        :param pulumi.Input[str] src_ip: Endpoint client IP address.
        :param pulumi.Input[str] src_mac: Endpoint client MAC address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ad_groups is not None:
            pulumi.set(__self__, "ad_groups", ad_groups)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if ftcl_uid is not None:
            pulumi.set(__self__, "ftcl_uid", ftcl_uid)
        if info is not None:
            pulumi.set(__self__, "info", info)
        if src_ip is not None:
            pulumi.set(__self__, "src_ip", src_ip)
        if src_mac is not None:
            pulumi.set(__self__, "src_mac", src_mac)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="adGroups")
    def ad_groups(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client AD logon groups.
        """
        return pulumi.get(self, "ad_groups")

    @ad_groups.setter
    def ad_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ad_groups", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Endpoint client ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="ftclUid")
    def ftcl_uid(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint FortiClient UID.
        """
        return pulumi.get(self, "ftcl_uid")

    @ftcl_uid.setter
    def ftcl_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftcl_uid", value)

    @property
    @pulumi.getter
    def info(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client information.
        """
        return pulumi.get(self, "info")

    @info.setter
    def info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "info", value)

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client IP address.
        """
        return pulumi.get(self, "src_ip")

    @src_ip.setter
    def src_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_ip", value)

    @property
    @pulumi.getter(name="srcMac")
    def src_mac(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client MAC address.
        """
        return pulumi.get(self, "src_mac")

    @src_mac.setter
    def src_mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_mac", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _ClientState:
    def __init__(__self__, *,
                 ad_groups: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ftcl_uid: Optional[pulumi.Input[str]] = None,
                 info: Optional[pulumi.Input[str]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 src_mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Client resources.
        :param pulumi.Input[str] ad_groups: Endpoint client AD logon groups.
        :param pulumi.Input[int] fosid: Endpoint client ID.
        :param pulumi.Input[str] ftcl_uid: Endpoint FortiClient UID.
        :param pulumi.Input[str] info: Endpoint client information.
        :param pulumi.Input[str] src_ip: Endpoint client IP address.
        :param pulumi.Input[str] src_mac: Endpoint client MAC address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ad_groups is not None:
            pulumi.set(__self__, "ad_groups", ad_groups)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if ftcl_uid is not None:
            pulumi.set(__self__, "ftcl_uid", ftcl_uid)
        if info is not None:
            pulumi.set(__self__, "info", info)
        if src_ip is not None:
            pulumi.set(__self__, "src_ip", src_ip)
        if src_mac is not None:
            pulumi.set(__self__, "src_mac", src_mac)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="adGroups")
    def ad_groups(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client AD logon groups.
        """
        return pulumi.get(self, "ad_groups")

    @ad_groups.setter
    def ad_groups(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ad_groups", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Endpoint client ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="ftclUid")
    def ftcl_uid(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint FortiClient UID.
        """
        return pulumi.get(self, "ftcl_uid")

    @ftcl_uid.setter
    def ftcl_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftcl_uid", value)

    @property
    @pulumi.getter
    def info(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client information.
        """
        return pulumi.get(self, "info")

    @info.setter
    def info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "info", value)

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client IP address.
        """
        return pulumi.get(self, "src_ip")

    @src_ip.setter
    def src_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_ip", value)

    @property
    @pulumi.getter(name="srcMac")
    def src_mac(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint client MAC address.
        """
        return pulumi.get(self, "src_mac")

    @src_mac.setter
    def src_mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_mac", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Client(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_groups: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ftcl_uid: Optional[pulumi.Input[str]] = None,
                 info: Optional[pulumi.Input[str]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 src_mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure endpoint control client lists. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        EndpointControl Client can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ad_groups: Endpoint client AD logon groups.
        :param pulumi.Input[int] fosid: Endpoint client ID.
        :param pulumi.Input[str] ftcl_uid: Endpoint FortiClient UID.
        :param pulumi.Input[str] info: Endpoint client information.
        :param pulumi.Input[str] src_ip: Endpoint client IP address.
        :param pulumi.Input[str] src_mac: Endpoint client MAC address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ClientArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure endpoint control client lists. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        EndpointControl Client can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:endpointcontrol/client:Client labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ClientArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_groups: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ftcl_uid: Optional[pulumi.Input[str]] = None,
                 info: Optional[pulumi.Input[str]] = None,
                 src_ip: Optional[pulumi.Input[str]] = None,
                 src_mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientArgs.__new__(ClientArgs)

            __props__.__dict__["ad_groups"] = ad_groups
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["ftcl_uid"] = ftcl_uid
            __props__.__dict__["info"] = info
            __props__.__dict__["src_ip"] = src_ip
            __props__.__dict__["src_mac"] = src_mac
            __props__.__dict__["vdomparam"] = vdomparam
        super(Client, __self__).__init__(
            'fortios:endpointcontrol/client:Client',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ad_groups: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            ftcl_uid: Optional[pulumi.Input[str]] = None,
            info: Optional[pulumi.Input[str]] = None,
            src_ip: Optional[pulumi.Input[str]] = None,
            src_mac: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Client':
        """
        Get an existing Client resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ad_groups: Endpoint client AD logon groups.
        :param pulumi.Input[int] fosid: Endpoint client ID.
        :param pulumi.Input[str] ftcl_uid: Endpoint FortiClient UID.
        :param pulumi.Input[str] info: Endpoint client information.
        :param pulumi.Input[str] src_ip: Endpoint client IP address.
        :param pulumi.Input[str] src_mac: Endpoint client MAC address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientState.__new__(_ClientState)

        __props__.__dict__["ad_groups"] = ad_groups
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["ftcl_uid"] = ftcl_uid
        __props__.__dict__["info"] = info
        __props__.__dict__["src_ip"] = src_ip
        __props__.__dict__["src_mac"] = src_mac
        __props__.__dict__["vdomparam"] = vdomparam
        return Client(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adGroups")
    def ad_groups(self) -> pulumi.Output[Optional[str]]:
        """
        Endpoint client AD logon groups.
        """
        return pulumi.get(self, "ad_groups")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Endpoint client ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="ftclUid")
    def ftcl_uid(self) -> pulumi.Output[str]:
        """
        Endpoint FortiClient UID.
        """
        return pulumi.get(self, "ftcl_uid")

    @property
    @pulumi.getter
    def info(self) -> pulumi.Output[str]:
        """
        Endpoint client information.
        """
        return pulumi.get(self, "info")

    @property
    @pulumi.getter(name="srcIp")
    def src_ip(self) -> pulumi.Output[str]:
        """
        Endpoint client IP address.
        """
        return pulumi.get(self, "src_ip")

    @property
    @pulumi.getter(name="srcMac")
    def src_mac(self) -> pulumi.Output[str]:
        """
        Endpoint client MAC address.
        """
        return pulumi.get(self, "src_mac")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

