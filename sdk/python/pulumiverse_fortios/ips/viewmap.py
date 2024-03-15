# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ViewmapArgs', 'Viewmap']

@pulumi.input_type
class ViewmapArgs:
    def __init__(__self__, *,
                 fosid: Optional[pulumi.Input[int]] = None,
                 id_policy_id: Optional[pulumi.Input[int]] = None,
                 policy_id: Optional[pulumi.Input[int]] = None,
                 vdom_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 which: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Viewmap resource.
        :param pulumi.Input[int] fosid: View ID.
        :param pulumi.Input[int] id_policy_id: ID-based policy ID.
        :param pulumi.Input[int] policy_id: Policy ID.
        :param pulumi.Input[int] vdom_id: VDOM ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] which: Policy.
        """
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if id_policy_id is not None:
            pulumi.set(__self__, "id_policy_id", id_policy_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if vdom_id is not None:
            pulumi.set(__self__, "vdom_id", vdom_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if which is not None:
            pulumi.set(__self__, "which", which)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        View ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="idPolicyId")
    def id_policy_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID-based policy ID.
        """
        return pulumi.get(self, "id_policy_id")

    @id_policy_id.setter
    def id_policy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id_policy_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[int]]:
        """
        Policy ID.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="vdomId")
    def vdom_id(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM ID.
        """
        return pulumi.get(self, "vdom_id")

    @vdom_id.setter
    def vdom_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom_id", value)

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

    @property
    @pulumi.getter
    def which(self) -> Optional[pulumi.Input[str]]:
        """
        Policy.
        """
        return pulumi.get(self, "which")

    @which.setter
    def which(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "which", value)


@pulumi.input_type
class _ViewmapState:
    def __init__(__self__, *,
                 fosid: Optional[pulumi.Input[int]] = None,
                 id_policy_id: Optional[pulumi.Input[int]] = None,
                 policy_id: Optional[pulumi.Input[int]] = None,
                 vdom_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 which: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Viewmap resources.
        :param pulumi.Input[int] fosid: View ID.
        :param pulumi.Input[int] id_policy_id: ID-based policy ID.
        :param pulumi.Input[int] policy_id: Policy ID.
        :param pulumi.Input[int] vdom_id: VDOM ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] which: Policy.
        """
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if id_policy_id is not None:
            pulumi.set(__self__, "id_policy_id", id_policy_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if vdom_id is not None:
            pulumi.set(__self__, "vdom_id", vdom_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if which is not None:
            pulumi.set(__self__, "which", which)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        View ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="idPolicyId")
    def id_policy_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID-based policy ID.
        """
        return pulumi.get(self, "id_policy_id")

    @id_policy_id.setter
    def id_policy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id_policy_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[int]]:
        """
        Policy ID.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="vdomId")
    def vdom_id(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM ID.
        """
        return pulumi.get(self, "vdom_id")

    @vdom_id.setter
    def vdom_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom_id", value)

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

    @property
    @pulumi.getter
    def which(self) -> Optional[pulumi.Input[str]]:
        """
        Policy.
        """
        return pulumi.get(self, "which")

    @which.setter
    def which(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "which", value)


class Viewmap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 id_policy_id: Optional[pulumi.Input[int]] = None,
                 policy_id: Optional[pulumi.Input[int]] = None,
                 vdom_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 which: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        configure ips view-map Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Ips ViewMap can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] fosid: View ID.
        :param pulumi.Input[int] id_policy_id: ID-based policy ID.
        :param pulumi.Input[int] policy_id: Policy ID.
        :param pulumi.Input[int] vdom_id: VDOM ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] which: Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ViewmapArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        configure ips view-map Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Ips ViewMap can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:ips/viewmap:Viewmap labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ViewmapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ViewmapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 id_policy_id: Optional[pulumi.Input[int]] = None,
                 policy_id: Optional[pulumi.Input[int]] = None,
                 vdom_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 which: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ViewmapArgs.__new__(ViewmapArgs)

            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["id_policy_id"] = id_policy_id
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["vdom_id"] = vdom_id
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["which"] = which
        super(Viewmap, __self__).__init__(
            'fortios:ips/viewmap:Viewmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            id_policy_id: Optional[pulumi.Input[int]] = None,
            policy_id: Optional[pulumi.Input[int]] = None,
            vdom_id: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            which: Optional[pulumi.Input[str]] = None) -> 'Viewmap':
        """
        Get an existing Viewmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] fosid: View ID.
        :param pulumi.Input[int] id_policy_id: ID-based policy ID.
        :param pulumi.Input[int] policy_id: Policy ID.
        :param pulumi.Input[int] vdom_id: VDOM ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] which: Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ViewmapState.__new__(_ViewmapState)

        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["id_policy_id"] = id_policy_id
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["vdom_id"] = vdom_id
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["which"] = which
        return Viewmap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        View ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="idPolicyId")
    def id_policy_id(self) -> pulumi.Output[int]:
        """
        ID-based policy ID.
        """
        return pulumi.get(self, "id_policy_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[int]:
        """
        Policy ID.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="vdomId")
    def vdom_id(self) -> pulumi.Output[int]:
        """
        VDOM ID.
        """
        return pulumi.get(self, "vdom_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def which(self) -> pulumi.Output[str]:
        """
        Policy.
        """
        return pulumi.get(self, "which")

