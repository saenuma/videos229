package slideshow

func RegisterAll(tmplStore *map[string]string) {
	(*tmplStore)["slideshows/1"] = method1Conf
	(*tmplStore)["slideshows/2"] = method2conf
}
