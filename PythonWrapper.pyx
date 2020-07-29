#cython: language_level=3

import ctypes
import time

lib = ctypes.cdll.LoadLibrary("./GoChest.so")
lib.FindChangepoints.argtypes = [ctypes.c_void_p, ctypes.c_int, ctypes.c_float]
lib.FindChangepoints.revtypes = ctypes.c_void_p


def find_changepoints(sequence, minimum_distance):

    double_array = (ctypes.c_double * len(sequence))(*sequence)
    c_length = ctypes.c_int(len(double_array) * 8)
    c_minimum_distance = ctypes.c_float(minimum_distance)


    changepoints = lib.FindChangepoints(double_array, c_length, c_minimum_distance)

    time.sleep(1)

    print("chpts")

    return 0.0

