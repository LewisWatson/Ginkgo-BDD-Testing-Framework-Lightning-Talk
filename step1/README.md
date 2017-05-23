# Step 1

## Given
See provided [keys.go](keys.go) which contains a barebones go function for retrieving JSON encoded keys along with a max age header attribute from a url.

## When
Create Ginkgo suite

```bash
$ ginkgo bootstrap
Generating ginkgo test suite bootstrap for keys in:
        step1_suite_test.go
```

```bash
$ ginkgo generate keys
Generating ginkgo test for Keys in:
  keys_test.go
```

## Then

Create a simple test to ensure no errors are returned

[Continue to step 2](../step2)