# University of Dayton

## CPS 544: Advanced UNIX/Linux Programming (3 hours)

### Fall 2024

- Instructor: Kyle M. Tarplee, Ph.D., PMI-ACP
- Email: <tarpleek1@udayton.edu>
- Cell Phone: 970-231-0663

Google chat (to the email above) is the instructor's preferred contact method.  Please be mindful to keep related conversations in the same thread (email or chat).  It is helpful for the instructor (and for the student) to have the prior conversation handy when responding to questions.  Starting a new thread (changing the subject line in email) but continuing the conversation from a prior email is considered poor professional conduct.  

## Course Description

Prepares students for developing software in the UNIX/Linux environment using the Go programming language.

Possible topics include libraries and system calls, shells, system structures and internals, inter-process communication (pipes and signals), communicating sequential processes, network programming (client-server model and sockets), configuration and compilation management, pattern matching and filters, shell programming, automatic program generation, and GUI programming.

Assignments are designed to provide students with pragmatic exposure to these tools as well as issues faced by modern practitioners. CPS 544 is a programming-intensive course.

<!-- CPS 544 catalog: The study of advanced topics related to software development in the UNIX/Linux environment. Topics may include advanced inter-process communication (IPC), systems administration, virtualization and security -->

<!-- CPS 444 catalog: Prepares students for developing software in the UNIX/Linux environment using the C programming language. Topics include system libraries and system calls, shells, system structures and internals, inter-process communication (pipes and signals), network programming (client-server model and sockets), pattern matching and filters, shell programming, automatic program generation, and GUI programming. Prerequisite(s): CPS 356. -->

## Goals/Learning Objectives

- Nearly complete understanding of the Go language syntax and semantics (yet only partial knowledge of the Go standard library).  This is possible because Go prides itself on simplicity.
  - Proficient in developing highly concurrent applications in Go
- Working knowledge of how to develop applications within UNIX/Linux operating systems
  - Perform everyday tasks on the UNIX command line
  - Basic understanding of GNU Make
  - Proficient in writing small shell scripts
  - Performs common operations with Git
  - Can submit an issue and associated merge request on Github via the fork model
  - Proficient in writing Markdown style documentation
- Partial understanding of the interface between applications and operating systems facilities/services

## Prerequisites

