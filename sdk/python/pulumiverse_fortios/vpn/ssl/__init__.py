# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from ... import _utilities
import typing
# Export this package's modules as members:
from .client import *
from .settings import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.vpn.ssl.web as __web
    web = __web
else:
    web = _utilities.lazy_import('pulumiverse_fortios.vpn.ssl.web')
