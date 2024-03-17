# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['H2qposuproviderArgs', 'H2qposuprovider']

@pulumi.input_type
class H2qposuproviderArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 friendly_names: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 osu_method: Optional[pulumi.Input[str]] = None,
                 osu_nai: Optional[pulumi.Input[str]] = None,
                 server_uri: Optional[pulumi.Input[str]] = None,
                 service_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a H2qposuprovider resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]] friendly_names: OSU provider friendly name. The structure of `friendly_name` block is documented below.
        :param pulumi.Input[str] icon: OSU provider icon.
        :param pulumi.Input[str] name: OSU provider ID.
        :param pulumi.Input[str] osu_method: OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        :param pulumi.Input[str] osu_nai: OSU NAI.
        :param pulumi.Input[str] server_uri: Server URI.
        :param pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]] service_descriptions: OSU service name. The structure of `service_description` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if friendly_names is not None:
            pulumi.set(__self__, "friendly_names", friendly_names)
        if icon is not None:
            pulumi.set(__self__, "icon", icon)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if osu_method is not None:
            pulumi.set(__self__, "osu_method", osu_method)
        if osu_nai is not None:
            pulumi.set(__self__, "osu_nai", osu_nai)
        if server_uri is not None:
            pulumi.set(__self__, "server_uri", server_uri)
        if service_descriptions is not None:
            pulumi.set(__self__, "service_descriptions", service_descriptions)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="friendlyNames")
    def friendly_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]]:
        """
        OSU provider friendly name. The structure of `friendly_name` block is documented below.
        """
        return pulumi.get(self, "friendly_names")

    @friendly_names.setter
    def friendly_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]]):
        pulumi.set(self, "friendly_names", value)

    @property
    @pulumi.getter
    def icon(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider icon.
        """
        return pulumi.get(self, "icon")

    @icon.setter
    def icon(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icon", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osuMethod")
    def osu_method(self) -> Optional[pulumi.Input[str]]:
        """
        OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        """
        return pulumi.get(self, "osu_method")

    @osu_method.setter
    def osu_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "osu_method", value)

    @property
    @pulumi.getter(name="osuNai")
    def osu_nai(self) -> Optional[pulumi.Input[str]]:
        """
        OSU NAI.
        """
        return pulumi.get(self, "osu_nai")

    @osu_nai.setter
    def osu_nai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "osu_nai", value)

    @property
    @pulumi.getter(name="serverUri")
    def server_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Server URI.
        """
        return pulumi.get(self, "server_uri")

    @server_uri.setter
    def server_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_uri", value)

    @property
    @pulumi.getter(name="serviceDescriptions")
    def service_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]]:
        """
        OSU service name. The structure of `service_description` block is documented below.
        """
        return pulumi.get(self, "service_descriptions")

    @service_descriptions.setter
    def service_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]]):
        pulumi.set(self, "service_descriptions", value)

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
class _H2qposuproviderState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 friendly_names: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 osu_method: Optional[pulumi.Input[str]] = None,
                 osu_nai: Optional[pulumi.Input[str]] = None,
                 server_uri: Optional[pulumi.Input[str]] = None,
                 service_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering H2qposuprovider resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]] friendly_names: OSU provider friendly name. The structure of `friendly_name` block is documented below.
        :param pulumi.Input[str] icon: OSU provider icon.
        :param pulumi.Input[str] name: OSU provider ID.
        :param pulumi.Input[str] osu_method: OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        :param pulumi.Input[str] osu_nai: OSU NAI.
        :param pulumi.Input[str] server_uri: Server URI.
        :param pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]] service_descriptions: OSU service name. The structure of `service_description` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if friendly_names is not None:
            pulumi.set(__self__, "friendly_names", friendly_names)
        if icon is not None:
            pulumi.set(__self__, "icon", icon)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if osu_method is not None:
            pulumi.set(__self__, "osu_method", osu_method)
        if osu_nai is not None:
            pulumi.set(__self__, "osu_nai", osu_nai)
        if server_uri is not None:
            pulumi.set(__self__, "server_uri", server_uri)
        if service_descriptions is not None:
            pulumi.set(__self__, "service_descriptions", service_descriptions)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="friendlyNames")
    def friendly_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]]:
        """
        OSU provider friendly name. The structure of `friendly_name` block is documented below.
        """
        return pulumi.get(self, "friendly_names")

    @friendly_names.setter
    def friendly_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderFriendlyNameArgs']]]]):
        pulumi.set(self, "friendly_names", value)

    @property
    @pulumi.getter
    def icon(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider icon.
        """
        return pulumi.get(self, "icon")

    @icon.setter
    def icon(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icon", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osuMethod")
    def osu_method(self) -> Optional[pulumi.Input[str]]:
        """
        OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        """
        return pulumi.get(self, "osu_method")

    @osu_method.setter
    def osu_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "osu_method", value)

    @property
    @pulumi.getter(name="osuNai")
    def osu_nai(self) -> Optional[pulumi.Input[str]]:
        """
        OSU NAI.
        """
        return pulumi.get(self, "osu_nai")

    @osu_nai.setter
    def osu_nai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "osu_nai", value)

    @property
    @pulumi.getter(name="serverUri")
    def server_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Server URI.
        """
        return pulumi.get(self, "server_uri")

    @server_uri.setter
    def server_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_uri", value)

    @property
    @pulumi.getter(name="serviceDescriptions")
    def service_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]]:
        """
        OSU service name. The structure of `service_description` block is documented below.
        """
        return pulumi.get(self, "service_descriptions")

    @service_descriptions.setter
    def service_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['H2qposuproviderServiceDescriptionArgs']]]]):
        pulumi.set(self, "service_descriptions", value)

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


