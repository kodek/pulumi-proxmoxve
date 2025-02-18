// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type File struct {
	pulumi.CustomResourceState

	// The content type
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The datastore id
	DatastoreId pulumi.StringOutput `pulumi:"datastoreId"`
	// The file modification date
	FileModificationDate pulumi.StringOutput `pulumi:"fileModificationDate"`
	// The file name
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// The file size in bytes
	FileSize pulumi.IntOutput `pulumi:"fileSize"`
	// The file tag
	FileTag pulumi.StringOutput `pulumi:"fileTag"`
	// The node name
	NodeName pulumi.StringOutput `pulumi:"nodeName"`
	// The source file
	SourceFile FileSourceFilePtrOutput `pulumi:"sourceFile"`
	// The raw source
	SourceRaw FileSourceRawPtrOutput `pulumi:"sourceRaw"`
}

// NewFile registers a new resource with the given unique name, arguments, and options.
func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreId == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreId'")
	}
	if args.NodeName == nil {
		return nil, errors.New("invalid value for required argument 'NodeName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource File
	err := ctx.RegisterResource("proxmoxve:Storage/file:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFile gets an existing File resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("proxmoxve:Storage/file:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering File resources.
type fileState struct {
	// The content type
	ContentType *string `pulumi:"contentType"`
	// The datastore id
	DatastoreId *string `pulumi:"datastoreId"`
	// The file modification date
	FileModificationDate *string `pulumi:"fileModificationDate"`
	// The file name
	FileName *string `pulumi:"fileName"`
	// The file size in bytes
	FileSize *int `pulumi:"fileSize"`
	// The file tag
	FileTag *string `pulumi:"fileTag"`
	// The node name
	NodeName *string `pulumi:"nodeName"`
	// The source file
	SourceFile *FileSourceFile `pulumi:"sourceFile"`
	// The raw source
	SourceRaw *FileSourceRaw `pulumi:"sourceRaw"`
}

type FileState struct {
	// The content type
	ContentType pulumi.StringPtrInput
	// The datastore id
	DatastoreId pulumi.StringPtrInput
	// The file modification date
	FileModificationDate pulumi.StringPtrInput
	// The file name
	FileName pulumi.StringPtrInput
	// The file size in bytes
	FileSize pulumi.IntPtrInput
	// The file tag
	FileTag pulumi.StringPtrInput
	// The node name
	NodeName pulumi.StringPtrInput
	// The source file
	SourceFile FileSourceFilePtrInput
	// The raw source
	SourceRaw FileSourceRawPtrInput
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	// The content type
	ContentType *string `pulumi:"contentType"`
	// The datastore id
	DatastoreId string `pulumi:"datastoreId"`
	// The node name
	NodeName string `pulumi:"nodeName"`
	// The source file
	SourceFile *FileSourceFile `pulumi:"sourceFile"`
	// The raw source
	SourceRaw *FileSourceRaw `pulumi:"sourceRaw"`
}

// The set of arguments for constructing a File resource.
type FileArgs struct {
	// The content type
	ContentType pulumi.StringPtrInput
	// The datastore id
	DatastoreId pulumi.StringInput
	// The node name
	NodeName pulumi.StringInput
	// The source file
	SourceFile FileSourceFilePtrInput
	// The raw source
	SourceRaw FileSourceRawPtrInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

// FileArrayInput is an input type that accepts FileArray and FileArrayOutput values.
// You can construct a concrete instance of `FileArrayInput` via:
//
//	FileArray{ FileArgs{...} }
type FileArrayInput interface {
	pulumi.Input

	ToFileArrayOutput() FileArrayOutput
	ToFileArrayOutputWithContext(context.Context) FileArrayOutput
}

type FileArray []FileInput

func (FileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (i FileArray) ToFileArrayOutput() FileArrayOutput {
	return i.ToFileArrayOutputWithContext(context.Background())
}

func (i FileArray) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileArrayOutput)
}

// FileMapInput is an input type that accepts FileMap and FileMapOutput values.
// You can construct a concrete instance of `FileMapInput` via:
//
//	FileMap{ "key": FileArgs{...} }
type FileMapInput interface {
	pulumi.Input

	ToFileMapOutput() FileMapOutput
	ToFileMapOutputWithContext(context.Context) FileMapOutput
}

type FileMap map[string]FileInput

func (FileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (i FileMap) ToFileMapOutput() FileMapOutput {
	return i.ToFileMapOutputWithContext(context.Background())
}

func (i FileMap) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileMapOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

// The content type
func (o FileOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The datastore id
func (o FileOutput) DatastoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.DatastoreId }).(pulumi.StringOutput)
}

// The file modification date
func (o FileOutput) FileModificationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.FileModificationDate }).(pulumi.StringOutput)
}

// The file name
func (o FileOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// The file size in bytes
func (o FileOutput) FileSize() pulumi.IntOutput {
	return o.ApplyT(func(v *File) pulumi.IntOutput { return v.FileSize }).(pulumi.IntOutput)
}

// The file tag
func (o FileOutput) FileTag() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.FileTag }).(pulumi.StringOutput)
}

// The node name
func (o FileOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.NodeName }).(pulumi.StringOutput)
}

// The source file
func (o FileOutput) SourceFile() FileSourceFilePtrOutput {
	return o.ApplyT(func(v *File) FileSourceFilePtrOutput { return v.SourceFile }).(FileSourceFilePtrOutput)
}

// The raw source
func (o FileOutput) SourceRaw() FileSourceRawPtrOutput {
	return o.ApplyT(func(v *File) FileSourceRawPtrOutput { return v.SourceRaw }).(FileSourceRawPtrOutput)
}

type FileArrayOutput struct{ *pulumi.OutputState }

func (FileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (o FileArrayOutput) ToFileArrayOutput() FileArrayOutput {
	return o
}

func (o FileArrayOutput) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return o
}

func (o FileArrayOutput) Index(i pulumi.IntInput) FileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *File {
		return vs[0].([]*File)[vs[1].(int)]
	}).(FileOutput)
}

type FileMapOutput struct{ *pulumi.OutputState }

func (FileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (o FileMapOutput) ToFileMapOutput() FileMapOutput {
	return o
}

func (o FileMapOutput) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return o
}

func (o FileMapOutput) MapIndex(k pulumi.StringInput) FileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *File {
		return vs[0].(map[string]*File)[vs[1].(string)]
	}).(FileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileInput)(nil)).Elem(), &File{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileArrayInput)(nil)).Elem(), FileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileMapInput)(nil)).Elem(), FileMap{})
	pulumi.RegisterOutputType(FileOutput{})
	pulumi.RegisterOutputType(FileArrayOutput{})
	pulumi.RegisterOutputType(FileMapOutput{})
}
