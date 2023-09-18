// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// ## Import
//
// An existing database service may be imported by `<name>@<zone>`
//
// ```sh
//
//	$ pulumi import exoscale:index/database:Database \
//
// ```
//
//	exoscale_database.my_database \
//
//	my-database@ch-gva-2
type Database struct {
	pulumi.CustomResourceState

	// CA Certificate required to reach a DBaaS service through a TLS-protected connection.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// The creation date of the database service.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The disk size of the database service.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// *grafana* database service type specific arguments. Structure is documented below.
	Grafana DatabaseGrafanaPtrOutput `pulumi:"grafana"`
	// *kafka* database service type specific arguments. Structure is documented below.
	Kafka DatabaseKafkaPtrOutput `pulumi:"kafka"`
	// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
	MaintenanceDow pulumi.StringOutput `pulumi:"maintenanceDow"`
	// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
	MaintenanceTime pulumi.StringOutput `pulumi:"maintenanceTime"`
	// *mysql* database service type specific arguments. Structure is documented below.
	Mysql DatabaseMysqlPtrOutput `pulumi:"mysql"`
	// ❗ The name of the database service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of CPUs of the database service.
	NodeCpus pulumi.IntOutput `pulumi:"nodeCpus"`
	// The amount of memory of the database service.
	NodeMemory pulumi.IntOutput `pulumi:"nodeMemory"`
	// The number of nodes of the database service.
	Nodes pulumi.IntOutput `pulumi:"nodes"`
	// *opensearch* database service type specific arguments. Structure is documented below.
	Opensearch DatabaseOpensearchPtrOutput `pulumi:"opensearch"`
	// *pg* database service type specific arguments. Structure is documented below.
	Pg DatabasePgPtrOutput `pulumi:"pg"`
	// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
	Plan pulumi.StringOutput `pulumi:"plan"`
	// *redis* database service type specific arguments. Structure is documented below.
	Redis DatabaseRedisPtrOutput `pulumi:"redis"`
	// The current state of the database service.
	State pulumi.StringOutput `pulumi:"state"`
	// The database service protection boolean flag against termination/power-off.
	TerminationProtection pulumi.BoolOutput         `pulumi:"terminationProtection"`
	Timeouts              DatabaseTimeoutsPtrOutput `pulumi:"timeouts"`
	// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
	Type pulumi.StringOutput `pulumi:"type"`
	// The date of the latest database service update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("exoscale:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("exoscale:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// CA Certificate required to reach a DBaaS service through a TLS-protected connection.
	CaCertificate *string `pulumi:"caCertificate"`
	// The creation date of the database service.
	CreatedAt *string `pulumi:"createdAt"`
	// The disk size of the database service.
	DiskSize *int `pulumi:"diskSize"`
	// *grafana* database service type specific arguments. Structure is documented below.
	Grafana *DatabaseGrafana `pulumi:"grafana"`
	// *kafka* database service type specific arguments. Structure is documented below.
	Kafka *DatabaseKafka `pulumi:"kafka"`
	// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
	MaintenanceDow *string `pulumi:"maintenanceDow"`
	// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
	MaintenanceTime *string `pulumi:"maintenanceTime"`
	// *mysql* database service type specific arguments. Structure is documented below.
	Mysql *DatabaseMysql `pulumi:"mysql"`
	// ❗ The name of the database service.
	Name *string `pulumi:"name"`
	// The number of CPUs of the database service.
	NodeCpus *int `pulumi:"nodeCpus"`
	// The amount of memory of the database service.
	NodeMemory *int `pulumi:"nodeMemory"`
	// The number of nodes of the database service.
	Nodes *int `pulumi:"nodes"`
	// *opensearch* database service type specific arguments. Structure is documented below.
	Opensearch *DatabaseOpensearch `pulumi:"opensearch"`
	// *pg* database service type specific arguments. Structure is documented below.
	Pg *DatabasePg `pulumi:"pg"`
	// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
	Plan *string `pulumi:"plan"`
	// *redis* database service type specific arguments. Structure is documented below.
	Redis *DatabaseRedis `pulumi:"redis"`
	// The current state of the database service.
	State *string `pulumi:"state"`
	// The database service protection boolean flag against termination/power-off.
	TerminationProtection *bool             `pulumi:"terminationProtection"`
	Timeouts              *DatabaseTimeouts `pulumi:"timeouts"`
	// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
	Type *string `pulumi:"type"`
	// The date of the latest database service update.
	UpdatedAt *string `pulumi:"updatedAt"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type DatabaseState struct {
	// CA Certificate required to reach a DBaaS service through a TLS-protected connection.
	CaCertificate pulumi.StringPtrInput
	// The creation date of the database service.
	CreatedAt pulumi.StringPtrInput
	// The disk size of the database service.
	DiskSize pulumi.IntPtrInput
	// *grafana* database service type specific arguments. Structure is documented below.
	Grafana DatabaseGrafanaPtrInput
	// *kafka* database service type specific arguments. Structure is documented below.
	Kafka DatabaseKafkaPtrInput
	// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
	MaintenanceDow pulumi.StringPtrInput
	// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
	MaintenanceTime pulumi.StringPtrInput
	// *mysql* database service type specific arguments. Structure is documented below.
	Mysql DatabaseMysqlPtrInput
	// ❗ The name of the database service.
	Name pulumi.StringPtrInput
	// The number of CPUs of the database service.
	NodeCpus pulumi.IntPtrInput
	// The amount of memory of the database service.
	NodeMemory pulumi.IntPtrInput
	// The number of nodes of the database service.
	Nodes pulumi.IntPtrInput
	// *opensearch* database service type specific arguments. Structure is documented below.
	Opensearch DatabaseOpensearchPtrInput
	// *pg* database service type specific arguments. Structure is documented below.
	Pg DatabasePgPtrInput
	// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
	Plan pulumi.StringPtrInput
	// *redis* database service type specific arguments. Structure is documented below.
	Redis DatabaseRedisPtrInput
	// The current state of the database service.
	State pulumi.StringPtrInput
	// The database service protection boolean flag against termination/power-off.
	TerminationProtection pulumi.BoolPtrInput
	Timeouts              DatabaseTimeoutsPtrInput
	// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
	Type pulumi.StringPtrInput
	// The date of the latest database service update.
	UpdatedAt pulumi.StringPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// *grafana* database service type specific arguments. Structure is documented below.
	Grafana *DatabaseGrafana `pulumi:"grafana"`
	// *kafka* database service type specific arguments. Structure is documented below.
	Kafka *DatabaseKafka `pulumi:"kafka"`
	// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
	MaintenanceDow *string `pulumi:"maintenanceDow"`
	// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
	MaintenanceTime *string `pulumi:"maintenanceTime"`
	// *mysql* database service type specific arguments. Structure is documented below.
	Mysql *DatabaseMysql `pulumi:"mysql"`
	// ❗ The name of the database service.
	Name *string `pulumi:"name"`
	// *opensearch* database service type specific arguments. Structure is documented below.
	Opensearch *DatabaseOpensearch `pulumi:"opensearch"`
	// *pg* database service type specific arguments. Structure is documented below.
	Pg *DatabasePg `pulumi:"pg"`
	// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
	Plan string `pulumi:"plan"`
	// *redis* database service type specific arguments. Structure is documented below.
	Redis *DatabaseRedis `pulumi:"redis"`
	// The database service protection boolean flag against termination/power-off.
	TerminationProtection *bool             `pulumi:"terminationProtection"`
	Timeouts              *DatabaseTimeouts `pulumi:"timeouts"`
	// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
	Type string `pulumi:"type"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// *grafana* database service type specific arguments. Structure is documented below.
	Grafana DatabaseGrafanaPtrInput
	// *kafka* database service type specific arguments. Structure is documented below.
	Kafka DatabaseKafkaPtrInput
	// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
	MaintenanceDow pulumi.StringPtrInput
	// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
	MaintenanceTime pulumi.StringPtrInput
	// *mysql* database service type specific arguments. Structure is documented below.
	Mysql DatabaseMysqlPtrInput
	// ❗ The name of the database service.
	Name pulumi.StringPtrInput
	// *opensearch* database service type specific arguments. Structure is documented below.
	Opensearch DatabaseOpensearchPtrInput
	// *pg* database service type specific arguments. Structure is documented below.
	Pg DatabasePgPtrInput
	// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
	Plan pulumi.StringInput
	// *redis* database service type specific arguments. Structure is documented below.
	Redis DatabaseRedisPtrInput
	// The database service protection boolean flag against termination/power-off.
	TerminationProtection pulumi.BoolPtrInput
	Timeouts              DatabaseTimeoutsPtrInput
	// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
	Type pulumi.StringInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

func (i *Database) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: i.ToDatabaseOutputWithContext(ctx).OutputState,
	}
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