<!-- - CPS 356 Operating Systems (for undergraduates, CPS 444) -->
- CPS 501 Advanced Programming and Data Structures <!-- (for graduates, CPS 544) -->
- Experience programming in any statically typed language (e.g., C/C++, Java, C#)
- Exposure to filesystems, concurrency, and multi-threaded systems is not required but preferred (e.g., a course in Operating Systems)

## Organization

### Lecture

Tuesday and Thursday 1705-1820 ET (virtual unless otherwise noted below).  First day of classes is August $19^{th}$.  First lecture Tuessday, August $20^{th}$.

During class time (which is virtual) students are expected to keep their camera on at all times and be ready to field questions (and thus un-mute) at a moments notice.  Therefore students must ensure they are in a suitable location to participate in the course in this manor.

### Office hours

- Monday 1300-1400 ET
- Wednesday 1400-1500 ET

Please come to the beginning of office hours.  If no students are present after 10 minutes the instructor may cancel that office hour.

Additional in-person office hours will be held in front of the Computer Science department if needed and/or requested.

### Exams

Exams are held in-person at Jessie Hathcock Hall 050.

- First mid-term exam - Tuesday, October $1^{st}$ during lecture
- Second mid-term exam - Tuesday, November $12^{th}$ during lecture
- Final exam - Tuesday, December $10^{th}$ at 1630-1820 ET

## Resources

### Textbooks

- (Required) The Go Programming Language, $1^{st}$ Edition, by Alan A.A. Donovan and Brian W. Kernighan, Addison-Wesley Professional Computing Series, 2015, ISBN: 978-0-13-419044-0
- (Required) Shell Programming in Unix, Linux, and OS X (Developer's Library), $4^{th}$ Edition, by  Stephen Kochan and Patrick Wood, ISBN 978-0134496009
- (Optional) The Linux Programming Interface, A Linux and UNIX System Programming Handbook, by Michael Kerrisk, 2010, ISBN 978-1-59327-220-3
- [Course Content](https://github.com/ktarplee-courses/cps544-content) includes the following
  - Slides
  - Review Sheets
  - Syllabus
  - Schedule
- [Go Programming Language example code](https://github.com/ktarplee-courses/gopl.io)

### Computing

Students are required to have a laptop for this course.  The laptop must be capable of running a small Linux virtual machine.

## Evaluation

### Grade Breakdown by Component

- Homework/projects: 40%
- Quizzes: 0% (practice for exams)
- First mid-term exam: 15%
- Second mid-term exam: 20%
- Final exam: 25%

Exams are cumulative and may cover material in any or all of the text, lectures, and projects.  As such, it is paramount that students diligently read the text.

If students have extenuating circumstances preventing homework from being turned in on time or attending an examination, please talk to the instructor ASAP (ideally before the due date) so accommodations can be made on a case-by-case basis.

All assignments will be posted on Canvas but turned in on GitHub to give students exposure to Git and GitHub Actions. This approach also gives students near constant feedback as they work on assignments.  To pass this course **all** assignments must pass **all** the acceptance tests by the end of the semester.

Assignments are worth 10 points each.  Late work will lose one point for each day it is late, capped at a 5-point deduction.  Solutions to assignments will be reviewed in class and once that happens the penalty jumps to 5 points (e.g., a maximum score of 50%).  If you would like an assignment regraded it **must** pass all the CI tests perfectly before being reconsidered (the maximum points is still 5).  No late work will be accepted after the last day of lecture.  Please reach out for help on assignments prior to the due date.

### Standard Grade Scheme

| Percentage | Letter |
| ---------- | ------ |
| ≥93% | A  |
| ≥90% | A- |
| ≥87% | B+ |
| ≥83% | B  |
| ≥80% | B- |
| ≥70% | C  |
| ≥0%  | F  |

<!-- | Percentage | Letter |
| ---------- | ------ |
| ≥93% | A  |
| ≥90% | A- |
| ≥87% | B+ |
| ≥83% | B  |
| ≥80% | B- |
| ≥77% | C+ |
| ≥73% | C  |
| ≥70% | C- |
| ≥60% | D  |
| ≥0%  | F  | -->

C- and D are not allowed by UD's graduate school thus a C- maps to a C and a D maps to an F.

<!-- Remember, your grades do not define you, your relationship with Christ does! -->

## Attendance Policy

Class attendance is fundamental to the teaching/learning process and any absence from a class results in a loss of learning for the student and learning community.  **As per UD's [policy](https://catalog.udayton.edu/undergraduate/generalinformation/academicinformation/classattendancepolicy/), attendance in the first week of class is mandatory**.  Lectures are generally not recorded to encourage students to attend the lectures and take scrupulous notes.  Attendance will be taken and for each lecture missed, without an acceptable excuse, will result in the **loss of one percentage point from your overall grade in the course**.  For example, if a students misses two lectures and would normally receive a 91% (A-) in the class, the student will instead receive a 89% (B+) for the course.  The course attendance is only taken into account for midterm and final grades.  This has the unfortunate consequence that the grade reflected in the Canvas gradebook higher than the grade ultimately recorded in Banner as the official grade.

## Academic Honesty

Students must abide by [UD's graduate school academic honor code](https://catalog.udayton.edu/graduate/generalinformation/academicinformation/academicdishonesty/).

**Use of Large Language Models (e.g., ChatGPT) are prohibited.**  The purpose of this course is to train the student's natural neural network between their ears (e.g., their brain).  Not teach it to rely on artificial neural networks.

All work must be completed **individually** for this course.  There is no group work, however working together to understand the problem and discuss the solutions is encouraged.  However, writing code together or sharing code is strictly prohibited.  Please attend office hours to ask the instructor for help on assignments.

For any given assignment, if two or more submissions look too similar, then all parties involved will be deemed to have plagiarized.  Students must realize that the instructor cannot differentiate between the student(s) that copied the work and the one(s) that did the work themselves.  This is why it is important that students protect their work from their friends to avoid any suspicion of plagiarism.  **Students are strongly discouraged from allowing others to directly look at their code.**

Anyone caught cheating on an assignment will be questioned about their submission.  If cheating is deemed to have occur then zero points will be awarded for the assignment and 10 percentage points will be subtracted from their overall grade.  This penalty will be applied for each occurrence of cheating.

Anyone caught cheating on an exam will receive a zero on the exam.

## Student Feedback Tool

The University will ask for your anonymous feedback regarding instruction in this course through the online Student Feedback Tool (SFT). Your candid, respectful opinions and constructive suggestions have an impact on the quality of teaching at UD. Instructions for how to complete SFT will be sent to your UD email account towards the end of the term and you will be given time in class to complete it. If you encounter technical problems accessing SFT, contact the UDit Service Center at 937-229-3888 or <itservicecenter@udayton.edu>. To learn more about SFT, visit <https://go.udayton.edu/sft>.

## Course Topics

Below is a general list of topics covered in this course grouped by topic.  There is also chronological view available in the [schedule](./schedule.md).

- UNIX
  - Brief history of UNIX and Linux
  - UNIX basics (shells, ssh, ls, cat, grep, which, top, and environment variables)
  - Advanced UNIX usage (package management, shell customization)
  - UNIX network programming from the command line (netcat/nc, socat, HTTP client/server)
  - Basic building blocks of Linux (linux kernel, processes, and /proc)

- POSIX Shell Scripting
  - Variables, flow control, functions, arithmetic
  - I/O redirection, file I/O
  - Signals
  - Concurrency

- Go Programming
  - Overview of Go (program structure including package management, command line interface, strings, branching and looping) (Go chapter 1 and 10)
  - Go program structure and basic types (naming, numbers, constants, and variables) (Go chapter 2 and 3)
  - Go composite types (arrays, slices, maps, structs, tags, JSON) (Go chapter 4)
  - Go functions (declaration, recursion, return values, error handling, anonymous functions, variadic functions, deferred function calls, panic/recover, and testing) (Go chapter 5 and 11)
  - Go objects (interfaces, struct embedding, encapsulation, type assertions, switch statements) (Go chapter 6 and 7)
  - Network programming (HTTP client/server)
  - Concurrency (goroutines, channels, race conditions, mutual exclusion, lazy initialization, race detection, relationship to threads) (Go chapter 8 and 9)
