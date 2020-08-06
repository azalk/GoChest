from setuptools import setup
from Cython.Build import cythonize

setup(
    name="GoChest",
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
    ext_modules=cythonize("GoChest.pyx"),
    package_data={'': ['GoChest.so']},
    include_package_data=True
)