func (i DatabaseArray) ToOutput(ctx context.Context) pulumix.Output[[]*Database] {
	return pulumix.Output[[]*Database]{
		OutputState: i.ToDatabaseArrayOutputWithContext(ctx).OutputState,
	}
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

func (i DatabaseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Database] {
	return pulumix.Output[map[string]*Database]{
		OutputState: i.ToDatabaseMapOutputWithContext(ctx).OutputState,
	}
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*Database] {
	return pulumix.Output[*Database]{
		OutputState: o.OutputState,
	}
}

// CA Certificate required to reach a DBaaS service through a TLS-protected connection.
func (o DatabaseOutput) CaCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CaCertificate }).(pulumi.StringOutput)
}

// The creation date of the database service.
func (o DatabaseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The disk size of the database service.
func (o DatabaseOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// *grafana* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Grafana() DatabaseGrafanaPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseGrafanaPtrOutput { return v.Grafana }).(DatabaseGrafanaPtrOutput)
}

// *kafka* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Kafka() DatabaseKafkaPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseKafkaPtrOutput { return v.Kafka }).(DatabaseKafkaPtrOutput)
}

// The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
func (o DatabaseOutput) MaintenanceDow() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.MaintenanceDow }).(pulumi.StringOutput)
}

// The time of day to perform the automated database service maintenance (`HH:MM:SS`)
func (o DatabaseOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.MaintenanceTime }).(pulumi.StringOutput)
}

