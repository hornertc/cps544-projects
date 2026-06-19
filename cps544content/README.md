# CPS 544 Course Content

This repository contains the [syllabus](./syllabus.md), [schedule](./schedule.md), and all the content for the course.  To view the slides, see [slides](#slides).  There are also review sheets [here](./review/).

## Slides

Slides can be found [here](./slides/)

### Viewing Slides

Here are several choices, ranging from the most basic to the most intricate they are:

1. View the raw markdown as a text file.  *Because you are a software developer and things don't have to look pretty for you to understand them.*
1. View the rendered markdown in the web browser by simply navigating to the GitHub.com site [here](https://github.com/ktarplee-courses/cps544-content/blob/main/slides/go/ch1-intro.md).  *They will not look like slides but they will be pretty.*
1. Use the markdown viewer in VSCode to view the slides.   *They will not look like slides but they will be pretty.*
1. Install NodeJS (with `npm`) then install [reveal-md](https://github.com/webpro/reveal-md) with `npm install -g reveal-md` and finally run `reveal-md slides/` or to watch changes and not open the browser run `reveal-md slides/ -w --disable-auto-open`.  *This is how the instructor shows slides during the lecture.*

#### Alternatives

1. Install the VSCode extension `evilz.vscode-reveal` to view them in slide form within VSCode.  This extension has strange quirks due to an out of date reveal.js and lack of maintenance.

### Exporting Slides to PDF

To export all slides to PDF format, you can use the provided [export-slides](./bin/export-slides) script like so `export-slides ./slides ./pdfs`.

*The script is useful for creating offline copies or preparing printable versions of the slides.*

### Development

- Highlight JS (code blocks) styles can be found [here](https://github.com/highlightjs/highlight.js/tree/main/src/styles)

## Contributing

Feel free to contribute corrections and content to this repository.  Students do not have write access to this repository but can create issues, forks, and pull requests in this repository.
