package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	fmtPackage     = protogen.GoImportPath("fmt")
	contextPackage = protogen.GoImportPath("context")
	ioPackage      = protogen.GoImportPath("io")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")
	statusPackage  = protogen.GoImportPath("google.golang.org/grpc/status")
	codesPackage   = protogen.GoImportPath("google.golang.org/grpc/codes")
	stubsPackage   = protogen.GoImportPath("github.com/pydio/cells/v4/common/server/stubs")
)

// generateFile generates a _grpc.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_grpc.stub.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-client-stub v", version)
	g.P("// - protoc             ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("// Requires gRPC-Go v1.32.0 or later.")
	g.P("const _ = ", grpcPackage.Ident("SupportPackageIsVersion7")) // When changing, update version number above.
	g.P()
	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	serviceName := service.GoName
	server := serviceName + "Server"
	stubServer := serviceName + "Stub"

	g.Annotate(stubServer, service.Location)
	g.P("type " + stubServer + " struct {")
	g.P(server)
	g.P("}")

	g.P("func (s *", stubServer, ") Invoke(ctx ", contextPackage.Ident("Context"), ", method string, args interface{}, reply interface{}, opts ...", grpcPackage.Ident("CallOption"), ") error {")
	g.P(fmtPackage.Ident("Println"), `("Serving", `, "method, args, reply, opts)")
	g.P("var e error")
	g.P("switch method {")
	for _, method := range service.Methods {
		if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
			continue
		}
		g.P("case ", "\"/", service.Desc.FullName(), "/", method.GoName, "\":")
		g.P("resp, er := s.", server, ".", method.GoName, "(ctx, args.(*", method.Desc.Input().Name(), "))")
		g.P("if er == nil {")
		g.P("e = ", stubsPackage.Ident("AssignToInterface"), "(resp, reply)")
		g.P("} else {")
		g.P("e = er")
		g.P("}")

	}
	g.P("default:")
	g.P("e = fmt.Errorf(method + \" not implemented\")")
	g.P("}")
	g.P("return e")
	g.P("}")

	g.P("func (s *", stubServer, ") NewStream(ctx ", contextPackage.Ident("Context"), ", desc *", grpcPackage.Ident("StreamDesc"), ", method string, opts ...", grpcPackage.Ident("CallOption"), ") (", grpcPackage.Ident("ClientStream"), ", error) {")
	g.P(fmtPackage.Ident("Println"), "(\"Serving\", method)")
	g.P("switch method {")
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() {
			g.P("case ", "\"/", service.Desc.FullName(), "/", method.GoName, "\":")
			g.P("st := &", stubServer, "_", method.GoName, "Streamer{}")
			g.P("st.Init(ctx)")
			g.P("go s.", server, ".", method.GoName, "(st)")
			g.P("return st, nil")
		} else if method.Desc.IsStreamingServer() {
			g.P("case ", "\"/", service.Desc.FullName(), "/", method.GoName, "\":")
			g.P("st := &", stubServer, "_", method.GoName, "Streamer{}")
			g.P("st.Init(ctx, func(i interface{}) error {")
			g.P("defer func() {")
			g.P("close(st.RespChan)")
			g.P("}()")
			g.P("return s.", server, ".", method.GoName, "(i.(*", method.Desc.Input().Name(), "), st)")
			//g.P("return nil")
			g.P("})")
			g.P("return st, nil")
		}
	}
	g.P("}")
	g.P("return nil, fmt.Errorf(method + \"  not implemented\")")
	g.P("}")

	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() {
			g.P("type ", stubServer, "_", method.GoName, "Streamer struct {")
			g.P(stubsPackage.Ident("BidirServerStreamerCore"))
			g.P("}")

			g.P("func (s *", stubServer, "_", method.GoName, "Streamer) Recv() (*", method.Input.Desc.Name(), ", error) {")
			g.P("if req, o := <-s.ReqChan; o {")
			g.P("return req.(*", method.Input.Desc.Name(), "), nil")
			g.P("}else{")
			g.P("return nil, ", ioPackage.Ident("EOF"))
			g.P("}")
			g.P("}")

			g.P("func (s *", stubServer, "_", method.GoName, "Streamer) Send(response *", method.Output.Desc.Name(), ") error {")
			g.P("s.RespChan <- response")
			g.P("return nil")
			g.P("}")

			if !method.Desc.IsStreamingServer() {
				// Add SendAndClose method
				g.P("func (s *", stubServer, "_", method.GoName, "Streamer) SendAndClose(*", method.Output.Desc.Name(), ") error{")
				g.P("return nil")
				g.P("}")
			}

		} else if method.Desc.IsStreamingServer() {
			g.P("type ", stubServer, "_", method.GoName, "Streamer struct {")
			g.P(stubsPackage.Ident("ClientServerStreamerCore"))
			g.P("}")

			g.P("func (s *", stubServer, "_", method.GoName, "Streamer) Send(response *", method.Output.Desc.Name(), ") error {")
			g.P("s.RespChan <- response")
			g.P("return nil")
			g.P("}")
		}
	}

	/*
		for _, method := range service.Methods {
			g.Annotate(multiServer+"."+method.GoName, method.Location)

			g.P(method.Comments.Leading)
			g.P("func (m ", multiServer, ") ", serverSignature(g, method), " {")
			if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
				g.P("for _, mm := range m {")
				g.P("if mm.Name() == ", serviceContextPackage.Ident("GetServiceName"), "(ctx) {")
				g.P("return mm.", method.GoName, "(ctx, r)")
				g.P("}")
				g.P("}")
				g.P("return nil, ", statusPackage.Ident("Errorf"), "(", codesPackage.Ident("Unimplemented"), ", \"method ", method.GoName, " not implemented\")")

			} else if !method.Desc.IsStreamingClient() {
				g.P("for _, mm := range m {")
				g.P("if mm.Name() == ", serviceContextPackage.Ident("GetServiceName"), "(s.Context()) {")
				g.P("return mm.", method.GoName, "(r, s)")
				g.P("}")
				g.P("}")
				g.P("return ", statusPackage.Ident("Errorf"), "(", codesPackage.Ident("Unimplemented"), ", \"method ", method.GoName, " not implemented\")")
			} else {
				g.P("for _, mm := range m {")
				g.P("if mm.Name() == ", serviceContextPackage.Ident("GetServiceName"), "(s.Context()) {")
				g.P("return mm.", method.GoName, "(s)")
				g.P("}")
				g.P("}")
				g.P("return ", statusPackage.Ident("Errorf"), "(", codesPackage.Ident("Unimplemented"), ", \"method ", method.GoName, " not implemented\")")
			}
			g.P("}")
		}

		g.P("func (m ", multiServer, ") mustEmbedUnimplemented", server, "() {}")

		g.P("func Register", multiServer, "(s grpc.ServiceRegistrar, srv ", namedServer, ") {")
		g.P("addr := ", fmtPackage.Ident("Sprintf"), "(\"%p\", s)")
		g.P("m, ok := multi", server, "s[addr]")
		g.P("if !ok {")
		g.P("m = ", multiServer, "{}")
		g.P("multi", server, "s[addr] = m")
		g.P("Register", server, "(s, m)")
		g.P("}")
		g.P("m = append(m, srv)")
		g.P("}")

	*/
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	var reqArgs []string
	ret := "error"
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, "ctx "+g.QualifiedGoIdent(contextPackage.Ident("Context")))
		ret = "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "r *"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, "s "+method.Parent.GoName+"_"+method.GoName+"Server")
	}
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}
