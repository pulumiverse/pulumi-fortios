# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from ... import _utilities
import typing
# Export this package's modules as members:
from .filter import *
from .overridefilter import *
from .overridesetting import *
from .setting import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.log.syslogd.v2 as __v2
    v2 = __v2
    import pulumiverse_fortios.log.syslogd.v3 as __v3
    v3 = __v3
    import pulumiverse_fortios.log.syslogd.v4 as __v4
    v4 = __v4
else:
    v2 = _utilities.lazy_import('pulumiverse_fortios.log.syslogd.v2')
    v3 = _utilities.lazy_import('pulumiverse_fortios.log.syslogd.v3')
    v4 = _utilities.lazy_import('pulumiverse_fortios.log.syslogd.v4')

