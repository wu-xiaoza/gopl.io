package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
    return []ent.Edge{
        // Create an inverse-edge called "owner" of type `User`
        // and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		//From represents a reversed-edge between two vertices that has a back-reference to its source edge.
        edge.From("owner", User.Type).
            Ref("cars").
            // setting the edge to unique, ensure
			// that a car can have only one owner.
			//Unique sets the edge type to be unique. Basically, it's limited the ent to be one of the two: one2one or one2many. one2one applied if the inverse-edge is also unique.
            Unique(),
    }
}