## Introduction

Code to automate the detection of `.md` files and convert them to `.html`.

## Development

- Install the `requirements-dev.txt` file.
- Initialize pre-commit, run the command `pre-commit install`.

### Test

Run, in the same path of this `README.md` file:

```bash
make test
```

## TODO

Note. This section describes the desired logic, the current implementation can be different.

### Logic

The process has three main parts:

1. Get path names of the files to convert.
2. Calculate the path names for the html results.
3. Convert files.

Now, we are going to review each part in more detail.

#### Get paths of the files to convert

##### Logic

- Get all files' path names in all subfolders with `.md` extension .
- Create output path for the result file.

##### Input

Main path to analyze.

Example:

```bash
/home/md-folder
```

##### Output

A `.txt` file called `input.txt` with detected path names.

Example:

- File path name: `/tmp/html-from-md/input.txt`.
- File content:

```bash
/home/md-folder/file-1.md
/home/md-folder/folder-1/file-2.md
```

#### Calculate the path names for the html results

##### Logic

- Change the files `.md` extensions to `.html`.
- Combine main output path with the path names of the files to convert.

##### Input

- Main path where the `.html` files must be stored.
- Path name of the `.txt` file where the files to convert has been stored.

##### Output

A `.txt` file called `output.txt` with the path names.

Example:

- File path name: `/tmp/html-from-md/output.txt`.
- File content:

```bash
/tmp/html-from-md/src/file-1.html
/tmp/html-from-md/src/folder-1/file-2.html
```

#### Convert files

##### Logic

- Assert `input.txt` and `output.txt` have equal number of rows.
- For each combination of the rows in the `input.txt` and `output.txt` files:
  - Create output paths if do not exist.
  - Convert the `.md` file to `.html`.

##### Input

- Path names of the `input.txt` and `output.txt` files.

##### Output

- Converted files.

### Design decisions explanation

#### Why to save the results to .txt files instead of in memory

This allow to separate the responsibility of each part of the process.
