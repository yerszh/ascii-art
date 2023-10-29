# ascii-art
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

## Details
- The program works with only one argument [STRING] and banners like shadow and thinkertoy.
- Maximum number of characters is 30.
- Characters are separated by a new line \n.
- This project was written in Go and respects the good practices.


## Run Locally
* Clone the project

```bash
  git clone git@git.01.alem.school:yzhumyro/ascii-art.git
```

* Go to the project directory

```bash
  cd ascii-art
```

* Run

``` bash
  go run main.go "string" [banner]
```

## Usage

``` bash
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . "Hello There" | cat -e
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $                               
```
