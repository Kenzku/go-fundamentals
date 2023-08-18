// 1. this is the package declaration.
// it can be either "main" or the name of the current "folder"
// the current folder is "menu". so we use "package menu"
// 2. all members in this file, like vars, functions, types etc are visible in the same package (folder) by default
// unless the member starts a Capital case letter, then it means totally public (outside the package)
// so if you use "import" -> members starting with lower case letter, would not be visible outside the package
// and members starting with Capital letter, would be available - like "public" in Java :)
// 3. there is no "private" scope

package menu

// since "menu" starts with lower case letter, it is only visible in the same package
var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}
