---
name: commit-work
description: "Create high-quality git commits: review/stage intended changes, split into logical commits, and write clear commit messages (including Conventional Commits). Use when the user asks to commit, craft a commit message, stage changes, or split work into multiple commits."
---

# Commit work

## Goal
Make commits that are easy to review and safe to ship:
- only intended changes are included
- commits are logically scoped (split when needed)
- commit messages describe what changed and why

## Workflow (checklist)
1) Inspect the working tree before staging
   - `git status`
   - `git diff` (unstaged)
   - If many changes: `git diff --stat`
2) Decide commit boundaries (split if needed)
   - Split by: feature vs refactor, backend vs frontend, formatting vs logic, tests vs prod code, dependency bumps vs behavior changes.
   - If changes are mixed in one file, plan to use patch staging.
3) Stage only what belongs in the next commit
   - Prefer patch staging for mixed changes: `git add -p`
   - To unstage a hunk/file: `git restore --staged -p` or `git restore --staged <path>`
4) Review what will actually be committed
   - `git diff --cached`
   - Sanity checks:
     - no secrets or tokens
     - no accidental debug logging
     - no unrelated formatting churn
5) Describe the staged change in 1-2 sentences (before writing the message)
   - "What changed?" + "Why?"
   - If you cannot describe it cleanly, the commit is probably too big or mixed; go back to step 2.
6) Write the commit message
   - The commit message should have the following:
     - First line
     - Blank line
     - Body
   - Prefer an editor for multi-line messages: `git commit -v`
   - See Appendix 2: Writing good CL descriptions for examples of good commit messages.
7) Run the smallest relevant verification
   - Run the repo's fastest meaningful check (unit tests, lint, or build) before moving on.
8) Repeat for the next commit until the working tree is clean

## Deliverable
Provide:
- the final commit message(s)
- a short summary per commit (what/why)
- the commands used to stage/review (at minimum: `git diff --cached`, plus any tests run)

# Appendix 1: Small CLs

## What is Small?

In general, the right size for a CL is **one self-contained change**. This means
that:

-   The CL makes a minimal change that addresses **just one thing**. This is
    usually just one part of a feature, rather than a whole feature at once. In
    general it's better to err on the side of writing CLs that are too small vs.
    CLs that are too large. Work with your reviewer to find out what an
    acceptable size is.
-   The CL should include related test code.
-   Everything the reviewer needs to understand about the CL (except future
    development) is in the CL, the CL's description, the existing codebase, or a
    CL they've already reviewed.
-   The system will continue to work well for its users and for the developers
    after the CL is checked in.
-   The CL is not so small that its implications are difficult to understand. If
    you add a new API, you should include a usage of the API in the same CL so
    that reviewers can better understand how the API will be used. This also
    prevents checking in unused APIs.

There are no hard and fast rules about how large is "too large." 100 lines is
usually a reasonable size for a CL, and 1000 lines is usually too large, but
it's up to the judgment of your reviewer. The number of files that a change is
spread across also affects its "size." A 200-line change in one file might be
okay, but spread across 50 files it would usually be too large.

Keep in mind that although you have been intimately involved with your code from
the moment you started to write it, the reviewer often has no context. What
seems like an acceptably-sized CL to you might be overwhelming to your reviewer.
When in doubt, write CLs that are smaller than you think you need to write.
Reviewers rarely complain about getting CLs that are too small.

## When are Large CLs Okay?

There are a few situations in which large changes aren't as bad:

-   You can usually count deletion of an entire file as being just one line of
    change, because it doesn't take the reviewer very long to review.

## Separate Out Refactorings

It's usually best to do refactorings in a separate CL from feature changes or
bug fixes. For example, moving and renaming a class should be in a different CL
from fixing a bug in that class. It is much easier for reviewers to understand
the changes introduced by each CL when they are separate.

Small cleanups such as fixing a local variable name can be included inside of a
feature change or bug fix CL, though. It's up to the judgment of developers and
reviewers to decide when a refactoring is so large that it will make the review
more difficult if included in your current CL.

## Keep related test code in the same CL

CLs should include related test code. Remember that smallness
here refers the conceptual idea that the CL should be focused and is not a
simplistic function on line count.

Tests are expected for all Google changes.

A CL that adds or changes logic should be accompanied by new or updated tests
for the new behavior. Pure refactoring CLs (that aren't intended to change
behavior) should also be covered by tests; ideally, these tests already exist,
but if they don't, you should add them.

*Independent* test modifications can go into separate CLs first, similar to the
refactorings guidelines. That includes:

*   Validating pre-existing, submitted code with new tests.
    *   Ensures that important logic is covered by tests.
    *   Increases confidence in subsequent refactorings on affected code. For
        example, if you want to refactor code that isn't already covered by
        tests, submitting test CLs *before* submitting refactoring CLs can
        validate that the tested behavior is unchanged before and after the
        refactoring.
