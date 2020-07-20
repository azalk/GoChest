from pybindgen import retval, param, Module

mod = Module('GoChest')
mod.add_include('"libGoChest.h"')
mod.add_function('Sum', retval('int'), [param('int', 'a'), param('int', 'b')])
# mod.add_function('test', None, [])
mod.generate(open("GoChest.c", "w"))
