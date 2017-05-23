# Step 2

## Given
Output of step 1

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

[Continue to step 3](../step3)