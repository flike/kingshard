# Color

Style terminal text.

## Installation

```sh
go get github.com/labstack/gommon/color
```

## Windows?

Try [cmder](http://bliker.github.io/cmder) or https://github.com/mattn/go-colorable

## [Usage](https://github.com/labstack/gommon/blob/master/color/color_test.go)

```sh
import github.com/labstack/gommon/color
```

### Colored text

```go
fmt.Println(color.Black("black"))
fmt.Println(color.Red("red"))
fmt.Println(color.Green("green"))
fmt.Println(color.Yellow("yellow"))
fmt.Println(color.Blue("blue"))
fmt.Println(color.Magenta("magenta"))
fmt.Println(color.Cyan("cyan"))
fmt.Println(color.White("white"))
fmt.Println(color.Grey("grey"))
```
![Colored Text](http://i.imgur.com/8RtY1QR.png)

### Colored background

```go
fmt.Println(color.BlackBg("black background", color.Wht))
fmt.Println(color.RedBg("red background"))
fmt.Println(color.GreenBg("green background"))
fmt.Println(color.YellowBg("yellow background"))
fmt.Println(color.BlueBg("blue background"))
fmt.Println(color.MagentaBg("magenta background"))
fmt.Println(color.CyanBg("cyan background"))
fmt.Println(color.WhiteBg("white background"))
```
![Colored Background](http://i.imgur.com/SrrS6lw.png)

### Emphasis

```go
fmt.Println(color.Bold("bold"))
fmt.Println(color.Dim("dim"))
fmt.Println(color.Italic("italic"))
fmt.Println(color.Underline("underline"))
fmt.Println(color.Inverse("inverse"))
fmt.Println(color.Hidden("hidden"))
fmt.Println(color.Strikeout("strikeout"))
```
![Emphasis](http://i.imgur.com/3RSJBbc.png)

### Mix and match

```go
fmt.Println(color.Green("bold green with white background", color.B, color.WhtBg))
fmt.Println(color.Red("underline red", color.U))
fmt.Println(color.Yellow("dim yellow", color.D))
fmt.Println(color.Cyan("inverse cyan", color.In))
fmt.Println(color.Blue("bold underline dim blue", color.B, color.U, color.D))
```
![Mix and match](http://i.imgur.com/jWGq9Ca.png)

### Enable/Disable the package

```go
color.Disable()
color.Enable()
```

### New instance

```go
c := New()
c.Green("green")
```
