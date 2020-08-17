#cython: language_level=3

from libc.stdint cimport uintptr_t
import ctypes
import os
from PyChestBuild.identify import get_lib_name

resource_package = os.path.dirname(__file__)
resource_path = "PyChestBuild/" + get_lib_name()


lib = ctypes.cdll.LoadLibrary(os.path.join(resource_package, resource_path))
lib.FindChangepoints.argtypes = [ctypes.c_void_p, ctypes.c_int, ctypes.c_float]
lib.FindChangepoints.restype = ctypes.c_void_p

def init():
    pass

def find_changepoints(sequence, minimum_distance):

    double_array = (ctypes.c_double * len(sequence))(*sequence)
    c_length = ctypes.c_int(len(double_array) * 8)
    c_minimum_distance = ctypes.c_float(minimum_distance)

    c_changepoints = <char *><void *><uintptr_t>lib.FindChangepoints(double_array, c_length, c_minimum_distance)
    changepoints = []
    changepoint_count = int.from_bytes(c_changepoints[0:8], byteorder='little', signed=True)
    for i in range(1, changepoint_count + 1):
        changepoints.append(int.from_bytes(c_changepoints[i * 8:(i + 1) * 8], byteorder='little', signed=True))

    return changepoints

