// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/translation.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Dataset metadata that is specific to translation.
type TranslationDatasetMetadata struct {
	// Required. The BCP-47 language code of the source language.
	SourceLanguageCode string `protobuf:"bytes,1,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Required. The BCP-47 language code of the target language.
	TargetLanguageCode   string   `protobuf:"bytes,2,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationDatasetMetadata) Reset()         { *m = TranslationDatasetMetadata{} }
func (m *TranslationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TranslationDatasetMetadata) ProtoMessage()    {}
func (*TranslationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3df1fd6bde1101e, []int{0}
}

func (m *TranslationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationDatasetMetadata.Unmarshal(m, b)
}
func (m *TranslationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TranslationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationDatasetMetadata.Merge(m, src)
}
func (m *TranslationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TranslationDatasetMetadata.Size(m)
}
func (m *TranslationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationDatasetMetadata proto.InternalMessageInfo

func (m *TranslationDatasetMetadata) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslationDatasetMetadata) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

// Evaluation metrics for the dataset.
type TranslationEvaluationMetrics struct {
	// Output only. BLEU score.
	BleuScore float64 `protobuf:"fixed64,1,opt,name=bleu_score,json=bleuScore,proto3" json:"bleu_score,omitempty"`
	// Output only. BLEU score for base model.
	BaseBleuScore        float64  `protobuf:"fixed64,2,opt,name=base_bleu_score,json=baseBleuScore,proto3" json:"base_bleu_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationEvaluationMetrics) Reset()         { *m = TranslationEvaluationMetrics{} }
func (m *TranslationEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*TranslationEvaluationMetrics) ProtoMessage()    {}
func (*TranslationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3df1fd6bde1101e, []int{1}
}

func (m *TranslationEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationEvaluationMetrics.Unmarshal(m, b)
}
func (m *TranslationEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *TranslationEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationEvaluationMetrics.Merge(m, src)
}
func (m *TranslationEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_TranslationEvaluationMetrics.Size(m)
}
func (m *TranslationEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationEvaluationMetrics proto.InternalMessageInfo

func (m *TranslationEvaluationMetrics) GetBleuScore() float64 {
	if m != nil {
		return m.BleuScore
	}
	return 0
}

func (m *TranslationEvaluationMetrics) GetBaseBleuScore() float64 {
	if m != nil {
		return m.BaseBleuScore
	}
	return 0
}

// Model metadata that is specific to translation.
type TranslationModelMetadata struct {
	// The resource name of the model to use as a baseline to train the custom
	// model. If unset, we use the default base model provided by Google
	// Translate. Format:
	// `projects/{project_id}/locations/{location_id}/models/{model_id}`
	BaseModel string `protobuf:"bytes,1,opt,name=base_model,json=baseModel,proto3" json:"base_model,omitempty"`
	// Output only. Inferred from the dataset.
	// The source languge (The BCP-47 language code) that is used for training.
	SourceLanguageCode string `protobuf:"bytes,2,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Output only. The target languge (The BCP-47 language code) that is used for
	// training.
	TargetLanguageCode   string   `protobuf:"bytes,3,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationModelMetadata) Reset()         { *m = TranslationModelMetadata{} }
func (m *TranslationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TranslationModelMetadata) ProtoMessage()    {}
func (*TranslationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3df1fd6bde1101e, []int{2}
}

func (m *TranslationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationModelMetadata.Unmarshal(m, b)
}
func (m *TranslationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationModelMetadata.Marshal(b, m, deterministic)
}
func (m *TranslationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationModelMetadata.Merge(m, src)
}
func (m *TranslationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TranslationModelMetadata.Size(m)
}
func (m *TranslationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationModelMetadata proto.InternalMessageInfo

func (m *TranslationModelMetadata) GetBaseModel() string {
	if m != nil {
		return m.BaseModel
	}
	return ""
}

func (m *TranslationModelMetadata) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslationModelMetadata) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

// Annotation details specific to translation.
type TranslationAnnotation struct {
	// Output only . The translated content.
	TranslatedContent    *TextSnippet `protobuf:"bytes,1,opt,name=translated_content,json=translatedContent,proto3" json:"translated_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TranslationAnnotation) Reset()         { *m = TranslationAnnotation{} }
func (m *TranslationAnnotation) String() string { return proto.CompactTextString(m) }
func (*TranslationAnnotation) ProtoMessage()    {}
func (*TranslationAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3df1fd6bde1101e, []int{3}
}

func (m *TranslationAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationAnnotation.Unmarshal(m, b)
}
func (m *TranslationAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationAnnotation.Marshal(b, m, deterministic)
}
func (m *TranslationAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationAnnotation.Merge(m, src)
}
func (m *TranslationAnnotation) XXX_Size() int {
	return xxx_messageInfo_TranslationAnnotation.Size(m)
}
func (m *TranslationAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationAnnotation proto.InternalMessageInfo

func (m *TranslationAnnotation) GetTranslatedContent() *TextSnippet {
	if m != nil {
		return m.TranslatedContent
	}
	return nil
}

func init() {
	proto.RegisterType((*TranslationDatasetMetadata)(nil), "google.cloud.automl.v1.TranslationDatasetMetadata")
	proto.RegisterType((*TranslationEvaluationMetrics)(nil), "google.cloud.automl.v1.TranslationEvaluationMetrics")
	proto.RegisterType((*TranslationModelMetadata)(nil), "google.cloud.automl.v1.TranslationModelMetadata")
	proto.RegisterType((*TranslationAnnotation)(nil), "google.cloud.automl.v1.TranslationAnnotation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/translation.proto", fileDescriptor_c3df1fd6bde1101e)
}

var fileDescriptor_c3df1fd6bde1101e = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x8a, 0xd3, 0x50,
	0x10, 0x86, 0x49, 0x16, 0x84, 0x1e, 0x11, 0x35, 0xe8, 0xb2, 0x96, 0x15, 0xa5, 0x82, 0xee, 0x55,
	0xb2, 0x51, 0xbc, 0x89, 0xde, 0xb4, 0x55, 0xbc, 0xd9, 0xc2, 0x92, 0x5d, 0x7a, 0x21, 0x85, 0x30,
	0x4d, 0xc6, 0x18, 0x3c, 0x3d, 0x13, 0x92, 0x49, 0xf1, 0x19, 0x7c, 0x07, 0x5f, 0xc0, 0x67, 0xf0,
	0x09, 0x7c, 0x0a, 0xaf, 0x7d, 0x0a, 0x39, 0xe7, 0xa4, 0xbb, 0x87, 0xd2, 0xdc, 0x85, 0xf9, 0xff,
	0xef, 0xcc, 0x3f, 0x93, 0x11, 0x67, 0x25, 0x51, 0x29, 0x31, 0xca, 0x25, 0x75, 0x45, 0x04, 0x1d,
	0xd3, 0x46, 0x46, 0xdb, 0x38, 0xe2, 0x06, 0x54, 0x2b, 0x81, 0x2b, 0x52, 0x61, 0xdd, 0x10, 0x53,
	0x70, 0x6c, 0x9d, 0xa1, 0x71, 0x86, 0xd6, 0x19, 0x6e, 0xe3, 0xf1, 0x69, 0xff, 0x02, 0xd4, 0x55,
	0x04, 0x4a, 0x11, 0x1b, 0xa8, 0xb5, 0xd4, 0xf8, 0x99, 0xa3, 0x7e, 0xa9, 0x50, 0x16, 0xd9, 0x1a,
	0xbf, 0xc2, 0xb6, 0xa2, 0xa6, 0x37, 0xbc, 0x1a, 0x08, 0x50, 0x00, 0x43, 0x56, 0x31, 0x6e, 0xfa,
	0x97, 0x26, 0x3f, 0x3c, 0x31, 0xbe, 0xbe, 0x4d, 0xf5, 0x01, 0x18, 0x5a, 0xe4, 0x05, 0x32, 0x68,
	0x6b, 0xf0, 0x56, 0x3c, 0x6a, 0xa9, 0x6b, 0x72, 0xcc, 0x24, 0xa8, 0xb2, 0x83, 0x12, 0xb3, 0x9c,
	0x0a, 0x3c, 0xf1, 0x9e, 0x7b, 0x67, 0xa3, 0xd9, 0xd1, 0xdf, 0xa9, 0x9f, 0x06, 0xd6, 0x70, 0xd1,
	0xeb, 0x73, 0x2a, 0x50, 0x63, 0x0c, 0x4d, 0x89, 0xbc, 0x87, 0xf9, 0x0e, 0x66, 0x0d, 0x2e, 0x36,
	0x41, 0x71, 0xea, 0x64, 0xf9, 0xb8, 0x05, 0xd9, 0x99, 0xaf, 0x05, 0x72, 0x53, 0xe5, 0x6d, 0xf0,
	0x54, 0x88, 0xb5, 0xc4, 0x2e, 0x6b, 0x73, 0x6a, 0x6c, 0x06, 0x2f, 0x1d, 0xe9, 0xca, 0x95, 0x2e,
	0x04, 0x2f, 0xc5, 0xfd, 0x35, 0xb4, 0x98, 0x39, 0x1e, 0xdf, 0x78, 0xee, 0xe9, 0xf2, 0x6c, 0xe7,
	0x9b, 0xfc, 0xf4, 0xc4, 0x89, 0xd3, 0x67, 0x41, 0x05, 0xca, 0x9b, 0x89, 0x75, 0x0f, 0xfd, 0xc8,
	0x46, 0x57, 0xed, 0x9c, 0xe9, 0x48, 0x57, 0x8c, 0x2d, 0x38, 0x1f, 0x58, 0x88, 0x99, 0xec, 0xe0,
	0x2e, 0xce, 0x07, 0x76, 0x71, 0x64, 0x89, 0x03, 0x6b, 0xf8, 0x26, 0x1e, 0x3b, 0xf1, 0xa6, 0x37,
	0x7f, 0x3f, 0x48, 0x45, 0xb0, 0xbb, 0x20, 0x2c, 0xb2, 0x9c, 0x14, 0xa3, 0x62, 0x93, 0xf1, 0xee,
	0xeb, 0x17, 0xe1, 0xe1, 0x4b, 0x0a, 0xaf, 0xf1, 0x3b, 0x5f, 0xa9, 0xaa, 0xae, 0x91, 0xd3, 0x87,
	0xb7, 0xf8, 0xdc, 0xd2, 0xb3, 0xdf, 0x9e, 0x18, 0xe7, 0xb4, 0x19, 0xa0, 0x67, 0x0f, 0x9c, 0x24,
	0x97, 0xfa, 0x62, 0x2e, 0xbd, 0xcf, 0xef, 0x7b, 0x6f, 0x49, 0x7a, 0xa2, 0x90, 0x9a, 0x32, 0x2a,
	0x51, 0x99, 0x7b, 0x8a, 0xac, 0x04, 0x75, 0xd5, 0xee, 0xdf, 0xde, 0x3b, 0xfb, 0xf5, 0xcb, 0x3f,
	0xfe, 0x64, 0xf1, 0xb9, 0x69, 0x35, 0xed, 0x98, 0x16, 0x17, 0xe1, 0x32, 0xfe, 0xb3, 0x13, 0x56,
	0x46, 0x58, 0x19, 0x41, 0xae, 0x96, 0xf1, 0x3f, 0xff, 0x89, 0x15, 0x92, 0xc4, 0x28, 0x49, 0x62,
	0x99, 0x24, 0x59, 0xc6, 0xeb, 0x3b, 0xa6, 0xed, 0x9b, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66,
	0x49, 0xdb, 0x1f, 0x72, 0x03, 0x00, 0x00,
}
