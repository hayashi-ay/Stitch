# Stitch

## Fetch program usage
Usage: ./fetch [--metadata] urls

Fetch web pages from the given urls.

If `--metadata` is specified, the metadata of the given pages are printed. In that case no retrieval occurs. 

Following is a sample output of fetch program.

```bash
/app # ./fetch https://hayapenguin.com https://www.google.com
/app # ls
fetch                 go.mod                go.sum                hayapenguin.com.html  main.go               main_test.go          www.google.com.html
/app # ./fetch --metadata https://hayapenguin.com https://www.google.com
site: hayapenguin.com
num_links: 1
images: 0
last_fetch: Thu Apr 07 2022 07:35 UTC 
-------
site: www.google.com
num_links: 17
images: 1
last_fetch: Thu Apr 07 2022 07:35 UTC 
```

## How to set up Docker container
You just need to run following two commands to run container.

```bash
docker build -t stitch .
docker run -it --rm stitch
```

Then, prompt shows and you can run `fetch` program there.

```bash
hayashi:Stitch[main *]$ docker run -it --rm stitch
/app # Type anything you wants.
```

# Feature Improvement
- Handle urls with file paths

The program is currently not able to handle an url with a file path(ex. https://hayapenguin.com/preseed.cfg). I can achieve this by escaping slash. 

- Improve Error Management

Error messages are displayed as it is in the internal function call in some parts, it should be worded appropriately.
Also, whether to continue processing for each error is unclear and needs to be reconsidered as a whole.

- Use goroutine to speed up fetch program

The program is currently executed synchronously but fetch process is independent for each url, so it can be improved by execute by parallel.

- Many other improvement
I use `os.ReadFile` but it loads file in memory at once, so perhaps I may read little by little. 
I wrote all cods in `main.go, but it should be structured.

# Note
I don't use [flag package](https://pkg.go.dev/flag) because it does not support flags starts with double dash(`--`).
