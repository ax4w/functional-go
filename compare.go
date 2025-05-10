package functionalgo

type ComparisonResult int

const (
	LT ComparisonResult = iota - 1 // Less than
	EQ                             // Equal
	GT                             // Greater than
)

func (c ComparisonResult) String() string {
	switch c {
	case LT:
		return "LT"
	case EQ:
		return "EQ"
	case GT:
		return "GT"
	default:
		return "Unknown"
	}
}

func Compare[T any](a, b T) ComparisonResult {
	switch any(a).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		va := any(a)
		vb := any(b)

		if ia, ok := va.(int); ok {
			if ib, ok := vb.(int); ok {
				if ia < ib {
					return LT
				} else if ia > ib {
					return GT
				}
				return EQ
			}
		}

		if ia, ok := va.(int8); ok {
			if ib, ok := vb.(int8); ok {
				if ia < ib {
					return LT
				} else if ia > ib {
					return GT
				}
				return EQ
			}
		}

		if ia, ok := va.(int16); ok {
			if ib, ok := vb.(int16); ok {
				if ia < ib {
					return LT
				} else if ia > ib {
					return GT
				}
				return EQ
			}
		}

		if ia, ok := va.(int32); ok {
			if ib, ok := vb.(int32); ok {
				if ia < ib {
					return LT
				} else if ia > ib {
					return GT
				}
				return EQ
			}
		}

		if ia, ok := va.(int64); ok {
			if ib, ok := vb.(int64); ok {
				if ia < ib {
					return LT
				} else if ia > ib {
					return GT
				}
				return EQ
			}
		}

		if ua, ok := va.(uint); ok {
			if ub, ok := vb.(uint); ok {
				if ua < ub {
					return LT
				} else if ua > ub {
					return GT
				}
				return EQ
			}
		}

		if ua, ok := va.(uint8); ok {
			if ub, ok := vb.(uint8); ok {
				if ua < ub {
					return LT
				} else if ua > ub {
					return GT
				}
				return EQ
			}
		}

		if ua, ok := va.(uint16); ok {
			if ub, ok := vb.(uint16); ok {
				if ua < ub {
					return LT
				} else if ua > ub {
					return GT
				}
				return EQ
			}
		}

		if ua, ok := va.(uint32); ok {
			if ub, ok := vb.(uint32); ok {
				if ua < ub {
					return LT
				} else if ua > ub {
					return GT
				}
				return EQ
			}
		}

		if ua, ok := va.(uint64); ok {
			if ub, ok := vb.(uint64); ok {
				if ua < ub {
					return LT
				} else if ua > ub {
					return GT
				}
				return EQ
			}
		}

		if fa, ok := va.(float32); ok {
			if fb, ok := vb.(float32); ok {
				if fa < fb {
					return LT
				} else if fa > fb {
					return GT
				}
				return EQ
			}
		}

		if fa, ok := va.(float64); ok {
			if fb, ok := vb.(float64); ok {
				if fa < fb {
					return LT
				} else if fa > fb {
					return GT
				}
				return EQ
			}
		}

	case string:
		sa := any(a).(string)
		sb := any(b).(string)
		if sa < sb {
			return LT
		} else if sa > sb {
			return GT
		}
		return EQ

	case bool:
		ba := any(a).(bool)
		bb := any(b).(bool)
		if !ba && bb {
			return LT
		} else if ba && !bb {
			return GT
		}
		return EQ
	}
	if any(a) == any(b) {
		return EQ
	}
	return GT
}
