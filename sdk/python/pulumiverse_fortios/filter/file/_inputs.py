# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'ProfileRuleArgs',
    'ProfileRuleFileTypeArgs',
]

@pulumi.input_type
class ProfileRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 file_types: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleFileTypeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password_protected: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Action taken for matched file. Valid values: `log-only`, `block`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] direction: Traffic direction. On FortiOS versions 6.4.1-7.4.1: HTTP, FTP, SSH, CIFS only. On FortiOS versions >= 7.4.2: HTTP, FTP, SSH, CIFS, and MAPI only. Valid values: `incoming`, `outgoing`, `any`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileRuleFileTypeArgs']]] file_types: Select file type. The structure of `file_type` block is documented below.
        :param pulumi.Input[str] name: File-filter rule name.
        :param pulumi.Input[str] password_protected: Match password-protected files. Valid values: `yes`, `any`.
        :param pulumi.Input[str] protocol: Protocols to apply rule to. Valid values: `http`, `ftp`, `smtp`, `imap`, `pop3`, `mapi`, `cifs`, `ssh`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if file_types is not None:
            pulumi.set(__self__, "file_types", file_types)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password_protected is not None:
            pulumi.set(__self__, "password_protected", password_protected)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action taken for matched file. Valid values: `log-only`, `block`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Traffic direction. On FortiOS versions 6.4.1-7.4.1: HTTP, FTP, SSH, CIFS only. On FortiOS versions >= 7.4.2: HTTP, FTP, SSH, CIFS, and MAPI only. Valid values: `incoming`, `outgoing`, `any`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="fileTypes")
    def file_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleFileTypeArgs']]]]:
        """
        Select file type. The structure of `file_type` block is documented below.
        """
        return pulumi.get(self, "file_types")

    @file_types.setter
    def file_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleFileTypeArgs']]]]):
        pulumi.set(self, "file_types", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File-filter rule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="passwordProtected")
    def password_protected(self) -> Optional[pulumi.Input[str]]:
        """
        Match password-protected files. Valid values: `yes`, `any`.
        """
        return pulumi.get(self, "password_protected")

    @password_protected.setter
    def password_protected(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_protected", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocols to apply rule to. Valid values: `http`, `ftp`, `smtp`, `imap`, `pop3`, `mapi`, `cifs`, `ssh`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class ProfileRuleFileTypeArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: File type name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


