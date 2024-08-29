package condb

import "gitlab.dohome.technology/dohome-2020/go-servicex/validx"

func IsColtypeUuid(ct string) bool {
	return validx.IsColtypeUuid(ct)
}

func IsColtypeTime(ct string) bool {
	return validx.IsColtypeTime(ct)
}

func IsColtypeString(ct string) bool {
	return validx.IsColtypeString(ct)
}

func IsColtypeInt(ct string) bool {
	return validx.IsColtypeInt(ct)
}

func IsColtypeFloat(ct string) bool {
	return validx.IsColtypeFloat(ct)
}

func IsColtypeBool(ct string) bool {
	return validx.IsColtypeBool(ct)
}
