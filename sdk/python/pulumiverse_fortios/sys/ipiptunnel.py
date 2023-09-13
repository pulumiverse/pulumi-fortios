# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IpiptunnelArgs', 'Ipiptunnel']

@pulumi.input_type
class IpiptunnelArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 local_gw: pulumi.Input[str],
                 remote_gw: pulumi.Input[str],
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 use_sdwan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ipiptunnel resource.
        :param pulumi.Input[str] interface: Interface name that is associated with the incoming traffic from available options.
        :param pulumi.Input[str] local_gw: IPv4 address for the local gateway.
        :param pulumi.Input[str] remote_gw: IPv4 address for the remote gateway.
        :param pulumi.Input[str] auto_asic_offload: Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: IPIP Tunnel name.
        :param pulumi.Input[str] use_sdwan: Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "local_gw", local_gw)
        pulumi.set(__self__, "remote_gw", remote_gw)
        if auto_asic_offload is not None:
            pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if use_sdwan is not None:
            pulumi.set(__self__, "use_sdwan", use_sdwan)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Interface name that is associated with the incoming traffic from available options.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> pulumi.Input[str]:
        """
        IPv4 address for the local gateway.
        """
        return pulumi.get(self, "local_gw")

    @local_gw.setter
    def local_gw(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_gw", value)

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> pulumi.Input[str]:
        """
        IPv4 address for the remote gateway.
        """
        return pulumi.get(self, "remote_gw")

    @remote_gw.setter
    def remote_gw(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_gw", value)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_asic_offload")

    @auto_asic_offload.setter
    def auto_asic_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_asic_offload", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPIP Tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "use_sdwan")

    @use_sdwan.setter
    def use_sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_sdwan", value)

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
class _IpiptunnelState:
    def __init__(__self__, *,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 use_sdwan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipiptunnel resources.
        :param pulumi.Input[str] auto_asic_offload: Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Interface name that is associated with the incoming traffic from available options.
        :param pulumi.Input[str] local_gw: IPv4 address for the local gateway.
        :param pulumi.Input[str] name: IPIP Tunnel name.
        :param pulumi.Input[str] remote_gw: IPv4 address for the remote gateway.
        :param pulumi.Input[str] use_sdwan: Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auto_asic_offload is not None:
            pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if local_gw is not None:
            pulumi.set(__self__, "local_gw", local_gw)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote_gw is not None:
            pulumi.set(__self__, "remote_gw", remote_gw)
        if use_sdwan is not None:
            pulumi.set(__self__, "use_sdwan", use_sdwan)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_asic_offload")

    @auto_asic_offload.setter
    def auto_asic_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_asic_offload", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name that is associated with the incoming traffic from available options.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address for the local gateway.
        """
        return pulumi.get(self, "local_gw")

    @local_gw.setter
    def local_gw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gw", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPIP Tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address for the remote gateway.
        """
        return pulumi.get(self, "remote_gw")

    @remote_gw.setter
    def remote_gw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_gw", value)

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "use_sdwan")

    @use_sdwan.setter
    def use_sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_sdwan", value)

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


class Ipiptunnel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 use_sdwan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IP in IP Tunneling.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Ipiptunnel("trname",
            interface="port3",
            local_gw="1.1.1.1",
            remote_gw="2.2.2.2")
        ```

        ## Import

        System IpipTunnel can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/ipiptunnel:Ipiptunnel labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/ipiptunnel:Ipiptunnel labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_asic_offload: Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Interface name that is associated with the incoming traffic from available options.
        :param pulumi.Input[str] local_gw: IPv4 address for the local gateway.
        :param pulumi.Input[str] name: IPIP Tunnel name.
        :param pulumi.Input[str] remote_gw: IPv4 address for the remote gateway.
        :param pulumi.Input[str] use_sdwan: Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpiptunnelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IP in IP Tunneling.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Ipiptunnel("trname",
            interface="port3",
            local_gw="1.1.1.1",
            remote_gw="2.2.2.2")
        ```

        ## Import

        System IpipTunnel can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/ipiptunnel:Ipiptunnel labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/ipiptunnel:Ipiptunnel labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param IpiptunnelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpiptunnelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_asic_offload: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 use_sdwan: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpiptunnelArgs.__new__(IpiptunnelArgs)

            __props__.__dict__["auto_asic_offload"] = auto_asic_offload
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if local_gw is None and not opts.urn:
                raise TypeError("Missing required property 'local_gw'")
            __props__.__dict__["local_gw"] = local_gw
            __props__.__dict__["name"] = name
            if remote_gw is None and not opts.urn:
                raise TypeError("Missing required property 'remote_gw'")
            __props__.__dict__["remote_gw"] = remote_gw
            __props__.__dict__["use_sdwan"] = use_sdwan
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ipiptunnel, __self__).__init__(
            'fortios:sys/ipiptunnel:Ipiptunnel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_asic_offload: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            local_gw: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_gw: Optional[pulumi.Input[str]] = None,
            use_sdwan: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ipiptunnel':
        """
        Get an existing Ipiptunnel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_asic_offload: Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Interface name that is associated with the incoming traffic from available options.
        :param pulumi.Input[str] local_gw: IPv4 address for the local gateway.
        :param pulumi.Input[str] name: IPIP Tunnel name.
        :param pulumi.Input[str] remote_gw: IPv4 address for the remote gateway.
        :param pulumi.Input[str] use_sdwan: Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpiptunnelState.__new__(_IpiptunnelState)

        __props__.__dict__["auto_asic_offload"] = auto_asic_offload
        __props__.__dict__["interface"] = interface
        __props__.__dict__["local_gw"] = local_gw
        __props__.__dict__["name"] = name
        __props__.__dict__["remote_gw"] = remote_gw
        __props__.__dict__["use_sdwan"] = use_sdwan
        __props__.__dict__["vdomparam"] = vdomparam
        return Ipiptunnel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> pulumi.Output[str]:
        """
        Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auto_asic_offload")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Interface name that is associated with the incoming traffic from available options.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> pulumi.Output[str]:
        """
        IPv4 address for the local gateway.
        """
        return pulumi.get(self, "local_gw")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IPIP Tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> pulumi.Output[str]:
        """
        IPv4 address for the remote gateway.
        """
        return pulumi.get(self, "remote_gw")

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> pulumi.Output[str]:
        """
        Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "use_sdwan")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

