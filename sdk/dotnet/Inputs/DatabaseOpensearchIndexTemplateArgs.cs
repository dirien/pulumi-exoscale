// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Inputs
{

    public sealed class DatabaseOpensearchIndexTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("mappingNestedObjectsLimit")]
        public Input<int>? MappingNestedObjectsLimit { get; set; }

        [Input("numberOfReplicas")]
        public Input<int>? NumberOfReplicas { get; set; }

        [Input("numberOfShards")]
        public Input<int>? NumberOfShards { get; set; }

        public DatabaseOpensearchIndexTemplateArgs()
        {
        }
        public static new DatabaseOpensearchIndexTemplateArgs Empty => new DatabaseOpensearchIndexTemplateArgs();
    }
}
