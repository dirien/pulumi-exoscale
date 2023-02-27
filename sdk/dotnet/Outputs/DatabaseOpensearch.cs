// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Outputs
{

    [OutputType]
    public sealed class DatabaseOpensearch
    {
        public readonly Outputs.DatabaseOpensearchDashboards? Dashboards;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string? ForkFromService;
        /// <summary>
        /// Allows you to create glob style patterns and set a max number of indexes matching this pattern you want to keep. Creating indexes exceeding this value will cause the oldest one to get deleted. You could for example create a pattern looking like 'logs.?' and then create index logs.1, logs.2 etc, it will delete logs.1 once you create logs.6. Do note 'logs.?' does not apply to logs.10. Note: Setting max_index_count to 0 will do nothing and the pattern gets ignored.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseOpensearchIndexPattern> IndexPatterns;
        /// <summary>
        /// Template settings for all new indexes
        /// </summary>
        public readonly Outputs.DatabaseOpensearchIndexTemplate? IndexTemplate;
        /// <summary>
        /// Allow incoming connections from this list of CIDR address block, e.g. `["10.20.0.0/16"]`
        /// </summary>
        public readonly ImmutableArray<string> IpFilters;
        /// <summary>
        /// Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.
        /// </summary>
        public readonly bool? KeepIndexRefreshInterval;
        /// <summary>
        /// Maximum number of indexes to keep before deleting the oldest one (Minimum value is `0`)
        /// </summary>
        public readonly int? MaxIndexCount;
        public readonly string? RecoveryBackupName;
        public readonly string? Settings;
        /// <summary>
        /// OpenSearch major version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private DatabaseOpensearch(
            Outputs.DatabaseOpensearchDashboards? dashboards,

            string? forkFromService,

            ImmutableArray<Outputs.DatabaseOpensearchIndexPattern> indexPatterns,

            Outputs.DatabaseOpensearchIndexTemplate? indexTemplate,

            ImmutableArray<string> ipFilters,

            bool? keepIndexRefreshInterval,

            int? maxIndexCount,

            string? recoveryBackupName,

            string? settings,

            string? version)
        {
            Dashboards = dashboards;
            ForkFromService = forkFromService;
            IndexPatterns = indexPatterns;
            IndexTemplate = indexTemplate;
            IpFilters = ipFilters;
            KeepIndexRefreshInterval = keepIndexRefreshInterval;
            MaxIndexCount = maxIndexCount;
            RecoveryBackupName = recoveryBackupName;
            Settings = settings;
            Version = version;
        }
    }
}
