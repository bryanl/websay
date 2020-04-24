package say

import (
	"fmt"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func Say(message, bgColor, fgColor, t string) (string, error) {
	phrase := parsePhrase(message)
	cow, err := cowsay.NewCow(
		cowsay.Phrase(phrase),
		cowsay.Type(t))
	if err != nil {
		return "\n", fmt.Errorf("create cow: %w", err)
	}

	saying, err := cow.Say()
	if err != nil {
		return "", fmt.Errorf("cow say: %w", err)
	}

	return gen(saying, bgColor, fgColor), nil
}

func parsePhrase(message string) string {
	lines := make([]string, 0, 40)
	for _, line := range strings.Split(message, "\n") {
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func gen(msg, bgColor, fgColor string) string {
	return fmt.Sprintf(`
	<html lang="">
	<head>
		<title>Say</title>
		<style>
body {
  font-family: "Courier New", Courier, monospace;
  font-size: 40px;
  background: %s;
  color: %s;
}
.content {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content-item {
  max-width: 50%%;
}

.content-item--top {
  align-self: flex-start;
}

.content-item--bottom {
  align-self: flex-end;
}
		</style>
	</head>
	<body>
		<div class='content'>
			<div class='content-item content-item--top'></div>
			<div class='content-item'>
				<pre>%s</pre>
			</div>
			<div class='content-item content-item--bottom'></div>
		</div>
	</body>
	</html>`, bgColor, fgColor, msg)
}
