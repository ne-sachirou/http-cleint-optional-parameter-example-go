package example

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExamplePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The param1 property
    param1 *int64
    // The param2 property
    param2 *int64
}
// NewExamplePostRequestBody instantiates a new ExamplePostRequestBody and sets the default values.
func NewExamplePostRequestBody()(*ExamplePostRequestBody) {
    m := &ExamplePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExamplePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExamplePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExamplePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExamplePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExamplePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["param1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParam1(val)
        }
        return nil
    }
    res["param2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParam2(val)
        }
        return nil
    }
    return res
}
// GetParam1 gets the param1 property value. The param1 property
// returns a *int64 when successful
func (m *ExamplePostRequestBody) GetParam1()(*int64) {
    return m.param1
}
// GetParam2 gets the param2 property value. The param2 property
// returns a *int64 when successful
func (m *ExamplePostRequestBody) GetParam2()(*int64) {
    return m.param2
}
// Serialize serializes information the current object
func (m *ExamplePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("param1", m.GetParam1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("param2", m.GetParam2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExamplePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParam1 sets the param1 property value. The param1 property
func (m *ExamplePostRequestBody) SetParam1(value *int64)() {
    m.param1 = value
}
// SetParam2 sets the param2 property value. The param2 property
func (m *ExamplePostRequestBody) SetParam2(value *int64)() {
    m.param2 = value
}
type ExamplePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParam1()(*int64)
    GetParam2()(*int64)
    SetParam1(value *int64)()
    SetParam2(value *int64)()
}
