# gdndebug

## User Stories

A command `gdndebug` is available to me.
  As a Developer
  When I run `gdndebug`
  Then it exits with an exit code of 0

  Acceptance:
    * Run `gdndebug`
    * Run `echo $?`
    * Confirm that the exit code is 0

---

List namespace inode numbers for current process.
  As a Developer
  When I run `gdndebug namespaces`
  Then I see the list of namespace inode numbers for the current process printed to stdout

  Acceptance:
    * Run `ls -lah /proc/self/ns` and take note of the namespace inode numbers
    * Run `gdndebug namespaces`
    * Run `echo $?`
    * Confirm that the exit code is 0
    * See a list of namespace inode numbers printed to stdout
    * See that the namespace inode numbers match the numbers reported from the first step of acceptance

---

List namespace inode numbers for a different process (one that is running in a separate set of namespaces).
  As a Developer
  When I run `gdndebug namespaces -p <pid of some other process>
  Then I see the list of namespace inode numbers for that process printed to stdout

  Acceptance:
    * Run `unshare -xxxx /bin/sh`
    * Run `ls -lah /proc/self/ns` and take note of the namespace inode numbers
    * In another terminal session, run `gdndebug -p <pid of the unshared /bin/sh command>`
    * Run `echo $?`
    * Confirm that the exit code is 0
    * See a list of namespace inode numbers printed to stdout
    * See that the namespace inode numbers match the numbers reported from the second step of acceptance

---

List a count of independent namespace instances for each type of namespace
  As a Developer
  When I run `gdndebug namespaces -c`
  Then I see the list of available namespaces along with a count for each namesapce type printed to stdout

  Acceptance:
    * Run `gdndebug namespaces -c`
    * Run `echo $?`
    * Confirm that the exit code is 0
    * See a list of available namespaces along with a count for each namesapce type printed to stdout (I would expect the count for each type to be 1, but it may not be)
    * In another terminal session, run `unshare -xxxx /bin/sh`
    * In the original terminal session, run `gdndebug namespaces -c`
    * See that the count for each namespace type has incremented by 1


