package tox

// https://go.dev/play/p/tN8mxT_V9h

// --- ย้ายไป gms/scan-to.go

// func ScanTo(s interface{}, d interface{}) error {

// 	// validate to s it must not pointer
// 	if s == nil {
// 		return errorx.StatusInternalServerError(`tox.ScanTo, s is nil`)
// 	}

// 	// validate to d is pointer of struct
// 	if d == nil {
// 		return errorx.StatusInternalServerError(`tox.ScanTo, d is nil`)
// 	}
// 	if reflect.TypeOf(d).Kind() != reflect.Ptr {
// 		return errorx.StatusInternalServerError(`tox.ScanTo, d is not pointer`)
// 	}
// 	vd := reflect.ValueOf(d)
// 	if vd.IsNil() {
// 		return errorx.StatusInternalServerError(`tox.ScanTo, value of d is nil`)
// 	}
// 	vdk := reflect.Indirect(vd).Type().Kind()
// 	if vdk == reflect.Interface {
// 		// set value to interface
// 		reflect.Indirect(reflect.ValueOf(d)).Set(reflect.Indirect(reflect.ValueOf(s)))
// 		// vd.Set(reflect.ValueOf(s))
// 		// d = reflect.ValueOf(s)
// 		return nil
// 	}
// 	if vdk != reflect.Struct && vdk != reflect.Slice {
// 		return errorx.StatusInternalServerError(`tox.ScanTo, type of d is not struct or slice`)
// 	}

// 	return setValue(reflect.ValueOf(d).Elem(), s)
// }

// // หาชื่อคัลมล์ที่อยู่ใน tag
// func findColumn(e reflect.Value, colName, tagName string) *string {
// 	for i := 0; i < e.NumField(); i++ {
// 		f := e.Type().Field(i)
// 		if strings.EqualFold(f.Name, colName) {
// 			return &f.Name // e.FieldByName(f.Name)
// 		}
// 		if a, _ := f.Tag.Lookup(tagName); a != `` {
// 			if tagName == `json` {
// 				if strings.EqualFold(a, colName) {
// 					return &f.Name // e.FieldByName(f.Name)
// 				}
// 				continue
// 			}
// 			for _, att := range strings.Split(a, `;`) {
// 				if strings.HasPrefix(att, `column:`) {
// 					column := strings.Split(a, `:`)[1]
// 					if column != `` {
// 						if strings.EqualFold(column, colName) {
// 							return &f.Name // e.FieldByName(f.Name)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// // หาชื่อคอลัมล์
// func fieldByName(e reflect.Value, name string) reflect.Value {
// 	// field := e.FieldByNameFunc(func(s string) bool {
// 	// 	return strings.EqualFold(s, name)
// 	// })
// 	field := e.FieldByName(name)
// 	if !field.IsValid() {
// 		// https://www.golangprograms.com/how-to-get-struct-variable-information-using-reflect-package.html
// 		// find gorm column
// 		if fieldName := findColumn(e, name, `gorm`); fieldName != nil {
// 			return e.FieldByName(*fieldName)
// 		}
// 		// find sap column
// 		if fieldName := findColumn(e, name, `sap`); fieldName != nil {
// 			return e.FieldByName(*fieldName)
// 		}
// 		// find json column
// 		if fieldName := findColumn(e, name, `json`); fieldName != nil {
// 			return e.FieldByName(*fieldName)
// 		}
// 		// แปลงชื่อฟิลด์ให้เป็น snake-case
// 		for i := 0; i < e.NumField(); i++ {
// 			f := e.Type().Field(i)
// 			if f.Type.Kind() != reflect.Struct {
// 				if strings.EqualFold(stringx.SnakeCase(f.Name), name) {
// 					return e.FieldByName(f.Name)
// 				}
// 			}
// 		}
// 	}
// 	return field
// }

// func setValue(field reflect.Value, v interface{}) error {

// 	if !field.IsValid() {
// 		return nil // return (`field[%v] not found in struct`, k)
// 	}
// 	if !field.CanSet() {
// 		return nil // return (`cannot set field value[%v]`, k)
// 	}

// 	// v is nil
// 	if v == nil {
// 		return nil
// 	}

// 	// value of v
// 	rv := reflect.ValueOf(v)

// 	// field type kind
// 	ftk := field.Type().Kind()

