# Ginkgo BDD Testing Framework

## [Lightning Talks - May](https://www.meetup.com/Edinburgh-Golang-meetup/events/239797718/)

### [Edinburgh Golang Meetup](https://www.meetup.com/Edinburgh-Golang-meetup)

---

## What is BDD?

- **B**ehaviour **D**riven **D**evelopment
- A particular style of **T**est **D**riven **D**evelopment
- Focuses on expected *behaviour* of a system
- Uses non-technical language

Note:

Explain that BDD is a style of TDD that focuses on howw a system should be have.

TDD is out of scope for this talk.

---

## Example

```plain
Given: an apple
When: I eat the apple
Then: I expect to be left with the core
```

Note:

Doesn't matter if the apple lives inside MongoDB, or that the apple gets eaten by Golang application running inside a docker container. The test describes the behaviour and expected goal in using an english style description.

---

## When to use BDD

- Working with non-technical stakeholders
- Want descriptive, specification style tests
- Useful way to tease out edge cases.

Note:

## Working with non-technical stakeholders

Non-technical stakeholders at least have a chance of understanding and being able to talk about BDD style tests. Requirements could also be captured in this style and directly implemeneted as tests.

Also useful for moving some techincal discussions away from low level technicallities, to a higher level "what are we trying to achieve here" style discussion. Anti-bike shedding.

## Want descriptive, specification style tests

Well crafted BDD-style test tend to end up being very understandable and helps the team develop a common language for describing things.

BDD tools can leverage given, when, then style of the tests to provide context when test fail.

## Useful way to tease out edge cases