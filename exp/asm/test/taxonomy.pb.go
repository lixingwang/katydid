// Code generated by protoc-gen-gogo.
// source: taxonomy.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	taxonomy.proto

It has these top-level messages:
	TaxonomyDatabase
	Entry
	Records
	Reference
	Link
	GeneticCode
	CommentAndReference
	GenomeInformation
	ExternalInformationResources
*/
package main

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type GenomeInformationType int32

const (
	GenomeInformationType_NONE  GenomeInformationType = 1
	GenomeInformationType_BLAST GenomeInformationType = 2
)

var GenomeInformationType_name = map[int32]string{
	1: "NONE",
	2: "BLAST",
}
var GenomeInformationType_value = map[string]int32{
	"NONE":  1,
	"BLAST": 2,
}

func (x GenomeInformationType) Enum() *GenomeInformationType {
	p := new(GenomeInformationType)
	*p = x
	return p
}
func (x GenomeInformationType) String() string {
	return proto.EnumName(GenomeInformationType_name, int32(x))
}
func (x *GenomeInformationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GenomeInformationType_value, data, "GenomeInformationType")
	if err != nil {
		return err
	}
	*x = GenomeInformationType(value)
	return nil
}

type Subject int32

const (
	Subject_Unknown          Subject = 1
	Subject_OrganismSpecific Subject = 2
)

var Subject_name = map[int32]string{
	1: "Unknown",
	2: "OrganismSpecific",
}
var Subject_value = map[string]int32{
	"Unknown":          1,
	"OrganismSpecific": 2,
}

func (x Subject) Enum() *Subject {
	p := new(Subject)
	*p = x
	return p
}
func (x Subject) String() string {
	return proto.EnumName(Subject_name, int32(x))
}
func (x *Subject) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Subject_value, data, "Subject")
	if err != nil {
		return err
	}
	*x = Subject(value)
	return nil
}

type TaxonomyDatabase struct {
	Entries          []*Entry `protobuf:"bytes,1,rep" json:"Entries,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TaxonomyDatabase) Reset()         { *m = TaxonomyDatabase{} }
func (m *TaxonomyDatabase) String() string { return proto.CompactTextString(m) }
func (*TaxonomyDatabase) ProtoMessage()    {}

func (m *TaxonomyDatabase) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	TaxonomyId                   *string                         `protobuf:"bytes,1,opt" json:"TaxonomyId,omitempty"`
	InheritedBlastName           *string                         `protobuf:"bytes,2,opt" json:"InheritedBlastName,omitempty"`
	Rank                         *string                         `protobuf:"bytes,3,opt" json:"Rank,omitempty"`
	GeneticCode                  *GeneticCode                    `protobuf:"bytes,4,opt" json:"GeneticCode,omitempty"`
	OtherNames                   []string                        `protobuf:"bytes,5,rep" json:"OtherNames,omitempty"`
	Synonym                      *string                         `protobuf:"bytes,6,opt" json:"Synonym,omitempty"`
	Lineage                      []*Link                         `protobuf:"bytes,7,rep" json:"Lineage,omitempty"`
	CommentsAndReferences        []*CommentAndReference          `protobuf:"bytes,8,rep" json:"CommentsAndReferences,omitempty"`
	GenomeInformation            *GenomeInformation              `protobuf:"bytes,9,opt" json:"GenomeInformation,omitempty"`
	ExternalInformationResources []*ExternalInformationResources `protobuf:"bytes,10,rep" json:"ExternalInformationResources,omitempty"`
	Records                      []*Records                      `protobuf:"bytes,11,rep" json:"Records,omitempty"`
	XXX_unrecognized             []byte                          `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}

func (m *Entry) GetTaxonomyId() string {
	if m != nil && m.TaxonomyId != nil {
		return *m.TaxonomyId
	}
	return ""
}

func (m *Entry) GetInheritedBlastName() string {
	if m != nil && m.InheritedBlastName != nil {
		return *m.InheritedBlastName
	}
	return ""
}

func (m *Entry) GetRank() string {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return ""
}

func (m *Entry) GetGeneticCode() *GeneticCode {
	if m != nil {
		return m.GeneticCode
	}
	return nil
}

func (m *Entry) GetOtherNames() []string {
	if m != nil {
		return m.OtherNames
	}
	return nil
}

func (m *Entry) GetSynonym() string {
	if m != nil && m.Synonym != nil {
		return *m.Synonym
	}
	return ""
}

func (m *Entry) GetLineage() []*Link {
	if m != nil {
		return m.Lineage
	}
	return nil
}

func (m *Entry) GetCommentsAndReferences() []*CommentAndReference {
	if m != nil {
		return m.CommentsAndReferences
	}
	return nil
}

func (m *Entry) GetGenomeInformation() *GenomeInformation {
	if m != nil {
		return m.GenomeInformation
	}
	return nil
}

