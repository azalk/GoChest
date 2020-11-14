from setuptools import setup, Extension
from Cython.Build import cythonize
from PyChestBuild.identify import get_lib_name
import os

library_path = get_lib_name()

# If the user has compiled their own shared library, then we are using that one
if os.path.isfile("PyChestBuild/GoChest.so"):
    library_path = "GoChest.so"
    print("Using found GoChest.so over precompiled libraries")

if os.path.isfile("PyChestBuild/GoChest.dll"):
    library_path = "GoChest.dll"
    print("Using found GoChest.dll over precompiled libraries")

extensions = [
    Extension("PyChest", ["PyChest/PyChest.pyx"])
]

setup(
    name="PyChest",
    version="0.54",
    license="bsd-3-clause",
    description="Locating Changes in Highly Dependent Data with Unknown Number of Change Points",
    author="Lukas Zierahn",
    author_email="lukas@kappa-mm.de",
    url="",
    download_url="",
    keywords=["Changepoint Estimation", "Dependent Data", "Unknown Number of Change Points"],
    install_requires=["setuptools", "Cython"],
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Build Tools",
        "License :: OSI Approved :: BSD License",
        "Programming Language :: Python :: 3",
    ],
    ext_modules=cythonize(extensions),
    data_files=[('PyChestBuild', [f"PyChestBuild/{library_path}", "PyChestBuild/identify.py"])],
)
