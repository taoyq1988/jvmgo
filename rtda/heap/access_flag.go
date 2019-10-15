package heap

const (
	AccPublic       = 0x0001
	AccPrivate      = 0x0002
	AccProtected    = 0x0004
	AccStatic       = 0x0008
	AccFinal        = 0x0010
	AccSuper        = 0x0020
	AccSynchronized = 0x0020
	AccVolatile     = 0x0040
	AccBridge       = 0x0040
	AccTransient    = 0x0080
	AccVarargs      = 0x0080 // 5.0
	AccNative       = 0x0100
	AccInterface    = 0x0200
	AccAbstract     = 0x0400
	AccStrict       = 0x0800
	AccSynthetic    = 0x1000
	AccAnnotation   = 0x2000 // 5.0
	AccEnum         = 0x4000 // 5.0
)

type AccessFlag uint16

func (flags AccessFlag) IsPublic() bool       { return flags&AccPublic != 0 }
func (flags AccessFlag) IsPrivate() bool      { return flags&AccPrivate != 0 }
func (flags AccessFlag) IsProtected() bool    { return flags&AccProtected != 0 }
func (flags AccessFlag) IsStatic() bool       { return flags&AccStatic != 0 }
func (flags AccessFlag) IsFinal() bool        { return flags&AccFinal != 0 }
func (flags AccessFlag) IsSuper() bool        { return flags&AccSuper != 0 }
func (flags AccessFlag) IsSynchronized() bool { return flags&AccSynchronized != 0 }
func (flags AccessFlag) IsVolatile() bool     { return flags&AccVolatile != 0 }
func (flags AccessFlag) IsBridge() bool       { return flags&AccBridge != 0 }
func (flags AccessFlag) IsTransient() bool    { return flags&AccTransient != 0 }
func (flags AccessFlag) IsVarargs() bool      { return flags&AccVarargs != 0 }
func (flags AccessFlag) IsNative() bool       { return flags&AccNative != 0 }
func (flags AccessFlag) IsInterface() bool    { return flags&AccInterface != 0 }
func (flags AccessFlag) IsAbstract() bool     { return flags&AccAbstract != 0 }
func (flags AccessFlag) IsStrict() bool       { return flags&AccStrict != 0 }
func (flags AccessFlag) IsSynthetic() bool    { return flags&AccSynthetic != 0 }
func (flags AccessFlag) IsAnnotation() bool   { return flags&AccAnnotation != 0 }
func (flags AccessFlag) IsEnum() bool         { return flags&AccEnum != 0 }
