# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ManualkeyArgs', 'Manualkey']

@pulumi.input_type
class ManualkeyArgs:
    def __init__(__self__, *,
                 authentication: pulumi.Input[str],
                 encryption: pulumi.Input[str],
                 interface: pulumi.Input[str],
                 remote_gw: pulumi.Input[str],
                 authkey: Optional[pulumi.Input[str]] = None,
                 enckey: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 localspi: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 npu_offload: Optional[pulumi.Input[str]] = None,
                 remotespi: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Manualkey resource.
        :param pulumi.Input[str] authentication: Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        :param pulumi.Input[str] encryption: Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        :param pulumi.Input[str] interface: Name of the physical, aggregate, or VLAN interface.
        :param pulumi.Input[str] remote_gw: Peer gateway.
        :param pulumi.Input[str] authkey: Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] enckey: Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] local_gw: Local gateway.
        :param pulumi.Input[str] localspi: Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] name: IPsec tunnel name.
        :param pulumi.Input[str] npu_offload: Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] remotespi: Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "authentication", authentication)
        pulumi.set(__self__, "encryption", encryption)
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "remote_gw", remote_gw)
        if authkey is not None:
            pulumi.set(__self__, "authkey", authkey)
        if enckey is not None:
            pulumi.set(__self__, "enckey", enckey)
        if local_gw is not None:
            pulumi.set(__self__, "local_gw", local_gw)
        if localspi is not None:
            pulumi.set(__self__, "localspi", localspi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if npu_offload is not None:
            pulumi.set(__self__, "npu_offload", npu_offload)
        if remotespi is not None:
            pulumi.set(__self__, "remotespi", remotespi)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Input[str]:
        """
        Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: pulumi.Input[str]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Input[str]:
        """
        Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Name of the physical, aggregate, or VLAN interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> pulumi.Input[str]:
        """
        Peer gateway.
        """
        return pulumi.get(self, "remote_gw")

    @remote_gw.setter
    def remote_gw(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_gw", value)

    @property
    @pulumi.getter
    def authkey(self) -> Optional[pulumi.Input[str]]:
        """
        Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "authkey")

    @authkey.setter
    def authkey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authkey", value)

    @property
    @pulumi.getter
    def enckey(self) -> Optional[pulumi.Input[str]]:
        """
        Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "enckey")

    @enckey.setter
    def enckey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enckey", value)

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> Optional[pulumi.Input[str]]:
        """
        Local gateway.
        """
        return pulumi.get(self, "local_gw")

    @local_gw.setter
    def local_gw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gw", value)

    @property
    @pulumi.getter
    def localspi(self) -> Optional[pulumi.Input[str]]:
        """
        Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "localspi")

    @localspi.setter
    def localspi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "localspi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPsec tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="npuOffload")
    def npu_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "npu_offload")

    @npu_offload.setter
    def npu_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "npu_offload", value)

    @property
    @pulumi.getter
    def remotespi(self) -> Optional[pulumi.Input[str]]:
        """
        Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "remotespi")

    @remotespi.setter
    def remotespi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remotespi", value)

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
class _ManualkeyState:
    def __init__(__self__, *,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authkey: Optional[pulumi.Input[str]] = None,
                 enckey: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 localspi: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 npu_offload: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 remotespi: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Manualkey resources.
        :param pulumi.Input[str] authentication: Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        :param pulumi.Input[str] authkey: Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] enckey: Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] encryption: Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        :param pulumi.Input[str] interface: Name of the physical, aggregate, or VLAN interface.
        :param pulumi.Input[str] local_gw: Local gateway.
        :param pulumi.Input[str] localspi: Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] name: IPsec tunnel name.
        :param pulumi.Input[str] npu_offload: Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] remote_gw: Peer gateway.
        :param pulumi.Input[str] remotespi: Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if authkey is not None:
            pulumi.set(__self__, "authkey", authkey)
        if enckey is not None:
            pulumi.set(__self__, "enckey", enckey)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if local_gw is not None:
            pulumi.set(__self__, "local_gw", local_gw)
        if localspi is not None:
            pulumi.set(__self__, "localspi", localspi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if npu_offload is not None:
            pulumi.set(__self__, "npu_offload", npu_offload)
        if remote_gw is not None:
            pulumi.set(__self__, "remote_gw", remote_gw)
        if remotespi is not None:
            pulumi.set(__self__, "remotespi", remotespi)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter
    def authkey(self) -> Optional[pulumi.Input[str]]:
        """
        Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "authkey")

    @authkey.setter
    def authkey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authkey", value)

    @property
    @pulumi.getter
    def enckey(self) -> Optional[pulumi.Input[str]]:
        """
        Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "enckey")

    @enckey.setter
    def enckey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enckey", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the physical, aggregate, or VLAN interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> Optional[pulumi.Input[str]]:
        """
        Local gateway.
        """
        return pulumi.get(self, "local_gw")

    @local_gw.setter
    def local_gw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gw", value)

    @property
    @pulumi.getter
    def localspi(self) -> Optional[pulumi.Input[str]]:
        """
        Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "localspi")

    @localspi.setter
    def localspi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "localspi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPsec tunnel name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="npuOffload")
    def npu_offload(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "npu_offload")

    @npu_offload.setter
    def npu_offload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "npu_offload", value)

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> Optional[pulumi.Input[str]]:
        """
        Peer gateway.
        """
        return pulumi.get(self, "remote_gw")

    @remote_gw.setter
    def remote_gw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_gw", value)

    @property
    @pulumi.getter
    def remotespi(self) -> Optional[pulumi.Input[str]]:
        """
        Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "remotespi")

    @remotespi.setter
    def remotespi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remotespi", value)

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


