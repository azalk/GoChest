from setuptools import setup, Extension

main_module = Extension("GoChest",
                        sources=["GoChest.cpp"],
                        libraries=["python3.8"],
                        library_dirs=[".", "/Library/Frameworks/Python.framework/Versions/3.8/lib/python3.8/config-3.8-darwin"],
                        )

setup(
    name="GoChest",
    version="0.51",
    license="bsd-3-clause",
    description="Locating Changes in Highly Dependent Data with Unknown Number of Change Points",
    author="Lukas Zierahn",
    author_email="lukas@kappa-mm.de",
    url="",
    download_url="",
    keywords=["Changepoint Estimation", "Dependent Data", "Unkown Number of Change Points"],
    install_requires=["pybindgen", ],
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Build Tools",
        "License :: OSI Approved :: BSD License",
        "Programming Language :: Python :: 3",
    ],
    ext_modules=[main_module],
    package_data={"": ["libGoChest.so"]}
)
