// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    /// <summary>
    /// Manage Exoscale [DNS](https://community.exoscale.com/documentation/dns/) Domain Records.
    /// 
    /// Corresponding data source: exoscale_domain_record.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Exoscale = Pulumiverse.Exoscale;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDomain = new Exoscale.Domain("my_domain", new()
    ///     {
    ///         Name = "example.net",
    ///     });
    /// 
    ///     var myHost = new Exoscale.DomainRecord("my_host", new()
    ///     {
    ///         Domain = myDomain.Id,
    ///         Name = "my-host",
    ///         RecordType = "A",
    ///         Content = "1.2.3.4",
    ///     });
    /// 
    ///     var myHostAlias = new Exoscale.DomainRecord("my_host_alias", new()
    ///     {
    ///         Domain = myDomain.Id,
    ///         Name = "my-host-alias",
    ///         RecordType = "CNAME",
    ///         Content = myHost.Hostname,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Please refer to the examples
    /// directory for complete configuration examples.
    /// 
    /// ## Import
    /// 
    /// An existing DNS domain record may be imported by `&lt;ID&gt;`:
    /// 
    /// ```sh
    /// $ pulumi import exoscale:index/domainRecord:DomainRecord \
    /// ```
    /// 
    ///   exoscale_domain_record.my_host \
    /// 
    ///   f81d4fae-7dec-11d0-a765-00a0c91e6bf6
    /// </summary>
    [ExoscaleResourceType("exoscale:index/domainRecord:DomainRecord")]
    public partial class DomainRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The record value.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The normalized value of the record
        /// </summary>
        [Output("contentNormalized")]
        public Output<string> ContentNormalized { get; private set; } = null!;

        /// <summary>
        /// ❗ The parent exoscale.Domain to attach the record to.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The record *Fully Qualified Domain Name* (FQDN). Useful for aliasing `A`/`AAAA` records with `CNAME`.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The record name, Leave blank (`""`) to create a root record (similar to using `@` in a DNS zone file).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The record priority (for types that support it; minimum `0`).
        /// </summary>
        [Output("prio")]
        public Output<int> Prio { get; private set; } = null!;

        /// <summary>
        /// ❗ The record type (`A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `HINFO`, `MX`, `NAPTR`, `NS`, `POOL`, `SPF`, `SRV`, `SSHFP`, `TXT`, `URL`).
        /// </summary>
        [Output("recordType")]
        public Output<string> RecordType { get; private set; } = null!;

        /// <summary>
        /// The record TTL (seconds; minimum `0`; default: `3600`).
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a DomainRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainRecord(string name, DomainRecordArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/domainRecord:DomainRecord", name, args ?? new DomainRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainRecord(string name, Input<string> id, DomainRecordState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/domainRecord:DomainRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainRecord Get(string name, Input<string> id, DomainRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainRecord(name, id, state, options);
        }
    }

    public sealed class DomainRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The record value.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// ❗ The parent exoscale.Domain to attach the record to.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The record name, Leave blank (`""`) to create a root record (similar to using `@` in a DNS zone file).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The record priority (for types that support it; minimum `0`).
        /// </summary>
        [Input("prio")]
        public Input<int>? Prio { get; set; }

        /// <summary>
        /// ❗ The record type (`A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `HINFO`, `MX`, `NAPTR`, `NS`, `POOL`, `SPF`, `SRV`, `SSHFP`, `TXT`, `URL`).
        /// </summary>
        [Input("recordType", required: true)]
        public Input<string> RecordType { get; set; } = null!;

        /// <summary>
        /// The record TTL (seconds; minimum `0`; default: `3600`).
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public DomainRecordArgs()
        {
        }
        public static new DomainRecordArgs Empty => new DomainRecordArgs();
    }

    public sealed class DomainRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The record value.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The normalized value of the record
        /// </summary>
        [Input("contentNormalized")]
        public Input<string>? ContentNormalized { get; set; }

        /// <summary>
        /// ❗ The parent exoscale.Domain to attach the record to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The record *Fully Qualified Domain Name* (FQDN). Useful for aliasing `A`/`AAAA` records with `CNAME`.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The record name, Leave blank (`""`) to create a root record (similar to using `@` in a DNS zone file).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The record priority (for types that support it; minimum `0`).
        /// </summary>
        [Input("prio")]
        public Input<int>? Prio { get; set; }

        /// <summary>
        /// ❗ The record type (`A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `HINFO`, `MX`, `NAPTR`, `NS`, `POOL`, `SPF`, `SRV`, `SSHFP`, `TXT`, `URL`).
        /// </summary>
        [Input("recordType")]
        public Input<string>? RecordType { get; set; }

        /// <summary>
        /// The record TTL (seconds; minimum `0`; default: `3600`).
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public DomainRecordState()
        {
        }
        public static new DomainRecordState Empty => new DomainRecordState();
    }
}