func (m *Entry) GetExternalInformationResources() []*ExternalInformationResources {
	if m != nil {
		return m.ExternalInformationResources
	}
	return nil
}

func (m *Entry) GetRecords() []*Records {
	if m != nil {
		return m.Records
	}
	return nil
}

type Records struct {
	DatabaseName     *string      `protobuf:"bytes,1,opt" json:"DatabaseName,omitempty"`
	References       []*Reference `protobuf:"bytes,2,rep" json:"References,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Records) Reset()         { *m = Records{} }
func (m *Records) String() string { return proto.CompactTextString(m) }
func (*Records) ProtoMessage()    {}

func (m *Records) GetDatabaseName() string {
	if m != nil && m.DatabaseName != nil {
		return *m.DatabaseName
	}
	return ""
}

func (m *Records) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

type Reference struct {
	Name             *Link   `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Accession        *string `protobuf:"bytes,2,opt" json:"Accession,omitempty"`
	GI               *string `protobuf:"bytes,3,opt" json:"GI,omitempty"`
	Links            []*Link `protobuf:"bytes,4,rep" json:"Links,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Reference) Reset()         { *m = Reference{} }
func (m *Reference) String() string { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()    {}

func (m *Reference) GetName() *Link {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Reference) GetAccession() string {
	if m != nil && m.Accession != nil {
		return *m.Accession
	}
	return ""
}

func (m *Reference) GetGI() string {
	if m != nil && m.GI != nil {
		return *m.GI
	}
	return ""
}

func (m *Reference) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

type Link struct {
	Name             *string `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Link             *string `protobuf:"bytes,2,opt" json:"Link,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}

func (m *Link) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Link) GetLink() string {
	if m != nil && m.Link != nil {
		return *m.Link
	}
	return ""
}

type GeneticCode struct {
	TranslationTable *int32  `protobuf:"varint,1,opt" json:"TranslationTable,omitempty"`
	Type             *string `protobuf:"bytes,2,opt" json:"Type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GeneticCode) Reset()         { *m = GeneticCode{} }
func (m *GeneticCode) String() string { return proto.CompactTextString(m) }
func (*GeneticCode) ProtoMessage()    {}

func (m *GeneticCode) GetTranslationTable() int32 {
	if m != nil && m.TranslationTable != nil {
		return *m.TranslationTable
	}
	return 0
}

func (m *GeneticCode) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

type CommentAndReference struct {
	Heading          *string `protobuf:"bytes,1,opt" json:"Heading,omitempty"`
	Content          *string `protobuf:"bytes,2,opt" json:"Content,omitempty"`
	Reference        *Link   `protobuf:"bytes,3,opt" json:"Reference,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommentAndReference) Reset()         { *m = CommentAndReference{} }
func (m *CommentAndReference) String() string { return proto.CompactTextString(m) }
func (*CommentAndReference) ProtoMessage()    {}

func (m *CommentAndReference) GetHeading() string {
	if m != nil && m.Heading != nil {
		return *m.Heading
	}
	return ""
}

func (m *CommentAndReference) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *CommentAndReference) GetReference() *Link {
	if m != nil {
		return m.Reference
	}
	return nil
}

type GenomeInformation struct {
	Type             *GenomeInformationType `protobuf:"varint,1,opt,enum=main.GenomeInformationType" json:"Type,omitempty"`
	Programs         []*Link                `protobuf:"bytes,2,rep" json:"Programs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *GenomeInformation) Reset()         { *m = GenomeInformation{} }
func (m *GenomeInformation) String() string { return proto.CompactTextString(m) }
func (*GenomeInformation) ProtoMessage()    {}

func (m *GenomeInformation) GetType() GenomeInformationType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return GenomeInformationType_NONE
}

func (m *GenomeInformation) GetPrograms() []*Link {
	if m != nil {
		return m.Programs
	}
	return nil
}

type ExternalInformationResources struct {
	LinkOut          *Link    `protobuf:"bytes,1,opt" json:"LinkOut,omitempty"`
	Subject          *Subject `protobuf:"varint,2,opt,enum=main.Subject" json:"Subject,omitempty"`
	LinkOutProvider  *Link    `protobuf:"bytes,3,opt" json:"LinkOutProvider,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ExternalInformationResources) Reset()         { *m = ExternalInformationResources{} }
func (m *ExternalInformationResources) String() string { return proto.CompactTextString(m) }
func (*ExternalInformationResources) ProtoMessage()    {}

func (m *ExternalInformationResources) GetLinkOut() *Link {
	if m != nil {
		return m.LinkOut
	}
	return nil
}

func (m *ExternalInformationResources) GetSubject() Subject {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return Subject_Unknown
}

func (m *ExternalInformationResources) GetLinkOutProvider() *Link {
	if m != nil {
		return m.LinkOutProvider
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.GenomeInformationType", GenomeInformationType_name, GenomeInformationType_value)
	proto.RegisterEnum("main.Subject", Subject_name, Subject_value)
}
