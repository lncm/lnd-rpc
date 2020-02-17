from os import path

from setuptools import setup, find_packages

version = '0.9.0'
build = 5

with open(path.join(path.abspath(path.dirname(__file__)), 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name="lnd-rpc",

    version=f"{version}.post{build}",
    packages=find_packages(f"v{version}"),
    package_dir={"": f"v{version}"},

    description="gRPC bindings for various lnd versions",
    long_description=long_description,
    long_description_content_type='text/markdown',

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
        "google-api-core",
        "googleapis-common-protos",
        "grpcio",
        "protobuf",
    ]
)
