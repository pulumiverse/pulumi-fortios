# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['AnqpnetworkauthtypeArgs', 'Anqpnetworkauthtype']

@pulumi.input_type
class AnqpnetworkauthtypeArgs:
    def __init__(__self__, *,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Anqpnetworkauthtype resource.
        :param pulumi.Input[str] auth_type: Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        :param pulumi.Input[str] name: Authentication type name.
        :param pulumi.Input[str] url: Redirect URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        """
        Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Redirect URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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
class _AnqpnetworkauthtypeState:
    def __init__(__self__, *,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Anqpnetworkauthtype resources.
        :param pulumi.Input[str] auth_type: Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        :param pulumi.Input[str] name: Authentication type name.
        :param pulumi.Input[str] url: Redirect URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        """
        Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Redirect URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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


class Anqpnetworkauthtype(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure network authentication type.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wirelesscontroller.hotspot20.Anqpnetworkauthtype("trname",
            auth_type="acceptance-of-terms",
            url="www.example.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WirelessControllerHotspot20 AnqpNetworkAuthType can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/anqpnetworkauthtype:Anqpnetworkauthtype labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/anqpnetworkauthtype:Anqpnetworkauthtype labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        :param pulumi.Input[str] name: Authentication type name.
        :param pulumi.Input[str] url: Redirect URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AnqpnetworkauthtypeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure network authentication type.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wirelesscontroller.hotspot20.Anqpnetworkauthtype("trname",
            auth_type="acceptance-of-terms",
            url="www.example.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WirelessControllerHotspot20 AnqpNetworkAuthType can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/anqpnetworkauthtype:Anqpnetworkauthtype labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/anqpnetworkauthtype:Anqpnetworkauthtype labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AnqpnetworkauthtypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AnqpnetworkauthtypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AnqpnetworkauthtypeArgs.__new__(AnqpnetworkauthtypeArgs)

            __props__.__dict__["auth_type"] = auth_type
            __props__.__dict__["name"] = name
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(Anqpnetworkauthtype, __self__).__init__(
            'fortios:wirelesscontroller/hotspot20/anqpnetworkauthtype:Anqpnetworkauthtype',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Anqpnetworkauthtype':
        """
        Get an existing Anqpnetworkauthtype resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        :param pulumi.Input[str] name: Authentication type name.
        :param pulumi.Input[str] url: Redirect URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AnqpnetworkauthtypeState.__new__(_AnqpnetworkauthtypeState)

        __props__.__dict__["auth_type"] = auth_type
        __props__.__dict__["name"] = name
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return Anqpnetworkauthtype(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output[str]:
        """
        Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Authentication type name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Redirect URL.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

