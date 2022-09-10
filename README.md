![alt text](https://raw.githubusercontent.com/n25a/doc-hunt/1f3b6c85b8d1d7050299e08d05df1bc835bab243/images/logo.png)

# Doc-Hunt
Doc-Hunt is godoc checker in project.

## Usage
You can install it by following command.

```bash
$ go get github.com/n25a/doc-hunt
```

Then, you can run it by following command.

```bash
$ doc-hunt hunt --dir "/path/to/project"
```

## Options
You can use these flags to execute doc-hunt files or directory.

|      flag      |        description         |
|:--------------:|:--------------------------:|
| --exclude-path | Set directories to ignore  |
| --exclude-file |    Set files to ignore     |


## Reports
Doc-Hunt print a report in terminal.

If your code have some problems, it will print a report like this:
![alt text](https://raw.githubusercontent.com/n25a/doc-hunt/1f3b6c85b8d1d7050299e08d05df1bc835bab243/images/failed.png)

If your code have no problems, it will print a report like this:
![alt text](https://raw.githubusercontent.com/n25a/doc-hunt/1f3b6c85b8d1d7050299e08d05df1bc835bab243/images/Success.png)
