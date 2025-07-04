// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package forms

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms(in *jlexer.Lexer, out *FriendsInfoOut) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ID).UnmarshalText(data))
			}
		case "username":
			out.Username = string(in.String())
		case "firstname":
			out.FirstName = string(in.String())
		case "lastname":
			out.LastName = string(in.String())
		case "avatar_url":
			out.AvatarURL = string(in.String())
		case "university":
			out.University = string(in.String())
		case "is_online":
			out.IsOnline = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms(out *jwriter.Writer, in FriendsInfoOut) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.ID).MarshalText())
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"firstname\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastname\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"avatar_url\":"
		out.RawString(prefix)
		out.String(string(in.AvatarURL))
	}
	{
		const prefix string = ",\"university\":"
		out.RawString(prefix)
		out.String(string(in.University))
	}
	{
		const prefix string = ",\"is_online\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsOnline))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FriendsInfoOut) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FriendsInfoOut) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FriendsInfoOut) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FriendsInfoOut) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms(l, v)
}
func easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms1(in *jlexer.Lexer, out *FriendRequestDel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "friend_id":
			out.FriendID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms1(out *jwriter.Writer, in FriendRequestDel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"friend_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.FriendID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FriendRequestDel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FriendRequestDel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FriendRequestDel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FriendRequestDel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms1(l, v)
}
func easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms2(in *jlexer.Lexer, out *FriendRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "receiver_id":
			out.ReceiverID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms2(out *jwriter.Writer, in FriendRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"receiver_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ReceiverID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FriendRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FriendRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3994edd1EncodeQuickflowGatewayInternalDeliveryHttpForms2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FriendRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FriendRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3994edd1DecodeQuickflowGatewayInternalDeliveryHttpForms2(l, v)
}
