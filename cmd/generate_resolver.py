import re

first_lower = lambda s: s[:1].lower() + s[1:] if s else ''
first_upper = lambda s: s[:1].upper() + s[1:] if s else ''

if __name__ == '__main__':
    pattern = re.compile(r'type (\w+?) \{\s+?([^\}]+)\n\}')
    pattern1 = re.compile(r'(\w+?): (\w+|\[\w+\])(!)?')

    file = "/home/wangbo/go/src/pinjihui.com/pinjihui/schema/type/order.graphql"
    with open(file, 'r') as f:
        data = f.read()
        res = re.match(pattern, data)
        # print(res.groups())
        typeName = res.group(1)
        print("typeName:" + typeName)

        tmpl = """
package resolver

import (
	"pinjihui.com/pinjihui/model"
	"github.com/graph-gophers/graphql-go"
	"time"
)

type {$typeName}Resolver struct {
	m *model.{$typeName}
}
"""
        tmpl = tmpl.replace("{$typeName}", typeName)
        # print(res.group(2).split("\n"))
        # exit(0)
        fields = [re.match(pattern1, x.strip()).groups() for x in res.group(2).split("\n")]
        print(fields)
        for field in fields:
            methodTmpl = """
func (r *{$typeName}Resolver) {$method}() {$fieldType} {
	return &r.m.{$method}
}
"""
            gqlFieldType = field[1]
            isArray = False
            if gqlFieldType.startswith("["):
                gqlFieldType = field[1][1:-1]
                isArray = True

            if gqlFieldType in ["Int", "String", "Boolean", "Float"]:
                goFieldType = first_lower(field[1])
            elif gqlFieldType == "ID":
                goFieldType = "graphql.ID"
            elif gqlFieldType == "Time":
                goFieldType = "(graphql.Time, error)"
            else:
                goFieldType = first_lower(gqlFieldType) + "Resolver"

            if isArray:
                goFieldType = "[]*" + goFieldType

            if field[2] == None:
                if gqlFieldType == "Time":
                    goFieldType = "(*graphql.Time, error)"
                else:
                    goFieldType = "*" + goFieldType

            methodTmpl = methodTmpl.replace("{$method}", first_upper(field[0]))
            tmpl += methodTmpl.replace("{$fieldType}", goFieldType)

        tmplnew = tmpl.replace("{$typeName}", first_lower(typeName))
        print(tmplnew)
