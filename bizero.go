package errkratos

// BizEro 把 erk 包一层，但不实现 golang 的 error 接口，这样就能避免空指针也被当成是有 error 的，避免踩坑也避免时刻保持警惕，让开发者在写业务时保持放松状态
type BizEro struct {
	Erk *Erk
}

// NewBizEro 返回的 BizEro 只是把 erk 包一层
func NewBizEro(erk *Erk) *BizEro {
	return &BizEro{
		Erk: erk,
	}
}
