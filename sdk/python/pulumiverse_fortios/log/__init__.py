# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from .customfield import *
from .eventfilter import *
from .guidisplay import *
from .setting import *
from .syslog_setting import *
from .threatweight import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.log.disk as __disk
    disk = __disk
    import pulumiverse_fortios.log.fortianalyzer as __fortianalyzer
    fortianalyzer = __fortianalyzer
    import pulumiverse_fortios.log.fortiguard as __fortiguard
    fortiguard = __fortiguard
    import pulumiverse_fortios.log.memory as __memory
    memory = __memory
    import pulumiverse_fortios.log.nulldevice as __nulldevice
    nulldevice = __nulldevice
    import pulumiverse_fortios.log.syslogd as __syslogd
    syslogd = __syslogd
    import pulumiverse_fortios.log.tacacsaccounting as __tacacsaccounting
    tacacsaccounting = __tacacsaccounting
    import pulumiverse_fortios.log.webtrends as __webtrends
    webtrends = __webtrends
else:
    disk = _utilities.lazy_import('pulumiverse_fortios.log.disk')
    fortianalyzer = _utilities.lazy_import('pulumiverse_fortios.log.fortianalyzer')
    fortiguard = _utilities.lazy_import('pulumiverse_fortios.log.fortiguard')
    memory = _utilities.lazy_import('pulumiverse_fortios.log.memory')
    nulldevice = _utilities.lazy_import('pulumiverse_fortios.log.nulldevice')
    syslogd = _utilities.lazy_import('pulumiverse_fortios.log.syslogd')
    tacacsaccounting = _utilities.lazy_import('pulumiverse_fortios.log.tacacsaccounting')
    webtrends = _utilities.lazy_import('pulumiverse_fortios.log.webtrends')

