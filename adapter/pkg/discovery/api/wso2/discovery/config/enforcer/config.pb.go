// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/config/enforcer/config.proto

package enforcer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enforcer config model
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security     *Security     `protobuf:"bytes,1,opt,name=security,proto3" json:"security,omitempty"`
	Keystore     *CertStore    `protobuf:"bytes,2,opt,name=keystore,proto3" json:"keystore,omitempty"`
	Truststore   *CertStore    `protobuf:"bytes,3,opt,name=truststore,proto3" json:"truststore,omitempty"`
	AuthService  *Service      `protobuf:"bytes,4,opt,name=authService,proto3" json:"authService,omitempty"`
	JwtGenerator *JWTGenerator `protobuf:"bytes,5,opt,name=jwtGenerator,proto3" json:"jwtGenerator,omitempty"`
	Cache        *Cache        `protobuf:"bytes,7,opt,name=cache,proto3" json:"cache,omitempty"`
	JwtIssuer    *JWTIssuer    `protobuf:"bytes,8,opt,name=jwtIssuer,proto3" json:"jwtIssuer,omitempty"`
	Analytics    *Analytics    `protobuf:"bytes,9,opt,name=analytics,proto3" json:"analytics,omitempty"`
	Management   *Management   `protobuf:"bytes,10,opt,name=management,proto3" json:"management,omitempty"`
	RestServer   *RestServer   `protobuf:"bytes,11,opt,name=restServer,proto3" json:"restServer,omitempty"`
	Tracing      *Tracing      `protobuf:"bytes,12,opt,name=tracing,proto3" json:"tracing,omitempty"`
	Metrics      *Metrics      `protobuf:"bytes,13,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Filters      []*Filter     `protobuf:"bytes,14,rep,name=filters,proto3" json:"filters,omitempty"`
	Soap         *Soap         `protobuf:"bytes,15,opt,name=soap,proto3" json:"soap,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetSecurity() *Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Config) GetKeystore() *CertStore {
	if x != nil {
		return x.Keystore
	}
	return nil
}

func (x *Config) GetTruststore() *CertStore {
	if x != nil {
		return x.Truststore
	}
	return nil
}

func (x *Config) GetAuthService() *Service {
	if x != nil {
		return x.AuthService
	}
	return nil
}

func (x *Config) GetJwtGenerator() *JWTGenerator {
	if x != nil {
		return x.JwtGenerator
	}
	return nil
}

func (x *Config) GetCache() *Cache {
	if x != nil {
		return x.Cache
	}
	return nil
}

func (x *Config) GetJwtIssuer() *JWTIssuer {
	if x != nil {
		return x.JwtIssuer
	}
	return nil
}

func (x *Config) GetAnalytics() *Analytics {
	if x != nil {
		return x.Analytics
	}
	return nil
}

func (x *Config) GetManagement() *Management {
	if x != nil {
		return x.Management
	}
	return nil
}

func (x *Config) GetRestServer() *RestServer {
	if x != nil {
		return x.RestServer
	}
	return nil
}

func (x *Config) GetTracing() *Tracing {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Config) GetMetrics() *Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Config) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Config) GetSoap() *Soap {
	if x != nil {
		return x.Soap
	}
	return nil
}