class H2qposuprovider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 friendly_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderFriendlyNameArgs']]]]] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 osu_method: Optional[pulumi.Input[str]] = None,
                 osu_nai: Optional[pulumi.Input[str]] = None,
                 server_uri: Optional[pulumi.Input[str]] = None,
                 service_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderServiceDescriptionArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure online sign up (OSU) provider list.

        ## Import

        WirelessControllerHotspot20 H2QpOsuProvider can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderFriendlyNameArgs']]]] friendly_names: OSU provider friendly name. The structure of `friendly_name` block is documented below.
        :param pulumi.Input[str] icon: OSU provider icon.
        :param pulumi.Input[str] name: OSU provider ID.
        :param pulumi.Input[str] osu_method: OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        :param pulumi.Input[str] osu_nai: OSU NAI.
        :param pulumi.Input[str] server_uri: Server URI.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderServiceDescriptionArgs']]]] service_descriptions: OSU service name. The structure of `service_description` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[H2qposuproviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure online sign up (OSU) provider list.

        ## Import

        WirelessControllerHotspot20 H2QpOsuProvider can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param H2qposuproviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(H2qposuproviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 friendly_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderFriendlyNameArgs']]]]] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 osu_method: Optional[pulumi.Input[str]] = None,
                 osu_nai: Optional[pulumi.Input[str]] = None,
                 server_uri: Optional[pulumi.Input[str]] = None,
                 service_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderServiceDescriptionArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = H2qposuproviderArgs.__new__(H2qposuproviderArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["friendly_names"] = friendly_names
            __props__.__dict__["icon"] = icon
            __props__.__dict__["name"] = name
            __props__.__dict__["osu_method"] = osu_method
            __props__.__dict__["osu_nai"] = osu_nai
            __props__.__dict__["server_uri"] = server_uri
            __props__.__dict__["service_descriptions"] = service_descriptions
            __props__.__dict__["vdomparam"] = vdomparam
        super(H2qposuprovider, __self__).__init__(
            'fortios:wirelesscontroller/hotspot20/h2qposuprovider:H2qposuprovider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            friendly_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderFriendlyNameArgs']]]]] = None,
            icon: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            osu_method: Optional[pulumi.Input[str]] = None,
            osu_nai: Optional[pulumi.Input[str]] = None,
            server_uri: Optional[pulumi.Input[str]] = None,
            service_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderServiceDescriptionArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'H2qposuprovider':
        """
        Get an existing H2qposuprovider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderFriendlyNameArgs']]]] friendly_names: OSU provider friendly name. The structure of `friendly_name` block is documented below.
        :param pulumi.Input[str] icon: OSU provider icon.
        :param pulumi.Input[str] name: OSU provider ID.
        :param pulumi.Input[str] osu_method: OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        :param pulumi.Input[str] osu_nai: OSU NAI.
        :param pulumi.Input[str] server_uri: Server URI.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['H2qposuproviderServiceDescriptionArgs']]]] service_descriptions: OSU service name. The structure of `service_description` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _H2qposuproviderState.__new__(_H2qposuproviderState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["friendly_names"] = friendly_names
        __props__.__dict__["icon"] = icon
        __props__.__dict__["name"] = name
        __props__.__dict__["osu_method"] = osu_method
        __props__.__dict__["osu_nai"] = osu_nai
        __props__.__dict__["server_uri"] = server_uri
        __props__.__dict__["service_descriptions"] = service_descriptions
        __props__.__dict__["vdomparam"] = vdomparam
        return H2qposuprovider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="friendlyNames")
    def friendly_names(self) -> pulumi.Output[Optional[Sequence['outputs.H2qposuproviderFriendlyName']]]:
        """
        OSU provider friendly name. The structure of `friendly_name` block is documented below.
        """
        return pulumi.get(self, "friendly_names")

    @property
    @pulumi.getter
    def icon(self) -> pulumi.Output[str]:
        """
        OSU provider icon.
        """
        return pulumi.get(self, "icon")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        OSU provider ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osuMethod")
    def osu_method(self) -> pulumi.Output[str]:
        """
        OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
        """
        return pulumi.get(self, "osu_method")

    @property
    @pulumi.getter(name="osuNai")
    def osu_nai(self) -> pulumi.Output[str]:
        """
        OSU NAI.
        """
        return pulumi.get(self, "osu_nai")

    @property
    @pulumi.getter(name="serverUri")
    def server_uri(self) -> pulumi.Output[str]:
        """
        Server URI.
        """
        return pulumi.get(self, "server_uri")

    @property
    @pulumi.getter(name="serviceDescriptions")
    def service_descriptions(self) -> pulumi.Output[Optional[Sequence['outputs.H2qposuproviderServiceDescription']]]:
        """
        OSU service name. The structure of `service_description` block is documented below.
        """
        return pulumi.get(self, "service_descriptions")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

