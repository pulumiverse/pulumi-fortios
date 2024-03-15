# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApstatusArgs', 'Apstatus']

@pulumi.input_type
class ApstatusArgs:
    def __init__(__self__, *,
                 bssid: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ssid: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Apstatus resource.
        :param pulumi.Input[str] bssid: Access Point's (AP's) BSSID.
        :param pulumi.Input[int] fosid: AP ID.
        :param pulumi.Input[str] ssid: Access Point's (AP's) SSID.
        :param pulumi.Input[str] status: Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if bssid is not None:
            pulumi.set(__self__, "bssid", bssid)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if ssid is not None:
            pulumi.set(__self__, "ssid", ssid)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def bssid(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) BSSID.
        """
        return pulumi.get(self, "bssid")

    @bssid.setter
    def bssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bssid", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        AP ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def ssid(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) SSID.
        """
        return pulumi.get(self, "ssid")

    @ssid.setter
    def ssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssid", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
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
class _ApstatusState:
    def __init__(__self__, *,
                 bssid: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ssid: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Apstatus resources.
        :param pulumi.Input[str] bssid: Access Point's (AP's) BSSID.
        :param pulumi.Input[int] fosid: AP ID.
        :param pulumi.Input[str] ssid: Access Point's (AP's) SSID.
        :param pulumi.Input[str] status: Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if bssid is not None:
            pulumi.set(__self__, "bssid", bssid)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if ssid is not None:
            pulumi.set(__self__, "ssid", ssid)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def bssid(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) BSSID.
        """
        return pulumi.get(self, "bssid")

    @bssid.setter
    def bssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bssid", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        AP ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def ssid(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) SSID.
        """
        return pulumi.get(self, "ssid")

    @ssid.setter
    def ssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssid", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
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


class Apstatus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bssid: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ssid: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure access point status (rogue | accepted | suppressed).

        ## Import

        WirelessController ApStatus can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bssid: Access Point's (AP's) BSSID.
        :param pulumi.Input[int] fosid: AP ID.
        :param pulumi.Input[str] ssid: Access Point's (AP's) SSID.
        :param pulumi.Input[str] status: Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApstatusArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure access point status (rogue | accepted | suppressed).

        ## Import

        WirelessController ApStatus can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/apstatus:Apstatus labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ApstatusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApstatusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bssid: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 ssid: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApstatusArgs.__new__(ApstatusArgs)

            __props__.__dict__["bssid"] = bssid
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["ssid"] = ssid
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Apstatus, __self__).__init__(
            'fortios:wirelesscontroller/apstatus:Apstatus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bssid: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            ssid: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Apstatus':
        """
        Get an existing Apstatus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bssid: Access Point's (AP's) BSSID.
        :param pulumi.Input[int] fosid: AP ID.
        :param pulumi.Input[str] ssid: Access Point's (AP's) SSID.
        :param pulumi.Input[str] status: Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApstatusState.__new__(_ApstatusState)

        __props__.__dict__["bssid"] = bssid
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["ssid"] = ssid
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Apstatus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bssid(self) -> pulumi.Output[str]:
        """
        Access Point's (AP's) BSSID.
        """
        return pulumi.get(self, "bssid")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        AP ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def ssid(self) -> pulumi.Output[str]:
        """
        Access Point's (AP's) SSID.
        """
        return pulumi.get(self, "ssid")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

