# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PasswordpolicyArgs', 'Passwordpolicy']

@pulumi.input_type
class PasswordpolicyArgs:
    def __init__(__self__, *,
                 expire_days: Optional[pulumi.Input[int]] = None,
                 expire_status: Optional[pulumi.Input[str]] = None,
                 expired_password_renewal: Optional[pulumi.Input[str]] = None,
                 min_change_characters: Optional[pulumi.Input[int]] = None,
                 min_lower_case_letter: Optional[pulumi.Input[int]] = None,
                 min_non_alphanumeric: Optional[pulumi.Input[int]] = None,
                 min_number: Optional[pulumi.Input[int]] = None,
                 min_upper_case_letter: Optional[pulumi.Input[int]] = None,
                 minimum_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reuse_password: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Passwordpolicy resource.
        :param pulumi.Input[int] expire_days: Time in days before the user's password expires.
        :param pulumi.Input[str] expire_status: Enable/disable password expiration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] expired_password_renewal: Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] min_change_characters: Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        :param pulumi.Input[int] min_lower_case_letter: Minimum number of lowercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_non_alphanumeric: Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_number: Minimum number of numeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_upper_case_letter: Minimum number of uppercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] minimum_length: Minimum password length (8 - 128, default = 8).
        :param pulumi.Input[str] name: Password policy name.
        :param pulumi.Input[str] reuse_password: Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] warn_days: Time in days before a password expiration warning message is displayed to the user upon login.
        """
        if expire_days is not None:
            pulumi.set(__self__, "expire_days", expire_days)
        if expire_status is not None:
            pulumi.set(__self__, "expire_status", expire_status)
        if expired_password_renewal is not None:
            pulumi.set(__self__, "expired_password_renewal", expired_password_renewal)
        if min_change_characters is not None:
            pulumi.set(__self__, "min_change_characters", min_change_characters)
        if min_lower_case_letter is not None:
            pulumi.set(__self__, "min_lower_case_letter", min_lower_case_letter)
        if min_non_alphanumeric is not None:
            pulumi.set(__self__, "min_non_alphanumeric", min_non_alphanumeric)
        if min_number is not None:
            pulumi.set(__self__, "min_number", min_number)
        if min_upper_case_letter is not None:
            pulumi.set(__self__, "min_upper_case_letter", min_upper_case_letter)
        if minimum_length is not None:
            pulumi.set(__self__, "minimum_length", minimum_length)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if reuse_password is not None:
            pulumi.set(__self__, "reuse_password", reuse_password)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if warn_days is not None:
            pulumi.set(__self__, "warn_days", warn_days)

    @property
    @pulumi.getter(name="expireDays")
    def expire_days(self) -> Optional[pulumi.Input[int]]:
        """
        Time in days before the user's password expires.
        """
        return pulumi.get(self, "expire_days")

    @expire_days.setter
    def expire_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expire_days", value)

    @property
    @pulumi.getter(name="expireStatus")
    def expire_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable password expiration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expire_status")

    @expire_status.setter
    def expire_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_status", value)

    @property
    @pulumi.getter(name="expiredPasswordRenewal")
    def expired_password_renewal(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expired_password_renewal")

    @expired_password_renewal.setter
    def expired_password_renewal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_password_renewal", value)

    @property
    @pulumi.getter(name="minChangeCharacters")
    def min_change_characters(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        """
        return pulumi.get(self, "min_change_characters")

    @min_change_characters.setter
    def min_change_characters(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_change_characters", value)

    @property
    @pulumi.getter(name="minLowerCaseLetter")
    def min_lower_case_letter(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of lowercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_lower_case_letter")

    @min_lower_case_letter.setter
    def min_lower_case_letter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_lower_case_letter", value)

    @property
    @pulumi.getter(name="minNonAlphanumeric")
    def min_non_alphanumeric(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_non_alphanumeric")

    @min_non_alphanumeric.setter
    def min_non_alphanumeric(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_non_alphanumeric", value)

    @property
    @pulumi.getter(name="minNumber")
    def min_number(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of numeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_number")

    @min_number.setter
    def min_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_number", value)

    @property
    @pulumi.getter(name="minUpperCaseLetter")
    def min_upper_case_letter(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of uppercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_upper_case_letter")

    @min_upper_case_letter.setter
    def min_upper_case_letter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_upper_case_letter", value)

    @property
    @pulumi.getter(name="minimumLength")
    def minimum_length(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum password length (8 - 128, default = 8).
        """
        return pulumi.get(self, "minimum_length")

    @minimum_length.setter
    def minimum_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_length", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Password policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="reusePassword")
    def reuse_password(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "reuse_password")

    @reuse_password.setter
    def reuse_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reuse_password", value)

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
    @pulumi.getter(name="warnDays")
    def warn_days(self) -> Optional[pulumi.Input[int]]:
        """
        Time in days before a password expiration warning message is displayed to the user upon login.
        """
        return pulumi.get(self, "warn_days")

    @warn_days.setter
    def warn_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "warn_days", value)


@pulumi.input_type
class _PasswordpolicyState:
    def __init__(__self__, *,
                 expire_days: Optional[pulumi.Input[int]] = None,
                 expire_status: Optional[pulumi.Input[str]] = None,
                 expired_password_renewal: Optional[pulumi.Input[str]] = None,
                 min_change_characters: Optional[pulumi.Input[int]] = None,
                 min_lower_case_letter: Optional[pulumi.Input[int]] = None,
                 min_non_alphanumeric: Optional[pulumi.Input[int]] = None,
                 min_number: Optional[pulumi.Input[int]] = None,
                 min_upper_case_letter: Optional[pulumi.Input[int]] = None,
                 minimum_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reuse_password: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Passwordpolicy resources.
        :param pulumi.Input[int] expire_days: Time in days before the user's password expires.
        :param pulumi.Input[str] expire_status: Enable/disable password expiration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] expired_password_renewal: Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] min_change_characters: Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        :param pulumi.Input[int] min_lower_case_letter: Minimum number of lowercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_non_alphanumeric: Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_number: Minimum number of numeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_upper_case_letter: Minimum number of uppercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] minimum_length: Minimum password length (8 - 128, default = 8).
        :param pulumi.Input[str] name: Password policy name.
        :param pulumi.Input[str] reuse_password: Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] warn_days: Time in days before a password expiration warning message is displayed to the user upon login.
        """
        if expire_days is not None:
            pulumi.set(__self__, "expire_days", expire_days)
        if expire_status is not None:
            pulumi.set(__self__, "expire_status", expire_status)
        if expired_password_renewal is not None:
            pulumi.set(__self__, "expired_password_renewal", expired_password_renewal)
        if min_change_characters is not None:
            pulumi.set(__self__, "min_change_characters", min_change_characters)
        if min_lower_case_letter is not None:
            pulumi.set(__self__, "min_lower_case_letter", min_lower_case_letter)
        if min_non_alphanumeric is not None:
            pulumi.set(__self__, "min_non_alphanumeric", min_non_alphanumeric)
        if min_number is not None:
            pulumi.set(__self__, "min_number", min_number)
        if min_upper_case_letter is not None:
            pulumi.set(__self__, "min_upper_case_letter", min_upper_case_letter)
        if minimum_length is not None:
            pulumi.set(__self__, "minimum_length", minimum_length)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if reuse_password is not None:
            pulumi.set(__self__, "reuse_password", reuse_password)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if warn_days is not None:
            pulumi.set(__self__, "warn_days", warn_days)

    @property
    @pulumi.getter(name="expireDays")
    def expire_days(self) -> Optional[pulumi.Input[int]]:
        """
        Time in days before the user's password expires.
        """
        return pulumi.get(self, "expire_days")

    @expire_days.setter
    def expire_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expire_days", value)

    @property
    @pulumi.getter(name="expireStatus")
    def expire_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable password expiration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expire_status")

    @expire_status.setter
    def expire_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_status", value)

    @property
    @pulumi.getter(name="expiredPasswordRenewal")
    def expired_password_renewal(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expired_password_renewal")

    @expired_password_renewal.setter
    def expired_password_renewal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_password_renewal", value)

    @property
    @pulumi.getter(name="minChangeCharacters")
    def min_change_characters(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        """
        return pulumi.get(self, "min_change_characters")

    @min_change_characters.setter
    def min_change_characters(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_change_characters", value)

    @property
    @pulumi.getter(name="minLowerCaseLetter")
    def min_lower_case_letter(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of lowercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_lower_case_letter")

    @min_lower_case_letter.setter
    def min_lower_case_letter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_lower_case_letter", value)

    @property
    @pulumi.getter(name="minNonAlphanumeric")
    def min_non_alphanumeric(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_non_alphanumeric")

    @min_non_alphanumeric.setter
    def min_non_alphanumeric(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_non_alphanumeric", value)

    @property
    @pulumi.getter(name="minNumber")
    def min_number(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of numeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_number")

    @min_number.setter
    def min_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_number", value)

    @property
    @pulumi.getter(name="minUpperCaseLetter")
    def min_upper_case_letter(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of uppercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_upper_case_letter")

    @min_upper_case_letter.setter
    def min_upper_case_letter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_upper_case_letter", value)

    @property
    @pulumi.getter(name="minimumLength")
    def minimum_length(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum password length (8 - 128, default = 8).
        """
        return pulumi.get(self, "minimum_length")

    @minimum_length.setter
    def minimum_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_length", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Password policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="reusePassword")
    def reuse_password(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "reuse_password")

    @reuse_password.setter
    def reuse_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reuse_password", value)

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
    @pulumi.getter(name="warnDays")
    def warn_days(self) -> Optional[pulumi.Input[int]]:
        """
        Time in days before a password expiration warning message is displayed to the user upon login.
        """
        return pulumi.get(self, "warn_days")

    @warn_days.setter
    def warn_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "warn_days", value)


class Passwordpolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_days: Optional[pulumi.Input[int]] = None,
                 expire_status: Optional[pulumi.Input[str]] = None,
                 expired_password_renewal: Optional[pulumi.Input[str]] = None,
                 min_change_characters: Optional[pulumi.Input[int]] = None,
                 min_lower_case_letter: Optional[pulumi.Input[int]] = None,
                 min_non_alphanumeric: Optional[pulumi.Input[int]] = None,
                 min_number: Optional[pulumi.Input[int]] = None,
                 min_upper_case_letter: Optional[pulumi.Input[int]] = None,
                 minimum_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reuse_password: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configure user password policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Passwordpolicy("trname",
            expire_days=22,
            warn_days=13)
        ```

        ## Import

        User PasswordPolicy can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/passwordpolicy:Passwordpolicy labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/passwordpolicy:Passwordpolicy labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] expire_days: Time in days before the user's password expires.
        :param pulumi.Input[str] expire_status: Enable/disable password expiration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] expired_password_renewal: Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] min_change_characters: Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        :param pulumi.Input[int] min_lower_case_letter: Minimum number of lowercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_non_alphanumeric: Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_number: Minimum number of numeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_upper_case_letter: Minimum number of uppercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] minimum_length: Minimum password length (8 - 128, default = 8).
        :param pulumi.Input[str] name: Password policy name.
        :param pulumi.Input[str] reuse_password: Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] warn_days: Time in days before a password expiration warning message is displayed to the user upon login.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PasswordpolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure user password policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Passwordpolicy("trname",
            expire_days=22,
            warn_days=13)
        ```

        ## Import

        User PasswordPolicy can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/passwordpolicy:Passwordpolicy labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/passwordpolicy:Passwordpolicy labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param PasswordpolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PasswordpolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_days: Optional[pulumi.Input[int]] = None,
                 expire_status: Optional[pulumi.Input[str]] = None,
                 expired_password_renewal: Optional[pulumi.Input[str]] = None,
                 min_change_characters: Optional[pulumi.Input[int]] = None,
                 min_lower_case_letter: Optional[pulumi.Input[int]] = None,
                 min_non_alphanumeric: Optional[pulumi.Input[int]] = None,
                 min_number: Optional[pulumi.Input[int]] = None,
                 min_upper_case_letter: Optional[pulumi.Input[int]] = None,
                 minimum_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reuse_password: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PasswordpolicyArgs.__new__(PasswordpolicyArgs)

            __props__.__dict__["expire_days"] = expire_days
            __props__.__dict__["expire_status"] = expire_status
            __props__.__dict__["expired_password_renewal"] = expired_password_renewal
            __props__.__dict__["min_change_characters"] = min_change_characters
            __props__.__dict__["min_lower_case_letter"] = min_lower_case_letter
            __props__.__dict__["min_non_alphanumeric"] = min_non_alphanumeric
            __props__.__dict__["min_number"] = min_number
            __props__.__dict__["min_upper_case_letter"] = min_upper_case_letter
            __props__.__dict__["minimum_length"] = minimum_length
            __props__.__dict__["name"] = name
            __props__.__dict__["reuse_password"] = reuse_password
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["warn_days"] = warn_days
        super(Passwordpolicy, __self__).__init__(
            'fortios:user/passwordpolicy:Passwordpolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expire_days: Optional[pulumi.Input[int]] = None,
            expire_status: Optional[pulumi.Input[str]] = None,
            expired_password_renewal: Optional[pulumi.Input[str]] = None,
            min_change_characters: Optional[pulumi.Input[int]] = None,
            min_lower_case_letter: Optional[pulumi.Input[int]] = None,
            min_non_alphanumeric: Optional[pulumi.Input[int]] = None,
            min_number: Optional[pulumi.Input[int]] = None,
            min_upper_case_letter: Optional[pulumi.Input[int]] = None,
            minimum_length: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            reuse_password: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            warn_days: Optional[pulumi.Input[int]] = None) -> 'Passwordpolicy':
        """
        Get an existing Passwordpolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] expire_days: Time in days before the user's password expires.
        :param pulumi.Input[str] expire_status: Enable/disable password expiration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] expired_password_renewal: Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] min_change_characters: Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        :param pulumi.Input[int] min_lower_case_letter: Minimum number of lowercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_non_alphanumeric: Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_number: Minimum number of numeric characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] min_upper_case_letter: Minimum number of uppercase characters in password (0 - 128, default = 0).
        :param pulumi.Input[int] minimum_length: Minimum password length (8 - 128, default = 8).
        :param pulumi.Input[str] name: Password policy name.
        :param pulumi.Input[str] reuse_password: Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] warn_days: Time in days before a password expiration warning message is displayed to the user upon login.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PasswordpolicyState.__new__(_PasswordpolicyState)

        __props__.__dict__["expire_days"] = expire_days
        __props__.__dict__["expire_status"] = expire_status
        __props__.__dict__["expired_password_renewal"] = expired_password_renewal
        __props__.__dict__["min_change_characters"] = min_change_characters
        __props__.__dict__["min_lower_case_letter"] = min_lower_case_letter
        __props__.__dict__["min_non_alphanumeric"] = min_non_alphanumeric
        __props__.__dict__["min_number"] = min_number
        __props__.__dict__["min_upper_case_letter"] = min_upper_case_letter
        __props__.__dict__["minimum_length"] = minimum_length
        __props__.__dict__["name"] = name
        __props__.__dict__["reuse_password"] = reuse_password
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["warn_days"] = warn_days
        return Passwordpolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expireDays")
    def expire_days(self) -> pulumi.Output[int]:
        """
        Time in days before the user's password expires.
        """
        return pulumi.get(self, "expire_days")

    @property
    @pulumi.getter(name="expireStatus")
    def expire_status(self) -> pulumi.Output[str]:
        """
        Enable/disable password expiration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expire_status")

    @property
    @pulumi.getter(name="expiredPasswordRenewal")
    def expired_password_renewal(self) -> pulumi.Output[str]:
        """
        Enable/disable renewal of a password that already is expired. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "expired_password_renewal")

    @property
    @pulumi.getter(name="minChangeCharacters")
    def min_change_characters(self) -> pulumi.Output[int]:
        """
        Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
        """
        return pulumi.get(self, "min_change_characters")

    @property
    @pulumi.getter(name="minLowerCaseLetter")
    def min_lower_case_letter(self) -> pulumi.Output[int]:
        """
        Minimum number of lowercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_lower_case_letter")

    @property
    @pulumi.getter(name="minNonAlphanumeric")
    def min_non_alphanumeric(self) -> pulumi.Output[int]:
        """
        Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_non_alphanumeric")

    @property
    @pulumi.getter(name="minNumber")
    def min_number(self) -> pulumi.Output[int]:
        """
        Minimum number of numeric characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_number")

    @property
    @pulumi.getter(name="minUpperCaseLetter")
    def min_upper_case_letter(self) -> pulumi.Output[int]:
        """
        Minimum number of uppercase characters in password (0 - 128, default = 0).
        """
        return pulumi.get(self, "min_upper_case_letter")

    @property
    @pulumi.getter(name="minimumLength")
    def minimum_length(self) -> pulumi.Output[int]:
        """
        Minimum password length (8 - 128, default = 8).
        """
        return pulumi.get(self, "minimum_length")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Password policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reusePassword")
    def reuse_password(self) -> pulumi.Output[str]:
        """
        Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "reuse_password")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="warnDays")
    def warn_days(self) -> pulumi.Output[int]:
        """
        Time in days before a password expiration warning message is displayed to the user upon login.
        """
        return pulumi.get(self, "warn_days")

