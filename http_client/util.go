package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"
)

// GetParam 将结构体转换为请求参数Map
func GetParam[T any](obj T) (map[string]any, error) {
	obj_type := reflect.TypeOf(obj)
	obj_value := reflect.ValueOf(obj)
	if obj_type.Kind() != reflect.Struct {
		obj_type = obj_type.Elem()
		obj_value = obj_value.Elem()
	}
	switch obj_type.Kind() {
	case reflect.Struct: //结构体就遍历结构体的字段来
		{
			result := map[string]any{}
			for i := 0; i < obj_type.NumField(); i++ {
				field := obj_type.Field(i)
				tag_name := field.Tag.Get("parm") //处理Tag
				if tag_name == "" {
					tag_name = field.Name
				}
				result[tag_name] = obj_value.Field(i).Interface()
			}
			return result, nil
		}
	default:
		{
			//如果不是一个结构体就直接是一个pair的map
			return map[string]any{
				obj_type.Name(): obj,
			}, nil
		}
	}
}

// ParseFormTag 解析表单标签 type:str;name:data;file_name:aaa
func ParseFormTag(tag_val string) (map[string]string, error) {
	tagItems := strings.Split(tag_val, ";")
	result := make(map[string]string)
	for _, tagItem := range tagItems {
		tagKv := strings.Split(tagItem, ":")
		if len(tagKv) != 2 {
			return nil, errors.New("format error")
		}
		result[tagKv[0]] = tagKv[1]
	}
	return result, nil
}

// GetForm 将结构体转换为表单Map
func GetForm(obj any) (*HttpFormOpt, error) {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	var result = &HttpFormOpt{
		Files: make([]struct {
			FileName    string
			FieldName   string
			FileContent io.Reader
		}, 0),
		Fields: make([]struct {
			FieldName    string
			FieldContent string
		}, 0),
	}
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
		objValue = objValue.Elem()
	}
	if objType.Kind() == reflect.Struct || objValue.Elem().Type().Kind() == reflect.Struct {
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			tagMap, err := ParseFormTag(field.Tag.Get("form"))
			if err != nil {
				return nil, err
			}
			formType, ok := tagMap["type"]
			if !ok {
				return nil, errors.New("form tag not found")
			}
			formVal, ok := tagMap["name"]
			if !ok {
				return nil, errors.New("form tag not found")
			}
			if formType == "file" {
				var fileInfo = struct {
					FileName    string
					FieldName   string
					FileContent io.Reader
				}{
					FileName:    "",
					FieldName:   formVal,
					FileContent: nil,
				}
				if field.Type.Kind() != reflect.Struct {
					return nil, errors.New("file form is no struct")
				} else {
					//处理表单的文件类型
					fieldVal := reflect.ValueOf(objValue.Field(i).Interface())
					for j := 0; j < fieldVal.NumField(); j++ {
						fieldField := fieldVal.Type().Field(j)
						tag_name := fieldField.Tag.Get("form_file")
						if tag_name == "name" {
							name, ok := (fieldVal.Field(j).Interface()).(string)
							if !ok {
								return nil, errors.New("form_file field name error")
							}
							fileInfo.FileName = name
						}
						if tag_name == "content" {
							content, ok := fieldVal.Field(j).Interface().(io.Reader)
							if !ok {
								return nil, errors.New("form_file field content error")
							}
							fileInfo.FileContent = content
						}
					}
				}
				result.Files = append(result.Files, fileInfo)
			} else {
				info, err := json.Marshal(objValue.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				result.Fields = append(result.Fields, struct {
					FieldName    string
					FieldContent string
				}{FieldName: formVal, FieldContent: string(info)})
			}
		}
	}
	return result, nil
}
