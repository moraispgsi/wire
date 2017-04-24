package wire

import (
        "reflect"    
)

type Component interface {
    Process()
}



//returns a map o inputs, key is the variable name and value is its kind 
func GetComponentInputs(comp Component) map[string]reflect.Kind {
    
    result := make(map[string]reflect.Kind)

    s := reflect.ValueOf(comp).Elem()
    typeOfComponent := s.Type()
    
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        ft := f.Type()
        if ft.Kind() == reflect.Chan {
            
            if ft.ChanDir() == reflect.RecvDir {
                result[typeOfComponent.Field(i).Name] = ft.Elem().Kind()
                
            }
            
        }
    }
    
    return result
    
}

//returns a map o outputs, key is the variable name and value is its kind 
func GetComponentOutputs(comp Component) map[string]reflect.Kind {
    
    result := make(map[string]reflect.Kind)
    
    s := reflect.ValueOf(comp).Elem()
    typeOfComponent := s.Type()

    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        ft := f.Type()
        if ft.Kind() == reflect.Chan {
            
            if ft.ChanDir() == reflect.SendDir {
                result[typeOfComponent.Field(i).Name] = ft.Elem().Kind()
                
            }
            
        }
    }
    
    return result
    
}
