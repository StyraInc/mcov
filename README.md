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

## OPA Versions Cheatsheet

Below lists additions to the Rego language, as presented in the OPA capabilities file for each version.
Features that may have an impact on the minimum compatible OPA version but are **not** covered by capabilities are
mentioned separately.

### v0.44.0

**New built-in functions**

- `strings.any_prefix_match`
- `strings.any_suffix_match`

### v0.43.0

**Not covered by capabilities**

- All `is_valid` functions made to return boolean false rather than throw errors

### v0.42.0

**Future keywords**

- `contains`
- `if`

**New built-in functions**

- `object.subset`

### v0.41.0

**New built-in functions**

- `graphql.is_valid`
- `graphql.parse`
- `graphql.parse_and_verify`
- `graphql.parse_query`
- `graphql.parse_schema`
- `units.parse`

### v0.40.0

**New built-in functions**

- `rego.metadata.chain`
- `rego.metadata.rule`

**Not covered by capabilities**

- Function mocking
- Assignment with `:=` allowed in all locations (rule heads, functions, object generating rules)

### v0.38.0

**Future keywords**

- `every`

**Not covered by capabilities**

- Metadata [annotations](https://www.openpolicyagent.org/docs/latest/policy-language/#annotations)

### v0.37.0

**New built-in functions**

- `object.union_n`
- `graph.reachable_paths`
- `indexof_n`

**Not covered by capabilities**

- `object.get`: accepting path argument as array

### v0.36.0

**New built-in functions**

- `crypto.hmac.md5`
- `crypto.hmac.sha1`
- `crypto.hmac.sha256`
- `crypto.hmac.sha512`
- `array.reverse`
- `strings.reverse`

**Miscellaneous**

- `allow_net` capability added

### v0.35.0

**New built-in functions**

- `net.lookup_ip_addr`

### v0.34.0

**Future keywords**

- `in`

**New built-in functions**

- `internal.member_2` (`in` operator)
- `internal.member_3` (`in` operator)
- `print`

### v0.33.0

**New built-in functions**

- `crypto.x509.parse_rsa_private_key`
