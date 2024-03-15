# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RemotestorageArgs', 'Remotestorage']

@pulumi.input_type
class RemotestorageArgs:
    def __init__(__self__, *,
                 local_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Remotestorage resource.
        :param pulumi.Input[str] local_cache_id: ID that this device uses to connect to the remote device.
        :param pulumi.Input[str] remote_cache_id: ID of the remote device to which the device connects.
        :param pulumi.Input[str] remote_cache_ip: IP address of the remote device to which the device connects.
        :param pulumi.Input[str] status: Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if local_cache_id is not None:
            pulumi.set(__self__, "local_cache_id", local_cache_id)
        if remote_cache_id is not None:
            pulumi.set(__self__, "remote_cache_id", remote_cache_id)
        if remote_cache_ip is not None:
            pulumi.set(__self__, "remote_cache_ip", remote_cache_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="localCacheId")
    def local_cache_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID that this device uses to connect to the remote device.
        """
        return pulumi.get(self, "local_cache_id")

    @local_cache_id.setter
    def local_cache_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_cache_id", value)

    @property
    @pulumi.getter(name="remoteCacheId")
    def remote_cache_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_id")

    @remote_cache_id.setter
    def remote_cache_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_cache_id", value)

    @property
    @pulumi.getter(name="remoteCacheIp")
    def remote_cache_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_ip")

    @remote_cache_ip.setter
    def remote_cache_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_cache_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _RemotestorageState:
    def __init__(__self__, *,
                 local_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Remotestorage resources.
        :param pulumi.Input[str] local_cache_id: ID that this device uses to connect to the remote device.
        :param pulumi.Input[str] remote_cache_id: ID of the remote device to which the device connects.
        :param pulumi.Input[str] remote_cache_ip: IP address of the remote device to which the device connects.
        :param pulumi.Input[str] status: Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if local_cache_id is not None:
            pulumi.set(__self__, "local_cache_id", local_cache_id)
        if remote_cache_id is not None:
            pulumi.set(__self__, "remote_cache_id", remote_cache_id)
        if remote_cache_ip is not None:
            pulumi.set(__self__, "remote_cache_ip", remote_cache_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="localCacheId")
    def local_cache_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID that this device uses to connect to the remote device.
        """
        return pulumi.get(self, "local_cache_id")

    @local_cache_id.setter
    def local_cache_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_cache_id", value)

    @property
    @pulumi.getter(name="remoteCacheId")
    def remote_cache_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_id")

    @remote_cache_id.setter
    def remote_cache_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_cache_id", value)

    @property
    @pulumi.getter(name="remoteCacheIp")
    def remote_cache_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_ip")

    @remote_cache_ip.setter
    def remote_cache_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_cache_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class Remotestorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 local_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure a remote cache device as Web cache storage.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Remotestorage("trname",
            remote_cache_ip="0.0.0.0",
            status="disable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Wanopt RemoteStorage can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] local_cache_id: ID that this device uses to connect to the remote device.
        :param pulumi.Input[str] remote_cache_id: ID of the remote device to which the device connects.
        :param pulumi.Input[str] remote_cache_ip: IP address of the remote device to which the device connects.
        :param pulumi.Input[str] status: Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RemotestorageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure a remote cache device as Web cache storage.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Remotestorage("trname",
            remote_cache_ip="0.0.0.0",
            status="disable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Wanopt RemoteStorage can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param RemotestorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemotestorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 local_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_id: Optional[pulumi.Input[str]] = None,
                 remote_cache_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RemotestorageArgs.__new__(RemotestorageArgs)

            __props__.__dict__["local_cache_id"] = local_cache_id
            __props__.__dict__["remote_cache_id"] = remote_cache_id
            __props__.__dict__["remote_cache_ip"] = remote_cache_ip
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Remotestorage, __self__).__init__(
            'fortios:wanopt/remotestorage:Remotestorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            local_cache_id: Optional[pulumi.Input[str]] = None,
            remote_cache_id: Optional[pulumi.Input[str]] = None,
            remote_cache_ip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Remotestorage':
        """
        Get an existing Remotestorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] local_cache_id: ID that this device uses to connect to the remote device.
        :param pulumi.Input[str] remote_cache_id: ID of the remote device to which the device connects.
        :param pulumi.Input[str] remote_cache_ip: IP address of the remote device to which the device connects.
        :param pulumi.Input[str] status: Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RemotestorageState.__new__(_RemotestorageState)

        __props__.__dict__["local_cache_id"] = local_cache_id
        __props__.__dict__["remote_cache_id"] = remote_cache_id
        __props__.__dict__["remote_cache_ip"] = remote_cache_ip
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Remotestorage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="localCacheId")
    def local_cache_id(self) -> pulumi.Output[str]:
        """
        ID that this device uses to connect to the remote device.
        """
        return pulumi.get(self, "local_cache_id")

    @property
    @pulumi.getter(name="remoteCacheId")
    def remote_cache_id(self) -> pulumi.Output[str]:
        """
        ID of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_id")

    @property
    @pulumi.getter(name="remoteCacheIp")
    def remote_cache_ip(self) -> pulumi.Output[str]:
        """
        IP address of the remote device to which the device connects.
        """
        return pulumi.get(self, "remote_cache_ip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

