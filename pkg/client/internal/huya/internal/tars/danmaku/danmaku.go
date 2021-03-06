// Package danmaku comment
// This file was generated by tars2go 1.1.4
// Generated from danmaku.tars
package danmaku

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// SenderInfo struct implement
type SenderInfo struct {
	LUid      int64  `json:"lUid"`
	LImid     int64  `json:"lImid"`
	SNickName string `json:"sNickName"`
	IGender   int32  `json:"iGender"`
}

func (st *SenderInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *SenderInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int64(&st.LUid, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LImid, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SNickName, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IGender, 3, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *SenderInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require SenderInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *SenderInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.LUid, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LImid, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SNickName, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IGender, 3)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *SenderInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// ContentFormat struct implement
type ContentFormat struct {
	IFontColor  int32 `json:"iFontColor"`
	IFontSize   int32 `json:"iFontSize"`
	IPopupStyle int32 `json:"iPopupStyle"`
}

func (st *ContentFormat) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *ContentFormat) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.IFontColor, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IFontSize, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IPopupStyle, 2, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ContentFormat) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ContentFormat, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ContentFormat) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.IFontColor, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IFontSize, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IPopupStyle, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *ContentFormat) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// BulletFormat struct implement
type BulletFormat struct {
	IFontColor      int32 `json:"iFontColor"`
	IFontSize       int32 `json:"iFontSize"`
	ITextSpeed      int32 `json:"iTextSpeed"`
	ITransitionType int32 `json:"iTransitionType"`
	IPopupStyle     int32 `json:"iPopupStyle"`
}

func (st *BulletFormat) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *BulletFormat) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.IFontColor, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IFontSize, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ITextSpeed, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ITransitionType, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IPopupStyle, 4, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *BulletFormat) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require BulletFormat, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *BulletFormat) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.IFontColor, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IFontSize, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ITextSpeed, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ITransitionType, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IPopupStyle, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *BulletFormat) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// DecorationInfo struct implement
type DecorationInfo struct {
	IAppId    int32  `json:"iAppId"`
	IViewType int32  `json:"iViewType"`
	VData     []int8 `json:"vData"`
}

func (st *DecorationInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *DecorationInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.IAppId, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IViewType, 1, false)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(2, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			st.VData = make([]int8, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = _is.Read_int8(&st.VData[i0], 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SIMPLE_LIST {

			err, _ = _is.SkipTo(codec.BYTE, 0, true)
			if err != nil {
				return err
			}

			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			err = _is.Read_slice_int8(&st.VData, length, true)
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *DecorationInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require DecorationInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *DecorationInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.IAppId, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IViewType, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.SIMPLE_LIST, 2)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.VData)), 0)
	if err != nil {
		return err
	}

	err = _os.Write_slice_int8(st.VData)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *DecorationInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// UidNickName struct implement
type UidNickName struct {
	LUid      int64  `json:"lUid"`
	SNickName string `json:"sNickName"`
}

func (st *UidNickName) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *UidNickName) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int64(&st.LUid, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SNickName, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *UidNickName) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require UidNickName, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *UidNickName) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.LUid, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SNickName, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *UidNickName) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// MessageNotice struct implement
type MessageNotice struct {
	TUserInfo         SenderInfo       `json:"tUserInfo"`
	LTid              int64            `json:"lTid"`
	LSid              int64            `json:"lSid"`
	SContent          string           `json:"sContent"`
	IShowMode         int32            `json:"iShowMode"`
	TFormat           ContentFormat    `json:"tFormat"`
	TBulletFormat     BulletFormat     `json:"tBulletFormat"`
	ITermType         int32            `json:"iTermType"`
	VDecorationPrefix []DecorationInfo `json:"vDecorationPrefix"`
	VDecorationSuffix []DecorationInfo `json:"vDecorationSuffix"`
	VAtSomeone        []UidNickName    `json:"vAtSomeone"`
	LPid              int64            `json:"lPid"`
}

func (st *MessageNotice) ResetDefault() {
	st.TUserInfo.ResetDefault()
	st.TFormat.ResetDefault()
	st.TBulletFormat.ResetDefault()
}

//ReadFrom reads  from _is and put into struct.
func (st *MessageNotice) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = st.TUserInfo.ReadBlock(_is, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LTid, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LSid, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SContent, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.IShowMode, 4, false)
	if err != nil {
		return err
	}

	err = st.TFormat.ReadBlock(_is, 5, false)
	if err != nil {
		return err
	}

	err = st.TBulletFormat.ReadBlock(_is, 6, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ITermType, 7, false)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(8, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			st.VDecorationPrefix = make([]DecorationInfo, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.VDecorationPrefix[i0].ReadBlock(_is, 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SIMPLE_LIST {
			err = fmt.Errorf("not support simple_list type")
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	err, have, ty = _is.SkipToNoCheck(9, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			st.VDecorationSuffix = make([]DecorationInfo, length)
			for i1, e1 := int32(0), length; i1 < e1; i1++ {

				err = st.VDecorationSuffix[i1].ReadBlock(_is, 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SIMPLE_LIST {
			err = fmt.Errorf("not support simple_list type")
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	err, have, ty = _is.SkipToNoCheck(10, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			st.VAtSomeone = make([]UidNickName, length)
			for i2, e2 := int32(0), length; i2 < e2; i2++ {

				err = st.VAtSomeone[i2].ReadBlock(_is, 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SIMPLE_LIST {
			err = fmt.Errorf("not support simple_list type")
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	err = _is.Read_int64(&st.LPid, 11, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *MessageNotice) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MessageNotice, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *MessageNotice) WriteTo(_os *codec.Buffer) error {
	var err error

	err = st.TUserInfo.WriteBlock(_os, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LTid, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LSid, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SContent, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.IShowMode, 4)
	if err != nil {
		return err
	}

	err = st.TFormat.WriteBlock(_os, 5)
	if err != nil {
		return err
	}

	err = st.TBulletFormat.WriteBlock(_os, 6)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ITermType, 7)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 8)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.VDecorationPrefix)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.VDecorationPrefix {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	err = _os.WriteHead(codec.LIST, 9)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.VDecorationSuffix)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.VDecorationSuffix {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	err = _os.WriteHead(codec.LIST, 10)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.VAtSomeone)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.VAtSomeone {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	err = _os.Write_int64(st.LPid, 11)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *MessageNotice) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
