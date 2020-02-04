from setuptools import setup, find_packages

version = "0.4.1"

setup(
    version=f"{version}.post1",
    packages=find_packages(f"v{version}"),
    package_dir={"": f"v{version}"},
)
