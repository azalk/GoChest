from setuptools import setup, Extension
from Cython.Build import cythonize
from PyChestBuild.identify import get_lib_name

extensions = [
    Extension("PyChest", ["PyChest/PyChest.pyx"])
]

setup(
    name="PyChest",
    version="0.51",
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
    packages=['PyChestBuild'],
    package_data={'PyChestBuild': [get_lib_name()]},
)
