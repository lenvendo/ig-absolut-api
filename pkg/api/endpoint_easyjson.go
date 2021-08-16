// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

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

func easyjson3a9ed663DecodeApiPkgApi(in *jlexer.Lexer, out *UserRegResponse) {
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
		case "status":
			out.Status = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi(out *jwriter.Writer, in UserRegResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Status != "" {
		const prefix string = ",\"status\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserRegResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserRegResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserRegResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserRegResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi1(in *jlexer.Lexer, out *UserRegRequest) {
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
		case "phone":
			out.Phone = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi1(out *jwriter.Writer, in UserRegRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Phone != "" {
		const prefix string = ",\"phone\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserRegRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserRegRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserRegRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserRegRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi1(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi2(in *jlexer.Lexer, out *UserProfileResponse) {
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
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(User)
				}
				(*out.User).UnmarshalEasyJSON(in)
			}
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
func easyjson3a9ed663EncodeApiPkgApi2(out *jwriter.Writer, in UserProfileResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.User != nil {
		const prefix string = ",\"user\":"
		first = false
		out.RawString(prefix[1:])
		(*in.User).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserProfileResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserProfileResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserProfileResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserProfileResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi2(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi3(in *jlexer.Lexer, out *UserProfileRequest) {
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
		case "session":
			out.Session = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi3(out *jwriter.Writer, in UserProfileRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Session != "" {
		const prefix string = ",\"session\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Session))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserProfileRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserProfileRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserProfileRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserProfileRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi3(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi4(in *jlexer.Lexer, out *UserConfirmResponse) {
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
		case "session":
			out.Session = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi4(out *jwriter.Writer, in UserConfirmResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Session != "" {
		const prefix string = ",\"session\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Session))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserConfirmResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserConfirmResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserConfirmResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserConfirmResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi4(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi5(in *jlexer.Lexer, out *UserConfirmRequest) {
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
		case "code":
			out.Code = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi5(out *jwriter.Writer, in UserConfirmRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Code != "" {
		const prefix string = ",\"code\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserConfirmRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserConfirmRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserConfirmRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserConfirmRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi5(l, v)
}
func easyjson3a9ed663DecodeApiPkgApi6(in *jlexer.Lexer, out *User) {
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
			out.Id = string(in.String())
		case "isConfirmed":
			out.IsConfirmed = bool(in.Bool())
		case "createdAt":
			out.CreatedAt = string(in.String())
		case "updatedAt":
			out.UpdatedAt = string(in.String())
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
func easyjson3a9ed663EncodeApiPkgApi6(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	if in.IsConfirmed {
		const prefix string = ",\"isConfirmed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsConfirmed))
	}
	if in.CreatedAt != "" {
		const prefix string = ",\"createdAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CreatedAt))
	}
	if in.UpdatedAt != "" {
		const prefix string = ",\"updatedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UpdatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeApiPkgApi6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeApiPkgApi6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeApiPkgApi6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeApiPkgApi6(l, v)
}