// 	// struct
// 	if ftk == reflect.Struct {
// 		fvm := Map(v)
// 		for ks, vs := range fvm {
// 			// if ks == `last_update_time` {
// 			// 	logx.Infoln(ks)
// 			// }
// 			fiels := fieldByName(field, ks)
// 			if !fiels.IsValid() {
// 				// ถ้าไม่เจอ field ให้หาคอลัมล์ที่เป็น embeded struct เช่น dbx.BaseTable
// 				for i := 0; i < field.NumField(); i++ {
// 					f := field.Type().Field(i)
// 					if f.Type.Kind() == reflect.Struct {
// 						// logx.Infof(`--------- %v`, f.Name)         //  BaseTable
// 						found := false
// 						fielEmbedded := reflect.New(f.Type).Elem() // dbx.BaseTable
// 						for j := 0; j < f.Type.NumField(); j++ {
// 							if strings.EqualFold(stringx.SnakeCase(f.Type.Field(j).Name), ks) {
// 								// logx.Infoln(f)                                        // {BaseTable  dbx.BaseTable  0 [0] true}
// 								// logx.Infoln(f.Type)                                   // dbx.BaseTable
// 								// logx.Infoln(f.Type.Field(j))                          // {ID  uuid.UUID bson:"_id" gorm:"type:uuid;primaryKey" json:"id" 0 [0] false}
// 								// logx.Infoln(f.Type.Field(j).Name)                     // ID
// 								// logx.Infoln(f.Type.Field(j).Type)                     // uuid.UUID
// 								// logx.Infoln(f.Type.FieldByName(f.Type.Field(j).Name)) // {ID  uuid.UUID bson:"_id" gorm:"type:uuid;primaryKey" json:"id" 0 [0] false} true
// 								// fielm := reflect.New(f.Type.Field(j).Type)
// 								fielm := reflect.New(f.Type.Field(j).Type).Elem()
// 								if ex := setValue(fielm, vs); ex != nil {
// 									return ex
// 								}
// 								// กำหนดค่ากลับเข้าไปใน fielEmbedded
// 								fiele := fielEmbedded.FieldByName(f.Type.Field(j).Name)
// 								fiele.Set(fielm)
// 								found = true
// 							}
// 						}
// 						if found {
// 							// กำหนด fielEmbedded กลับเข้าไปใน field
// 							fielx := field.FieldByName(f.Name)
// 							fielx.Set(fielEmbedded)
// 							break
// 						}
// 					}
// 				}
// 			}
// 			// เจอ field
// 			if ex := setValue(fiels, vs); ex != nil {
// 				return ex
// 			}
// 		}
// 		if fvm == nil {
// 			// initial empty slice
// 			for i := 0; i < field.Type().NumField(); i++ {
// 				fi := field.Type().Field(i)
// 				if fi.Type.Kind() == reflect.Slice {
// 					fiels := fieldByName(field, fi.Name)
// 					if ex := setValue(fiels, reflect.Zero(fi.Type).Interface()); ex != nil {
// 						return ex
// 					}
// 				}
// 			}
// 		}
// 		return nil
// 	}

// 	// slice
// 	if ftk == reflect.Slice {
// 		// x2 := field.Type() // []*string
// 		// x2 := field.Type().Elem() // *string
// 		fielx := reflect.MakeSlice(field.Type(), 0, rv.Len())
// 		for i := 0; i < rv.Len(); i++ {
// 			fielx = reflect.Append(fielx, reflect.Zero(field.Type().Elem()))
// 			if ex := setValue(fielx.Index(i), rv.Index(i).Interface()); ex != nil {
// 				return ex
// 			}
// 		}
// 		field.Set(fielx)
// 		return nil
// 	}

// 	// set value to field
// 	if ftk == rv.Type().Kind() {
// 		field.Set(rv)
// 		return nil
// 	}

// 	switch ftk {
// 	case reflect.Interface:
// 		field.Set(rv)

// 	case reflect.Float32, reflect.Float64:
// 		field.SetFloat(Float(v))

// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		field.SetInt(Int64(v))

// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		field.SetUint(Uint64(v))

// 	case reflect.Bool:
// 		field.SetBool(Bool(v))

// 	case reflect.String:
// 		field.SetString(String(v))

// 	case reflect.Array:
// 		if field.Type() == reflect.TypeOf((*uuid.UUID)(nil)).Elem() {
// 			vx := Uuid(v)
// 			if vx != nil {
// 				field.Set(reflect.ValueOf(*vx))
// 			}
// 		} else if field.Type() == reflect.TypeOf((*primitive.ObjectID)(nil)).Elem() {
// 			vx, _ := primitive.ObjectIDFromHex(String(v))
// 			if vx != primitive.NilObjectID {
// 				field.Set(reflect.ValueOf(vx))
// 			}

