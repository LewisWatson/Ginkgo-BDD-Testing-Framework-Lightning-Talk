# Ginkgo BDD Testing Framework

## [Lightning Talks - May](https://www.meetup.com/Edinburgh-Golang-meetup/events/239797718/)

### [Edinburgh Golang Meetup](https://www.meetup.com/Edinburgh-Golang-meetup)

---

## Behaviour Driven Development

- A particular style of **T**est **D**riven **D**evelopment
- Focuses on expected *behaviour* of a system
- Uses more natural english style language

Note:

Explain that BDD is a style of TDD that focuses on how a system should behave. Adopts an outside in view for writing behavioural specifications.

TDD is out of scope for this talk.

+++

## Example

```plain
Scenario: In order to maximise our enjoyment of apples, we need to eat them two at a time.
Given: 5 apples
When: I eat 2 apples
Then: I should have 3 apples
Then: I should have 2 apple cores.
```

Note:

By convention BDD generally focus on desired outcomes using less technical, and more natual language.

Doesn't matter if the apple lives inside MongoDB, or that the apple gets eaten by a Golang microsevice. The test describes the behaviour and expected goal using an english style description. Although there is nothing to stop us using some technical langauge, depending on the intended audience.

---

## When to use BDD

- Working with stakeholders with differeing specialities
- For more descriptive, specification style tests
- Useful way to structure tests and tease out edge cases

Note:

## Working with stakeholders with differing specialities

Non-technical stakeholders at least have a chance of understanding and being able to talk about BDD style tests. Requirements could also be captured in this style and directly implemeneted as tests.

Also useful for moving some techincal discussions away from low level technicallities, to a higher level "what are we trying to achieve here" style discussion. Anti-bike shedding.

## Want descriptive, specification style tests

Well crafted BDD-style test tend to end up being very understandable and helps the team develop a common language for describing things.

BDD tools can leverage given, when, then style of the tests to provide context when test fail.

## Useful way to tease out edge cases

At the very least it is a useful methodology for working through the behaviour of the system. Once you have some tests it is relatively easy to think of additional `when`'s for the `given`'s.

+++

## When not to use BDD

- If you want compact and succinct tests.
- If you want to quickly test something very low level and technical.

---

## Golang BDD

+++

## Golang BDD in the wild

- Cloud Foundry
  - Most of the projects
- Kubernetes
  - Only for integration and acceptance tests.
- https://github.com/LewisWatson/carshare-back/
- https://github.com/LewisWatson/firebase-jwt-auth

+++

## BDD options in Golang

Lots of options, but the main ones are

- [GoConvey](http://goconvey.co/)
- [Goblin](https://github.com/franela/goblin)
- [Godog](https://data-dog.github.io/godog/)
- [Ginkgo](https://onsi.github.io/ginkgo/)

--- 

## Demo

![ginko logo](https://onsi.github.io/ginkgo/images/ginkgo.png)