package scontext

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AnyMessage interface {
	ProtoReflect() protoreflect.Message
}

func TargetUserID(m AnyMessage) (string, bool) {
	msg := m.ProtoReflect()

	d := msg.Descriptor()

	e := d.Extensions().ByName("target_user_id")

	field := msg.Get(e).Int()
	if field == 0 {
		return "", false
	}
	target := d.Fields().Get(int(field))

	uid := msg.Get(target).String()

	return uid, uid == ""
}
