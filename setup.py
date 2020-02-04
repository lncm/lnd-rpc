from setuptools import setup, find_packages

# NOTE: keep this single-quoted!
version = '0.9.0'

setup(
    name="lnd-rpc",

    version=f"{version}.post1",
    packages=find_packages(f"v{version}"),
    package_dir={"": f"v{version}"},

    description="gRPC bindings for various lnd versions",
    url="https://github.com/lncm/lnd-rpc",
    keywords=["ln", "lnd", "grpc", "proto"],
    license="MIT",
    author="Damian Mee",
    author_email="bugs@meedamian.com",
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.5",
    ],
    install_requires=[
        "protobuf",
        "google-api-core"
    ]
)
