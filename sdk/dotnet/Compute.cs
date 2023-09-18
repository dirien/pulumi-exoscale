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
    /// !&gt; **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ComputeInstance instead (or refer to the ad-hoc migration guide).
    /// 
    /// Manage Exoscale Compute Instances.
    /// </summary>
    [ExoscaleResourceType("exoscale:index/compute:Compute")]
    public partial class Compute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ❗ A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinity_groups`).
        /// </summary>
        [Output("affinityGroupIds")]
        public Output<ImmutableArray<string>> AffinityGroupIds { get; private set; } = null!;

        /// <summary>
        /// ❗ A list of anti-affinity groups (names; at creation time only; conflicts with `affinity_group_ids`).
        /// </summary>
        [Output("affinityGroups")]
        public Output<ImmutableArray<string>> AffinityGroups { get; private set; } = null!;

        /// <summary>
        /// The instance disk size (GiB; at least `10`).
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen (" - ") characters; it can be changed to any character during a later update. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Request an IPv4 address on the default NIC
        /// </summary>
        [Output("ip4")]
        public Output<bool?> Ip4 { get; private set; } = null!;

        /// <summary>
        /// Request an IPv6 address on the default NIC
        /// </summary>
        [Output("ip6")]
        public Output<bool?> Ip6 { get; private set; } = null!;

        /// <summary>
        /// The instance (main network interface) IPv6 address (if enabled).
        /// </summary>
        [Output("ip6Address")]
        public Output<string> Ip6Address { get; private set; } = null!;

        [Output("ip6Cidr")]
        public Output<string> Ip6Cidr { get; private set; } = null!;

        /// <summary>
        /// The instance (main network interface) IPv4 address.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// ❗ The SSH keypair (name) to authorize in the instance.
        /// </summary>
        [Output("keyPair")]
        public Output<string?> KeyPair { get; private set; } = null!;

        /// <summary>
        /// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
        /// </summary>
        [Output("keyboard")]
        public Output<string?> Keyboard { get; private set; } = null!;

        /// <summary>
        /// The instance hostname. Please use the `hostname` argument instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The instance initial password and/or encrypted password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
        /// </summary>
        [Output("reverseDns")]
        public Output<string?> ReverseDns { get; private set; } = null!;

        /// <summary>
        /// A list of security groups (IDs; conflicts with `security_groups`).
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A list of security groups (names; conflicts with `security_group_ids`).
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
        /// </summary>
        [Output("size")]
        public Output<string?> Size { get; private set; } = null!;

        /// <summary>
        /// The instance state (`Running` or `Stopped`; default: `Running`)
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Map of tags (key/value). To remove all tags, set `tags = {}`.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// ❗ The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `template_id` attribute instead.
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;

        /// <summary>
        /// ❗ The compute instance template (ID). Usage of the `exoscale.getComputeTemplate` data source is recommended.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// cloud-init configuration.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// was the cloud-init configuration base64 encoded
        /// </summary>
        [Output("userDataBase64")]
        public Output<bool> UserDataBase64 { get; private set; } = null!;

        /// <summary>
        /// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `exoscale.getComputeTemplate` data source `username` attribute instead.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// ❗ The Exoscale Zone name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Compute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Compute(string name, ComputeArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/compute:Compute", name, args ?? new ComputeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Compute(string name, Input<string> id, ComputeState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/compute:Compute", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Compute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Compute Get(string name, Input<string> id, ComputeState? state = null, CustomResourceOptions? options = null)
        {
            return new Compute(name, id, state, options);
        }
    }

    public sealed class ComputeArgs : global::Pulumi.ResourceArgs
    {
        [Input("affinityGroupIds")]
        private InputList<string>? _affinityGroupIds;

        /// <summary>
        /// ❗ A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinity_groups`).
        /// </summary>
        public InputList<string> AffinityGroupIds
        {
            get => _affinityGroupIds ?? (_affinityGroupIds = new InputList<string>());
            set => _affinityGroupIds = value;
        }

        [Input("affinityGroups")]
        private InputList<string>? _affinityGroups;

        /// <summary>
        /// ❗ A list of anti-affinity groups (names; at creation time only; conflicts with `affinity_group_ids`).
        /// </summary>
        public InputList<string> AffinityGroups
        {
            get => _affinityGroups ?? (_affinityGroups = new InputList<string>());
            set => _affinityGroups = value;
        }

        /// <summary>
        /// The instance disk size (GiB; at least `10`).
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen (" - ") characters; it can be changed to any character during a later update. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Request an IPv4 address on the default NIC
        /// </summary>
        [Input("ip4")]
        public Input<bool>? Ip4 { get; set; }

        /// <summary>
        /// Request an IPv6 address on the default NIC
        /// </summary>
        [Input("ip6")]
        public Input<bool>? Ip6 { get; set; }

        /// <summary>
        /// ❗ The SSH keypair (name) to authorize in the instance.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        /// <summary>
        /// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
        /// </summary>
        [Input("keyboard")]
        public Input<string>? Keyboard { get; set; }

        /// <summary>
        /// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
        /// </summary>
        [Input("reverseDns")]
        public Input<string>? ReverseDns { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security groups (IDs; conflicts with `security_groups`).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security groups (names; conflicts with `security_group_ids`).
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The instance state (`Running` or `Stopped`; default: `Running`)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags (key/value). To remove all tags, set `tags = {}`.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ❗ The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `template_id` attribute instead.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// ❗ The compute instance template (ID). Usage of the `exoscale.getComputeTemplate` data source is recommended.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// cloud-init configuration.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// ❗ The Exoscale Zone name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ComputeArgs()
        {
        }
        public static new ComputeArgs Empty => new ComputeArgs();
    }

    public sealed class ComputeState : global::Pulumi.ResourceArgs
    {
        [Input("affinityGroupIds")]
        private InputList<string>? _affinityGroupIds;

        /// <summary>
        /// ❗ A list of anti-affinity groups (IDs; at creation time only; conflicts with `affinity_groups`).
        /// </summary>
        public InputList<string> AffinityGroupIds
        {
            get => _affinityGroupIds ?? (_affinityGroupIds = new InputList<string>());
            set => _affinityGroupIds = value;
        }

        [Input("affinityGroups")]
        private InputList<string>? _affinityGroups;

        /// <summary>
        /// ❗ A list of anti-affinity groups (names; at creation time only; conflicts with `affinity_group_ids`).
        /// </summary>
        public InputList<string> AffinityGroups
        {
            get => _affinityGroups ?? (_affinityGroups = new InputList<string>());
            set => _affinityGroups = value;
        }

        /// <summary>
        /// The instance disk size (GiB; at least `10`).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The displayed instance name. Note: if the `hostname` attribute is not set, this attribute is also used to set the OS' *hostname* during creation, so the value must contain only alphanumeric and hyphen (" - ") characters; it can be changed to any character during a later update. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The instance hostname, must contain only alphanumeric and hyphen (`-`) characters. If neither `display_name` or `hostname` attributes are set, a random value will be generated automatically. Note: updating this attribute's value requires to reboot the instance.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Request an IPv4 address on the default NIC
        /// </summary>
        [Input("ip4")]
        public Input<bool>? Ip4 { get; set; }

        /// <summary>
        /// Request an IPv6 address on the default NIC
        /// </summary>
        [Input("ip6")]
        public Input<bool>? Ip6 { get; set; }

        /// <summary>
        /// The instance (main network interface) IPv6 address (if enabled).
        /// </summary>
        [Input("ip6Address")]
        public Input<string>? Ip6Address { get; set; }

        [Input("ip6Cidr")]
        public Input<string>? Ip6Cidr { get; set; }

        /// <summary>
        /// The instance (main network interface) IPv4 address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// ❗ The SSH keypair (name) to authorize in the instance.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        /// <summary>
        /// The keyboard layout configuration (`de`, `de-ch`, `es`, `fi`, `fr`, `fr-be`, `fr-ch`, `is`, `it`, `jp`, `nl-be`, `no`, `pt`, `uk`, `us`; at creation time only).
        /// </summary>
        [Input("keyboard")]
        public Input<string>? Keyboard { get; set; }

        /// <summary>
        /// The instance hostname. Please use the `hostname` argument instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The instance initial password and/or encrypted password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The instance reverse DNS record (must end with a `.`; e.g: `my-instance.example.net.`).
        /// </summary>
        [Input("reverseDns")]
        public Input<string>? ReverseDns { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security groups (IDs; conflicts with `security_groups`).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security groups (names; conflicts with `security_group_ids`).
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The instance size (`Tiny`, `Small`, `Medium`, `Large`, etc.)
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The instance state (`Running` or `Stopped`; default: `Running`)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags (key/value). To remove all tags, set `tags = {}`.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ❗ The compute instance template (name). Only *featured* templates are available, if you want to reference *custom templates* use the `template_id` attribute instead.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// ❗ The compute instance template (ID). Usage of the `exoscale.getComputeTemplate` data source is recommended.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// cloud-init configuration.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// was the cloud-init configuration base64 encoded
        /// </summary>
        [Input("userDataBase64")]
        public Input<bool>? UserDataBase64 { get; set; }

        /// <summary>
        /// The user to use to connect to the instance. If you've referenced a *custom template* in the resource, use the `exoscale.getComputeTemplate` data source `username` attribute instead.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// ❗ The Exoscale Zone name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ComputeState()
        {
        }
        public static new ComputeState Empty => new ComputeState();
    }
}
