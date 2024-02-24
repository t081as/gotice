# gotice - go notice
Go copyright notice generator.

Gotice generates a third-party copyright notice file based on the go modules in the `go.mod` file of a project.

Use `gotice help` for further information.

## Installation/Start

### go install

```
> go install pkg.tk-software.de/gotice@latest
> gotice version
```

### go run

```
> go run pkg.tk-software.de/gotice@latest version
```

### Releases
Releases are available here: https://gitlab.com/tobiaskoch/gotice/-/releases

## Usage

### Initialization
```
gotice init
```

Gotice will create a `.gotice.json` file with default values in the current working directory.

### Generation
```
> gotice generate ./project1 ./project1/dist/NOTICE.txt
```

Gotice will generate the file `./project1/dist/NOTICE.txt` bases on the `go.mod` file in the directory `./project1`.
If there is a `.gotice.json` configuration file in the `./project1` directory it will be used.

## Configuration

**.gotice.json**

```
{
    "template":"built-in:txt",
    "rendering":"text"
}
```

### template
The template that shall be used for the generation of the notice file.
Supported values:

- **built-in:txt**: the build-in text template (default)
- **built-in:md**: the build-in markdown template
- **built-in:html:** the built-in html template
- **custom filename:** if a custom filename is provided gotice will try to locate the file in the project directory; relative paths are supported

#### Example
The following example shows a custom gotice template file:

```
NOTICES AND INFORMATION

This software incorporates material from third parties listed below.

{{range .}}-------------------------------------------------------------------
{{ .Path }} {{ .Version}}

{{ .LicenseText}}

-------------------------------------------------------------------
{{end}}
```

### rendering
The rendering engine that shall be used to render the template file.
Supported values:

- **text**: uses  `text/template` (default)
- **html**: uses `html/template`

## Contributing
see [CONTRIBUTING.md](CONTRIBUTING.md)

## Donating
Thanks for your interest in this project. You can show your appreciation and support further development by [donating](https://www.tk-software.de/donate).

## License
**GoMod** © 2023-2024 [Tobias Koch](https://www.tk-software.de). Released under a [BSD-style license](https://gitlab.com/tobiaskoch/gomod/-/blob/main/LICENSE).