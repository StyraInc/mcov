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

## Caveats

Note that not all features of Rego that have been added over time — or changes made to existing ones — are possible to track using capabilities alone. Use the version reported by mcov as a starting point — not as a replacement for testing compatibility!

## OPA Versions Cheatsheet

Below lists additions to the Rego language, as presented in the OPA capabilities file for each version.
Features that may have an impact on the minimum compatible OPA version but are **not** covered by capabilities are
mentioned separately.

### v0.59.0

**Features**

- `import rego.v1` (`rego_v1_import`)

### v0.57.0

**Features**

- Support for General References in Rule Heads (`rule_head_refs`)

**Not covered by capabilities**

- Short form `else` bodies

### v0.56.0

**New built-in functions**

- `numbers.range_step`

### v0.55.0

**New built-in functions**

- `crypto.parse_private_keys`

**Not covered by capabilities**

- Honor `default` keyword on functions

### v0.53.0

**New built-in functions**

- `crypto.x509.parse_keypair`

### v0.52.0

**New built-in functions**

- `crypto.hmac.equal`

### v0.50.0

**New built-in functions**

- `json.verify_schema`
- `json.match_schema`

### v0.48.0

**New built-in functions**

- `time.format`

### v0.47.0

**New built-in functions**

- `object.keys`
- `providers.aws.sign_req`

### v0.46.0

**Features**

- Refs in rule heads (`rule_head_ref_string_prefixes`)

**New built-in functions**

- `graphql.schema_is_valid`
- `net.cidr_is_valid`

**Not covered by capabilities**

- Entrypoint annotation
- `with`: Allow replacing functions with rules

### v0.45.0

**New built-in functions**

- `regex.replace`

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


## Community

For questions, discussions and announcements related to Styra products, services and open source projects, please join
the Styra community on [Slack](https://communityinviter.com/apps/styracommunity/signup)!