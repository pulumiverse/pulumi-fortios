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

__all__ = ['AcmeArgs', 'Acme']

@pulumi.input_type
class AcmeArgs:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Acme resource.
        :param pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]] accounts: ACME accounts list. The structure of `accounts` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]] interfaces: Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to connect to the ACME server.
        :param pulumi.Input[str] source_ip6: Source IPv6 address used to connect to the ACME server.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 is not None:
            pulumi.set(__self__, "source_ip6", source_ip6)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]]:
        """
        ACME accounts list. The structure of `accounts` block is documented below.
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]]):
        pulumi.set(self, "accounts", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]]:
        """
        Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv4 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv6 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip6")

    @source_ip6.setter
    def source_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip6", value)

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
class _AcmeState:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Acme resources.
        :param pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]] accounts: ACME accounts list. The structure of `accounts` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]] interfaces: Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to connect to the ACME server.
        :param pulumi.Input[str] source_ip6: Source IPv6 address used to connect to the ACME server.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 is not None:
            pulumi.set(__self__, "source_ip6", source_ip6)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]]:
        """
        ACME accounts list. The structure of `accounts` block is documented below.
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeAccountArgs']]]]):
        pulumi.set(self, "accounts", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]]:
        """
        Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AcmeInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv4 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv6 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip6")

    @source_ip6.setter
    def source_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip6", value)

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


class Acme(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeAccountArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeInterfaceArgs']]]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure ACME client. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        System Acme can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/acme:Acme labelname SystemAcme
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/acme:Acme labelname SystemAcme
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeAccountArgs']]]] accounts: ACME accounts list. The structure of `accounts` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeInterfaceArgs']]]] interfaces: Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to connect to the ACME server.
        :param pulumi.Input[str] source_ip6: Source IPv6 address used to connect to the ACME server.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AcmeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure ACME client. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        System Acme can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/acme:Acme labelname SystemAcme
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/acme:Acme labelname SystemAcme
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AcmeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AcmeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeAccountArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeInterfaceArgs']]]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AcmeArgs.__new__(AcmeArgs)

            __props__.__dict__["accounts"] = accounts
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["source_ip6"] = source_ip6
            __props__.__dict__["vdomparam"] = vdomparam
        super(Acme, __self__).__init__(
            'fortios:system/acme:Acme',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeAccountArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeInterfaceArgs']]]]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            source_ip6: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Acme':
        """
        Get an existing Acme resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeAccountArgs']]]] accounts: ACME accounts list. The structure of `accounts` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcmeInterfaceArgs']]]] interfaces: Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to connect to the ACME server.
        :param pulumi.Input[str] source_ip6: Source IPv6 address used to connect to the ACME server.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AcmeState.__new__(_AcmeState)

        __props__.__dict__["accounts"] = accounts
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["source_ip6"] = source_ip6
        __props__.__dict__["vdomparam"] = vdomparam
        return Acme(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accounts(self) -> pulumi.Output[Optional[Sequence['outputs.AcmeAccount']]]:
        """
        ACME accounts list. The structure of `accounts` block is documented below.
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.AcmeInterface']]]:
        """
        Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IPv4 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> pulumi.Output[str]:
        """
        Source IPv6 address used to connect to the ACME server.
        """
        return pulumi.get(self, "source_ip6")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

