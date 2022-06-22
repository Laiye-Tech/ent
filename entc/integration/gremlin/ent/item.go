// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
}

// FromResponse scans the gremlin response data into Item.
func (i *Item) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scani struct {
		ID   string `json:"id,omitempty"`
		Text string `json:"text,omitempty"`
	}
	if err := vmap.Decode(&scani); err != nil {
		return err
	}
	i.ID = scani.ID
	i.Text = scani.Text
	return nil
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("text=")
	builder.WriteString(i.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

// FromResponse scans the gremlin response data into Items.
func (i *Items) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scani []struct {
		ID   string `json:"id,omitempty"`
		Text string `json:"text,omitempty"`
	}
	if err := vmap.Decode(&scani); err != nil {
		return err
	}
	for _, v := range scani {
		*i = append(*i, &Item{
			ID:   v.ID,
			Text: v.Text,
		})
	}
	return nil
}

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
