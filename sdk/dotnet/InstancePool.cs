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
    /// ## Import
    /// 
    /// An existing instance pool may be imported by `&lt;ID&gt;@&lt;zone&gt;`
    /// 
    /// ```sh
    ///  $ pulumi import exoscale:index/instancePool:InstancePool \
    /// ```
    /// 
    ///  exoscale_instance_pool.my_instance_pool \
    /// 
    ///  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
    /// </summary>
    [ExoscaleResourceType("exoscale:index/instancePool:InstancePool")]
    public partial class InstancePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
        /// </summary>
        [Output("affinityGroupIds")]
        public Output<ImmutableArray<string>> AffinityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Output("deployTargetId")]
        public Output<string?> DeployTargetId { get; private set; } = null!;

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The managed instances disk size (GiB).
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// A list of exoscale*elastic*ip (IDs).
        /// </summary>
        [Output("elasticIpIds")]
        public Output<ImmutableArray<string>> ElasticIpIds { get; private set; } = null!;

        /// <summary>
        /// The string used to prefix managed instances name (default: `pool`).
        /// </summary>
        [Output("instancePrefix")]
        public Output<string?> InstancePrefix { get; private set; } = null!;

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The list of managed instances. Structure is documented below.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.InstancePoolInstance>> Instances { get; private set; } = null!;

        /// <summary>
        /// Enable IPv6 on managed instances (boolean; default: `false`).
        /// </summary>
        [Output("ipv6")]
        public Output<bool?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// The exoscale*ssh*key (name) to authorize in the managed instances.
        /// </summary>
        [Output("keyPair")]
        public Output<string?> KeyPair { get; private set; } = null!;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The instance pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of exoscale*private*network (IDs).
        /// </summary>
        [Output("networkIds")]
        public Output<ImmutableArray<string>> NetworkIds { get; private set; } = null!;

        /// <summary>
        /// A list of exoscale*security*group (IDs).
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The managed instances type. Please use the `instance_type` argument instead.
        /// </summary>
        [Output("serviceOffering")]
        public Output<string> ServiceOffering { get; private set; } = null!;

        /// <summary>
        /// The number of managed instances.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The exoscale*compute*template (ID) to use when creating the managed instances.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
        /// </summary>
        [Output("virtualMachines")]
        public Output<ImmutableArray<string>> VirtualMachines { get; private set; } = null!;

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstancePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstancePool(string name, InstancePoolArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/instancePool:InstancePool", name, args ?? new InstancePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstancePool(string name, Input<string> id, InstancePoolState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/instancePool:InstancePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstancePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstancePool Get(string name, Input<string> id, InstancePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new InstancePool(name, id, state, options);
        }
    }

    public sealed class InstancePoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("affinityGroupIds")]
        private InputList<string>? _affinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
        /// </summary>
        public InputList<string> AffinityGroupIds
        {
            get => _affinityGroupIds ?? (_affinityGroupIds = new InputList<string>());
            set => _affinityGroupIds = value;
        }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("elasticIpIds")]
        private InputList<string>? _elasticIpIds;

        /// <summary>
        /// A list of exoscale*elastic*ip (IDs).
        /// </summary>
        public InputList<string> ElasticIpIds
        {
            get => _elasticIpIds ?? (_elasticIpIds = new InputList<string>());
            set => _elasticIpIds = value;
        }

        /// <summary>
        /// The string used to prefix managed instances name (default: `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("instances")]
        private InputList<Inputs.InstancePoolInstanceArgs>? _instances;

        /// <summary>
        /// The list of managed instances. Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstancePoolInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.InstancePoolInstanceArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Enable IPv6 on managed instances (boolean; default: `false`).
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// The exoscale*ssh*key (name) to authorize in the managed instances.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The instance pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkIds")]
        private InputList<string>? _networkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs).
        /// </summary>
        public InputList<string> NetworkIds
        {
            get => _networkIds ?? (_networkIds = new InputList<string>());
            set => _networkIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The managed instances type. Please use the `instance_type` argument instead.
        /// </summary>
        [Input("serviceOffering")]
        public Input<string>? ServiceOffering { get; set; }

        /// <summary>
        /// The number of managed instances.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The exoscale*compute*template (ID) to use when creating the managed instances.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        /// <summary>
        /// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("virtualMachines")]
        private InputList<string>? _virtualMachines;

        /// <summary>
        /// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
        /// </summary>
        [Obsolete(@"Use the instances exported attribute instead.")]
        public InputList<string> VirtualMachines
        {
            get => _virtualMachines ?? (_virtualMachines = new InputList<string>());
            set => _virtualMachines = value;
        }

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstancePoolArgs()
        {
        }
        public static new InstancePoolArgs Empty => new InstancePoolArgs();
    }

    public sealed class InstancePoolState : global::Pulumi.ResourceArgs
    {
        [Input("affinityGroupIds")]
        private InputList<string>? _affinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs; may only be set at creation time).
        /// </summary>
        public InputList<string> AffinityGroupIds
        {
            get => _affinityGroupIds ?? (_affinityGroupIds = new InputList<string>());
            set => _affinityGroupIds = value;
        }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("elasticIpIds")]
        private InputList<string>? _elasticIpIds;

        /// <summary>
        /// A list of exoscale*elastic*ip (IDs).
        /// </summary>
        public InputList<string> ElasticIpIds
        {
            get => _elasticIpIds ?? (_elasticIpIds = new InputList<string>());
            set => _elasticIpIds = value;
        }

        /// <summary>
        /// The string used to prefix managed instances name (default: `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("instances")]
        private InputList<Inputs.InstancePoolInstanceGetArgs>? _instances;

        /// <summary>
        /// The list of managed instances. Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstancePoolInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.InstancePoolInstanceGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Enable IPv6 on managed instances (boolean; default: `false`).
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// The exoscale*ssh*key (name) to authorize in the managed instances.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The instance pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkIds")]
        private InputList<string>? _networkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs).
        /// </summary>
        public InputList<string> NetworkIds
        {
            get => _networkIds ?? (_networkIds = new InputList<string>());
            set => _networkIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The managed instances type. Please use the `instance_type` argument instead.
        /// </summary>
        [Input("serviceOffering")]
        public Input<string>? ServiceOffering { get; set; }

        /// <summary>
        /// The number of managed instances.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The exoscale*compute*template (ID) to use when creating the managed instances.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// [cloud-init](http://cloudinit.readthedocs.io/) configuration to apply to the managed instances.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("virtualMachines")]
        private InputList<string>? _virtualMachines;

        /// <summary>
        /// The list of managed instances (IDs). Please use the `instances.*.id` attribute instead.
        /// </summary>
        [Obsolete(@"Use the instances exported attribute instead.")]
        public InputList<string> VirtualMachines
        {
            get => _virtualMachines ?? (_virtualMachines = new InputList<string>());
            set => _virtualMachines = value;
        }

        /// <summary>
        /// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstancePoolState()
        {
        }
        public static new InstancePoolState Empty => new InstancePoolState();
    }
}
