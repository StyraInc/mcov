# mcov

Minimum compatible [Open Policy Agent](https://www.openpolicyagent.org/) (OPA) version, or mcov for short, is a tiny
tool that, as the name implies, reports the minimum compatible OPA version for any given
[Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) files you point it at. This can be used to:

* Ensure that your policies are compatible with the OPA version you are running in production
* Check the impact on version requirements for any given change to policy
* That's it! There's literally nothing more to see

## Usage

```shell
$ mcov policies/
v0.37.0
```
