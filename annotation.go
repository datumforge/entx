package entx

import "encoding/json"

// CascadeAnnotationName is a name for our cascading delete annotation
var CascadeAnnotationName = "DATUM_CASCADE"

// CascadeThroughAnnotationName is a name for our cascading through edge delete annotation
var CascadeThroughAnnotationName = "DATUM_CASCADE_THROUGH"

// SchemaGenAnnotationName is a name for our graphql schema generation annotation
var SchemaGenAnnotationName = "DATUM_SCHEMAGEN"

// QueryGenAnnotationName is a name for our graphql query generation annotation
var QueryGenAnnotationName = "DATUM_QUERYGEN"

// CascadeAnnotation is an annotation used to indicate that an edge should be cascaded
type CascadeAnnotation struct {
	Field string
}

// CascadeThroughAnnotation is an annotation used to indicate that an edge should be cascaded through
type CascadeThroughAnnotation struct {
	Schemas []ThroughCleanup
}

// ThroughCleanup is a struct used to indicate the field and through edge to cascade through
type ThroughCleanup struct {
	Field   string
	Through string
}

// SchemaGenAnnotation is an annotation used to indicate that schema generation should be skipped for this type
type SchemaGenAnnotation struct {
	Skip bool
}

// QueryGenAnnotation is an annotation used to indicate that query generation should be skipped for this type
type QueryGenAnnotation struct {
	Skip bool
}

// Name returns the name of the CascadeAnnotation
func (a CascadeAnnotation) Name() string {
	return CascadeAnnotationName
}

// Name returns the name of the CascadeThroughAnnotation
func (a CascadeThroughAnnotation) Name() string {
	return CascadeThroughAnnotationName
}

// Name returns the name of the SchemaGenAnnotation
func (a SchemaGenAnnotation) Name() string {
	return SchemaGenAnnotationName
}

// Name returns the name of the QueryGenAnnotation
func (a QueryGenAnnotation) Name() string {
	return QueryGenAnnotationName
}

// CascadeAnnotationField sets the field name of the edge containing the ID of a record from the current schema
func CascadeAnnotationField(fieldname string) *CascadeAnnotation {
	return &CascadeAnnotation{
		Field: fieldname,
	}
}

// CascadeThroughAnnotationField sets the field name of the edge containing the ID of a record from the current schema
func CascadeThroughAnnotationField(schemas []ThroughCleanup) *CascadeThroughAnnotation {
	return &CascadeThroughAnnotation{
		Schemas: schemas,
	}
}

// SchemaGenSkip sets whether the schema generation should be skipped for this type
func SchemaGenSkip(skip bool) *SchemaGenAnnotation {
	return &SchemaGenAnnotation{
		Skip: skip,
	}
}

// QueryGenSkip sets whether the query generation should be skipped for this type
func QueryGenSkip(skip bool) *QueryGenAnnotation {
	return &QueryGenAnnotation{
		Skip: skip,
	}
}

// Decode unmarshalls the CascadeAnnotation
func (a *CascadeAnnotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, a)
}

// Decode unmarshalls the CascadeThroughAnnotation
func (a *CascadeThroughAnnotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, a)
}

// Decode unmarshalls the SchemaGenAnnotation
func (a *SchemaGenAnnotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, a)
}

// Decode unmarshalls the QueryGenAnnotation
func (a *QueryGenAnnotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, a)
}
