from pybindgen import retval, param, Module

mod = Module('GoChest')

# mod.add_container("GoSlice", "float", "set")

mod.add_include('"libGoChest.h"')
mod.add_function('FindChangePoints', retval('int'), [param('GoSlice', 'sequence'), param('float', 'minimumDistance')])
mod.generate(open("GoChest.c", "w"))
