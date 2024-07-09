# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['DefaultArgs', 'Default']

@pulumi.input_type
class DefaultArgs:
    def __init__(__self__, *,
                 fgt_policy: Optional[pulumi.Input[str]] = None,
                 icl_policy: Optional[pulumi.Input[str]] = None,
                 isl_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Default resource.
        :param pulumi.Input[str] fgt_policy: Default FortiLink auto-config policy.
        :param pulumi.Input[str] icl_policy: Default ICL auto-config policy.
        :param pulumi.Input[str] isl_policy: Default ISL auto-config policy.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fgt_policy is not None:
            pulumi.set(__self__, "fgt_policy", fgt_policy)
        if icl_policy is not None:
            pulumi.set(__self__, "icl_policy", icl_policy)
        if isl_policy is not None:
            pulumi.set(__self__, "isl_policy", isl_policy)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="fgtPolicy")
    def fgt_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default FortiLink auto-config policy.
        """
        return pulumi.get(self, "fgt_policy")

    @fgt_policy.setter
    def fgt_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fgt_policy", value)

    @property
    @pulumi.getter(name="iclPolicy")
    def icl_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default ICL auto-config policy.
        """
        return pulumi.get(self, "icl_policy")

    @icl_policy.setter
    def icl_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icl_policy", value)

    @property
    @pulumi.getter(name="islPolicy")
    def isl_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default ISL auto-config policy.
        """
        return pulumi.get(self, "isl_policy")

    @isl_policy.setter
    def isl_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isl_policy", value)

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
class _DefaultState:
    def __init__(__self__, *,
                 fgt_policy: Optional[pulumi.Input[str]] = None,
                 icl_policy: Optional[pulumi.Input[str]] = None,
                 isl_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Default resources.
        :param pulumi.Input[str] fgt_policy: Default FortiLink auto-config policy.
        :param pulumi.Input[str] icl_policy: Default ICL auto-config policy.
        :param pulumi.Input[str] isl_policy: Default ISL auto-config policy.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fgt_policy is not None:
            pulumi.set(__self__, "fgt_policy", fgt_policy)
        if icl_policy is not None:
            pulumi.set(__self__, "icl_policy", icl_policy)
        if isl_policy is not None:
            pulumi.set(__self__, "isl_policy", isl_policy)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="fgtPolicy")
    def fgt_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default FortiLink auto-config policy.
        """
        return pulumi.get(self, "fgt_policy")

    @fgt_policy.setter
    def fgt_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fgt_policy", value)

    @property
    @pulumi.getter(name="iclPolicy")
    def icl_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default ICL auto-config policy.
        """
        return pulumi.get(self, "icl_policy")

    @icl_policy.setter
    def icl_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icl_policy", value)

    @property
    @pulumi.getter(name="islPolicy")
    def isl_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Default ISL auto-config policy.
        """
        return pulumi.get(self, "isl_policy")

    @isl_policy.setter
    def isl_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isl_policy", value)

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


class Default(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fgt_policy: Optional[pulumi.Input[str]] = None,
                 icl_policy: Optional[pulumi.Input[str]] = None,
                 isl_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure default auto-config QoS policy for FortiSwitch.

        ## Import

        SwitchControllerAutoConfig Default can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/autoconfig/default:Default labelname SwitchControllerAutoConfigDefault
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/autoconfig/default:Default labelname SwitchControllerAutoConfigDefault
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fgt_policy: Default FortiLink auto-config policy.
        :param pulumi.Input[str] icl_policy: Default ICL auto-config policy.
        :param pulumi.Input[str] isl_policy: Default ISL auto-config policy.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DefaultArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure default auto-config QoS policy for FortiSwitch.

        ## Import

        SwitchControllerAutoConfig Default can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/autoconfig/default:Default labelname SwitchControllerAutoConfigDefault
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/autoconfig/default:Default labelname SwitchControllerAutoConfigDefault
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DefaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fgt_policy: Optional[pulumi.Input[str]] = None,
                 icl_policy: Optional[pulumi.Input[str]] = None,
                 isl_policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultArgs.__new__(DefaultArgs)

            __props__.__dict__["fgt_policy"] = fgt_policy
            __props__.__dict__["icl_policy"] = icl_policy
            __props__.__dict__["isl_policy"] = isl_policy
            __props__.__dict__["vdomparam"] = vdomparam
        super(Default, __self__).__init__(
            'fortios:switchcontroller/autoconfig/default:Default',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fgt_policy: Optional[pulumi.Input[str]] = None,
            icl_policy: Optional[pulumi.Input[str]] = None,
            isl_policy: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Default':
        """
        Get an existing Default resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fgt_policy: Default FortiLink auto-config policy.
        :param pulumi.Input[str] icl_policy: Default ICL auto-config policy.
        :param pulumi.Input[str] isl_policy: Default ISL auto-config policy.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultState.__new__(_DefaultState)

        __props__.__dict__["fgt_policy"] = fgt_policy
        __props__.__dict__["icl_policy"] = icl_policy
        __props__.__dict__["isl_policy"] = isl_policy
        __props__.__dict__["vdomparam"] = vdomparam
        return Default(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fgtPolicy")
    def fgt_policy(self) -> pulumi.Output[str]:
        """
        Default FortiLink auto-config policy.
        """
        return pulumi.get(self, "fgt_policy")

    @property
    @pulumi.getter(name="iclPolicy")
    def icl_policy(self) -> pulumi.Output[str]:
        """
        Default ICL auto-config policy.
        """
        return pulumi.get(self, "icl_policy")

    @property
    @pulumi.getter(name="islPolicy")
    def isl_policy(self) -> pulumi.Output[str]:
        """
        Default ISL auto-config policy.
        """
        return pulumi.get(self, "isl_policy")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

