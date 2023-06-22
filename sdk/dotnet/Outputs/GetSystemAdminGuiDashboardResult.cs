// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class GetSystemAdminGuiDashboardResult
    {
        /// <summary>
        /// Number of columns.
        /// </summary>
        public readonly int Columns;
        /// <summary>
        /// Select menu ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Layout type.
        /// </summary>
        public readonly string LayoutType;
        /// <summary>
        /// Specify the name of the desired system admin.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Dashboard scope.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// Dashboard widgets. The structure of `widget` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAdminGuiDashboardWidgetResult> Widgets;

        [OutputConstructor]
        private GetSystemAdminGuiDashboardResult(
            int columns,

            int id,

            string layoutType,

            string name,

            string scope,

            ImmutableArray<Outputs.GetSystemAdminGuiDashboardWidgetResult> widgets)
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
