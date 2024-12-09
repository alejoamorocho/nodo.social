package firebase

import (
    "encoding/json"
    "time"
)

// FirestoreEvent es la estructura que representa un evento de Firestore
type FirestoreEvent struct {
    OldValue   FirestoreValue `json:"oldValue"`
    Value      FirestoreValue `json:"value"`
    UpdateMask struct {
        FieldPaths []string `json:"fieldPaths"`
    } `json:"updateMask"`
}

// DataTo convierte los datos del evento a una estructura
func (e *FirestoreEvent) DataTo(v interface{}) error {
    b, err := json.Marshal(e.Value.Fields)
    if err != nil {
        return err
    }
    return json.Unmarshal(b, v)
}

// FirestoreValue representa el valor de un documento en un evento de Firestore
type FirestoreValue struct {
    CreateTime time.Time        `json:"createTime"`
    Fields     map[string]Value `json:"fields"`
    Name       string          `json:"name"`
    UpdateTime time.Time        `json:"updateTime"`
}

// Value representa un valor en un documento de Firestore
type Value struct {
    StringValue  string  `json:"stringValue,omitempty"`
    IntegerValue int64   `json:"integerValue,omitempty"`
    DoubleValue  float64 `json:"doubleValue,omitempty"`
    BooleanValue bool    `json:"booleanValue,omitempty"`
}
