from datetime import datetime
import re

def go_type(self, val):
    if isinstance(val, bool):
        return "bool"
    if isinstance(val, str):
        if re.match(r'^\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(\+\d\d:\d\d|Z)$', val):
            return "time.Time"
        else:
            return "string"
    if isinstance(val, (int, float)):
        if isinstance(val, int) and -2147483648 < val < 2147483647:
            return "int"
        else:
            return "int64" if isinstance(val, int) else "float64"
    if isinstance(val, list):
        b = {}
        for value in val:
            if isinstance(value, list):
                for item in value:
                    b[str(type(item).__name__)] = type(item).__name__
        if len(b)==1:
            return "[]"+str(type(val[0]).__name__)
        else:
            return "[]interface{}"
    elif isinstance(val, dict):
        return "struct"
    else:
        return "interface{}"
# def python_to_golang_type(py_type):
#     golang_types = {
#         int: "int",
#         float: "float64",
#         str: "string",
#         bool: "bool",
#         list: "[]interface{}",
#         dict: "map[string]interface{}",
#         type(None): "interface{}",
#     }

#     return golang_types.get(py_type, "interface{}")



json_input = '{"listdata":[]}' 


import json
data = json.loads(json_input)

for key, value in data.items():
    print(key, go_type(value))

# b=str(type(data["listdata"]).__name__)
# print(b)

# b = {}
# for key, value in data.items():
#     if isinstance(value, list):
#         for item in value:
#             b[str(type(item).__name__)] = type(item).__name__
#     else:
#         b[key] = value

# if len(b)==1:
#     print(b)
