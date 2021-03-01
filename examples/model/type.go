package model
// 此文件是根据GoGoTypePoint自动生成

import (
    "time"
)

func Time(v time.Time) *time.Time { return &v }

func Uint(v uint) *uint { return &v }

func Bool(v bool) *bool { return &v }

func Int(v int) *int { return &v }

func Int32(v int32) *int32 { return &v }

func Int64(v int64) *int64 { return &v }

func Uint32(v uint32) *uint32 { return &v }

func Uint64(v uint64) *uint64 { return &v }

func Float32(v float32) *float32 { return &v }

func Float64(v float64) *float64 { return &v }

func String(v string) *string { return &v }

func PointTime(v *time.Time) time.Time {
    if v == nil {
        return time.Time{}
    }
    return *v
}

func PointUint(v *uint) uint {
    if v == nil {
        return 0
    }
    return *v
}

func PointBool(v *bool) bool {
    if v == nil {
        return false
    }
    return *v
}

func PointInt(v *int) int {
    if v == nil {
        return 0
    }
    return *v
}

func PointInt32(v *int32) int32 {
    if v == nil {
        return 0
    }
    return *v
}

func PointInt64(v *int64) int64 {
    if v == nil {
        return 0
    }
    return *v
}

func PointUint32(v *uint32) uint32 {
    if v == nil {
        return 0
    }
    return *v
}

func PointUint64(v *uint64) uint64 {
    if v == nil {
        return 0
    }
    return *v
}

func PointFloat32(v *float32) float32 {
    if v == nil {
        return 0
    }
    return *v
}

func PointFloat64(v *float64) float64 {
    if v == nil {
        return 0
    }
    return *v
}

func PointString(v *string) string {
    if v == nil {
        return ""
    }
    return *v
}