// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class AdminGuiDashboard
    {
        /// <summary>
        /// Number of columns.
        /// </summary>
        public readonly int? Columns;
        /// <summary>
        /// Dashboard ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Layout type. Valid values: `responsive`, `fixed`.
        /// </summary>
        public readonly string? LayoutType;
        /// <summary>
        /// Dashboard name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Dashboard scope. Valid values: `global`, `vdom`.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Dashboard widgets. The structure of `widget` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdminGuiDashboardWidget> Widgets;

        [OutputConstructor]
        private AdminGuiDashboard(
            int? columns,

            int? id,

            string? layoutType,

            string? name,

            string? scope,

            ImmutableArray<Outputs.AdminGuiDashboardWidget> widgets)
        {
            Columns = columns;
            Id = id;
            LayoutType = layoutType;
            Name = name;
            Scope = scope;
            Widgets = widgets;
        }
    }
}