class Manualkey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authkey: Optional[pulumi.Input[str]] = None,
                 enckey: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 localspi: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 npu_offload: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 remotespi: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPsec manual keys.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.vpnipsec.Manualkey("trname",
            authentication="md5",
            authkey="EE32CA121ECD772A-ECACAABA212345EC",
            enckey="-",
            encryption="null",
            interface="port4",
            local_gw="0.0.0.0",
            localspi="0x100",
            remote_gw="1.1.1.1",
            remotespi="0x100")
        ```

        ## Import

        VpnIpsec Manualkey can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authentication: Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        :param pulumi.Input[str] authkey: Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] enckey: Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] encryption: Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        :param pulumi.Input[str] interface: Name of the physical, aggregate, or VLAN interface.
        :param pulumi.Input[str] local_gw: Local gateway.
        :param pulumi.Input[str] localspi: Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] name: IPsec tunnel name.
        :param pulumi.Input[str] npu_offload: Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] remote_gw: Peer gateway.
        :param pulumi.Input[str] remotespi: Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManualkeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPsec manual keys.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.vpnipsec.Manualkey("trname",
            authentication="md5",
            authkey="EE32CA121ECD772A-ECACAABA212345EC",
            enckey="-",
            encryption="null",
            interface="port4",
            local_gw="0.0.0.0",
            localspi="0x100",
            remote_gw="1.1.1.1",
            remotespi="0x100")
        ```

        ## Import

        VpnIpsec Manualkey can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:vpnipsec/manualkey:Manualkey labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ManualkeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManualkeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authkey: Optional[pulumi.Input[str]] = None,
                 enckey: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 local_gw: Optional[pulumi.Input[str]] = None,
                 localspi: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 npu_offload: Optional[pulumi.Input[str]] = None,
                 remote_gw: Optional[pulumi.Input[str]] = None,
                 remotespi: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManualkeyArgs.__new__(ManualkeyArgs)

            if authentication is None and not opts.urn:
                raise TypeError("Missing required property 'authentication'")
            __props__.__dict__["authentication"] = authentication
            __props__.__dict__["authkey"] = None if authkey is None else pulumi.Output.secret(authkey)
            __props__.__dict__["enckey"] = None if enckey is None else pulumi.Output.secret(enckey)
            if encryption is None and not opts.urn:
                raise TypeError("Missing required property 'encryption'")
            __props__.__dict__["encryption"] = encryption
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            __props__.__dict__["local_gw"] = local_gw
            __props__.__dict__["localspi"] = localspi
            __props__.__dict__["name"] = name
            __props__.__dict__["npu_offload"] = npu_offload
            if remote_gw is None and not opts.urn:
                raise TypeError("Missing required property 'remote_gw'")
            __props__.__dict__["remote_gw"] = remote_gw
            __props__.__dict__["remotespi"] = remotespi
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authkey", "enckey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Manualkey, __self__).__init__(
            'fortios:vpnipsec/manualkey:Manualkey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication: Optional[pulumi.Input[str]] = None,
            authkey: Optional[pulumi.Input[str]] = None,
            enckey: Optional[pulumi.Input[str]] = None,
            encryption: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            local_gw: Optional[pulumi.Input[str]] = None,
            localspi: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            npu_offload: Optional[pulumi.Input[str]] = None,
            remote_gw: Optional[pulumi.Input[str]] = None,
            remotespi: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Manualkey':
        """
        Get an existing Manualkey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authentication: Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        :param pulumi.Input[str] authkey: Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] enckey: Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        :param pulumi.Input[str] encryption: Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        :param pulumi.Input[str] interface: Name of the physical, aggregate, or VLAN interface.
        :param pulumi.Input[str] local_gw: Local gateway.
        :param pulumi.Input[str] localspi: Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] name: IPsec tunnel name.
        :param pulumi.Input[str] npu_offload: Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] remote_gw: Peer gateway.
        :param pulumi.Input[str] remotespi: Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManualkeyState.__new__(_ManualkeyState)

        __props__.__dict__["authentication"] = authentication
        __props__.__dict__["authkey"] = authkey
        __props__.__dict__["enckey"] = enckey
        __props__.__dict__["encryption"] = encryption
        __props__.__dict__["interface"] = interface
        __props__.__dict__["local_gw"] = local_gw
        __props__.__dict__["localspi"] = localspi
        __props__.__dict__["name"] = name
        __props__.__dict__["npu_offload"] = npu_offload
        __props__.__dict__["remote_gw"] = remote_gw
        __props__.__dict__["remotespi"] = remotespi
        __props__.__dict__["vdomparam"] = vdomparam
        return Manualkey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[str]:
        """
        Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter
    def authkey(self) -> pulumi.Output[str]:
        """
        Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "authkey")

    @property
    @pulumi.getter
    def enckey(self) -> pulumi.Output[str]:
        """
        Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
        """
        return pulumi.get(self, "enckey")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[str]:
        """
        Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Name of the physical, aggregate, or VLAN interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> pulumi.Output[str]:
        """
        Local gateway.
        """
        return pulumi.get(self, "local_gw")

    @property
    @pulumi.getter
    def localspi(self) -> pulumi.Output[str]:
        """
        Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "localspi")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IPsec tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="npuOffload")
    def npu_offload(self) -> pulumi.Output[str]:
        """
        Enable/disable NPU offloading. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "npu_offload")

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> pulumi.Output[str]:
        """
        Peer gateway.
        """
        return pulumi.get(self, "remote_gw")

    @property
    @pulumi.getter
    def remotespi(self) -> pulumi.Output[str]:
        """
        Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
        """
        return pulumi.get(self, "remotespi")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

