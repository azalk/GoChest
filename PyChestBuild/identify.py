import platform
import sys


def get_lib_name():
    is_64bits = sys.maxsize > 2**32
    plat = platform.system().lower()
    name = ""

    if plat == "windows":
        if is_64bits:
            name = "windows.amd64.dll"
        else:
            name = "windows.386.dll"

    elif plat == "darwin":
        if is_64bits:
            name = "darwin.amd64.so"
        else:
            name = "darwin.386.so"

    elif plat == "linux":
        import os

        if os.uname()[4][:3] == "arm":
            if is_64bits:
                name = "linux.arm64.so"
            else:
                name = "linux.arm.so"
        else:
            if is_64bits:
                name = "linux.amd64.so"
            else:
                name = "linux.386.so"

    else:
        raise Exception("""
There is no pre compiled package for this Operating System.
Please compile the package from source. You can find instructions here: https://github.com/azalk/GoChest""")

    return "GoChest." + name
