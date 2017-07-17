from setuptools import setup, find_packages

setup(
    name='thrift-server',
    description='thrift python server',
    author='doctor',
    author_email="chenzhi@zhiweicloud.com",
    version=1.0,
    install_requires=['thrift'],
    packages=find_packages(),
    entry_points={
    }
)