// 		} else {
// 			return errorx.StatusInternalServerError(`kind of pointer not in case-4[%v]`, ftk)
// 		}

// 	case reflect.Ptr:

// 		// ชนิดคอลัมล์
// 		fek := field.Type().Elem().Kind()

// 		// v is not nil > ชนิดคอลัมล์ตรงกัน
// 		if field.Type() == reflect.PtrTo(reflect.TypeOf(v)) {
// 			switch fek {
// 			case reflect.Float32, reflect.Float64:
// 			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			case reflect.Bool:
// 			case reflect.String:
// 			default:
// 				fielx := reflect.New(reflect.TypeOf(v))
// 				fielx.Elem().Set(rv)
// 				field.Set(fielx)
// 				return nil
// 			}
// 		}

// 		// v is not nil > ชนิดคอลัมล์ไม่ตรงกัน > convert data
// 		switch fek {
// 		case reflect.Slice:
// 			// logx.Infoln(field.Type())               // *[]float64
// 			// logx.Infoln(field.Type().Elem())        // []float64
// 			// logx.Infoln(field.Type().Elem().Elem()) // float64
// 			fiels := reflect.MakeSlice(field.Type().Elem(), 0, rv.Len())
// 			for i := 0; i < rv.Len(); i++ {
// 				fiels = reflect.Append(fiels, reflect.Zero(field.Type().Elem().Elem()))
// 				if ex := setValue(fiels.Index(i), rv.Index(i)); ex != nil {
// 					return ex
// 				}
// 			}
// 			// reflect.New(reflect.TypeOf(fielx).Elem())
// 			fielx := reflect.New(field.Type().Elem())
// 			fielx.Elem().Set(fiels)
// 			field.Set(fielx)

// 		case reflect.Interface:
// 			fielx := reflect.New(field.Type().Elem())
// 			fielx.Elem().Set(rv)
// 			field.Set(fielx)

// 		case reflect.Float32, reflect.Float64:
// 			vx := FloatPtr(v)
// 			if vx != nil {
// 				fielx := reflect.New(field.Type().Elem())
// 				fielx.Elem().SetFloat(*vx)
// 				field.Set(fielx)
// 			}

// 		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 			vx := Int64Ptr(v)
// 			if vx != nil {
// 				fielx := reflect.New(field.Type().Elem())
// 				fielx.Elem().SetInt(*vx)
// 				field.Set(fielx)
// 			}

// 		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			vx := Uint64Ptr(v)
// 			if vx != nil {
// 				fielx := reflect.New(field.Type().Elem())
// 				fielx.Elem().SetUint(*vx)
// 				field.Set(fielx)
// 			}

// 		case reflect.Bool:
// 			vx := BoolPtr(v)
// 			if vx != nil {
// 				fielx := reflect.New(field.Type().Elem())
// 				fielx.Elem().SetBool(*vx)
// 				field.Set(fielx)
// 			}

// 		case reflect.String:
// 			vx := StringPtr(v)
// 			if vx != nil {
// 				fielx := reflect.New(field.Type().Elem())
// 				fielx.Elem().SetString(*vx)
// 				field.Set(fielx)
// 			}

// 		case reflect.Array:
// 			if field.Type().Elem() == reflect.TypeOf((*uuid.UUID)(nil)).Elem() {
// 				vx := Uuid(v)
// 				if vx != nil {
// 					fielx := reflect.New(field.Type().Elem())
// 					fielx.Elem().Set(reflect.ValueOf(*vx))
// 					field.Set(fielx)
// 				}
// 			} else {
// 				return errorx.StatusInternalServerError(`kind of reflect.Array not in case-3[%v]`, field.Type().Elem().Kind())
// 			}

// 		case reflect.Struct:
// 			if field.Type().Elem() == reflect.TypeOf((*time.Time)(nil)).Elem() {
// 				vx, _ := time.Parse(time.RFC3339, rv.String())
// 				if vx != (time.Time{}) {
// 					fielx := reflect.New(field.Type().Elem())
// 					fielx.Elem().Set(reflect.ValueOf(vx))
// 					field.Set(fielx)
// 				}
// 			} else {
// 				return errorx.StatusInternalServerError(`kind of reflect.Struct not in case-4[%v]`, field.Type().Elem())
// 			}

// 		default:
// 			return errorx.StatusInternalServerError(`kind of reflect.Ptr not in case-2[%v]`, field.Type().Elem().Kind())
// 		}
// 	default:
// 		return errorx.StatusInternalServerError(`kind of field not in case-1[%v]`, ftk)
// 	}

// 	return nil
// }
