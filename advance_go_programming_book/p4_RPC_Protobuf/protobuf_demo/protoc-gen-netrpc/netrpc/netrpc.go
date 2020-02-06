package netrpc

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type netrpc struct{ *generator.Generator }

func (p *netrpc) Name() string {
	return "netrpc"
}
func (p *netrpc) Init(g *generator.Generator) {
	p.Generator = g
	generator.RegisterPlugin(new(netrpc))
}
func (p *netrpc) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}
func (p *netrpc) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpc) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
}
func (p *netrpc) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}