// *mysql* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Mysql() DatabaseMysqlPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseMysqlPtrOutput { return v.Mysql }).(DatabaseMysqlPtrOutput)
}

// ❗ The name of the database service.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of CPUs of the database service.
func (o DatabaseOutput) NodeCpus() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.NodeCpus }).(pulumi.IntOutput)
}

// The amount of memory of the database service.
func (o DatabaseOutput) NodeMemory() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.NodeMemory }).(pulumi.IntOutput)
}

// The number of nodes of the database service.
func (o DatabaseOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.Nodes }).(pulumi.IntOutput)
}

// *opensearch* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Opensearch() DatabaseOpensearchPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseOpensearchPtrOutput { return v.Opensearch }).(DatabaseOpensearchPtrOutput)
}

// *pg* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Pg() DatabasePgPtrOutput {
	return o.ApplyT(func(v *Database) DatabasePgPtrOutput { return v.Pg }).(DatabasePgPtrOutput)
}

// The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
func (o DatabaseOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// *redis* database service type specific arguments. Structure is documented below.
func (o DatabaseOutput) Redis() DatabaseRedisPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseRedisPtrOutput { return v.Redis }).(DatabaseRedisPtrOutput)
}

// The current state of the database service.
func (o DatabaseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The database service protection boolean flag against termination/power-off.
func (o DatabaseOutput) TerminationProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolOutput { return v.TerminationProtection }).(pulumi.BoolOutput)
}

func (o DatabaseOutput) Timeouts() DatabaseTimeoutsPtrOutput {
	return o.ApplyT(func(v *Database) DatabaseTimeoutsPtrOutput { return v.Timeouts }).(DatabaseTimeoutsPtrOutput)
}

// ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date of the latest database service update.
func (o DatabaseOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o DatabaseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Database] {
	return pulumix.Output[[]*Database]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Database] {
	return pulumix.Output[map[string]*Database]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