*   Refactoring the test code (e.g. introduce helper functions).
*   Introducing larger test framework code (e.g. an integration test).

## Don't Break the Build

If you have several CLs that depend on each other, you need to find a way to
make sure the whole system keeps working after each CL is submitted. Otherwise
you might break the build for all your fellow developers for a few minutes
between your CL submissions (or even longer if something goes wrong unexpectedly
with your later CL submissions).

# Appendix 2: Writing good CL descriptions

A CL description is a public record of change, and it is important that it
communicates:

1.  **What** change is being made? This should summarize the major changes such
    that readers have a sense of what is being changed without needing to read
    the entire CL.

1.  **Why** are these changes being made? What contexts did you have as an
    author when making this change? Were there decisions you made that aren't
    reflected in the source code? etc.

The CL description will become a permanent part of our version control history
and will possibly be read by hundreds of people over the years.

Future developers will search for your CL based on its description. Someone in
the future might be looking for your change because of a faint memory of its
relevance but without the specifics handy. If all the important information is
in the code and not the description, it's going to be a lot harder for them to
locate your CL.

And then, after they find the CL, will they be able to understand *why* the
change was made? Reading source code may reveal what the software is doing but
it may not reveal why it exists, which can make it harder for future developers
to know whether they can move
[Chesterton's fence](https://abseil.io/resources/swe-book/html/ch03.html#understand_context).

A well-written CL description will help those future engineers -- sometimes,
including yourself!

## First Line

*   Short summary of what is being done.
*   Complete sentence, written in the imperative mood.
*   Follow by empty line.

The **first line** of a CL description should be a short summary of
*specifically* **what** *is being done by the CL*, followed by a blank line.
This is what appears in version control history summaries, so it should be
informative enough that future code searchers don't have to read your CL or its
whole description to understand what your CL actually *did* or how it differs
from other CLs. That is, the first line should stand alone, allowing readers to
skim through code history much faster.

Try to keep your first line short, focused, and to the point. The clarity and
utility to the reader should be the top concern.

By tradition, the first line of a CL description is a complete sentence, written
as though it were an order (an imperative sentence). For example, say
\"**Delete** the FizzBuzz RPC and **replace** it with the new system." instead
of \"**Deleting** the FizzBuzz RPC and **replacing** it with the new system."
You don't have to write the rest of the description as an imperative sentence,
though.

## Body is Informative {#informative}

The [first line](#first-line) should be a short, focused summary, while the rest
of the description should fill in the details and include any supplemental
information a reader needs to understand the changelist holistically. It might
include a brief description of the problem that's being solved, and why this is
the best approach. If there are any shortcomings to the approach, they should be
mentioned. If relevant, include background information such as bug numbers,
benchmark results, and links to design documents.

If you include links to external resources consider that they may not be visible
to future readers due to access restrictions or retention policies. Where
possible include enough context for reviewers and future readers to understand
the CL.

Even small CLs deserve a little attention to detail. Put the CL in context.

## Good CL Descriptions {#good}

Here are some examples of good descriptions.

### Functionality change {#functionality-change}

Example:

> RPC: Remove size limit on RPC server message freelist.
>
> Servers like FizzBuzz have very large messages and would benefit from reuse.
> Make the freelist larger, and add a goroutine that frees the freelist entries
> slowly over time, so that idle servers eventually release all freelist
> entries.

The first few words describe what the CL actually does. The rest of the
description talks about the problem being solved, why this is a good solution,
and a bit more information about the specific implementation.

### Refactoring {#refactoring}

Example:

> Construct a Task with a TimeKeeper to use its TimeStr and Now methods.
>
> Add a Now method to Task, so the borglet() getter method can be removed (which
> was only used by OOMCandidate to call borglet's Now method). This replaces the
> methods on Borglet that delegate to a TimeKeeper.
>
> Allowing Tasks to supply Now is a step toward eliminating the dependency on
> Borglet. Eventually, collaborators that depend on getting Now from the Task
> should be changed to use a TimeKeeper directly, but this has been an
> accommodation to refactoring in small steps.
>
> Continuing the long-range goal of refactoring the Borglet Hierarchy.

The first line describes what the CL does and how this is a change from the
past. The rest of the description talks about the specific implementation, the
context of the CL, that the solution isn't ideal, and possible future direction.
It also explains *why* this change is being made.

### Small CL that needs some context

Example:

> Create a Python3 build rule for status.py.
>
> This allows consumers who are already using this as in Python3 to depend on a
> rule that is next to the original status build rule instead of somewhere in
> their own tree. It encourages new consumers to use Python3 if they can,
> instead of Python2, and significantly simplifies some automated build file
> refactoring tools being worked on currently.

The first sentence describes what's actually being done. The rest of the
description explains *why* the change is being made and gives the reviewer a lot
of context.
