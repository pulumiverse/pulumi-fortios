# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]] = None,
                 behavior: Optional[pulumi.Input[str]] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 popularity: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[str]] = None,
                 risks: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]] = None,
                 technology: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]] applications: Application ID list. The structure of `application` block is documented below.
        :param pulumi.Input[str] behavior: Application behavior filter.
        :param pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]] categories: Application category ID list. The structure of `category` block is documented below.
        :param pulumi.Input[str] comment: Comment
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Application group name.
        :param pulumi.Input[str] popularity: Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        :param pulumi.Input[str] protocols: Application protocol filter.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]] risks: Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        :param pulumi.Input[str] technology: Application technology filter.
        :param pulumi.Input[str] type: Application group type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: Application vendor filter.
        """
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if behavior is not None:
            pulumi.set(__self__, "behavior", behavior)
        if categories is not None:
            pulumi.set(__self__, "categories", categories)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if popularity is not None:
            pulumi.set(__self__, "popularity", popularity)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if risks is not None:
            pulumi.set(__self__, "risks", risks)
        if technology is not None:
            pulumi.set(__self__, "technology", technology)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vendor is not None:
            pulumi.set(__self__, "vendor", vendor)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]]:
        """
        Application ID list. The structure of `application` block is documented below.
        """
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter
    def behavior(self) -> Optional[pulumi.Input[str]]:
        """
        Application behavior filter.
        """
        return pulumi.get(self, "behavior")

    @behavior.setter
    def behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "behavior", value)

    @property
    @pulumi.getter
    def categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]]:
        """
        Application category ID list. The structure of `category` block is documented below.
        """
        return pulumi.get(self, "categories")

    @categories.setter
    def categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]]):
        pulumi.set(self, "categories", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Application group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def popularity(self) -> Optional[pulumi.Input[str]]:
        """
        Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        """
        return pulumi.get(self, "popularity")

    @popularity.setter
    def popularity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "popularity", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[str]]:
        """
        Application protocol filter.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocols", value)

    @property
    @pulumi.getter
    def risks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]]:
        """
        Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        """
        return pulumi.get(self, "risks")

    @risks.setter
    def risks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]]):
        pulumi.set(self, "risks", value)

    @property
    @pulumi.getter
    def technology(self) -> Optional[pulumi.Input[str]]:
        """
        Application technology filter.
        """
        return pulumi.get(self, "technology")

    @technology.setter
    def technology(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "technology", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Application group type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
    def vendor(self) -> Optional[pulumi.Input[str]]:
        """
        Application vendor filter.
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]] = None,
                 behavior: Optional[pulumi.Input[str]] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 popularity: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[str]] = None,
                 risks: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]] = None,
                 technology: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]] applications: Application ID list. The structure of `application` block is documented below.
        :param pulumi.Input[str] behavior: Application behavior filter.
        :param pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]] categories: Application category ID list. The structure of `category` block is documented below.
        :param pulumi.Input[str] comment: Comment
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Application group name.
        :param pulumi.Input[str] popularity: Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        :param pulumi.Input[str] protocols: Application protocol filter.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]] risks: Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        :param pulumi.Input[str] technology: Application technology filter.
        :param pulumi.Input[str] type: Application group type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: Application vendor filter.
        """
        if applications is not None:
            pulumi.set(__self__, "applications", applications)
        if behavior is not None:
            pulumi.set(__self__, "behavior", behavior)
        if categories is not None:
            pulumi.set(__self__, "categories", categories)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if popularity is not None:
            pulumi.set(__self__, "popularity", popularity)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if risks is not None:
            pulumi.set(__self__, "risks", risks)
        if technology is not None:
            pulumi.set(__self__, "technology", technology)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vendor is not None:
            pulumi.set(__self__, "vendor", vendor)

    @property
    @pulumi.getter
    def applications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]]:
        """
        Application ID list. The structure of `application` block is documented below.
        """
        return pulumi.get(self, "applications")

    @applications.setter
    def applications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupApplicationArgs']]]]):
        pulumi.set(self, "applications", value)

    @property
    @pulumi.getter
    def behavior(self) -> Optional[pulumi.Input[str]]:
        """
        Application behavior filter.
        """
        return pulumi.get(self, "behavior")

    @behavior.setter
    def behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "behavior", value)

    @property
    @pulumi.getter
    def categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]]:
        """
        Application category ID list. The structure of `category` block is documented below.
        """
        return pulumi.get(self, "categories")

    @categories.setter
    def categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupCategoryArgs']]]]):
        pulumi.set(self, "categories", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Application group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def popularity(self) -> Optional[pulumi.Input[str]]:
        """
        Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        """
        return pulumi.get(self, "popularity")

    @popularity.setter
    def popularity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "popularity", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[str]]:
        """
        Application protocol filter.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocols", value)

    @property
    @pulumi.getter
    def risks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]]:
        """
        Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        """
        return pulumi.get(self, "risks")

    @risks.setter
    def risks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRiskArgs']]]]):
        pulumi.set(self, "risks", value)

    @property
    @pulumi.getter
    def technology(self) -> Optional[pulumi.Input[str]]:
        """
        Application technology filter.
        """
        return pulumi.get(self, "technology")

    @technology.setter
    def technology(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "technology", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Application group type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
    def vendor(self) -> Optional[pulumi.Input[str]]:
        """
        Application vendor filter.
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupApplicationArgs']]]]] = None,
                 behavior: Optional[pulumi.Input[str]] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupCategoryArgs']]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 popularity: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[str]] = None,
                 risks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRiskArgs']]]]] = None,
                 technology: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure firewall application groups.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.application.Group("trname",
            categories=[fortios.application.GroupCategoryArgs(
                id=2,
            )],
            comment="group1 test",
            type="category")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application Group can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:application/group:Group labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:application/group:Group labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupApplicationArgs']]]] applications: Application ID list. The structure of `application` block is documented below.
        :param pulumi.Input[str] behavior: Application behavior filter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupCategoryArgs']]]] categories: Application category ID list. The structure of `category` block is documented below.
        :param pulumi.Input[str] comment: Comment
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Application group name.
        :param pulumi.Input[str] popularity: Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        :param pulumi.Input[str] protocols: Application protocol filter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRiskArgs']]]] risks: Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        :param pulumi.Input[str] technology: Application technology filter.
        :param pulumi.Input[str] type: Application group type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: Application vendor filter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure firewall application groups.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.application.Group("trname",
            categories=[fortios.application.GroupCategoryArgs(
                id=2,
            )],
            comment="group1 test",
            type="category")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application Group can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:application/group:Group labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:application/group:Group labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupApplicationArgs']]]]] = None,
                 behavior: Optional[pulumi.Input[str]] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupCategoryArgs']]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 popularity: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[str]] = None,
                 risks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRiskArgs']]]]] = None,
                 technology: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["applications"] = applications
            __props__.__dict__["behavior"] = behavior
            __props__.__dict__["categories"] = categories
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["popularity"] = popularity
            __props__.__dict__["protocols"] = protocols
            __props__.__dict__["risks"] = risks
            __props__.__dict__["technology"] = technology
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vendor"] = vendor
        super(Group, __self__).__init__(
            'fortios:application/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            applications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupApplicationArgs']]]]] = None,
            behavior: Optional[pulumi.Input[str]] = None,
            categories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupCategoryArgs']]]]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            popularity: Optional[pulumi.Input[str]] = None,
            protocols: Optional[pulumi.Input[str]] = None,
            risks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRiskArgs']]]]] = None,
            technology: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vendor: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupApplicationArgs']]]] applications: Application ID list. The structure of `application` block is documented below.
        :param pulumi.Input[str] behavior: Application behavior filter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupCategoryArgs']]]] categories: Application category ID list. The structure of `category` block is documented below.
        :param pulumi.Input[str] comment: Comment
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Application group name.
        :param pulumi.Input[str] popularity: Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        :param pulumi.Input[str] protocols: Application protocol filter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRiskArgs']]]] risks: Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        :param pulumi.Input[str] technology: Application technology filter.
        :param pulumi.Input[str] type: Application group type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: Application vendor filter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["applications"] = applications
        __props__.__dict__["behavior"] = behavior
        __props__.__dict__["categories"] = categories
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["popularity"] = popularity
        __props__.__dict__["protocols"] = protocols
        __props__.__dict__["risks"] = risks
        __props__.__dict__["technology"] = technology
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vendor"] = vendor
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def applications(self) -> pulumi.Output[Optional[Sequence['outputs.GroupApplication']]]:
        """
        Application ID list. The structure of `application` block is documented below.
        """
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter
    def behavior(self) -> pulumi.Output[str]:
        """
        Application behavior filter.
        """
        return pulumi.get(self, "behavior")

    @property
    @pulumi.getter
    def categories(self) -> pulumi.Output[Optional[Sequence['outputs.GroupCategory']]]:
        """
        Application category ID list. The structure of `category` block is documented below.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Application group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def popularity(self) -> pulumi.Output[str]:
        """
        Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        """
        return pulumi.get(self, "popularity")

    @property
    @pulumi.getter
    def protocols(self) -> pulumi.Output[str]:
        """
        Application protocol filter.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def risks(self) -> pulumi.Output[Optional[Sequence['outputs.GroupRisk']]]:
        """
        Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        """
        return pulumi.get(self, "risks")

    @property
    @pulumi.getter
    def technology(self) -> pulumi.Output[str]:
        """
        Application technology filter.
        """
        return pulumi.get(self, "technology")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Application group type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vendor(self) -> pulumi.Output[str]:
        """
        Application vendor filter.
        """
        return pulumi.get(self, "vendor")