var File_wso2_discovery_config_enforcer_config_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x1a, 0x29, 0x77,
	0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x77, 0x73, 0x6f, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x73, 0x6f,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x77, 0x73, 0x6f, 0x32, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2f, 0x73, 0x6f, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x07,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x45,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x49, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x6a,
	0x77, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2e, 0x4a, 0x57, 0x54, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x0c, 0x6a, 0x77, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a,
	0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x6a, 0x77,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x4a,
	0x57, 0x54, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x09, 0x6a, 0x77, 0x74, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x4a, 0x0a, 0x0a,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x73,
	0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x04,
	0x73, 0x6f, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x61, 0x70,
	0x52, 0x04, 0x73, 0x6f, 0x61, 0x70, 0x42, 0x92, 0x01, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x0b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_config_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_config_proto_rawDescData = file_wso2_discovery_config_enforcer_config_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_config_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_config_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_config_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_config_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_config_proto_goTypes = []interface{}{
	(*Config)(nil),       // 0: wso2.discovery.config.enforcer.Config
	(*Security)(nil),     // 1: wso2.discovery.config.enforcer.Security
	(*CertStore)(nil),    // 2: wso2.discovery.config.enforcer.CertStore
	(*Service)(nil),      // 3: wso2.discovery.config.enforcer.Service
	(*JWTGenerator)(nil), // 4: wso2.discovery.config.enforcer.JWTGenerator
	(*Cache)(nil),        // 5: wso2.discovery.config.enforcer.Cache
	(*JWTIssuer)(nil),    // 6: wso2.discovery.config.enforcer.JWTIssuer
	(*Analytics)(nil),    // 7: wso2.discovery.config.enforcer.Analytics
	(*Management)(nil),   // 8: wso2.discovery.config.enforcer.Management
	(*RestServer)(nil),   // 9: wso2.discovery.config.enforcer.RestServer
	(*Tracing)(nil),      // 10: wso2.discovery.config.enforcer.Tracing
	(*Metrics)(nil),      // 11: wso2.discovery.config.enforcer.Metrics
	(*Filter)(nil),       // 12: wso2.discovery.config.enforcer.Filter
	(*Soap)(nil),         // 13: wso2.discovery.config.enforcer.Soap
}
var file_wso2_discovery_config_enforcer_config_proto_depIdxs = []int32{
	1,  // 0: wso2.discovery.config.enforcer.Config.security:type_name -> wso2.discovery.config.enforcer.Security
	2,  // 1: wso2.discovery.config.enforcer.Config.keystore:type_name -> wso2.discovery.config.enforcer.CertStore
	2,  // 2: wso2.discovery.config.enforcer.Config.truststore:type_name -> wso2.discovery.config.enforcer.CertStore
	3,  // 3: wso2.discovery.config.enforcer.Config.authService:type_name -> wso2.discovery.config.enforcer.Service
	4,  // 4: wso2.discovery.config.enforcer.Config.jwtGenerator:type_name -> wso2.discovery.config.enforcer.JWTGenerator
	5,  // 5: wso2.discovery.config.enforcer.Config.cache:type_name -> wso2.discovery.config.enforcer.Cache
	6,  // 6: wso2.discovery.config.enforcer.Config.jwtIssuer:type_name -> wso2.discovery.config.enforcer.JWTIssuer
	7,  // 7: wso2.discovery.config.enforcer.Config.analytics:type_name -> wso2.discovery.config.enforcer.Analytics
	8,  // 8: wso2.discovery.config.enforcer.Config.management:type_name -> wso2.discovery.config.enforcer.Management
	9,  // 9: wso2.discovery.config.enforcer.Config.restServer:type_name -> wso2.discovery.config.enforcer.RestServer
	10, // 10: wso2.discovery.config.enforcer.Config.tracing:type_name -> wso2.discovery.config.enforcer.Tracing
	11, // 11: wso2.discovery.config.enforcer.Config.metrics:type_name -> wso2.discovery.config.enforcer.Metrics
	12, // 12: wso2.discovery.config.enforcer.Config.filters:type_name -> wso2.discovery.config.enforcer.Filter
	13, // 13: wso2.discovery.config.enforcer.Config.soap:type_name -> wso2.discovery.config.enforcer.Soap
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_config_proto_init() }
func file_wso2_discovery_config_enforcer_config_proto_init() {
	if File_wso2_discovery_config_enforcer_config_proto != nil {
		return
	}
	file_wso2_discovery_config_enforcer_cert_proto_init()
	file_wso2_discovery_config_enforcer_service_proto_init()
	file_wso2_discovery_config_enforcer_jwt_generator_proto_init()
	file_wso2_discovery_config_enforcer_jwt_issuer_proto_init()
	file_wso2_discovery_config_enforcer_cache_proto_init()
	file_wso2_discovery_config_enforcer_analytics_proto_init()
	file_wso2_discovery_config_enforcer_security_proto_init()
	file_wso2_discovery_config_enforcer_management_proto_init()
	file_wso2_discovery_config_enforcer_rest_server_proto_init()
	file_wso2_discovery_config_enforcer_filter_proto_init()
	file_wso2_discovery_config_enforcer_tracing_proto_init()
	file_wso2_discovery_config_enforcer_metrics_proto_init()
	file_wso2_discovery_config_enforcer_soap_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_config_enforcer_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_config_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_config_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_config_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_config_proto = out.File
	file_wso2_discovery_config_enforcer_config_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_config_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_config_proto_depIdxs = nil
}
