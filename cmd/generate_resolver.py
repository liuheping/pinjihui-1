import re
import random
import os

methodTmpl = """
func (r *{$typeName}Resolver) {$method}() {$fieldType} {
    {$value}
    return {$res}
}
"""
tmplHeader = """
package resolver

{$import}

type {$typeName}Resolver struct {
    //m *model.{$typeName}
}
"""
importTmplOrigin = 'import (\n    //"pinjihui.com/pinjihui/model"'

first_lower = lambda s: s[:1].lower() + s[1:] if s else ''
first_upper = lambda s: s[:1].upper() + s[1:] if s else ''


pattern = re.compile(r'type (\w+) \{\s+?([^\}]+)\n\}')
fieldPattern = re.compile(r'(\w+?): (\w+|\[(\w+)(!)?\])(!)?')
enumPattern = re.compile(r'enum (\w+) \{([^\}]+)\}')
basicTypeMap = {"Int": "int32", "Float": 'float64', "String": "string", "Boolean": "bool"}

def initValue(typeName):
    typeMapValue = {"Int": "int32(3)", "Float": "0.0", "String": '"test string"', "Boolean": "false"}
    return typeMapValue[typeName]

def hump2underline(hunp_str):
    '''
    驼峰形式字符串转成下划线形式
    :param hunp_str: 驼峰形式字符串
    :return: 字母全小写的下划线形式字符串
    '''
    # 匹配正则，匹配小写字母和大写字母的分界位置
    p = re.compile(r'([a-z]|\d)([A-Z])')
    # 这里第二个参数使用了正则分组的后向引用
    sub = re.sub(p, r'\1_\2', hunp_str).lower()
    return sub

def generate(file):
    tmpl = tmplHeader
    def shouldReturnPoint(isRequire, goFieldType):
        return isRequire == None or goFieldType.endswith("Resolver")

    with open(file, 'r') as f:
        data = f.read()
        res = re.match(pattern, data)
        if res == None:
            return -1
        # print(res.group(0))
        typeName = res.group(1)
        print("typeName:" + typeName)

        enumsTypes = re.findall(enumPattern, data)
        enums = [x[0] for x in enumsTypes]
        enumsMap = {e[0]:e[1].strip().split("\n    ") for e in enumsTypes}
        print("enums:", enumsMap)

        fields = [re.match(fieldPattern, x.strip()).groups() for x in res.group(2).split("\n")]
        print(fields)
        shouldImportTime = False
        should_import_graphql_go = False
        for field in fields:
            methodTmplNew = methodTmpl
            gqlFieldType = field[1]
            isArray = False
            if gqlFieldType.startswith("["):
                gqlFieldType = field[2]
                isArray = True

            if gqlFieldType in ["Int", "String", "Boolean", "Float"]:
                value = initValue(gqlFieldType)
                goFieldType = basicTypeMap[gqlFieldType]
            elif gqlFieldType in enums:
                goFieldType = "string"
                random.shuffle(enumsMap[gqlFieldType])
                value = '"' + enumsMap[gqlFieldType][0] + '"'
            elif gqlFieldType == "ID":
                goFieldType = "graphql.ID"
                value = 'graphql.ID("xjauwkahsi92h1j")'
                should_import_graphql_go = True
            elif gqlFieldType == "Time":
                shouldImportTime = True
                should_import_graphql_go = True
                goFieldType = "graphql.Time, error"
                value = 'time.Parse(time.RFC3339, "2018-04-01 12:04:56.539453")'
            else:
                goFieldType = first_lower(gqlFieldType) + "Resolver"
                value = goFieldType + "{}"

            if isArray:
                if shouldReturnPoint(field[3], goFieldType):
                    goFieldType = "*" + goFieldType
                goFieldType = "[]" + goFieldType
                value = "make("+goFieldType+", 3)"

            if shouldReturnPoint(field[4], goFieldType):
                goFieldType = "*" + goFieldType

            res = "res"
            if gqlFieldType == "Time":
                goFieldType = "(" + goFieldType + ")"
                value = "res, err := " + value
                res = "graphql.Time{Time: res}, err"
            else:
                value = "res := " + value

            if shouldReturnPoint(field[4], goFieldType):
                res = '&' + res

            method = "ID" if field[0] == 'id' else first_upper(field[0])
            methodTmplNew = methodTmplNew.replace("{$method}", method)
            methodTmplNew = methodTmplNew.replace("{$res}", res)
            methodTmplNew = methodTmplNew.replace("{$value}", value)
            tmpl += methodTmplNew.replace("{$fieldType}", goFieldType)

        importTmpl = importTmplOrigin
        if shouldImportTime:
            importTmpl += '\n    "time"'
        if should_import_graphql_go:
            importTmpl += '\n    "github.com/graph-gophers/graphql-go"'
        importTmpl += "\n)"
        tmpl = tmpl.replace("{$import}", first_lower(importTmpl))
        tmpl = tmpl.replace("{$typeName}", first_lower(typeName))

        resolverFile = "../resolver/" + hump2underline(typeName) + ".go"
        with open(resolverFile, 'w') as rf:
            rf.write(tmpl)

    return 0   

if __name__ == '__main__':

    dirname = "../schema/type"
    typeFiles = [os.path.join(dirname, name) for name in os.listdir(dirname) 
                if name.endswith(".graphql") and not name.startswith("user")]
    for file in typeFiles:
        generate(file)
